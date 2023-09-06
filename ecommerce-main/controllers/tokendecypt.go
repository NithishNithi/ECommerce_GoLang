package controllers

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ExtractCustomerID(jwtToken string, secretKey string) (string, error) {

	// Parse the JWT token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err // Handle token parsing errors
	}

	// Check if the token is valid
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Extract the customer ID from the claims
			customerID, ok := claims["customerid"].(string)
			if ok {
				return customerID, nil
			}
		}
	}

	return "", fmt.Errorf("Invalid or expired JWT token")
}
