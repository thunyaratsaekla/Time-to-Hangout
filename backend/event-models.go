package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Event Organizer
type Event_org struct {
	Name_eo    string     `json:"name"`
	Organizer  string     `json:"organizer"`
	Location   string     `json:"location"`
	About      string     `json:"about"`
	Point      string     `json:"point"`
	Status     string     `json:"status"`
	Time_start *time.Time `json:"timestart"`
	Time_end   *time.Time `json:"timeend"`
	Image      string     `json:"image"`
	gorm.Model
}

// Envent User
type Event_user struct {
	Name_eu   string     `json:"name"`
	Person    string     `json:"person"`
	Location  string     `json:"location"`
	About     string     `json:"about"`
	Point     string     `json:"point"`
	User_id   uint       `json:"user_id"`
	Status    string     `json:"status"`
	TimeStart *time.Time `json:"timestart"`
	TimeEnd   *time.Time `json:"timeend"`
	Image     string     `json:"image"`
	gorm.Model
}

type My_event struct {
	User_id  uint `json:"user_id"`
	Event_id uint `json:"event_id"`
	gorm.Model
}

func GetEvents(db *gorm.DB, c *fiber.Ctx) error {
	var Events []Event_org
	db.Find(&Events)

	for i := range Events {
		if Events[i].Time_start != nil {
			t := Events[i].Time_start.Local().Truncate(time.Minute)
			Events[i].Time_start = &t
		}
		if Events[i].Time_end != nil {
			t := Events[i].Time_end.Local().Truncate(time.Minute)
			Events[i].Time_end = &t
		}
	}

	return c.JSON(Events)
}

func GetEvent(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var Event Event_org
	if err := db.First(&Event, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}
	if Event.Time_start != nil {
		t := Event.Time_start.Local().Truncate(time.Minute)
		Event.Time_start = &t
	}
	if Event.Time_end != nil {
		t := Event.Time_end.Local().Truncate(time.Minute)
		Event.Time_end = &t
	}

	return c.JSON(Event)
}

func CreateEvent(db *gorm.DB, c *fiber.Ctx) error {
	Event := new(Event_org)
	if err := c.BodyParser(Event); err != nil {
		return err
	}
	db.Create(&Event)
	return c.JSON(Event)
}

func UpdateEvent(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	Event := new(Event_org)
	if err := db.First(&Event, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}
	if err := c.BodyParser(Event); err != nil {
		return err
	}
	db.Save(&Event)
	return c.JSON(Event)
}

func DeleteEvent(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	if err := db.Delete(&Event_org{}, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}
	return c.SendString("Event successfully deleted")
}

func JoinEvent(db *gorm.DB, c *fiber.Ctx) error {
	type JoinRequest struct {
		UserID  uint `json:"user_id"`
		EventID uint `json:"event_id"`
	}
	var req JoinRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	// ป้องกัน joinEvent ซ้ำ
	var exist My_event
	if err := db.Where("user_id = ? AND event_id = ?", req.UserID, req.EventID).First(&exist).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Already joined"})
	}
	joined := My_event{User_id: req.UserID, Event_id: req.EventID}
	db.Create(&joined)
	return c.JSON(joined)
}

func GetJoinedEvents(db *gorm.DB, c *fiber.Ctx) error {
	userID := c.Params("user_id")
	var joined []My_event
	db.Where("user_id = ?", userID).Find(&joined)

	// ดึงข้อมูล Event เต็มๆ
	var events []Event_org
	var eventIDs []uint
	for _, j := range joined {
		eventIDs = append(eventIDs, j.Event_id)
	}
	if len(eventIDs) > 0 {
		db.Where("id IN ?", eventIDs).Find(&events)
	}
	return c.JSON(events)
}

func GetEventsUser(db *gorm.DB, c *fiber.Ctx) error {
	var Events []Event_user
	db.Find(&Events)

	for i := range Events {
		if Events[i].TimeStart != nil {
			t := Events[i].TimeStart.Local().Truncate(time.Minute)
			Events[i].TimeStart = &t
		}
		if Events[i].TimeEnd != nil {
			t := Events[i].TimeEnd.Local().Truncate(time.Minute)
			Events[i].TimeEnd = &t
		}
	}

	return c.JSON(Events)
}

func GetEventUser(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var Event Event_user
	if err := db.First(&Event, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}
	if Event.TimeStart != nil {
		t := Event.TimeStart.Local().Truncate(time.Minute)
		Event.TimeStart = &t
	}
	if Event.TimeEnd != nil {
		t := Event.TimeEnd.Local().Truncate(time.Minute)
		Event.TimeEnd = &t
	}

	return c.JSON(Event)
}

func CreateEventUser(db *gorm.DB, c *fiber.Ctx) error {
	Event := new(Event_user)
	if err := c.BodyParser(Event); err != nil {
		return err
	}
	db.Create(&Event)
	return c.JSON(Event)
}

func UpdateEventUser(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	Event := new(Event_user)
	if err := db.First(&Event, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}
	if err := c.BodyParser(Event); err != nil {
		return err
	}
	db.Save(&Event)
	return c.JSON(Event)
}

func DeleteEventUser(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	if err := db.Delete(&Event_user{}, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}
	return c.SendString("Event successfully deleted")
}

func GivePointToJoinedUsers(db *gorm.DB, c *fiber.Ctx) error {
	eventID := c.Params("event_id")
	var event Event_org
	if err := db.First(&event, eventID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Event not found"})
	}

	// ดึง point ของ event นี้ (ควรเป็น int)
	point := 0
	fmt.Sscanf(event.Point, "%d", &point)

	// ดึง user_id ทั้งหมดที่ join event นี้
	var joined []My_event
	db.Where("event_id = ?", eventID).Find(&joined)

	// แจก point ให้ user ทุกคน
	for _, j := range joined {
		var user User
		if err := db.First(&user, j.User_id).Error; err == nil {
			user.Point += point
			db.Save(&user)
		}
	}
	return c.JSON(fiber.Map{"status": "success"})
}

func UpdateEventStatus(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var event Event_user
	if err := db.First(&event, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}

	// เปลี่ยนสถานะเป็น Finished
	event.Status = "Finished"
	if err := db.Save(&event).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update event status",
		})
	}

	return c.JSON(fiber.Map{
        "status":  "success",
        "message": "Event status updated to Finished",
        "event":   event,
	})
}
