package types

type RegisterUserPayLoad struct {
	FileName  string `json:"firstname"`
	LastName  string `json:"latname"`
	Email     string `json:"email"`
	Passoword string `json:"password"`
}
