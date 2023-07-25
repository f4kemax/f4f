package handlers
import (
  "f4f/models"
  "golang.org/x/crypto/bcrypt"
  "net/http"
  "encoding/json"
  "gorm.io/gorm"
  _ "github.com/go-sql-driver/mysql"
)
func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
  if r.Method == "POST" {
    var user models.Users
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil{
      w.WriteHeader(http.StatusBadRequest)
      answer:=models.Answer{Msg: "The client request contains a syntax error and cannot be fulfilled by the server"}
      json.NewEncoder(w).Encode(answer)
      return
    }
    if len(user.Login)==0 || len(user.PasswordHash)==0{
      w.WriteHeader(http.StatusBadRequest)
      answer:=models.Answer{Msg: "Login or password field is empty"}
      json.NewEncoder(w).Encode(answer)
      return
    }
    token, err:=authorization(db, user.Login, user.PasswordHash)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      answer:=models.Answer{Msg: "Wrong password or login"}
      json.NewEncoder(w).Encode(answer)
      return
    }
    answer:=models.Answer{Msg: token}
    json.NewEncoder(w).Encode(answer)
  } else {//Этого произойти не должно
    w.WriteHeader(http.StatusBadRequest)//400 Запрос клиента содержит синтаксическую ошибку и не может быть выполнен сервером.
    answer:=models.Answer{Msg: "Method not post"}
    json.NewEncoder(w).Encode(answer)
    return
  }
}
}
func authorization(db *gorm.DB, u string, p string)(string, error){
  var user models.Users
	if err := db.Where("login = ?", u).First(&user).Error; err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(p))
	if err != nil {
		return "", err// Пароль не совпадает или возникла ошибка
	}
	return user.TokenJWT, nil
}
