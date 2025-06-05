package go_datamask

import "testing"

func TestMask(t *testing.T) {
	tests := []struct {
		value        string
		visibleCount int
		dir          Direction
		maskChar     rune
		want         string
	}{
		{"1234567890", 4, FromEnd, '*', "******7890"},
		{"abcdefg", 3, FromStart, '#', "abc####"},
		{"abcdefgh", 4, FromMiddle, '*', "**cdef**"},
		{"short", 10, FromEnd, '*', "*****"},      // visibleCount > length, all masked
		{"visible", 0, FromStart, '*', "*******"}, // visibleCount <= 0, all masked
	}

	for _, tt := range tests {
		got := Mask(tt.value, tt.visibleCount, tt.dir, tt.maskChar)
		if got != tt.want {
			t.Errorf("Mask(%q, %d, %v, %q) = %q; want %q", tt.value, tt.visibleCount, tt.dir, tt.maskChar, got, tt.want)
		}
	}
}

func TestFormatPreservingMask(t *testing.T) {
	input := "4111 1111 1111 1234"
	want := "**** **** **** 1234"
	got := FormatPreservingMask(input, 4, '*')
	if got != want {
		t.Errorf("FormatPreservingMask(%q) = %q; want %q", input, got, want)
	}
}

func TestMaskCardFormatted(t *testing.T) {
	card := "4111 1111 1111 1234"
	want := "**** **** **** 1234"
	got := MaskCardFormatted(card)
	if got != want {
		t.Errorf("MaskCardFormatted(%q) = %q; want %q", card, got, want)
	}
}

func TestMaskIBANFormatted(t *testing.T) {
	iban := "TR320010009999901234567890"
	want := "TR********************7890"
	got := MaskIBANFormatted(iban)
	if got != want {
		t.Errorf("MaskIBANFormatted(%q) = %q; want %q", iban, got, want)
	}
}

func TestMaskEmail(t *testing.T) {
	tests := []struct {
		email string
		want  string
	}{
		{"ahmet@example.com", "a****@example.com"},
		{"@domain.com", "*@domain.com"},
		{"invalidemail", "************"},
	}

	for _, tt := range tests {
		got := MaskEmail(tt.email)
		if got != tt.want {
			t.Errorf("MaskEmail(%q) = %q; want %q", tt.email, got, tt.want)
		}
	}
}

func TestMaskPhone(t *testing.T) {
	tests := []struct {
		phone string
		want  string
	}{
		{"+90 555 123 4567", "+90 5** *** 4567"},
		{"5551234567", "5551234567"}, // pattern not matched, unchanged
	}

	for _, tt := range tests {
		got := MaskPhone(tt.phone)
		if got != tt.want {
			t.Errorf("MaskPhone(%q) = %q; want %q", tt.phone, got, tt.want)
		}
	}
}
