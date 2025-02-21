package database

import (
	"demo/models"

	"gorm.io/gorm"
)

type IContactDB interface {
	Create(*models.Contact) (*models.Contact, error)
}

func NewContactDB(db *gorm.DB) IContactDB {
	return &ContactDB{db: db}
}

type ContactDB struct {
	db *gorm.DB
}

func (c *ContactDB) Create(contact *models.Contact) (*models.Contact, error) {
	if c.db != nil {
		c.db.AutoMigrate(&models.Contact{})
	}
	tx := c.db.Create((contact))
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}
