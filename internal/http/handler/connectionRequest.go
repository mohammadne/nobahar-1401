package handler

import "github.com/gofiber/fiber/v2"

// connectGroups connects two groups
// if user is not admin, returns 400
func (g Group) connectGroups(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getConnectionRequests returns all connection requests for the user
// newest to the oldest
// if user is not admin, returns 400
func (g Group) getConnectionRequests(ctx *fiber.Ctx) error {
	panic("implement me")
}

// acceptConnectionRequest accepts a connection request
// if user is not admin, returns 400
func (g Group) acceptConnectionRequest(ctx *fiber.Ctx) error {
	panic("implement me")
}
