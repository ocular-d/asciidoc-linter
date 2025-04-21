package linter

import (
	"strings"
    "github.com/ocular-d/asciidoc-linter/internal/linter/ignore"
    "github.com/ocular-d/asciidoc-linter/rules"
    "testing"
)

func TestLinter_LintFile(t *testing.T) {
    tests := []struct {
        name           string
        fileContent    string
        disabledRules  []string
        expectedErrors int
    }{
        {
            name: "AD002 disabled globally",
            fileContent: `// lint-disable AD002
== Heading One
Some content`,
            disabledRules: []string{"AD002"},
            expectedErrors: 0, // no AD002 error should be reported
        },
        {
            name: "AD001 ignored on specific line",
            fileContent: `// lint-ignore AD001
== Heading One
Some content
== Heading Two
No blank line here`,
            disabledRules:  []string{},
            expectedErrors: 2, // AD002 should still be triggered
        },
        {
            name: "No ignores, should fail AD002",
            fileContent: `== Heading One
Some content
== Heading Two
No blank line here`,
            disabledRules:  []string{},
            expectedErrors: 2, // AD002 should fail here due to missing blank line
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            linter := NewLinter()
            linter.RegisterRule(rules.HeadingSpacingRule{}) // AD001
            linter.RegisterRule(rules.HeadingSurroundRule{}) // AD002
            linter.Ignorer = ignore.NewIgnoreManager()

            // Split the file content into lines
            lines := splitToLines(tt.fileContent)

            // Mock LintFile process
            linter.Ignorer.ProcessDirectives(lines)

            errorCount := 0
            for _, rule := range linter.Rules {
                results := rule.Apply("test.adoc", lines)
                for _, res := range results {
                    if !linter.Ignorer.IsRuleIgnored(res.Line, rule.Name()) && !linter.Ignorer.IsRuleDisabled(rule.Name()) {
                        errorCount++
                    }
                }
            }

            if errorCount != tt.expectedErrors {
                t.Errorf("Expected %d errors but got %d", tt.expectedErrors, errorCount)
            }
        })
    }
}

func splitToLines(content string) []string {
    return strings.Split(content, "\n")
}
