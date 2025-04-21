package linter

import (
    "github.com/ocular-d/asciidoclint/internal/linter/ignore"
    "github.com/ocular-d/asciidoclint/rules"
    "fmt"
    "os"
    "strings"
)

type Linter struct {
    Rules []rules.Rule
    Ignorer *ignore.IgnoreManager
}

func NewLinter() *Linter {
    return &Linter{
        Rules:   []rules.Rule{},
        Ignorer: ignore.NewIgnoreManager(),
    }
}

func (l *Linter) RegisterRule(rule rules.Rule) {
    l.Rules = append(l.Rules, rule)
}

func (l *Linter) LintFile(fileName string) {
    content, err := os.ReadFile(fileName)
    if err != nil {
        fmt.Printf("❌ Error reading file: %s\n", err)
        return
    }

    lines := strings.Split(string(content), "\n")
    anyIssues := false

    // Process directives to set up ignored rules
    l.Ignorer.ProcessDirectives(lines)

    for _, rule := range l.Rules {
        ruleID := rule.Name()
        if l.Ignorer.IsRuleDisabled(ruleID) {
            continue
        }

        results := rule.Apply(fileName, lines)
        for _, res := range results {
            if l.Ignorer.IsRuleIgnored(res.Line, ruleID) {
                continue
            }

            anyIssues = true
            fmt.Printf("❌ %s:%d [%s] %s\n", res.File, res.Line, res.RuleName, res.Message)
        }
    }

    if !anyIssues {
        fmt.Printf("✅ %s passed all rules!\n", fileName)
    }
}

