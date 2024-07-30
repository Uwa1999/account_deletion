// // handlers/register.go
// package handlers

// import (
// 	"app/db"
// 	"app/models"
// 	"math"
// 	"strconv"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// )

// func Register(c *fiber.Ctx) error {
// 	var user models.User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
// 	}
// 	if user.Password != user.Confirmpassword {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Password not match"})
// 	}

// 	// Check if the username already exists
// 	if UsernameExists(user.Username) {
// 		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Username already exists"})
// 	}

// 	// Insert user into the database
// 	result := db.DB.Create(&user)

// 	if result.Error != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error inserting user into the database"})
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
// }

// func SoftDeleteUser(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	// Check if the user exists
// 	var user models.User
// 	if err := db.DB.First(&user, id).Error; err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
// 	}

// 	// Soft delete the user by setting a 'deleted_at' timestamp
// 	if err := db.DB.Model(&user).Update("deletedAt", time.Now()).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
// 	}

// 	return c.JSON(fiber.Map{"message": "User soft deleted"})
// }

// func UpdateUser(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var user models.User

// 	// Find the user by ID
// 	result := db.DB.First(&user, id)
// 	if result.Error != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
// 	}

// 	// Parse the request body to update user data
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
// 	}

// 	// Check if the username already exists
// 	if UsernameExists(user.Username) {
// 		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Username already exists"})
// 	}

// 	// Update user data in the database
// 	db.DB.Save(&user)

// 	return c.JSON(fiber.Map{"message": "User updated successfully"})
// }

// func ViewUserHandler(c *fiber.Ctx) error {
// 	// Parse pagination parameters
// 	page, err := strconv.Atoi(c.Query("page", "1"))
// 	if err != nil || page <= 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page number"})
// 	}

// 	perPage, err := strconv.Atoi(c.Query("per_page", "10"))
// 	if err != nil || perPage <= 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid per_page value"})
// 	}

// 	if perPage > 100 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "per_page value exceeds maximum limit of 100"})
// 	}

// 	// Store the current page number
// 	currentPage := page

// 	offset := (page - 1) * perPage

// 	// Extract the search keyword from the query parameters
// 	searchKeyword := c.Query("keyword", "")

// 	// Fetch users from the database with pagination and search
// 	var users []models.ViewUser
// 	query := db.DB
// 	if searchKeyword != "" {
// 		query = query.Where("username LIKE ?", "%"+searchKeyword+"%")
// 	}
// 	result := query.Limit(perPage).Offset(offset).Find(&users)
// 	if result.Error != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
// 	}

// 	// Count all data per page
// 	var totalCount int64
// 	totalCountQuery := db.DB.Model(&models.ViewUser{})
// 	if searchKeyword != "" {
// 		totalCountQuery = totalCountQuery.Where("username LIKE ?", "%"+searchKeyword+"%")
// 	}
// 	totalCountQuery.Count(&totalCount)

// 	// Calculate total pages
// 	totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))

// 	// Return the user data as a JSON response along with the total count and search keyword
// 	return c.JSON(fiber.Map{"users": users, "totalRecords": totalCount, "page": currentPage, "perPage": perPage, "totalPages": totalPages, "search": searchKeyword})
// }

// // UsernameExists checks if the given username already exists in the database
// func UsernameExists(username string) bool {
// 	var count int64
// 	db.DB.Model(&models.User{}).Where("username = ?", username).Count(&count)
// 	return count > 0
// }

// // Get All Clients NGO
