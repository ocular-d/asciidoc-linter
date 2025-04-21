package rules

import (
    "strings"
)

type HeadingSurroundRule struct{}

func (r HeadingSurroundRule) Name() string {
    return "AD002"
}

func (r HeadingSurroundRule) Description() string {
    return "Headings must be followed by a blank line (unless followed by attributes)"
}

func (r HeadingSurroundRule) Apply(fileName string, lines []string) []LintResult {
    var results []LintResult

    for i, line := range lines {
        if isHeading(line) {
            // Check next non-blank line after heading
            for j := i + 1; j < len(lines); j++ {
                trimmed := strings.TrimSpace(lines[j])
                if trimmed == "" {
                    // Blank line found → all good
                    break
                } else if strings.HasPrefix(trimmed, ":") {
                    // Attribute line → ignore
                    break
                } else {
                    // Content line immediately after heading with no spacing
                    results = append(results, LintResult{
                        File:     fileName,
                        Line:     j + 1,
                        RuleName: r.Name(),
                        Message:  r.Description(),
                    })
                    break
                }
            }
        }
    }

    return results
}
