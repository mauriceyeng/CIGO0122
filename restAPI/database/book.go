package database

import "github.com/jinzhu/gorm"

func GetBooks(db *gorm.DB) ([]models.Book, error) {

	books := []models.Book{}
	query := db.Select("books.*")
	query.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}
