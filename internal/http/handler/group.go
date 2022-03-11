package handler

import "github.com/gofiber/fiber/v2"

type Group struct {
}

func (g Group) Register(app *fiber.App) {
	app.Post("/api/v1/groups", g.create)
	app.Get("/api/v1/groups", g.getAll)
	app.Get("/api/v1/groups/my", g.getMyGroups)
	app.Post("/api/v1/join_requests", g.join)
	app.Get("/api/v1/join_requests", g.getJoinRequests)
	app.Get("/api/v1/join_requests/group", g.getJoinRequestsByGroup)
	app.Post("/api/v1/join_requests/accept", g.acceptJoinRequest)
	app.Post("/api/v1/connection_requests", g.connectGroups)
	app.Get("/api/v1/connection_requests", g.getConnectionRequests)
	app.Get("/api/v1/connection_requests/accept", g.acceptConnectionRequest)
}

func (g Group) create(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getAll returns all groups from newest to oldest
func (g Group) getAll(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getMyGroups returns all groups that the user is a member of
func (g Group) getMyGroups(ctx *fiber.Ctx) error {
	panic("implement me")
}

// join requests to join a group
func (g Group) join(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getJoinRequests returns user's all join requests to other groups
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

// connectGroups connects two groups
// if user is not admin, returns 400
func (g Group) connectGroups(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getConnectionRequests returns all connection requests for the user
// newest to oldest
// if user is not admin, returns 400
func (g Group) getConnectionRequests(ctx *fiber.Ctx) error {
	panic("implement me")
}

// acceptConnectionRequest accepts a connection request
// if user is not admin, returns 400
func (g Group) acceptConnectionRequest(ctx *fiber.Ctx) error {
	panic("implement me")
}
