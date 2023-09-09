package routes

import (
	"github.com/dassudip2001/webapp/database"
	"github.com/dassudip2001/webapp/models"
	"github.com/gofiber/fiber/v2"
)

// The User struct represents a user with ID, username, email, and password fields.
// @property {uint} ID - The ID property is of type uint, which represents an unsigned integer. It is
// used to uniquely identify each user in the system.
// @property {string} Username - The `Username` property is a string that represents the username of a
// user. It is used to uniquely identify a user and is often used for authentication and authorization
// purposes.
// @property {string} Email - The `Email` property is a string that represents the email address of the
// user.
// @property {string} Password - The "Password" property is a string that represents the user's
// password.
type User struct {
	ID       uint   `json:"id""`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// The `createResponceUser` function is a helper function that takes a `models.User` object as input
// and creates a new `User` object with the same values. It is used to convert a `models.User` object
// to a `User` object, which is a simplified version of the user model that will be returned as a JSON
// response.
func createResponceUser(userModel models.User) User {
	return User{
		ID:       userModel.ID,
		Username: userModel.Username,
		Email:    userModel.Email,
		Password: userModel.Password,
	}
}



// The function `CreateUser` creates a new user by parsing the request body, saving it to the database,
// and returning the created user as a JSON response.
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responceUser:=createResponceUser(user)
	return c.JSON(responceUser)
}

// The function retrieves all users from the database and returns them as a JSON response.
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)
	return c.JSON(users)
}

// 
// The GetUser function retrieves a user from the database based on the provided ID and returns it as a
// JSON response.
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	
	var user models.User
	database.Database.Db.Find(&user, id)
	return c.JSON(user)
}

// 
// The function retrieves a user from the database based on their username and returns it as a JSON
// response.
func GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	var user models.User
	database.Database.Db.Where("username = ?", username).Find(&user)
	return c.JSON(user)
}
// The UpdateUser function updates a user in the database based on the provided ID and returns the
// updated user as a JSON response.
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.Database.Db.First(&user, id)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Save(&user)
	return c.JSON(user)
}

// The DeleteUser function deletes a user from the database based on the provided ID and returns the
// deleted user as a JSON response.
// error handling is not done here

func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")
    
    var user models.User
    if err := database.Database.Db.First(&user, id).Error; err != nil {
        
        return c.Status(fiber.StatusInternalServerError).SendString("User not found " )
    }

    if err := database.Database.Db.Delete(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Error while deleting user")
    }

    return c.SendString("User successfully deleted")
}
