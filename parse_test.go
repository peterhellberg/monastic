package monastic

import "testing"

func TestParse(t *testing.T) {
	for i, tt := range testCases {
		got, err := Parse(tt.s)

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if int(got) != tt.n {
			t.Fatalf(`[%d] Parse(s) = %d, want %d`, i, got, tt.n)
		}

		if got.String() != tt.s {
			t.Fatalf(`[%d] got.String() != %q`, tt.s)
		}
	}
}

func TestAllNumbers(t *testing.T) {
	for n := 1; n <= 9999; n++ {
		s := Cipher(n).String()

		got, err := Parse(s)
		if err != nil {
			t.Fatalf(`unexpected error: %v`, err)
		}

		if int(got) != n {
			t.Fatalf("Parse(Cipher(%d).String()) = %d, want %d", n, got, n)
		}

		if got.String() != s {
			t.Fatalf(`got.String() != %q`, s)
		}
	}
}
