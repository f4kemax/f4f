package handlers
import (
  "f4f/models"
  "net/http"
  "github.com/gorilla/mux"
  "strconv"
  "encoding/json"
  "gorm.io/gorm"
)
func GetMessageId(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    user, ok := r.Context().Value("Users").(models.Users)
        if !ok {
          w.WriteHeader(http.StatusInternalServerError)
          answer:=models.Answer{Msg: "Error 500 please try again later"}
          json.NewEncoder(w).Encode(answer)
          return
        }
      id := mux.Vars(r)//получает id
      s := id["id"]
      i, _ := strconv.Atoi(s)
      err:=user.ChekReqId(i, db)
      if err!=nil{
        w.WriteHeader(http.StatusForbidden)
        answer:=models.Answer{Msg: "No access"}
        json.NewEncoder(w).Encode(answer)
        return
      }
      message:=models.Messages{}
      err=message.GetIdMessage(i, db)
      if err!=nil{
        w.WriteHeader(http.StatusNotFound)
        answer:=models.Answer{Msg: "The requested resource was not found on the server"}
        json.NewEncoder(w).Encode(answer)
        return
      }
      json.NewEncoder(w).Encode(message)
}
}
