package linter

import (
    "github.com/ocular-d/asciidoclint/internal/linter/ignore"
    "testing"
)

func TestIgnoreManager_ProcessDirectives(t *testing.T) {
    lines := []string{
        "// lint-disable AD002",
        "// lint-ignore AD001",
        "// lint-enable AD001",
        "// lint-enable AD002",
        "== Heading One",
        "Some content here",
        "// lint-ignore AD002",
        "== Heading Two",
    }

    im := ignore.NewIgnoreManager() // ✅ use the package prefix here
    im.ProcessDirectives(lines)

    tests := []struct {
        ruleName   string
        line       int
        isIgnored  bool
        isDisabled bool
    }{
        {"AD002", 4, false, false}, //Line 4: == Heading One — AD002 is not ignored or disabled
        {"AD001", 5, false, false}, //Line 5: "Some content here" — AD001 is not ignored (the ignore was ineffective)
        {"AD002", 7, true, false}, //Line 7: == Heading Two — correctly ignored by directive on line 6
    }

    for _, test := range tests {
        t.Run(test.ruleName, func(t *testing.T) {
            if im.IsRuleIgnored(test.line, test.ruleName) != test.isIgnored {
                t.Errorf("Expected ignore status %v for rule %s on line %d", test.isIgnored, test.ruleName, test.line)
            }
            if im.IsRuleDisabled(test.ruleName) != test.isDisabled {
                t.Errorf("Expected disable status %v for rule %s", test.isDisabled, test.ruleName)
            }
        })
    }
}
