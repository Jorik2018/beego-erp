package controllers

import (
	"crudDemo/helpers"
	"crudDemo/models"
	requestStruct "crudDemo/requstStruct"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

type Users struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type JwtClaim struct {
	Email  string `json:"user_email"`
	UserID int    `json:"user_id"`
	jwt.StandardClaims
}

var secretKey = []byte("devendra_secretkey")

func (c *UserController) Login() {
	var user requestStruct.LoginUser
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		c.CustomAbort(http.StatusBadRequest, "Invalid JSON format")
		return
	}
	loginUserData := models.LoginUsers(user)
	tokenExpire := time.Now().Add(1 * time.Hour)
	claims := &JwtClaim{Email: user.Email, UserID: loginUserData.UserId, StandardClaims: jwt.StandardClaims{
		ExpiresAt: tokenExpire.Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, fmt.Sprintf("Error signing token: %s", err.Error()))
		return
	}

	data := map[string]interface{}{"User_Data": token.Claims, "Token": tokenString}

	c.Data["json"] = map[string]interface{}{"data": data}
	c.ServeJSON()
}

func (c *UserController) AuthMiddleware() {
	tokenString := c.Ctx.Input.Header("Authorization")
	if tokenString == "" {
		c.CustomAbort(http.StatusBadRequest, "Token missing. Please provide a valid token")
		return
	}

	tokenString = tokenString[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected login method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		c.CustomAbort(http.StatusUnauthorized, fmt.Sprintf("Invalid token: %s", err.Error()))
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.CustomAbort(http.StatusBadRequest, "Invalid claims in the token")
		return
	}
	c.SetSession("login_user", claims)
	c.Finish()

}

func (u *UserController) RegisterUser() {
	var user requestStruct.InsertUser
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	result, _ := models.RegisterUser(user)
	if result != nil {
		helpers.ApiSuccessResponse(&u.Controller, "", "Register Successfully User Please Login Now")
	}
	helpers.ApiFailedResponse(&u.Controller, "Please Try Again")
}

func (u *UserController) LoginUser() {
	var user requestStruct.LoginUser
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	result := models.LoginUser(user)
	if result != 0 {
		helpers.ApiSuccessResponse(&u.Controller, "", "Login Successfully User")
	}
	helpers.ApiFailedResponse(&u.Controller, "Invalid Email and Password Please Try Again")
}
