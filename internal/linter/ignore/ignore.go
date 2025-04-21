package ignore

import (
	"regexp"
	"strings"
)

type IgnoreManager struct {
	DisabledRules map[string]bool
	LineIgnores   map[int]map[string]bool
}

func NewIgnoreManager() *IgnoreManager {
	return &IgnoreManager{
		DisabledRules: make(map[string]bool),
		LineIgnores:   make(map[int]map[string]bool),
	}
}

var (
	disableRegex = regexp.MustCompile(`(?i)lint-disable\s+(\S+)`)
	enableRegex  = regexp.MustCompile(`(?i)lint-enable\s+(\S+)`)
	ignoreRegex  = regexp.MustCompile(`(?i)lint-ignore\s+(\S+)`)
)

// ProcessDirectives parses the file lines and registers rule disable/enable/ignore directives.
func (im *IgnoreManager) ProcessDirectives(lines []string) {
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Handle lint-disable <RULE>
		if m := disableRegex.FindStringSubmatch(trimmed); m != nil {
			rule := m[1]
			im.DisabledRules[rule] = true
			continue
		}

		// Handle lint-enable <RULE>
		if m := enableRegex.FindStringSubmatch(trimmed); m != nil {
			rule := m[1]
			delete(im.DisabledRules, rule)
			continue
		}

		// Handle lint-ignore <RULE> — apply to next meaningful line
		if m := ignoreRegex.FindStringSubmatch(trimmed); m != nil {
			rule := m[1]

			// Search for next non-blank, non-comment line
			for j := i + 1; j < len(lines); j++ {
				next := strings.TrimSpace(lines[j])
				if next == "" || strings.HasPrefix(next, "//") {
					continue
				}

				// Apply ignore for that line
				if im.LineIgnores[j] == nil {
					im.LineIgnores[j] = make(map[string]bool)
				}
				im.LineIgnores[j][rule] = true
				break
			}
		}
	}
}

// IsRuleIgnored returns true if a rule is ignored on a specific line.
func (im *IgnoreManager) IsRuleIgnored(line int, rule string) bool {
	if ignores, ok := im.LineIgnores[line]; ok {
		return ignores[rule]
	}
	return false
}

// IsRuleDisabled returns true if a rule is globally disabled.
func (im *IgnoreManager) IsRuleDisabled(rule string) bool {
	return im.DisabledRules[rule]
}
