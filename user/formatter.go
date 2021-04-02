package user

type UserFormater struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Occupation string `json:"occupation"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user User,token string)UserFormater{
	formatter := UserFormater{
		Id :user.ID,
		Name:user.Name,
		Occupation: user.Occupation,
		Email: user.Email,
		Token:token,
	}

	return formatter
}