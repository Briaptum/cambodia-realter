package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"cambodia-realter/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthController handles authentication-related endpoints
type AuthController struct{}

// NewAuthController creates a new auth controller
func NewAuthController() *AuthController {
	return &AuthController{}
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// User represents a user
type User struct {
	ID    string   `json:"id"`
	Email string   `json:"email"`
	Role  string   `json:"role"`
	Sites []string `json:"sites"`
}

// Auth0TokenResponse represents the response from Auth0 token endpoint
type Auth0TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// Auth0User represents a user from Auth0 Management API
type Auth0User struct {
	UserID      string                 `json:"user_id"`
	Email       string                 `json:"email"`
	AppMetadata map[string]interface{} `json:"app_metadata"`
}

// Auth0Error represents an error from Auth0
type Auth0Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// Login handles custom login form submission
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Get Auth0 configuration from environment
	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	auth0ClientID := os.Getenv("AUTH0_CLIENT_ID")
	auth0ClientSecret := os.Getenv("AUTH0_CLIENT_SECRET")

	if auth0Domain == "" || auth0ClientID == "" || auth0ClientSecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth0 configuration missing"})
		return
	}

	// Step 1: Authenticate with Auth0 using Resource Owner Password Grant
	fmt.Printf("Attempting Auth0 authentication for user: %s\n", req.Email)
	_, err := ac.authenticateWithAuth0(auth0Domain, auth0ClientID, auth0ClientSecret, req.Email, req.Password)
	if err != nil {
		fmt.Printf("❌ Auth0 authentication failed: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Authentication failed: %v", err)})
		return
	}
	fmt.Printf("✅ Auth0 authentication successful\n")

	// Step 2: Get user information from Auth0 Management API
	fmt.Printf("Fetching user details from Auth0 Management API...\n")
	user, err := ac.getUserFromAuth0(auth0Domain, auth0ClientID, auth0ClientSecret, req.Email)
	if err != nil {
		fmt.Printf("❌ Failed to get user information: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get user information: %v", err)})
		return
	}
	fmt.Printf("✅ User details fetched successfully: %+v\n", user)

	// Step 3: Check if user has required permissions
	fmt.Printf("Checking user permissions...\n")
	if !ac.hasRequiredPermissions(user) {
		fmt.Printf("❌ Access denied: User %s does not have required permissions\n", req.Email)
		fmt.Printf("User app_metadata: %+v\n", user.AppMetadata)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Insufficient permissions"})
		return
	}
	fmt.Printf("✅ User has required permissions\n")

	// Step 4: Create our own JWT token
	jwtToken, err := ac.createJWTToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token: jwtToken,
		User: User{
			ID:    user.UserID,
			Email: user.Email,
			Role:  ac.getUserRole(user),
			Sites: ac.getUserSites(user),
		},
	})
}

// Logout handles user logout
func (ac *AuthController) Logout(c *gin.Context) {
	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")

	if auth0Domain == "" || clientID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth0 configuration missing"})
		return
	}

	// Build Auth0 logout URL
	logoutURL := url.URL{
		Scheme: "https",
		Host:   auth0Domain,
		Path:   "/v2/logout",
	}

	siteURL := os.Getenv("SITE_URL")
	if siteURL == "" {
		siteURL = "http://localhost:3000" // Default for local development
	}

	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("returnTo", siteURL)

	logoutURL.RawQuery = params.Encode()

	c.Redirect(http.StatusTemporaryRedirect, logoutURL.String())
}

// Profile returns the current user's profile
func (ac *AuthController) Profile(c *gin.Context) {
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert claims to Auth0Claims type
	auth0Claims, ok := claims.(*middleware.Auth0Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
		return
	}

	// Extract role and sites from app metadata
	var role string
	var sites []string
	if auth0Claims.AppMetadata != nil {
		if roleVal, ok := auth0Claims.AppMetadata["role"].(string); ok {
			role = roleVal
		}
		if sitesVal, ok := auth0Claims.AppMetadata["sites"].([]interface{}); ok {
			for _, site := range sitesVal {
				if siteStr, ok := site.(string); ok {
					sites = append(sites, siteStr)
				}
			}
		}
	}

	user := User{
		ID:    auth0Claims.Sub,
		Email: auth0Claims.Email,
		Role:  role,
		Sites: sites,
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// authenticateWithAuth0 authenticates user with Auth0 using Resource Owner Password Grant
func (ac *AuthController) authenticateWithAuth0(domain, clientID, clientSecret, email, password string) (*Auth0TokenResponse, error) {
	url := fmt.Sprintf("https://%s/oauth/token", domain)

	payload := map[string]interface{}{
		"grant_type":    "password",
		"username":      email,
		"password":      password,
		"client_id":     clientID,
		"client_secret": clientSecret,
		"scope":         "openid profile email",
		"realm":         "Username-Password-Authentication",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var auth0Err Auth0Error
		json.NewDecoder(resp.Body).Decode(&auth0Err)
		return nil, fmt.Errorf("auth0 error: %s", auth0Err.ErrorDescription)
	}

	var tokenResp Auth0TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

// getUserFromAuth0 gets user information from Auth0 Management API
func (ac *AuthController) getUserFromAuth0(domain, clientID, clientSecret, email string) (*Auth0User, error) {
	// First get a Management API token
	mgmtToken, err := ac.getManagementToken(domain, clientID, clientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to get management token: %w", err)
	}

	// Then get user by email using Management API
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/api/v2/users-by-email?email=%s", domain, url.QueryEscape(email)), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+mgmtToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get user from Auth0: %s", string(body))
	}

	var users []Auth0User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &users[0], nil
}

// getManagementToken gets an access token for the Management API
func (ac *AuthController) getManagementToken(domain, clientID, clientSecret string) (string, error) {
	url := fmt.Sprintf("https://%s/oauth/token", domain)

	payload := map[string]interface{}{
		"grant_type":    "client_credentials",
		"client_id":     clientID,
		"client_secret": clientSecret,
		"audience":      fmt.Sprintf("https://%s/api/v2/", domain),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get management token: %s", string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("no access token in response")
	}

	return accessToken, nil
}

// hasRequiredPermissions checks if user has admin role or site access
func (ac *AuthController) hasRequiredPermissions(user *Auth0User) bool {
	role := ac.getUserRole(user)
	sites := ac.getUserSites(user)
	siteID := os.Getenv("SITE_ID")

	// Check for admin role
	if role == "admin" {
		return true
	}

	// Check for site access if SITE_ID is configured
	if siteID != "" {
		for _, site := range sites {
			if site == siteID {
				return true
			}
		}
	}

	return false
}

// getUserRole extracts role from user's app metadata
func (ac *AuthController) getUserRole(user *Auth0User) string {
	if user.AppMetadata == nil {
		return ""
	}
	if role, ok := user.AppMetadata["role"].(string); ok {
		return role
	}
	return ""
}

// getUserSites extracts sites from user's app metadata
func (ac *AuthController) getUserSites(user *Auth0User) []string {
	if user.AppMetadata == nil {
		return []string{}
	}
	if sites, ok := user.AppMetadata["sites"].([]interface{}); ok {
		var siteStrings []string
		for _, site := range sites {
			if siteStr, ok := site.(string); ok {
				siteStrings = append(siteStrings, siteStr)
			}
		}
		return siteStrings
	}
	return []string{}
}

// createJWTToken creates our own JWT token with user information
func (ac *AuthController) createJWTToken(user *Auth0User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "your-secret-key" // Fallback for development
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.UserID,
		"email": user.Email,
		"https://your-namespace.com/app_metadata": map[string]interface{}{
			"role":  ac.getUserRole(user),
			"sites": ac.getUserSites(user),
		},
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
