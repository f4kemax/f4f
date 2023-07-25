package models
import(
  "gorm.io/gorm"
)
func (m *Messages) DeleteIdMessage(id int, db *gorm.DB) error {
db.Delete(&m, id)
        return nil
}
