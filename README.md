
# go-datamask

A Go package for masking sensitive personal data such as Turkish identity numbers, credit card numbers, IBANs, emails, and phone numbers. Supports masking with format preservation and customizable options.

## Features

- Mask data from the start, end, or middle of the string
- Preserve formatting such as spaces in credit card numbers or IBANs while masking
- Built-in support for Turkish-specific formats:
  - Turkish Identity Number (TC Kimlik No)
  - Credit Card Number
  - IBAN
  - Email addresses
  - Phone numbers
- Customizable masking character and number of visible characters

## Installation

```bash
go get github.com/mustafamertkisa/go-datamask
```

## Example Usage

```go
func main() {
	// Simple masking examples
	println(Mask("1234567890", 4, FromEnd, '*'))          // ******7890
	println(Mask("abcdefg", 3, FromStart, '#'))           // abc####

	// Mask formatted credit card number
	card := "4111 1111 1111 1234"
	println(MaskCardFormatted(card))                       // **** **** **** 1234

	// Mask IBAN
	iban := "TR320010009999901234567890"
	println(MaskIBANFormatted(iban))                       // TR****************7890

	// Mask email address
	email := "ahmet@example.com"
	println(MaskEmail(email))                              // a****@example.com

	// Mask phone number
	phone := "+90 555 123 4567"
	println(MaskPhone(phone))                              // +90 5** *** 4567
}
```
