package main

type dir struct {
	Path     string
	Title    string
	Abstract string
	Subs     []*dir
}
