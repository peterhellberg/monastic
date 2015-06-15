package monastic

import "errors"

type cipher uint

var (
	ErrValueTooLarge = errors.New("value above 9999, not supported by cipher")
	ErrValueTooSmall = errors.New("value below 1, not supported by cipher")
)

func Cipher(n uint) (cipher, error) {
	switch {
	case n < 1:
		return cipher(0), ErrValueTooSmall
	case n > 9999:
		return cipher(0), ErrValueTooLarge
	}

	return cipher(n), nil
}

func (c cipher) Parts() []uint {
	n := uint(c)
	p := []uint{}

	if r := (n / 1000) * 1000; r > 0 {
		p = append(p, r)
	}

	if r := ((n / 100) % 10) * 100; r > 0 {
		p = append(p, r)
	}

	if r := ((n / 10) % 10) * 10; r > 0 {
		p = append(p, r)
	}

	if r := n % 10; r > 0 {
		p = append(p, r)
	}

	return p
}

func (c cipher) Pattern() [][]uint {
	pattern := [][]uint{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}

	for _, p := range c.Parts() {
		if pp, ok := patterns[p]; ok {
			for y, row := range pp {
				for x, val := range row {
					if val == 1 {
						pattern[y][x] = 1
					}
				}
			}
		}
	}

	return pattern
}

func (c cipher) String() string {
	var s string

	for _, row := range c.Pattern() {
		for _, val := range row {
			if val == 1 {
				s += "*"
			} else {
				s += " "
			}
		}
		s += "\n"
	}

	return s
}
