package controller

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"te-eme-backend/app/adapter/out/firestore"
	"te-eme-backend/app/shared/archetype/container"
	einar "te-eme-backend/app/shared/archetype/echo_server"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	container.InjectInboundAdapter(func() error {
		einar.Echo().POST("/api/login", login)
		return nil
	})
}

var SecretKey string = os.Getenv("JWT_AUTH_SECRET")

func login(c echo.Context) error {

	var ctx = context.Background()
	var userData map[string]string

	if err := c.Bind(&userData); err != nil {
		return err
	}

	user, err := firestore.FindUserByID(ctx, userData)

	// Verifica si el usuario existe (puedes reemplazar esta lógica con tu consulta a la base de datos)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": err,
		})
	}

	// Compara la contraseña hash con la contraseña proporcionada
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData["Password"])); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Contraseña incorrecta",
		})
	}

	// Crea un token JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 día
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	}

	// Configura el token JWT como una cookie
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Conectado correctamente",
	})

}
