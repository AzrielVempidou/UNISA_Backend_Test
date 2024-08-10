package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"UNISA_Server/config"
	"UNISA_Server/models"
	"UNISA_Server/utils"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// JWT secret key
var jwtSecret = []byte("your_secret_key")

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	// Validate role
	if user.Role != "user" && user.Role != "admin" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid role")
		return
	}

	// Check if username already exists
	var existingUser models.User
	err = config.MongoClient.Database("mydatabase").Collection("users").FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Username already exists")
		return
	} else if err != mongo.ErrNoDocuments {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error checking username")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error hashing password")
		return
	}
	user.Password = string(hashedPassword)

	// Insert user into the database
	_, err = config.MongoClient.Database("mydatabase").Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error registering user")
		return
	}

	utils.CreateResponse(w, true, http.StatusCreated, "User registered successfully", nil)
}

// Login handles user login and returns a JWT token
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	var user models.User
	err = config.MongoClient.Database("mydatabase").Collection("users").FindOne(context.TODO(), bson.M{"username": credentials.Username}).Decode(&user)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generate JWT token
	token, err := generateJWT(user.ID, user.Role)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	// Prepare the response
	response := map[string]interface{}{
		"Status_Code":  200,
		"access_token": token,
		"message":      "Login successful",
		"status":       true,
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


// generateJWT generates a JWT token
func generateJWT(userID primitive.ObjectID, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   userID.Hex(),
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
