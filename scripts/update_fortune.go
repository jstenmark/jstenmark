package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	readmePath := flag.String("readme", "README.md", "Path to the README file")
	flag.Parse()

	cmd := exec.Command("fortune")
	fortuneOutput, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running fortune: %v\n", err)
		return
	}

	mdFile := *readmePath
	header := "### Fortune cookie of the day"

	// Read the file content using os.ReadFile
	fileContent, err := os.ReadFile(mdFile)
	if err != nil && !os.IsNotExist(err) { // Handle errors other than file not existing
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Prepare the new fortune content
	newFortune := fmt.Sprintf("```\n%s```\n", string(fortuneOutput))

	// Create or update content under the header
	var updatedContent string
	// Use the header as part of the regex
	re := regexp.MustCompile(fmt.Sprintf("(?s)(%s)(\\n+)(```)(.*?)(```)(\\n*)", regexp.QuoteMeta(header)))

	if strings.Contains(string(fileContent), header) {
		// Replace the existing fortune content under the header
		updatedContent = re.ReplaceAllString(string(fileContent), header+"$2"+newFortune)
	} else {
		// If header doesn't exist, append the header and new fortune
		updatedContent = string(fileContent) + "\n\n" + header + "\n\n" + newFortune
	}

	// Write the updated content back to the file using os.WriteFile
	err = os.WriteFile(mdFile, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("Fortune updated in markdown file successfully!")
}
