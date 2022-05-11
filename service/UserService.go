package service

import (
	"golangAPI/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-redis/redis/v8"
)

// @Summary Get AllUser123
// @Tags User
// @version 1.0
// @produce application/json
// @Router / [GET]
func FindAllUsers(c *gin.Context) {
	users := pojo.FindAllUsers()
	c.JSON(http.StatusOK, users)
}

// @Summary get user by id
// @Tags User
// @version 1.0
// @produce application/json
// @param id path int true "id"
// @Router /{id} [GET]
func FindByUserId(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User ->", user)
	c.JSON(http.StatusOK, user)
}

// @Summary Post
// @Tags User
// @version 1.0
// @produce application/json
// @param data body pojo.User true "id_name_password_email"
// @Router / [Post]
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}

// @Summary Delete
// @Tags User
// @version 1.0
// @produce application/json
// @param id path int true "id"
// @Router /{id} [DELETE]
func DeleteUser(c *gin.Context) {
	user := pojo.Deleteuser(c.Param("id"))
	if !user {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully!")
}

// @Summary Update
// @Tags User
// @version 1.0
// @produce application/json
// @param id path int true "id"
// @param data body pojo.User true "id_name_password_email"
// @Router /{id} [PUT]
func PutUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = pojo.Updateuser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, user)
}
