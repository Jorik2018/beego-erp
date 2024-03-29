package middelware

import (
	"beego-erp/models"
	requestStruct "beego-erp/requstStruct"
	"fmt"
	"net/http"
	"time"

	"github.com/beego/beego"
	"github.com/dgrijalva/jwt-go"
)

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

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	var user requestStruct.LoginUser
	if err := c.ParseForm(&user); err != nil {
		c.CustomAbort(http.StatusBadRequest, "Invalid JSON format")
		return
	}

	userData := models.LoginUsers(user)

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JwtClaim{Email: user.Email, UserID: userData.UserId, StandardClaims: jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
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

func (c *AuthController) AuthMiddleware() {
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
}
