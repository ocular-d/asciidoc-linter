package rules

import (
	//"fmt"
	"strings"
)

type LintResult struct {
	File     string
	Line     int
	RuleName string
	Message  string
}

type Rule interface {
	Name() string
	Description() string
	Apply(fileName string, lines []string) []LintResult
}

type HeadingSpacingRule struct{}

func (r HeadingSpacingRule) Name() string {
	return "AD001"
}

func (r HeadingSpacingRule) Description() string {
	return "No more than one blank line before a heading"
}

func (r HeadingSpacingRule) Apply(fileName string, lines []string) []LintResult {
	var results []LintResult

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if isHeading(line) && i >= 2 {
			if strings.TrimSpace(lines[i-1]) == "" && strings.TrimSpace(lines[i-2]) == "" {
				results = append(results, LintResult{
					File:     fileName,
					Line:     i + 1,
					RuleName: r.Name(),
					Message:  r.Description(),
				})
			}
		}
	}

	return results
}

func isHeading(line string) bool {
	trimmed := strings.TrimSpace(line)
	return strings.HasPrefix(trimmed, "=") // covers =, ==, ===, etc.
}
