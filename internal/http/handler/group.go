package handler

import "github.com/gofiber/fiber/v2"

type Group struct {
}

func (g Group) Register(r fiber.Router) {
	r.Post("/groups", g.create)
	r.Get("/groups", g.getAll)
	r.Get("/groups/my", g.getMyGroups)
	r.Post("/join_requests", g.join)
	r.Get("/join_requests", g.getJoinRequests)
	r.Get("/join_requests/group", g.getJoinRequestsByGroup)
	r.Post("/join_requests/accept", g.acceptJoinRequest)
	r.Post("/connection_requests", g.connectGroups)
	r.Get("/connection_requests", g.getConnectionRequests)
	r.Get("/connection_requests/accept", g.acceptConnectionRequest)
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
