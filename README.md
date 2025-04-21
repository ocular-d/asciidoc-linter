# asciidoc-linter
TBD later

To do: Rewrite the README in AsciiDoc.

## Ignore

```adoc
// lint-disable AD002

== Heading
Some content without a blank line, but no warning because it's ignored.

// lint-enable AD002
```

Or for a single file

```adoc
// lint-ignore AD002
== Heading
Some content
```

### Strategy

We'll preprocess the lines and keep track of active ignore states per line:

- `// lint-ignore AD002` → ignore this rule for the next line

- `// lint-disable AD002` → disable rule globally until re-enabled

- `// lint-enable AD002` → re-enable it