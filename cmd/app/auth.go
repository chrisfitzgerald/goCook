package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"os"
)

const (
	TokenCookieName  = "random-state-string"
	oauthStateString = "random-state-string"
)

var JWTKey = os.Getenv("JWT_SECRET_KEY")

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func createJWT(email string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTKey))
}

func googleLoginHandler(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func googleCallbackHandler(c echo.Context) error {
	log.Println("Google callback handler started")
	code := c.QueryParam("code")
	if code == "" {
		return c.String(http.StatusBadRequest, "Code is missing")
	}

	token, err := googleOauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
	}

	client := googleOauthConfig.Client(c.Request().Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get user info: "+err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.String(http.StatusInternalServerError, "Failed to get user info: "+resp.Status)
	}

	var userInfo struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to decode user info: "+err.Error())
	}

	jwtToken, err := createJWT(userInfo.Email)
	if err != nil {
		log.Printf("Failed to create JWT: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to create JWT: "+err.Error())
	}
	log.Printf("JWT created for email: %s", userInfo.Email)

	cookie := &http.Cookie{
		Name:     TokenCookieName,
		Value:    jwtToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,  // Set to true if using HTTPS
		SameSite: http.SameSiteLaxMode,
		MaxAge:   24 * 60 * 60, // 24 hours
	}
	c.SetCookie(cookie)
	log.Printf("JWT cookie set: %+v", cookie)

	// Set the user in the context
	claims := &jwt.RegisteredClaims{
		Subject: userInfo.Email,
	}
	c.Set("user", claims)

	log.Println("Redirecting to home page after successful authentication")
	return c.Redirect(http.StatusSeeOther, "/")
}

func logoutHandler(c echo.Context) error {
	log.Println("Logout handler called")
	cookie := new(http.Cookie)
	cookie.Name = TokenCookieName
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false // Set to true if using HTTPS
	cookie.SameSite = http.SameSiteStrictMode

	c.SetCookie(cookie)
	log.Println("Cookie cleared, redirecting to home")
	return c.Redirect(http.StatusSeeOther, "/")
}
