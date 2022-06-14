package dto

import "ecommerce/models"

type User struct {
	ID string
	First_Name string
	Last_Name string  
	Email string      
	Phone string       
	Token string       
	Refresh_Token string
}

func GetUserDetails(entityUser models.User) User{
	return User{
		ID: entityUser.User_ID,
		First_Name: *entityUser.First_Name,
		Last_Name: *entityUser.Last_Name,
		Email: *entityUser.Email,
		Phone: *entityUser.Phone,
		Token: *entityUser.Token,
		Refresh_Token: *entityUser.Refresh_Token,
	}
}

type UserFullDetails struct {
	ID string
	First_Name string
	Last_Name string  
	Email string      
	Phone string       
	Token string       
	Refresh_Token string
}