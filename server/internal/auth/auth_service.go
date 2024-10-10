// internal/auth/auth_service.go
package auth

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/code-guy21/PingUp/server/models"
    "github.com/code-guy21/PingUp/server/repositories"
    "github.com/code-guy21/PingUp/server/utils"
)

// RegisterUser hashes the user's password and saves the user in the database.
func RegisterUser(user *models.User) error {
    // Hash the user's password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

    // Store the user in the database
    return repositories.CreateUser(user)
}

// AuthenticateUser verifies the user's credentials and generates a JWT token if successful.
func AuthenticateUser(email, password string) (string, error) {
    // Find the user by email
    user, err := repositories.GetUserByEmail(email)
    if err != nil {
        return "", err
    }

    // Compare the provided password with the stored hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", err
    }

    // Generate a JWT token if authentication succeeds
    return utils.GenerateJWT(user.ID)
}
