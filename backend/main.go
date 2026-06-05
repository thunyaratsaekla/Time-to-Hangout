package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"   // or the Docker service name if running in another container
	port     = 5432          // default PostgreSQL port
	user     = "myuser"      // as defined in docker-compose.yml
	password = "mypassword"  // as defined in docker-compose.yml
	dbname   = "me2database" // as defined in docker-compose.yml
)

func AuthRequired(c *fiber.Ctx) error {
	//JWT token from cookie
	cookie := c.Cookies("jwt")
	//Secret key
	jwtSecretKey := "KeySecret"
	//Check jwt token
	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	//If error or token invalid
	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	//Get claim from token
	claim := token.Claims.(jwt.MapClaims)
	fmt.Println(claim)
	return c.Next()
}

func main() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database")
	}
	// Migrate the schema
	//สร้าง table ใน db ถ้าไม่มี
	db.AutoMigrate(&Event_org{}, &User{}, &My_event{}, &Event_user{})
	fmt.Println("Database migration completed!")

	app := fiber.New()

	app.Use(cors.New())

	//Organizer
	//API Get all events
	app.Get("/events", func(c *fiber.Ctx) error {
		return GetEvents(db, c)
	})

	//API Get one events from id
	app.Get("/event/:id", func(c *fiber.Ctx) error {
		return GetEvent(db, c)
	})

	//API Create event
	app.Post("/event", func(c *fiber.Ctx) error {
		fmt.Println("POST /event called")
		return CreateEvent(db, c)
	})

	//API Update one events from id
	app.Put("/event/:id", func(c *fiber.Ctx) error {
		return UpdateEvent(db, c)
	})

	//API Delete one events from id
	app.Delete("/event/:id", func(c *fiber.Ctx) error {
		return DeleteEvent(db, c)
	})

	//API Register user
	app.Post("/register", func(c *fiber.Ctx) error {
		return CreateUser(db, c)
	})

	//API Login user
	app.Post("/login", func(c *fiber.Ctx) error {
		return LoginUser(db, c)
	})

	app.Get("users", func(c *fiber.Ctx) error {
		return GetUsers(db, c)
	})

	app.Get("user/:id", func(c *fiber.Ctx) error {
		return GetUser(db, c)
	})

	app.Put("user/:id", func(c *fiber.Ctx) error {
		return UpdateUser(db, c)
	})

	app.Post("/join-event", func(c *fiber.Ctx) error {
		return JoinEvent(db, c)
	})
	app.Get("/joined-events/:user_id", func(c *fiber.Ctx) error {
		return GetJoinedEvents(db, c)
	})

	//User
	//API Get all events
	app.Get("/events-user", func(c *fiber.Ctx) error {
		return GetEventsUser(db, c)
	})

	//API Get one events from id
	app.Get("/event-user/:id", func(c *fiber.Ctx) error {
		return GetEventUser(db, c)
	})

	//API Create event
	app.Post("/event-user", func(c *fiber.Ctx) error {
		fmt.Println("POST /event called")
		return CreateEventUser(db, c)
	})

	//API Update one events from id
	app.Put("/event-user/:id", func(c *fiber.Ctx) error {
		return UpdateEventUser(db, c)
	})

	//API Delete one events from id
	app.Delete("/event-user/:id", func(c *fiber.Ctx) error {
		return DeleteEventUser(db, c)
	})

	app.Post("/give-point/:event_id", func(c *fiber.Ctx) error {
		return GivePointToJoinedUsers(db, c)
	})

	app.Get("/events-user/:user_id", func(c *fiber.Ctx) error {
		return GetEventsUserByUserID(db, c)
	})

	app.Put("/event-user/finish/:id", func(c *fiber.Ctx) error {
		return UpdateEventStatus(db, c)
	})

	app.Listen(":8000")
}
