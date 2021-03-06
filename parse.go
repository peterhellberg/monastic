package monastic

func Parse(s string) (Cipher, error) {
	if len(s) != 56 {
		return 0, ErrInvalidInputString
	}

	return NewCipher(ones(s) + tens(s) + hundreds(s) + thousands(s))
}

func ones(s string) int {
	if len(s) != 56 {
		return 0
	}

	if s[3:7] == "****" && s[11:15] == "*   " && s[19:23] == "*   " {
		return 1
	}

	if s[3:7] == "*   " && s[11:15] == "*   " && s[19:23] == "****" {
		return 2
	}

	if s[3:7] == "*   " && s[11:15] == "**  " && s[19:23] == "* * " {
		return 3
	}

	if s[3:7] == "*  *" && s[11:15] == "* * " && s[19:23] == "**  " {
		return 4
	}

	if s[3:7] == "****" && s[11:15] == "* * " && s[19:23] == "**  " {
		return 5
	}

	if s[3:7] == "*  *" && s[11:15] == "*  *" && s[19:23] == "*  *" {
		return 6
	}

	if s[3:7] == "****" && s[11:15] == "*  *" && s[19:23] == "*  *" {
		return 7
	}

	if s[3:7] == "*  *" && s[11:15] == "*  *" && s[19:23] == "****" {
		return 8
	}

	if s[3:7] == "****" && s[11:15] == "*  *" && s[19:23] == "****" {
		return 9
	}

	return 0
}

func tens(s string) int {
	if len(s) != 56 {
		return 0
	}

	if s[0:4] == "****" && s[8:12] == "   *" && s[16:20] == "   *" {
		return 10
	}

	if s[0:4] == "   *" && s[8:12] == "   *" && s[16:20] == "****" {
		return 20
	}

	if s[0:4] == "   *" && s[8:12] == "  **" && s[16:20] == " * *" {
		return 30
	}

	if s[0:4] == "*  *" && s[8:12] == " * *" && s[16:20] == "  **" {
		return 40
	}

	if s[0:4] == "****" && s[8:12] == " * *" && s[16:20] == "  **" {
		return 50
	}

	if s[0:4] == "*  *" && s[8:12] == "*  *" && s[16:20] == "*  *" {
		return 60
	}

	if s[0:4] == "****" && s[8:12] == "*  *" && s[16:20] == "*  *" {
		return 70
	}

	if s[0:4] == "*  *" && s[8:12] == "*  *" && s[16:20] == "****" {
		return 80
	}

	if s[0:4] == "****" && s[8:12] == "*  *" && s[16:20] == "****" {
		return 90
	}

	return 0
}

func hundreds(s string) int {
	if len(s) != 56 {
		return 0
	}

	if s[35:39] == "*   " && s[43:47] == "*   " && s[51:55] == "****" {
		return 100
	}

	if s[35:39] == "****" && s[43:47] == "*   " && s[51:55] == "*   " {
		return 200
	}

	if s[35:39] == "* * " && s[43:47] == "**  " && s[51:55] == "*   " {
		return 300
	}

	if s[35:39] == "**  " && s[43:47] == "* * " && s[51:55] == "*  *" {
		return 400
	}

	if s[35:39] == "**  " && s[43:47] == "* * " && s[51:55] == "****" {
		return 500
	}

	if s[35:39] == "*  *" && s[43:47] == "*  *" && s[51:55] == "*  *" {
		return 600
	}

	if s[35:39] == "*  *" && s[43:47] == "*  *" && s[51:55] == "****" {
		return 700
	}

	if s[35:39] == "****" && s[43:47] == "*  *" && s[51:55] == "*  *" {
		return 800
	}

	if s[35:39] == "****" && s[43:47] == "*  *" && s[51:55] == "****" {
		return 900
	}

	return 0
}

func thousands(s string) int {
	if len(s) != 56 {
		return 0
	}

	if s[32:36] == "   *" && s[40:44] == "   *" && s[48:52] == "****" {
		return 1000
	}

	if s[32:36] == "****" && s[40:44] == "   *" && s[48:52] == "   *" {
		return 2000
	}

	if s[24:28] == "*  *" && s[32:36] == " * *" && s[40:44] == "  **" {
		return 3000
	}

	if s[32:36] == "  **" && s[40:44] == " * *" && s[48:52] == "*  *" {
		return 4000
	}

	if s[32:36] == "  **" && s[40:44] == " * *" && s[48:52] == "****" {
		return 5000
	}

	if s[32:36] == "*  *" && s[40:44] == "*  *" && s[48:52] == "*  *" {
		return 6000
	}

	if s[32:36] == "*  *" && s[40:44] == "*  *" && s[48:52] == "****" {
		return 7000
	}

	if s[32:36] == "****" && s[40:44] == "*  *" && s[48:52] == "*  *" {
		return 8000
	}

	if s[32:36] == "****" && s[40:44] == "*  *" && s[48:52] == "****" {
		return 9000
	}

	return 0
}
