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
	rt.router.POST("/session", rt.wrap(rt.doLogin))

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

	// setMyUserName
	rt.router.PUT("/users/:username", rt.wrap(rt.setMyUserName))

	// getMyStream
	rt.router.GET("/streams/:streamId", rt.wrap(rt.getMyStream))

	// likePhoto
	rt.router.POST("/likes/", rt.wrap(rt.likePhoto))

	// unlikePhoto
	rt.router.DELETE("/likes/:photoId", rt.wrap(rt.unlikePhoto))

	// commentPhoto
	rt.router.POST("/comments/", rt.wrap(rt.commentPhoto))

	// uncommentPhoto
	rt.router.DELETE("/comments/:commentId", rt.wrap(rt.uncommentPhoto))

	// getPhotos
	rt.router.GET("/photos/", rt.wrap(rt.getPhotos))

	return rt.router
}
