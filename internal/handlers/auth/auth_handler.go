package handlers

import (
	services "Orderly/internal/services/auth"
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

type AuthHandler struct {
	OAuth2Config *oauth2.Config
	OIDCVerifier *oidc.IDTokenVerifier
	AuthService  *services.AuthService
}

func NewAuthHandler(oauth2Config *oauth2.Config, verifier *oidc.IDTokenVerifier, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		OAuth2Config: oauth2Config,
		OIDCVerifier: verifier,
		AuthService:  authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {

	state, err := generateRandomState()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.SetCookie("oauthstate", state, 600, "/", "", false, true)

	url := h.OAuth2Config.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) AuthCallback(c *gin.Context) {
	state := c.Query("state")
	cookieState, err := c.Cookie("oauthstate")
	if err != nil || cookieState != state {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	code := c.Query("code")
	token, err := h.OAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token: " + err.Error()})
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No id_token field in oauth2 token."})
		return
	}

	idToken, err := h.OIDCVerifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify ID Token: " + err.Error()})
		return
	}

	var claims struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
		Phone   string `json:"phone"`
	}
	if err := idToken.Claims(&claims); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims: " + err.Error()})
		return
	}

	user, err := h.AuthService.RegisterOrLogin(claims.Email, claims.Name, claims.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register or log in user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User authenticated successfully",
		"user":    user,
	})
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
