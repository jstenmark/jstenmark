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
	header       = "---\n#### :cookie: Fortune cookie of the day"
	fortuneBlock = "```smalltalk\n%s\n```"
	fortuneRegex = "(?s)(%s)(\\n+)(```smalltalk)(.*?)(```)(\\n*)"
)

func main() {
	readmePath := flag.String("readme", "README.md", "Path to the README file")
	flag.Parse()

	fortuneOutput, err := runFortune(flag.Args())
	handleError(err, "running fortune")

	echoboxOutput, err := runEchobox(string(fortuneOutput))
	handleError(err, "running echobox")

	err = updateReadme(*readmePath, string(echoboxOutput))
	handleError(err, "updating README")

	fmt.Println("Fortune updated in markdown file successfully!\n", string(fortuneOutput))
}

func updateReadme(mdFile string, fortune string) error {
	fileContent, err := os.ReadFile(mdFile)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("reading file: %w", err)
	}

	newFortune := fmt.Sprintf(fortuneBlock, strings.TrimRight(fortune, "\n"))

	var updatedContent string
	re := regexp.MustCompile(fmt.Sprintf(fortuneRegex, regexp.QuoteMeta(header)))

	if strings.Contains(string(fileContent), header) {
		// Replace the existing fortune content
		updatedContent = re.ReplaceAllString(string(fileContent), header+"$2"+newFortune)
	} else {
		// Append new content if header doesn't exist
		updatedContent = string(fileContent) + "\n" + header + "\n" + newFortune + "\n"
	}

	// Write the updated content back to the file
	return os.WriteFile(mdFile, []byte(updatedContent), 0644)
}

func runFortune(args []string) ([]byte, error) {
	fortuneCmd := exec.Command("fortune", args...)
	return fortuneCmd.Output()
}

func runEchobox(input string) ([]byte, error) {
	echoboxCmd := exec.Command("echobox", "-S", "curved", "-s", "2")
	echoboxCmd.Stdin = strings.NewReader(input)
	return echoboxCmd.Output()
}

func handleError(err error, context string) {
	if err != nil {
		log.Fatalf("Error %s: %v", context, err)
	}
}
