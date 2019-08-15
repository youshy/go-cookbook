# Go Cookbook

Aka - 'Learning new language and documenting it'

## To run

`go run <filename>`

## To build

`go build <filename>`

## Notes

* There can't be multiple `main()` function in Go project

* Bypassing package? Use `_` in front of the package name

* `:=` - short assignment statement (instead of using `var`). Cannot reassign!

* `*` - that's a pointer (same as in C). Example:

```go
var p *string
```

this is a pointer to a string then.

## Printf verbs

| Verb | Desc |
|---|---|
| `%d` | decimal integer |
| `%x`, `%o`, `%b` | integer in hexadecimal, octal, binary |
| `%f`, `%g`, `%e` | floating-point number |
| `%t` | boolean |
| `%c` | rune (Unicode code point) |
| `%s` | string |
| `%q` | quoted string "abc" or rune 'c' |
| `%v` | any value in a natural format |
| `%T` | type of any value |
| `%%` | literal percert sign |
