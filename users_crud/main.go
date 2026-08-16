package main

import (
	"fmt";
	"errors";
)

type User struct {
	ID int
	Name string
	Email  string
}


func main() {

	users := []User{
		{
			ID: 12,
			Name: "Israel L.",
			Email: "johndoe@mail.com",
		},
		{
			ID: 13,
			Name: "Javier L.",
			Email: "joanede@mail.com",
		},
		{
			ID: 14,
			Name: "Pako L.",
			Email: "juande@mail.com",
		},
				{
			ID: 15,
			Name: "Pako L.",
			Email: "juande@mail.com",
		},
				{
			ID: 16,
			Name: "Pako L.",
			Email: "juande@mail.com",
		},
				{
			ID: 17,
			Name: "Pako L.",
			Email: "juande@mail.com",
		},
	}

	result, err := findUser(users, 12)

	if err != nil {
		return
	}
	fmt.Println(result)

}


func createUser(users []User, NewUser User) []User {

	users = append(users, NewUser)
	return users
}

func updateUser(users []User, ID int, NewEmail string, NewName string) (*User, error){ 

	for i := 0; i < len(users); i ++ {
		if users[i].ID == ID {

			userToUpdate := &users[i] 

			userToUpdate.Email = NewEmail
			userToUpdate.Name = NewName
			
			return userToUpdate, nil
		}
	}
	return nil, errors.New("user not found")
}

func deleteUser(users []User, ID int) ([]User, error){

	for i := 0; i < len(users); i ++ {
		
		if (users[i].ID == ID) {
			users = append(users[:i], users[i+1:]...)
			return users, nil
		} 
	}
	return nil, errors.New("user not found")
}

func findUser(users []User, ID int) (*User, error){ 

	for i, _ := range users{

		if users[i].ID == ID {
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}