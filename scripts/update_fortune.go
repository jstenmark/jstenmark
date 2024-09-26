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

const (
	defaultFile    = "README.md"
	header         = "---\n#### :cookie: Fortune cookie of the day"
	contentBlock   = "```smalltalk\n%s\n```"
	contentPattern = "(?s)(%s)(\\n+)(```smalltalk)(.*?)(```)(\\n*)"
)

var (
	fortuneArgs  = []string{"computers", "linux", "linuxcookie"}
	contentRegex = regexp.MustCompile(fmt.Sprintf(contentPattern, regexp.QuoteMeta(header)))
)

func main() {
	filePath := flag.String("f", defaultFile, "Path to markdown file")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = fortuneArgs
	}

	fileContent, err := readFile(*filePath)
	handleError(err, "reading file")

	fortuneOutput, err := runFortune(args)
	handleError(err, "running fortune")

	boxOutput := formatInBox(string(fortuneOutput))

	err = updateFile(fileContent, *filePath, string(boxOutput))
	handleError(err, "updating file")

	fmt.Println("Fortune updated successfully\n", string(boxOutput))
}

func updateFile(fileContent string, filePath string, content string) error {
	formattedContent := fmt.Sprintf(contentBlock, strings.TrimSpace(content))
	updatedContent := updateFortuneContent(fileContent, formattedContent)

	if !strings.HasSuffix(updatedContent, "\n") {
		updatedContent += "\n"
	}

	return os.WriteFile(filePath, []byte(updatedContent), 0644)
}

func updateFortuneContent(fileContent string, newContent string) string {
	if contentRegex.MatchString(fileContent) {
		return contentRegex.ReplaceAllString(fileContent, header+"$2"+newContent)
	}
	return fileContent + "\n" + header + "\n" + newContent + "\n"
}

func runFortune(args []string) ([]byte, error) {
	return exec.Command("fortune", args...).Output()
}

func readFile(fileName string) (string, error) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil // Return empty content if file does not exist
		}
		return "", fmt.Errorf("reading file %s: %w", fileName, err)
	}
	return string(fileContent), nil
}

func handleError(err error, context string) {
	if err != nil {
		log.Fatalf("Error %s: %v", context, err)
	}
}

func formatInBox(text string) string {
	text = convertTabs(text, 4)

	lines := strings.Split(text, "\n")

	// Determine the maximum line length
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	// Create the top and bottom borders using rounded corners
	topBorder := "╭" + strings.Repeat("─", maxLength+2) + "╮"
	bottomBorder := "╰" + strings.Repeat("─", maxLength+2) + "╯"
	sideChar := "│"

	// Build the boxed content with padding
	var boxContent strings.Builder
	boxContent.WriteString(topBorder + "\n")
	for _, line := range lines {
		boxContent.WriteString(sideChar + " " + line + strings.Repeat(" ", maxLength-len(line)) + " " + sideChar + "\n")
	}
	boxContent.WriteString(bottomBorder)

	return boxContent.String()
}

func convertTabs(text string, tabWidth int) string {
	return strings.ReplaceAll(text, "\t", strings.Repeat(" ", tabWidth))
}
