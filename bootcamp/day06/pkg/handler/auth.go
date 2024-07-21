package handler

import (
	"errors"
	"time"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/azeek21/blog/views/components"
	"github.com/azeek21/blog/views/layouts"
	"github.com/gin-gonic/gin"
)

var SIGNIN_ERROR = errors.New("Incorrect email or password! Plaease try again, or create a new accout if you don't already have one.")

func (h Handler) SignUpPage(ctx *gin.Context) {
	utils.RenderTempl(ctx, 200, layouts.SignUpPage())
}

func (h Handler) SignOutPage(ctx *gin.Context) {
	utils.RenderTempl(ctx, 200, layouts.SignOutPage())
}

func (h Handler) SignInPage(ctx *gin.Context) {
	utils.RenderTempl(ctx, 200, layouts.SignInPage())
}

func (h Handler) SignOut(ctx *gin.Context) {
	origin := ctx.Request.Header.Get("Origin")
	ctx.Header("HX-Location", "/")
	ctx.SetCookie(models.AUTH_COOKIE_NAME, "", 0, "/", origin, false, true)
	ctx.Status(200)
}

func (h Handler) SignIn(ctx *gin.Context) {
	creds := models.SignInDTO{}
	err := ctx.ShouldBind(&creds)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}
	user, err := h.service.UserService.GetUserByEmail(creds.Email)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, SIGNIN_ERROR.Error()))
		ctx.Abort()
		return
	}

	isPasswordMatch, err := h.service.PasswordService.VerifyPasswordAgainstHash(creds.Password, user.Password)
	if err != nil || !isPasswordMatch {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, SIGNIN_ERROR.Error()))
		ctx.Abort()
		return
	}

	token, err := h.service.CreateJwt(user)

	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	origin := ctx.Request.Header.Get("Origin")
	ctx.Header("HX-Location", "/")
	ctx.SetCookie(models.AUTH_COOKIE_NAME, token, int((time.Hour * 720).Seconds()), "/", origin, false, true)
}

func (h Handler) SignUp(ctx *gin.Context) {
	var creds models.SignUpDTO
	err := ctx.ShouldBind(&creds)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	if creds.Password != creds.PasswordRepeat {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, "Password and it's repeat should match"))
		ctx.Abort()
		return
	}

	existingUser, err := h.service.GetUserByEmail(creds.Email)
	if err == nil && creds.Email == existingUser.Email {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, "This use with this emai, already exists. Try Sign in if it's you."))
		ctx.Abort()
		return
	}

	phash, err := h.service.CreateHash(creds.Password)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	user, err := h.service.CreateUser(&models.User{
		Password: phash,
		Email:    creds.Email,
		FullName: creds.Name,
		Username: creds.Username,
		RoleCode: "user",
	})

	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	token, err := h.service.CreateJwt(user)

	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	origin := ctx.Request.Header.Get("Origin")
	ctx.Header("HX-Location", "/")
	ctx.SetCookie(models.AUTH_COOKIE_NAME, token, int((time.Hour * 720).Seconds()), "/", origin, false, true)
}
