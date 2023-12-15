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

	// doLogin
	rt.router.POST("/session", rt.wrap(rt.createSession))

	// uploadPhoto

	// deletePhoto

	// followUser
	rt.router.POST("/follow/", rt.wrap(rt.followUser))

	// unfollowUser
	rt.router.DELETE("/follow/:username", rt.wrap(rt.unfollowUser))

	// banUser
	rt.router.POST("/ban/", rt.wrap(rt.banUser))

	// unbanUser
	rt.router.DELETE("/ban/:username", rt.wrap(rt.unbanUser))

	// getUserProfile
	rt.router.GET("/users/:username", rt.wrap(rt.getUser))

	// setUsername
	rt.router.PUT("/users/:username", rt.wrap(rt.setUsername))

	// getMyStream

	// likePhoto

	// unlikePhoto

	// commentPhoto

	// uncommentPhoto

	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
