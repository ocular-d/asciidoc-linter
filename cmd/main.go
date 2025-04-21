package main

import (
    "github.com/ocular-d/asciidoc-linter/internal/linter"
    "github.com/ocular-d/asciidoc-linter/rules"
    "flag"
    "fmt"
)

func main() {
    flag.Parse()
    args := flag.Args()

    if len(args) == 0 {
        fmt.Println("Usage: asciidoc-linter <file1.adoc> <file2.adoc> ...")
        return
    }

    lint := linter.NewLinter()
    lint.RegisterRule(rules.HeadingSpacingRule{}) // AD001
	lint.RegisterRule(rules.HeadingSurroundRule{})    // AD002

    for _, file := range args {
        lint.LintFile(file)
    }
}
