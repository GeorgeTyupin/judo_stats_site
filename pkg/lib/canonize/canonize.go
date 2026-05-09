package canonize

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Canonize(s string) (string, error) {
	stripper := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		norm.NFC,
	)

	result, _, err := transform.String(stripper, s)
	if err != nil {
		return "", err
	}

	return result, nil
}
