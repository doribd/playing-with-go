// This package name isn't main, which means it will be imported into another package which will be used as the main.
package leftpad

import (
	"strings"
	"unicode/utf8"
)

// we can define variables within a package to be used within the package
var default_character = ' '

// you can define package visibility to be used only within the package or
// you can define publiic visibility which means that it will be accessible also from outside the package
// capital letter at the begining indicates that this is a public accessability
// lower case will indicate only package visibility
func Format(s string, size int) string {
	return FormatRune(s, size, default_character)
}

func FormatRune(s string, size int, r rune) string {
	length := utf8.RuneCountInString(s)
	if length >= size {
		return s
	}
	output := strings.Repeat(string(r), size-length) + s
	return output
}
