package main

import "testing"

func TestTransliterate(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "basic_cyrillic",
			input: "Привет",
			want:  "Privet",
		},
		{
			name:  "mixed_with_spaces",
			input: "Тест 123",
			want:  "Test123",
		},
		{
			name:  "keeps_unmapped",
			input: "A-Б.",
			want:  "A-B.",
		},
		{
			name:  "soft_hard_signs_removed",
			input: "подъезд",
			want:  "podezd",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := Transliterate(tc.input); got != tc.want {
				t.Fatalf("Transliterate(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
