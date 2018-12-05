package viewmodels


type UsersVM struct {
	Id			string		`json:"id"`
	FirstName	string		`json:"firstName"`
	LastName	string		`json:"lastName"`
	UserName	string		`json:"userName"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
}