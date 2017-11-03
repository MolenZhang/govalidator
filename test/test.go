package test

//go:generate gen_validator -type User
type User struct {
	ID    string `valid:"UUID"`
	Name  string `valid:"length(2|16),required"`
	Age   int    `valid:"required,range(5|20)"`
	Sex   string
	Job   string `valid:"-"`
	Email string `valid:"Email"`
}
