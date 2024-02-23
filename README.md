# go-hexcolor

Package **hexcolor** provides tools for working with **hex color codes**, for the Go programming language.

**Hex Color Codes** look like these:

* `#1B2A34`
* `1E5AA8`
* `#fd0304`
* `00852b`

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-hexcolor

[![GoDoc](https://godoc.org/github.com/reiver/go-hexcolor?status.svg)](https://godoc.org/github.com/reiver/go-hexcolor)

## Example

Here is a simple example of parsing a hex color code:

```golang

var hexcode string = "#8A928D"

red, green, blue, err := hexcolor.Parse(hexcode)
```

And here is a simple example of creating a hex color code:

```golang
var red   uint8 = 0xB4
var green uint8 = 0x00
var blue  uint8 = 0x00

code := hexcolor.Format(red, green, blue)
```

## Import

To import package **hexcolor** use `import` code like the follownig:
```
import "github.com/reiver/go-hexcolor"
```

## Installation

To install package **hexcolor** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-hexcolor
```

## Author

Package **hexcolor** was written by [Charles Iliya Krempeaux](http://changelog.ca)
