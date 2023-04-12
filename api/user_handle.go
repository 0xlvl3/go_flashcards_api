package api

import (
	"context"

	"github.com/0xlvl3/go_flashcards_api/db"
	"github.com/0xlvl3/go_flashcards_api/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)

	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)

}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		Username:       "lvl3",
		TotalCardCount: 3,
	}
	return c.JSON(u)
}
