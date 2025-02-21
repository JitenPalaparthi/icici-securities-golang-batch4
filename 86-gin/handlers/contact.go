package handlers

import (
	"demo/database"
	"demo/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IContactHandler interface {
	Create(ctx *gin.Context)
}

type Contacthandler struct {
	IContactDB database.IContactDB
}

func NewContacthandler(icontactdb database.IContactDB) IContactHandler {
	return &Contacthandler{IContactDB: icontactdb}
}

func (ch *Contacthandler) Create(ctx *gin.Context) {
	contact := new(models.Contact)
	err := ctx.Bind(contact)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
	}
	contact.Lastmodified = time.Now().Unix()
	contact.Status = "active"
	err = contact.Validate()
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
	}
	contact, err = ch.IContactDB.Create(contact)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusCreated, contact)

}
