package handler

import "github.com/gofiber/fiber/v2"

// join requests to join a group
// if user is in a group, return 400
func (g Group) join(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getJoinRequests returns user's all join requests to other groups
// newest to the oldest
func (g Group) getJoinRequests(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getJoinRequestsByGroup returns all join requests for user's group.
// newest to oldest
// if user is not admin, returns 400
func (g Group) getJoinRequestsByGroup(ctx *fiber.Ctx) error {
	panic("implement me")
}

// acceptJoinRequest accepts a join request
// if user is not admin, returns 400
func (g Group) acceptJoinRequest(ctx *fiber.Ctx) error {
	panic("implement me")
}
