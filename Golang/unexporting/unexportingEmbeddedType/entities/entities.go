package entities

type user struct {
	Name, Email string
}

type Admin struct {
	user
	Rights int
}
