package regex_parser

import (
	"fmt"
	"regexp"
	"strings"
)

func Parse(line, inputPattern string) (map[string]string, error) {
	// Escape parentheses in the input pattern to ensure they are treated as literal characters in the regex.
	escapedPattern := strings.ReplaceAll(inputPattern, "(", `\(`)
	escapedPattern = strings.ReplaceAll(escapedPattern, ")", `\)`)

	// Dynamically replace placeholders in the escaped pattern with regex patterns for capturing groups.
	placeholderPattern := regexp.MustCompile(`%(\w+)`)
	regexPattern := placeholderPattern.ReplaceAllString(escapedPattern, `(?P<$1>[^,)]+)`)


	// Compile the final regex pattern.
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil, err
	}

	// Match the line with the compiled regex.
	matches := re.FindStringSubmatch(line)
	if matches == nil {
		return nil, fmt.Errorf("no matches found")
	}

	// Map to hold the parsed values.
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i > 0 && name != "" {
			result[name] = matches[i]
		}
	}

	return result, nil
}