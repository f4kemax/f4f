package models
/*type UsersInterface interface{
  UsersAllFromDB(db *gorm.DB) error
	ReturnUsers() *Users
  ReturnLogin() string
Proverka(id int,db *gorm.DB) error
}
type MessagesInterface interface{
	PostMessage(u *Users, db *gorm.DB) error
  GetIdMessage(id int, db *gorm.DB) error
  PutIdMessage(id int, db *gorm.DB) error
  DeleteIdMessage(id int, db *gorm.DB) error
}*/
type Users struct {
	Id           int
	Login        string `json:"login"`
  PasswordHash string `json:"password"`
	TokenJWT     string
	Email        string `json:"email"`
  LastName     string `json:"lastname"`
  FirstName    string `json:"firstname"`
}
type Messages struct {
	Id      int
	Login   string
	Header  string `json:"header"`
	Message string `json:"message"`
	Created string
	Updated string
}
type Answer struct {
	Msg string `json:"msg"`
}
