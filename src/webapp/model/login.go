package model

//User defines attributes of user
type User struct {
	Name     string
	Email    string
	Password string
	Age      int
}

//NewUser return list of User
func (u *User) NewUser() []User {
	users := []User{
		{
			Name:     "Foo",
			Email:    "foo@fun.com",
			Password: "foobruh",
			Age:      15,
		},
		{
			Name:     "Bar",
			Email:    "bar@fun.com",
			Password: "barbruh",
			Age:      20,
		},
	}
	return users
}

//IsValidUser validates the user
func (u *User) IsValidUser(email string, pass string) bool {
	for _, u := range u.NewUser() {
		if u.Email == email && u.Password == pass {
			return true
		}
	}
	return false
}
