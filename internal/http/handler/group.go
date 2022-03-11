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

// create a new group
// if user is in a group, return 400
func (g Group) create(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getAll returns all groups from newest to oldest
func (g Group) getAll(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getMyGroups returns all groups that the user is a member of
// members must be sorted by join date from newest to the oldest
// if user is not in any groups, return 400
func (g Group) getMyGroups(ctx *fiber.Ctx) error {
	panic("implement me")
}
