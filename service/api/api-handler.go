package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	rt.router.POST("/session", rt.wrap(rt.createSession)) // doLogin
	rt.router.GET("/users/:username", rt.getUser)         // getUserProfile
	rt.router.POST("/follow/", rt.wrap(rt.followUser))    // followUser

	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
