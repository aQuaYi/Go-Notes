package main

import (
	tm "DesignModel/TemplateMethod/templateMethod"
)

func main() {
	d1 := tm.NewCharDisplay('H')
	d1.Display()

	d2 := tm.NewStringDisplay("Hello, World!")
	d2.Display()
}
