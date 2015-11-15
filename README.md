# Susumu
Susumu is a tiny progress bar package.

## Features
- Arbitrary max value
- Percent display
- Settable unit display
- Settable width in characters
- Smooth 1/8-char increments

## Demo
[![Susumu demo](https://asciinema.org/a/7e54gx2njw4mydhj1gve4tnel.png)](https://asciinema.org/a/7e54gx2njw4mydhj1gve4tnel?autoplay=1)

## Example
```go
package main

import (
	"fmt"

	"github.com/jsageryd/susumu"
)

func main() {
	b := &susumu.Bar{Width: 80, Max: 256, Position: 128, Unit: "KiB"}
	err := b.Draw()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	b.Position = 192
	err = b.Draw()
	if err != nil {
		panic(err)
	}
	fmt.Println()
}
```

### Output
```
50% ██████████████████████████████                               128 / 256 KiB
75% █████████████████████████████████████████████                192 / 256 KiB
```

## Licence
Copyright (c) 2015 Johan Sageryd <j@1616.se>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
