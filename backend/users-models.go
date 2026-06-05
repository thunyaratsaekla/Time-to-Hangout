package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `gorm:"unique"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Gender     string `json:"gender"`
	Birthday   string `json:"birthday"`
	Occupation string `json:"occupation"`
	Salary     string `json:"salary"`
	Point      int    `json:"point" gorm:"default:0"`
	Role       string `json:"role"`
	Status     string `json:"status"`
	gorm.Model
}

func CreateUser(db *gorm.DB, c *fiber.Ctx) error {
	// user := new(User)
	var user User
	//Parse request body to user struct
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}
	//Check if email already exists
	var existingUser User
	err = db.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		//If already Email in database
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Email is already in use",
		})
	} else if err != gorm.ErrRecordNotFound {
		//If error is not record not found
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error checking for existing email",
		})

	}

	//Hash password
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error hashing password",
		})
	}
	//Convert hashed password to string
	user.Password = string(hashedpassword)

	//Create user
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error saving user to database",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User registered successfully",
	})
}

// Login User
func LoginUser(db *gorm.DB, c *fiber.Ctx) error {
	var input User
	var user User
	//Parse request body to user struct
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	fmt.Println(input.Email)
	fmt.Println(input.Password)

	//Check if email exists
	err = db.Where("email = ?", input.Email).First(&user).Error

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid email or password",
		})
	}

	//Check password database and input are same
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid email or password",
		})
	}

	//Jwt token
	//Secret key
	jwtSecret := "KeySecret"
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	//Generate encoded token
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//Set Cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Minute * 30),
		HTTPOnly: true,
	})

	//Login success
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "success",
		"message":  "User Logined successfully",
		"user_id":  user.ID,
	})
}

func GetUsers(db *gorm.DB, c *fiber.Ctx) error {
	var Users []User
	db.Find(&Users)
	return c.JSON(Users)
}

func GetUser(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var User User
	if err := db.First(&User, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Users not found",
		})
	}
	return c.JSON(User)
}

func UpdateUser(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	User := new(User)
	if err := db.First(&User, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	if err := c.BodyParser(User); err != nil {
		return err
	}
	db.Save(&User)
	return c.JSON(User)
}

func GetEventsUserByUserID(db *gorm.DB, c *fiber.Ctx) error {
	userID := c.Params("user_id")
	var events []Event_user
	db.Where("user_id = ?", userID).Find(&events)
	return c.JSON(events)
}
