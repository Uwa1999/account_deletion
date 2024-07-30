// package handlers

// import (
// 	"app/db"
// 	"app/models"
// 	"fmt"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// )

// func Login(c *fiber.Ctx) error {
// 	var loginReq models.User
// 	if err := c.BodyParser(&loginReq); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
// 	}

// 	// Query the database for the user
// 	var user models.User
// 	result := db.DB.Where("username = ?", loginReq.Username).First(&user)
// 	if result.Error != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Account Not Found"})
// 	}

// 	// Check if the user is blocked
// 	if !user.BlockedUntil.IsZero() && user.BlockedUntil.After(time.Now()) {
// 		timeLeft := user.BlockedUntil.Sub(time.Now())
// 		formattedTimeLeft := formatDuration(timeLeft)
// 		if formattedTimeLeft == "00:00" {
// 			// Reset login attempts if time left is zero
// 			user.LoginAttempts = 0
// 			user.BlockedUntil = time.Time{}
// 			db.DB.Save(&user)
// 		} else {
// 			return c.Status(fiber.StatusLocked).JSON(fiber.Map{"error": "Account locked. Try again later", "retry_in": formattedTimeLeft})

// 		}
// 	}

// 	// Check the password
// 	if user.Password != loginReq.Password {
// 		// Check if login attempts threshold is reached
// 		if user.LoginAttempts < 3 {
// 			user.LoginAttempts++
// 			db.DB.Save(&user)
// 		}

// 		if user.LoginAttempts >= 3 {

// 			// Calculate the time when the user should be unblocked
// 			user.BlockedUntil = time.Now().Add(10 * time.Second)
// 			db.DB.Save(&user)
// 			timeLeft := user.BlockedUntil.Sub(time.Now())
// 			formattedTimeLeft := formatDuration(timeLeft)
// 			return c.Status(fiber.StatusLocked).JSON(fiber.Map{"error": "Account locked. Try again later", "retry_in": formattedTimeLeft})
// 		}

// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
// 	}

// 	// Reset login attempts if login is successful
// 	user.LoginAttempts = 0

// 	// Reset BlockedUntil time
// 	user.BlockedUntil = time.Time{}
// 	db.DB.Save(&user)

// 	return c.JSON(fiber.Map{"message": "Login successful"})
// }

// func formatDuration(duration time.Duration) string {
// 	// hours := int(duration.Hours())
// 	minutes := int(duration.Minutes()) % 60
// 	seconds := int(duration.Seconds()) % 60
// 	return fmt.Sprintf("%02d:%02d", minutes, seconds)
// }
