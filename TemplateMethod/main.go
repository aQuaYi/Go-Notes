package main

import (
	tm "DesignModel/TemplateMethod/templateMethod"
)

func main() {
	d1 := new(tm.AbstractDisplay)
	d1.Inst = &tm.CharDisplay{Ch: 'H'}

	d1.Display()
}
