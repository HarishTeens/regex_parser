package regex_parser

import (
	"fmt"
	"regexp"
)

func parse(line, inputPattern string) (map[string]string, error) {
	// First, dynamically replace placeholders in the input pattern with regex patterns.
	placeholderPattern := regexp.MustCompile(`%(\w+)`)
	// For each placeholder found, replace it with a regex named capture group.
	regexPattern := placeholderPattern.ReplaceAllString(inputPattern, `(?P<$1>[^,)]+)`)

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
