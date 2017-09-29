# fansi

> ANSI utilities

## Install

To get started, you need a working Go environment. Once available, grab the package here:

```
$ go get github.com/vutran/fansi
```

## Usage

### Cursor

To be documented.

### Display

To be documented.

### Colors

Simple functions to apply a foreground or background color to a string of text.

```
package main

import (
	"github.com/vutran/fansi/colors"
)

func main() {
	msg := colors.Blue("Hello, world")
	msg = colors.YellowBg(msg)

	fmt.Print(msg)
}
```

## License

MIT © [Vu Tran](https://github.com/vutran/srgnt)
