package main

type superUser struct {
	Name string
	Age int
}

var _ User = &user{}

type user struct{
	FIO, Adress, Phone string
}

func (u *user) ChangeFIO(newFio string) {
	u.FIO = newFio
}

func (u *user) ChangeAddress(newAddress string) {
	u.Adress = newAddress
}

type User interface {
	ChangeFIO(newFio string)
	ChangeAddress(newAddress string)
}

func NewUser(fio, address, phone string) User{
	u := user{FIO: fio,
		      Adress: address,
			  Phone: phone,
		}
		return &u
}

func main(){


}
