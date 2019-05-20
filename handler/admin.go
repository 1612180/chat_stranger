package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/1612180/chat_stranger/log"
	"github.com/1612180/chat_stranger/models"
	"github.com/1612180/chat_stranger/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service *service.AdminService
}

func NewAdminHandler(service *service.AdminService) *AdminHandler {
	return &AdminHandler{
		service: service,
	}
}

func (adminHandler *AdminHandler) FetchAll(c *gin.Context) {
	admins, errs := adminHandler.service.FetchAll()
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	res := Response(true, ":)")
	res["Admins"] = admins
	c.JSON(http.StatusOK, res)
}

func (adminHandler *AdminHandler) Find(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	admin, errs := adminHandler.service.Find(uint(id))
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	res := Response(true, ":)")
	res["Admin"] = admin
	c.JSON(http.StatusOK, res)
}

func (adminHandler *AdminHandler) Create(c *gin.Context) {
	var adminUpload models.AdminUpload
	if err := c.ShouldBindJSON(&adminUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	id, errs := adminHandler.service.Create(&adminUpload)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	res := Response(true, ":)")
	res["AdminID"] = id
	c.JSON(http.StatusOK, res)
}

func (adminHandler *AdminHandler) UpdateInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	var adminUpload models.AdminUpload
	if err = c.ShouldBindJSON(&adminUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	if errs := adminHandler.service.UpdateInfo(uint(id), &adminUpload); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, ":)"))
}

func (adminHandler *AdminHandler) UpdatePassword(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
	}

	var authentication models.Authentication
	if err = c.ShouldBindJSON(&authentication); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	if errs := adminHandler.service.UpdatePassword(uint(id), &authentication); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, ":)"))
}

func (adminHandler *AdminHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	if errs := adminHandler.service.Delete(uint(id)); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, ":)"))
}

func (adminHandler *AdminHandler) Authenticate(c *gin.Context) {
	var authentication models.Authentication
	if err := c.ShouldBindJSON(&authentication); err != nil {
		log.ServerLog(err)
		res := Response(false, ":(")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if errs := adminHandler.service.Authenticate(&authentication); len(errs) != 0 {
		log.ServerLogs(errs)
		res := Response(false, "Username or password is incorrect")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.CredentialClaims{
		authentication.Name,
		"Admin",
		jwt.StandardClaims{},
	})

	tokenStr, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		res := Response(false, ":(")
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := Response(true, "Login OK")
	res["Token"] = tokenStr
	c.JSON(http.StatusOK, res)
}
