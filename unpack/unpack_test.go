package unpack_test

import (
	"testing"

	"string-unpacker/unpack"
)

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "Standard case",
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
			wantErr:  false,
		},
		{
			name:     "No digits case",
			input:    "abcd",
			expected: "abcd",
			wantErr:  false,
		},
		{
			name:     "Empty input",
			input:    "",
			expected: "",
			wantErr:  false,
		},
		{
			name:     "Simple escape sequence 1",
			input:    `qwe\4\5`,
			expected: `qwe45`,
			wantErr:  false,
		},
		{
			name:     "Simple escape sequence 2",
			input:    `qwe\45`,
			expected: `qwe44444`, // \4 экранирует 4, затем 5 повторяет эту 4 пять раз
			wantErr:  false,
		},
		{
			name:     "Escape backslash case",
			input:    `qwe\\5`,
			expected: `qwe\\\\\`, // Экранированная \ повторяется 5 раз
			wantErr:  false,
		},
		{
			name:     "Unicode support",
			input:    `Привет4 мир2`,
			expected: `Приветттт мирр`,
			wantErr:  false,
		},
		{
			name:    "Error: digit in beginning",
			input:   "45",
			wantErr: true,
		},
		{
			name:    "Error: incomplete escape at end",
			input:   `abc\`,
			wantErr: true,
		},
		{
			name:    "Error: digit after digit",
			input:   "a34",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpack.String(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("String() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
