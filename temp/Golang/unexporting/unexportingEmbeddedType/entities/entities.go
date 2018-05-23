package entities

type user struct {
	Name, Email string
}

//Admin 是user的升级
type Admin struct {
	user
	Rights int
}
