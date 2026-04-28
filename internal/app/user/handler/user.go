package handler

import (
	"simple-blog-system/internal/app/user/model"
	"simple-blog-system/internal/app/user/payload"
	"simple-blog-system/internal/app/user/port"
	"simple-blog-system/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	userService port.IUserService
}

func New(userService port.IUserService) port.IUserHandler {
	return &handler{
		userService: userService,
	}
}

// @BasePath /v1

// @Summary Register User
// @Description Register User
// @Tags user
// @Accept json
// @Produce json
// @Param user body payload.RegisterRequest true "Param Register"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /public-api/user/register [post]
func (h *handler) Register(c *gin.Context) {
	// 1. Pakai RegisterRequest supaya sinkron dengan JSON dari Frontend
	var req payload.RegisterRequest

	// 2. Gunakan ShouldBindJSON untuk menangkap body JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseError(c, err)
		return
	}

	// 3. Validasi input
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		helper.ResponseError(c, err)
		return
	}

	// 4. Mapping dari Request ke Model yang dibutuhkan Service
	dataUser := model.AuthUserModel{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := h.userService.Register(c.Request.Context(), dataUser)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	helper.ResponseData(c, &helper.Response{
		Message: "register successfully",
		Data:    res,
	})
}

// @Summary Login User
// @Description Login User
// @Tags user
// @Accept json
// @Produce json
// @Param user body payload.RegisterRequest true "Param Login"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /public-api/user/login [post]
func (h *handler) Login(c *gin.Context) {
	var req payload.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseError(c, err)
		return
	}

	dataUser := model.AuthUserModel{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := h.userService.Login(c.Request.Context(), dataUser)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	helper.ResponseData(c, &helper.Response{
		Message: "login successfully",
		Data:    res,
	})
}

// @Summary Get User
// @Description Get User
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /api/profile [get]
func (h *handler) GetUser(c *gin.Context) {
	username := c.GetString("username")
	res, err := h.userService.GetUser(c.Request.Context(), username)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	helper.ResponseData(c, &helper.Response{
		Message: "get user successfully",
		Data:    res,
	})
}

func (h *handler) UpdateUser(c *gin.Context) {
	type RequestUpdate struct {
		Username string `json:"username" binding:"required"`
	}

	var req RequestUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseError(c, err)
		return
	}

	username := c.GetString("username")
	
	res, err := h.userService.UpdateUser(c.Request.Context(), username, req.Username)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	helper.ResponseData(c, &helper.Response{
		Message: "update username successfully",
		Data:    res,
	})
}