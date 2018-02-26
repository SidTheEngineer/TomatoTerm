package main

import tui "github.com/gizak/termui"

func main() {
	defer tui.Close()
	tui.Loop()
}
