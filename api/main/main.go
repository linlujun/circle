package main

import (
	"circleTest/api/session"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}


func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/comment/topic/:topicid", AddComment)
	router.POST("/user/:username", Login)
	router.POST("/info/user/:username", SetUserInfo)
	router.POST("/info/password/user/:username", SetUserPwd)

	router.POST("/circle/user/:username", JoninACircle)

	router.DELETE("/user/:username", DelUser)

	router.GET("/user/:username", GetUserInfo)

	router.POST("/topic", AddNewTopic)
	router.GET("/topic/:topicid", GetTopicInfo)
	//router.PUT("/topic/:topicid", ModifyTopic)
	router.DELETE("/topic/:topicid/:circleid", DelTopic)

	router.GET("/topics/circle/:cid", ListCircleTopics)
	router.GET("/topics/user/:uname", ListUserTopics)
	router.GET("/circles", ListCircles)
	router.GET("/comments/:tid", ListComments)

	
	router.POST("/upload/pics", UploadHandler)

	router.POST("/admin/circle", CreateCircle)
	router.DELETE("/admin/circle/:cid", DelCircle)
	router.PUT("/admin/circle/:cid", SetCircleDesc)
	return router
}

func Prepare() {
	session.LoadSessionsFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
