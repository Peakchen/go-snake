package main

type graphics interface {
	rect(x, y, w, h int, argb8 uint32)
	text(utf8 []byte, x, y int, clip rectangle, argb uint32)
	present() error
}

type rectangle struct {
	x, y, w, h int
}

func rect(x, y, w, h int) rectangle {
	return rectangle{x: x, y: y, w: w, h: h}
}
