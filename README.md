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

### Style

Simple functions to apply a style to a string of text.

```
package main

import (
	"github.com/vutran/fansi/styles"
)

func main() {
	msg: = styles.Bold("Hello, world")
	fmt.Print(msg)
}
```

### Colors

Simple functions to apply a foreground or background color to a string of text.

```
package main

import (
	"github.com/vutran/fansi/colors"
)

func main() {
	msg := colors.Blue("Hello, world")
	fmt.Print(msg)
}
```

### Mix and Match

You can mix and decorate your however you prefer.

```
package main

import (
	"github.com/vutran/fansi/colors"
	"github.com/vutran/fansi/styles"
)

func main() {
	// bold, blue text
	msg := styles.Bold(colors.Blue("Hello, world"))
	fmt.Print(msg)
}
```

## License

MIT © [Vu Tran](https://github.com/vutran/srgnt)
