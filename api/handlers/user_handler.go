package handlers

import (
	"errors"
	"net/http"

	dbstores "github.com/alijabbar034/hotelManagement/api/db_stores"
	"github.com/alijabbar034/hotelManagement/types"
	"github.com/alijabbar034/hotelManagement/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userStorer dbstores.User_Storer
}

func NewUserHandler(user_storer dbstores.User_Storer) *UserHandler {
	return &UserHandler{userStorer: user_storer}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user types.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	if err := types.ValidateInput(user); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	usr, err := h.userStorer.GetUserByEmail(user.Email)
	if err != nil {

		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	if usr != nil {
		utils.ErrorHandler(c, http.StatusConflict, errors.New("User already exists"))
		return
	}
	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	user.Password = pass

	use := types.NewUser(user)
	id, err := h.userStorer.RegisterUser(use)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendToken(c, id)

}

func (h *UserHandler) LoginUserHandler(c *gin.Context) {
	var login types.LoginData
	if err := c.BindJSON(&login); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	if err := types.ValidateLoginData(login); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	usr, err := h.userStorer.GetUserByEmail(login.Email)
	if err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	if usr == nil {
		utils.ErrorHandler(c, http.StatusNotFound, errors.New("User not found"))
		return
	}

	er := utils.ComapreHashPassword(login.Password, usr.Password)
	if er != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, errors.New("Invalid login password or email"))
		return
	}
	utils.SendToken(c, usr.ID)
}

func (h *UserHandler) GetByIdHandler(c *gin.Context) {

	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)
	user, err := h.userStorer.GetUserById(id)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	if user == nil {
		utils.ErrorHandler(c, http.StatusNotFound, errors.New("User not found"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *UserHandler) GetAllUserHandler(c *gin.Context) {

	users, err := h.userStorer.GetAllUsers()
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {

	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)
	count, err := h.userStorer.DeleteUser(id)
	if err != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"count":   count,
		"message": "Delete user successfully"})
}

func (h *UserHandler) LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", 1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully"})
}

func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	reqUsr, _ := c.Get("user")
	usr := reqUsr.(types.User)
	var user types.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}
	updateCount, eror := h.userStorer.UpdateUser(user, usr.ID)
	if eror != nil {
		utils.ErrorHandler(c, http.StatusInternalServerError, eror)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"count":   updateCount,
		"message": "Update user successfully",
	})
}

func (h *UserHandler) GetProfileHandler(c *gin.Context) {
	reqUser, _ := c.Get("user")
	user := reqUser.(types.User)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
