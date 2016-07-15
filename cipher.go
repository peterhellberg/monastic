/*

Package monastic is an implementation of The Ciphers of the Monks

The cipher is described on http://www.davidaking.org/Ciphers.htm

Installation

Just go get the package:

    go get -u github.com/peterhellberg/monastic

Usage

A small usage example

		package main

		import (
			"fmt"

			"github.com/peterhellberg/monastic"
		)

		func main() {
			if c, err := monastic.NewCipher(4444); err == nil {
				fmt.Println(c)
			}

			// Output:
			// *  *  *
			//  * * *
			//   ***
			//    *
			//   ***
			//  * * *
			// *  *  *
		}

*/
package monastic

import "errors"

// Cipher of the Monks
type Cipher uint

var (
	// ErrValueTooLarge is returned if the value is above 9999
	ErrValueTooLarge = errors.New("value above 9999, not supported by cipher")

	// ErrValueTooSmall is returned if the value is below 1
	ErrValueTooSmall = errors.New("value below 1, not supported by cipher")

	// ErrInvalidInputString is returned if len(s) != 56
	ErrInvalidInputString = errors.New("invalid input string")
)

// NewCipher returns a valid Cipher unless value is out of bounds.
func NewCipher(n int) (Cipher, error) {
	switch {
	case n < 1:
		return Cipher(0), ErrValueTooSmall
	case n > 9999:
		return Cipher(0), ErrValueTooLarge
	}

	return Cipher(n), nil
}

// Parts return the parts used by the cipher
func (c Cipher) Parts() []uint {
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

// Pattern returns the pattern used to draw the cipher
func (c Cipher) Pattern() [][]uint {
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

// String returns the string representation of the cipher value
func (c Cipher) String() string {
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
