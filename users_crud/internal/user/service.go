package internal 

func CreateUser(users []User, NewUser User) []User {

	users = append(users, NewUser)
	return users
}

func UpdateUser(users []User, ID int, NewEmail string, NewName string) (*User, error){ 

	userToUpdate, err := FindUser(users, ID)

	if err != nil {
		return nil, err
	}

	userToUpdate.Email = NewEmail
	userToUpdate.Name = NewName

	return userToUpdate, nil
}

func DeleteUser(users []User, ID int) ([]User, error){

	for i := range users {
		if users[i].ID == ID {
			users =  append(users[:i], users[i+1:]...)
			return users, nil 
		}
	}
	return users, ErrUserNotFound
}

func FindUser(users []User, ID int) (*User, error){ 

	for i := range users{

		if users[i].ID == ID {
			return &users[i], nil
		}
	}
	return nil, ErrUserNotFound
}