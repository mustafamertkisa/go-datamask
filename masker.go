package go_datamask

import (
	"regexp"
	"strings"
	"unicode"
)

// Direction defines the masking direction: from start, from end or from middle
type Direction int

const (
	FromStart Direction = iota
	FromEnd
	FromMiddle
)

// Mask masks a string by replacing characters with maskChar,
// showing visibleCount characters according to the specified Direction.
// If visibleCount is invalid, entire string is masked.
func Mask(value string, visibleCount int, dir Direction, maskChar rune) string {
	length := len(value)
	if visibleCount >= length || visibleCount <= 0 {
		return strings.Repeat(string(maskChar), length)
	}

	switch dir {
	case FromStart:
		visible := value[:visibleCount]
		return visible + strings.Repeat(string(maskChar), length-visibleCount)
	case FromEnd:
		visible := value[length-visibleCount:]
		return strings.Repeat(string(maskChar), length-visibleCount) + visible
	case FromMiddle:
		// Shift visible to the right side a bit if needed
		start := (length - visibleCount) / 2
		if start+visibleCount > length {
			start = length - visibleCount
		}
		prefix := start
		suffix := length - visibleCount - prefix
		return strings.Repeat(string(maskChar), prefix) + value[prefix:prefix+visibleCount] + strings.Repeat(string(maskChar), suffix)
	default:
		return strings.Repeat(string(maskChar), length)
	}
}

// FormatPreservingMask masks only the digits in input string while preserving non-digit characters (such as spaces).
// It keeps visibleCount digits visible from the end and masks the rest with maskChar.
func FormatPreservingMask(input string, visibleCount int, maskChar rune) string {
	var digitsOnly []rune
	for _, r := range input {
		if unicode.IsDigit(r) {
			digitsOnly = append(digitsOnly, r)
		}
	}

	masked := Mask(string(digitsOnly), visibleCount, FromEnd, maskChar)

	var result strings.Builder
	maskIndex := 0
	for _, r := range input {
		if unicode.IsDigit(r) {
			result.WriteRune(rune(masked[maskIndex]))
			maskIndex++
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// MaskCardFormatted masks credit card numbers by preserving formatting (spaces) and showing only last 4 digits
func MaskCardFormatted(card string) string {
	return FormatPreservingMask(card, 4, '*')
}

// MaskIBANFormatted masks IBAN numbers by preserving formatting and showing only last 4 digits
func MaskIBANFormatted(iban string) string {
	return FormatPreservingMask(iban, 4, '*')
}

// MaskEmail masks an email address by showing only the first character of the username
// and replacing the rest with asterisks, keeping the domain intact.
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return strings.Repeat("*", len(email))
	}
	name := parts[0]
	domain := parts[1]

	if len(name) == 1 {
		return "*" + "@" + domain
	} else if len(name) > 1 {
		return string(name[0]) + strings.Repeat("*", len(name)-1) + "@" + domain
	}
	return strings.Repeat("*", len(email))
}

// MaskPhone masks phone numbers in the format "+90 555 123 4567",
// showing country code, first digit, and last four digits while masking middle digits.
func MaskPhone(phone string) string {
	re := regexp.MustCompile(`(\+?\d{1,3})\s?(\d)(\d{2})\s?(\d{3})\s?(\d{4})`)
	matches := re.FindStringSubmatch(phone)
	if len(matches) == 6 {
		return matches[1] + " " + matches[2] + "** *** " + matches[5]
	}
	// Return input unchanged if pattern doesn't match
	return phone
}
