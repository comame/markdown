# markdown

An simplefied markdown parser written in go.

# Markdown spec (non-exhaustive)

## Assumptions

- Blocks are bounded by blank lines
- Blocks except lists does not contain blocks

## Constraints

- Escaping by backslash are always done (including text in Code blocks, Autolinks etc.)
- Thematic breaks is only `---`
- Spaces at the end of lines are ignored
- Closing sequences of `#` in ATX headings are ignored (treated as regular text)
- Setext Headings are ignored
- Link reference definitions are ignored
- Reference Links are ignored
