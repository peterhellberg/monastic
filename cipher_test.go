package monastic

import "testing"

func TestCipher(t *testing.T) {
	for _, tt := range []struct {
		n   uint
		err error
	}{
		{0, ErrValueTooSmall},
		{1, nil},
		{9, nil},
		{9999, nil},
		{10000, ErrValueTooLarge},
	} {
		if _, err := Cipher(tt.n); err != tt.err {
			t.Fatalf("unexpected error: %s", err)
		}
	}
}

func TestParts(t *testing.T) {
	for _, tt := range []struct {
		n uint
		p []uint
	}{
		{1, []uint{1}},
		{12, []uint{10, 2}},
		{123, []uint{100, 20, 3}},
		{1234, []uint{1000, 200, 30, 4}},
		{4567, []uint{4000, 500, 60, 7}},
		{9999, []uint{9000, 900, 90, 9}},
	} {
		c, err := Cipher(tt.n)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got := c.Parts(); !equalUints(got, tt.p) {
			t.Fatalf("unexpected parts: %v, want %v", got, tt.p)
		}
	}
}

func TestString(t *testing.T) {
	for _, tt := range []struct {
		n uint
		s string
	}{
		{1, "   ****\n   *   \n   *   \n   *   \n   *   \n   *   \n   *   \n"},
		{10, "****   \n   *   \n   *   \n   *   \n   *   \n   *   \n   *   \n"},
		{100, "   *   \n   *   \n   *   \n   *   \n   *   \n   *   \n   ****\n"},
		{1000, "   *   \n   *   \n   *   \n   *   \n   *   \n   *   \n****   \n"},

		{2, "   *   \n   *   \n   ****\n   *   \n   *   \n   *   \n   *   \n"},
		{20, "   *   \n   *   \n****   \n   *   \n   *   \n   *   \n   *   \n"},
		{200, "   *   \n   *   \n   *   \n   *   \n   ****\n   *   \n   *   \n"},
		{2000, "   *   \n   *   \n   *   \n   *   \n****   \n   *   \n   *   \n"},

		{3, "   *   \n   **  \n   * * \n   *  *\n   *   \n   *   \n   *   \n"},
		{30, "   *   \n  **   \n * *   \n*  *   \n   *   \n   *   \n   *   \n"},
		{300, "   *   \n   *   \n   *   \n   *  *\n   * * \n   **  \n   *   \n"},
		{3000, "   *   \n   *   \n   *   \n*  *   \n * *   \n  **   \n   *   \n"},

		{4, "   *  *\n   * * \n   **  \n   *   \n   *   \n   *   \n   *   \n"},
		{40, "*  *   \n * *   \n  **   \n   *   \n   *   \n   *   \n   *   \n"},
		{400, "   *   \n   *   \n   *   \n   *   \n   **  \n   * * \n   *  *\n"},
		{4000, "   *   \n   *   \n   *   \n   *   \n  **   \n * *   \n*  *   \n"},

		{5, "   ****\n   * * \n   **  \n   *   \n   *   \n   *   \n   *   \n"},
		{50, "****   \n * *   \n  **   \n   *   \n   *   \n   *   \n   *   \n"},
		{500, "   *   \n   *   \n   *   \n   *   \n   **  \n   * * \n   ****\n"},
		{5000, "   *   \n   *   \n   *   \n   *   \n  **   \n * *   \n****   \n"},

		{6, "   *  *\n   *  *\n   *  *\n   *   \n   *   \n   *   \n   *   \n"},
		{60, "*  *   \n*  *   \n*  *   \n   *   \n   *   \n   *   \n   *   \n"},
		{600, "   *   \n   *   \n   *   \n   *   \n   *  *\n   *  *\n   *  *\n"},
		{6000, "   *   \n   *   \n   *   \n   *   \n*  *   \n*  *   \n*  *   \n"},

		{7, "   ****\n   *  *\n   *  *\n   *   \n   *   \n   *   \n   *   \n"},
		{70, "****   \n*  *   \n*  *   \n   *   \n   *   \n   *   \n   *   \n"},
		{700, "   *   \n   *   \n   *   \n   *   \n   *  *\n   *  *\n   ****\n"},
		{7000, "   *   \n   *   \n   *   \n   *   \n*  *   \n*  *   \n****   \n"},

		{8, "   *  *\n   *  *\n   ****\n   *   \n   *   \n   *   \n   *   \n"},
		{80, "*  *   \n*  *   \n****   \n   *   \n   *   \n   *   \n   *   \n"},
		{800, "   *   \n   *   \n   *   \n   *   \n   ****\n   *  *\n   *  *\n"},
		{8000, "   *   \n   *   \n   *   \n   *   \n****   \n*  *   \n*  *   \n"},

		{9, "   ****\n   *  *\n   ****\n   *   \n   *   \n   *   \n   *   \n"},
		{90, "****   \n*  *   \n****   \n   *   \n   *   \n   *   \n   *   \n"},
		{900, "   *   \n   *   \n   *   \n   *   \n   ****\n   *  *\n   ****\n"},
		{9000, "   *   \n   *   \n   *   \n   *   \n****   \n*  *   \n****   \n"},

		{1111, "*******\n   *   \n   *   \n   *   \n   *   \n   *   \n*******\n"},
		{2222, "   *   \n   *   \n*******\n   *   \n*******\n   *   \n   *   \n"},
		{3333, "   *   \n  ***  \n * * * \n*  *  *\n * * * \n  ***  \n   *   \n"},
		{4444, "*  *  *\n * * * \n  ***  \n   *   \n  ***  \n * * * \n*  *  *\n"},
		{5555, "*******\n * * * \n  ***  \n   *   \n  ***  \n * * * \n*******\n"},
		{6666, "*  *  *\n*  *  *\n*  *  *\n   *   \n*  *  *\n*  *  *\n*  *  *\n"},
		{7777, "*******\n*  *  *\n*  *  *\n   *   \n*  *  *\n*  *  *\n*******\n"},
		{8888, "*  *  *\n*  *  *\n*******\n   *   \n*******\n*  *  *\n*  *  *\n"},
		{9999, "*******\n*  *  *\n*******\n   *   \n*******\n*  *  *\n*******\n"},

		{1993, "****   \n*  **  \n**** * \n   *  *\n   ****\n   *  *\n*******\n"},
		{4723, "   *   \n   **  \n**** * \n   *  *\n  **  *\n * *  *\n*  ****\n"},
		{6859, "*******\n * *  *\n  *****\n   *   \n*  ****\n*  *  *\n*  *  *\n"},
		{7085, "*  ****\n*  * * \n*****  \n   *   \n*  *   \n*  *   \n****   \n"},
		{9433, "   *   \n  ***  \n * * * \n*  *  *\n*****  \n*  * * \n****  *\n"},
	} {
		c, err := Cipher(tt.n)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got := c.String(); got != tt.s {
			t.Fatalf("\n\ngot string for %d:\n\n%s\nwant:\n\n%v", tt.n, got, tt.s)
		}
	}
}

func equalUints(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
