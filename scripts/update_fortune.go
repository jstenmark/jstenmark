// A script to add/update a markdown file with a "fortune" quote.
// Dependencies: fortunes, fortune-mod

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func printHelp() {
	fmt.Println(`Usage: update_fortune [options] [categories...]

Options:
  -f <path>    Specify the path to the markdown file to update (default: README.md).
  -b           Enable wrapping the fortune output in a decorative border (default: false).
  -h           Display this help message.

Categories:
  - Provide optional categories for the fortune command. These determine the type of fortune displayed.
  - Default categories: computers, linux, linuxcookie.
  - To view all available categories, run: "fortune -f".

Example:
  - Update a custom file with a bordered fortune and custom categories:
    go run update_fortune.go -f /app/FORTUNE.md -b ascii-art linux`)
}

const (
	defaultFile       = "README.md"
	defaultCategories = "ascii-art linux linuxcookie"
	header            = "---\n#### :cookie: Fortune cookie of the day"
	contentBlock      = "```smalltalk\n%s\n```"
	contentPattern    = "(?s)(%s)(\\n+)(```smalltalk)(.*?)(```)(\\n*)"
	tabWidth          = 4

	topLeftChar     = "╭"
	bottomLeftChar  = "╰"
	topRightChar    = "╮"
	bottomRightChar = "╯"
	sideChar        = "│"
	lineChar        = "─"
)

func main() {
	filePath := flag.String("f", defaultFile, "Path to the markdown file to update")
	useBoxOutput := flag.Bool("b", false, "Wrap the output in a border")
	showHelp := flag.Bool("h", false, "Show the help message")
	flag.Parse()

	if *showHelp {
		printHelp()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		args = strings.Split(defaultCategories, " ")
	}

	fileContent, err := readFile(*filePath)
	handleError(err, "Reading file")

	fortuneOutput, err := runFortune(args)
	handleError(err, "Running fortune")

	finalContent := string(fortuneOutput)
	if *useBoxOutput {
		finalContent = formatInBox(finalContent)
	}

	err = updateFile(fileContent, *filePath, finalContent)
	handleError(err, "Updating file")

	fmt.Println(finalContent)
}

func updateFile(fileContent string, filePath string, content string) error {
	formattedContent := fmt.Sprintf(contentBlock, strings.TrimSpace(content))
	updatedContent := updateFortuneContent(fileContent, formattedContent)

	// Avoid writes if content has not changed
	if fileContent == updatedContent {
		fmt.Println("No changes detected.")
		return nil
	}

	// Ensure file has a newline at the end
	if !strings.HasSuffix(updatedContent, "\n") {
		updatedContent += "\n"
	}

	return os.WriteFile(filePath, []byte(updatedContent), 0644)
}

func updateFortuneContent(fileContent string, newContent string) string {
	contentRegex := regexp.MustCompile(fmt.Sprintf(contentPattern, regexp.QuoteMeta(header)))
	if contentRegex.MatchString(fileContent) {
		return contentRegex.ReplaceAllString(fileContent, header+"$2"+newContent)
	}
	return fileContent + "\n" + header + "\n" + newContent + "\n"
}

func readFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", fmt.Errorf("Failed to read file %s: %w", fileName, err)
	}
	return string(content), nil
}

// Creates a decorative box with borders around the input text
func formatInBox(text string) string {
	text = convertTabs(text, tabWidth)
	lines := strings.Split(text, "\n")

	maxLength := maxLineLength(lines)

	// Generate borders and content.
	border := strings.Repeat(lineChar, maxLength+2)
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s%s%s\n", topLeftChar, border, topRightChar))

	// Add lines to the box
	for _, line := range lines {
		if line != "" {
			builder.WriteString(fmt.Sprintf("%s %s%s %s\n", sideChar, line, strings.Repeat(" ", maxLength-len(line)), sideChar))
		}
	}
	builder.WriteString(fmt.Sprintf("%s%s%s", bottomLeftChar, border, bottomRightChar))

	return builder.String()
}

func maxLineLength(lines []string) int {
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	return maxLength
}

func convertTabs(text string, tabWidth int) string {
	space := strings.Repeat(" ", tabWidth)
	return strings.ReplaceAll(text, "\t", space)
}

func runFortune(args []string) (string, error) {
	output, err := exec.Command("fortune", args...).Output()
	if err != nil {
		return "", fmt.Errorf("Failed to run fortune command: %w", err)
	}
	return string(output), nil
}

func handleError(err error, context string) {
	if err != nil {
		log.Fatalf("Error %s: %v", context, err)
	}
}
