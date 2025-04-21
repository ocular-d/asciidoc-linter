package main

import (
    "github.com/ocular-d/asciidoclint/internal/linter"
    "github.com/ocular-d/asciidoclint/rules"
    "flag"
    "fmt"
)

func main() {
    flag.Parse()
    args := flag.Args()

    if len(args) == 0 {
        fmt.Println("Usage: asciidoclint <file1.adoc> <file2.adoc> ...")
        return
    }

    lint := linter.NewLinter()
    lint.RegisterRule(rules.HeadingSpacingRule{}) // AD001
	lint.RegisterRule(rules.HeadingSurroundRule{})    // AD002

    for _, file := range args {
        lint.LintFile(file)
    }
}
