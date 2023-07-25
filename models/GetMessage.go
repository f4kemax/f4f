package models
import(
  "gorm.io/gorm"
  //_ "github.com/go-sql-driver/mysql"
)
func GetMessage(l string,db *gorm.DB) ([]Messages, error) {
  var messages []Messages
  if err := db.Where("login = ?", l).Find(&messages).Error; err != nil{
    return nil, err
}
return messages, nil
}
