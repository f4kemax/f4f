package models
import(
	"github.com/dgrijalva/jwt-go"
  "gorm.io/gorm"
  //_ "github.com/go-sql-driver/mysql"
)
var jwtSecret = []byte("f4keraven")
type CustomClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}
func (u *Users) UsersAllFromDB(db *gorm.DB)error{
  if err := db.Where("token_jwt = ?", u.TokenJWT).First(&u).Error; err != nil {
    return err
  }
  return nil
}
