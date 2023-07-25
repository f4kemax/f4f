package handlers
import (
  "f4f/models"
  "net/http"
  "github.com/gorilla/mux"
  "strconv"
  "encoding/json"
  "gorm.io/gorm"
)
func PutMessageId(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    user, ok := r.Context().Value("Users").(models.Users)
        if !ok {
          w.WriteHeader(http.StatusInternalServerError)
          answer:=models.Answer{Msg: "Error 500 please try again later"}
          json.NewEncoder(w).Encode(answer)
          return
        }
    message:=models.Messages{}
    err := json.NewDecoder(r.Body).Decode(&message)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      answer:=models.Answer{Msg: "Error reading data from request body"}
      json.NewEncoder(w).Encode(answer)
      return
    }
    if len(message.Header)==0 || len(message.Message)==0{
      w.WriteHeader(http.StatusBadRequest)
      answer:=models.Answer{Msg: "Title or message is empty"}
      json.NewEncoder(w).Encode(answer)
      return
    }
      id := mux.Vars(r)//получает id
      s := id["id"]
      i, _ := strconv.Atoi(s)
      err=user.ChekReqId(i, db)
      if err!=nil{
        w.WriteHeader(http.StatusForbidden)
        answer:=models.Answer{Msg: "No access"}
        json.NewEncoder(w).Encode(answer)
        return
      }
      err=message.PutIdMessage(i, db)
      if err!=nil{
        w.WriteHeader(http.StatusInternalServerError)
        answer:=models.Answer{Msg: "Repeat later"}//мб нету записи но это не точно или не удалось перезаписать
        json.NewEncoder(w).Encode(answer)
        return
      }
      json.NewEncoder(w).Encode(message)
}
}
