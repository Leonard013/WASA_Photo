package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.GET("/liveness", rt.liveness)

	// doLogin
	rt.router.POST("/session", rt.wrap(rt.createSession))

	// uploadPhoto
	rt.router.POST("/photos/", rt.wrap(rt.uploadPhoto))

	// deletePhoto
	rt.router.DELETE("/photos/:photoId", rt.wrap(rt.deletePhoto))

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
	rt.router.POST("/likes/", rt.wrap(rt.likePhoto))

	// unlikePhoto

	// commentPhoto

	// uncommentPhoto

	return rt.router
}
