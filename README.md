# monastic

[![Build Status](https://travis-ci.org/peterhellberg/monastic.svg?branch=master)](https://travis-ci.org/peterhellberg/monastic)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/monastic)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/monastic#license-mit)

Implementation of [The Ciphers of the Monks](https://en.wikipedia.org/wiki/The_Ciphers_of_the_Monks) in Go

![Ciphers](http://www.davidaking.org/Photos/Ciphers_clip_image002.jpg)

## Installation

    go get -u github.com/peterhellberg/monastic

## Usage

```go
package main

import (
	"fmt"

	"github.com/peterhellberg/monastic"
)

func main() {
	if c, err := monastic.Cipher(4444); err == nil {
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
```

## License (MIT)

Copyright (c) 2015 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
