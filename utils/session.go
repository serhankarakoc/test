package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

func GetSession(c *fiber.Ctx) (*session.Session, error) {
	sess, err := store.Get(c)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func IsAuthenticated(c *fiber.Ctx, store *session.Store) (bool, error) {
	sess, err := store.Get(c)
	if err != nil {
		return false, err
	}

	authenticated := sess.Get("authenticated")
	if authenticated == nil || authenticated == false {
		return false, nil
	}

	return true, nil
}

func GetUserID(c *fiber.Ctx, store *session.Store) (uint, error) {
	sess, err := store.Get(c)
	if err != nil {
		return 0, err
	}

	userID := sess.Get("user_id")
	if userID == nil {
		return 0, errors.New("user ID not found in session")
	}

	return userID.(uint), nil
}
