package main

import (
	"circleTest/api/dbops"
	"circleTest/api/defs"
	"circleTest/api/session"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

// 注册逻辑
//请求为json
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//读取请求体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	//反序列化json；此处对json内容不做检查，前端应完成检查
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	//写入DB users表
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd, 0); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	//生成session ID ，该函数同时会写入DB session表
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}
	//序列化消息体并发回前端
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

//登录逻辑
//json携带账号密码
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", res)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	//验证请求体
	uname := p.ByName("username")
	log.Printf("login url name:%s", uname)
	log.Printf("login body name:%s", ubody.Username)
	if uname != ubody.Username {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	log.Printf("%s", ubody.Username)
	pwd, err := dbops.GetUserCredential(ubody.Username)
	log.Printf("Login pwd:%s", pwd)
	log.Printf("Login body pwd:%s", ubody.Pwd)
	if err != nil || len(pwd) == 0 || pwd != ubody.Pwd {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	//分配session id 并存到cache—>DB
	id := session.GenerateNewSessionId(ubody.Username)
	si := &defs.SignedIn{Success: true, SessionId: id}
	if resp, err := json.Marshal(si); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

// NickName string `json:"nake_name"`
// Desc     string `json:"desc"`
// PicUrl   string `json:"pic_url"`
func SetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	n, ok := ValidateUser(w, r)
	if !ok || n != uname {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	//鉴权通过

	//解析请求
	res, _ := ioutil.ReadAll(r.Body)
	body := &defs.SetUserInfo{}
	if err := json.Unmarshal(res, body); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	err := dbops.SetUserInfo(uname, body.NickName, body.Desc, body.PicUrl)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
	}

	s := &defs.Success{Success: true}
	if resp, err := json.Marshal(s); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func SetUserPwd(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	n, ok := ValidateUser(w, r)
	if !ok || n != uname {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	//登录状态中

	//解析请求
	res, _ := ioutil.ReadAll(r.Body)
	body := &defs.SetUserPwd{}
	if err := json.Unmarshal(res, body); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	err := dbops.SetUserPwd(uname, body.OrgPwd, body.NewPwd)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	s := &defs.Success{Success: true}
	if resp, err := json.Marshal(s); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

//加入圈子
func JoninACircle(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//本人操作
	uname := p.ByName("username")
	n, ok := ValidateUser(w, r)
	if !ok || n != uname {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	//已经登录

	//解析请求
	res, _ := ioutil.ReadAll(r.Body)
	body := &defs.JoninACircle{}
	if err := json.Unmarshal(res, body); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	err := dbops.JoinCircle(uname, body.Cid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	s := &defs.Success{Success: true}
	if resp, err := json.Marshal(s); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

//Delete user
func DelUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	if uname != ubody.Username {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	err := dbops.DeleteUser(ubody.Username, ubody.Pwd)
	if err != nil {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	s := &defs.Success{Success: true}
	if resp, err := json.Marshal(s); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

//获取用户信息
//url:/user/:logi_nname

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !validateUserSession(r) {
		log.Printf("Unathorized user\n")
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	//已经登录
	uname := p.ByName("username")
	u, err := dbops.GetUserInfo(uname)
	if err != nil {
		log.Printf("Error in GetUserInfo:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(u); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

//贴子发表
// {
// 	"title":"wenzhang",
// 	"circle_id":"b87e8748-d55e-430a-a2c5-0ac00b131e20",
// 	"content":"23333333333333"
// }

func AddNewTopic(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname, ok := ValidateUser(w, r)
	if !ok {
		log.Printf("Unathorized user\n")
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	ntbody := &defs.NewTopic{}
	if err := json.Unmarshal(res, ntbody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	ti, err := dbops.CreateTopic(ntbody.Title, ntbody.Content, uname, ntbody.CircleId)
	log.Printf("author:%s,title:%s", uname, ntbody.Title)
	if err != nil {
		log.Printf("Error in new topic:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	if resp, err := json.Marshal(ti); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func DelTopic(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname, ok := ValidateUser(w, r)
	if !ok {
		log.Printf("Unathorized user\n")
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	tid := p.ByName("topicid")
	cid := p.ByName("circleid")
	err := dbops.DeleteTopic(tid, uname, cid)

	if err != nil {
		log.Printf("Error in delete topic:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	s := &defs.Success{Success: true}
	if resp, err := json.Marshal(s); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

//获取贴子详情
// /topic/:topicsid", GetTopicInfo

func GetTopicInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	tid := p.ByName("topicid")
	t, err := dbops.GetTopicInfo(tid)
	if err != nil {
		log.Printf("Error in TopicInfo:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(t); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

//modify topic

// func ModifyTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	uname, ok := ValidateUser(w, r)
// 	if !ok {
// 		log.Printf("Unathorized user\n")
// 		sendErrorResponse(w, defs.ErrorNotAuthUser)
// 		return
// 	}
// 	tid := p.ByName("topicid")
// }

//列出某个圈子的所有帖子

func ListCircleTopics(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	cid := p.ByName("cid")
	ts, err := dbops.ListCircleTopics(cid)
	if err != nil {
		log.Printf("Error in TopicInfo:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(ts); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func ListUserTopics(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	uname := p.ByName("uname")
	ts, err := dbops.ListUserTopics(uname)
	if err != nil {
		log.Printf("Error in TopicInfo:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(ts); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func ListCircles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	clist, err := dbops.ListCircleInfo()
	if err != nil {
		log.Printf("Error in ListCircleInfo:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(clist); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func ListComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tid := p.ByName("tid")
	clist, err := dbops.ListComments(tid)
	if err != nil {
		log.Printf("Error in ListComments:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(clist); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

// 	Content    string `json:"content"`
// 	PicUrl     string `json:"pic_url"`
// 	CommentId  string `json:"comment_id"`
//  CircleId   string `json:"circle_id"`
func AddComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname, ok := ValidateUser(w, r)
	if !ok {
		log.Printf("Unathorized user\n")
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	tid := p.ByName("topicid")
	res, _ := ioutil.ReadAll(r.Body)
	body := &defs.NewComment{}
	if err := json.Unmarshal(res, body); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	ci, err := dbops.AddNewComments(
		body.Content, body.PicUrl, tid, uname, body.CircleId, body.CommentId)

	if err != nil {
		log.Printf("Error in new comment:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	if resp, err := json.Marshal(ci); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

//**管理员接口******************************************************/

//创建圈子
func CreateCircle(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname, ok := ValidateUser(w, r)
	if !ok {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	//管理员
	if !IsAdmin(w, r) {
		sendErrorResponse(w, defs.ErrorRoleFaults)
		return
	}

	//登录状态中

	//解析请求
	res, _ := ioutil.ReadAll(r.Body)
	body := &defs.NewCircle{}
	if err := json.Unmarshal(res, body); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	ci, err := dbops.CreateCircle(body.Name, uname, body.Description)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	if resp, err := json.Marshal(ci); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

//删除圈子
func DelCircle(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, ok := ValidateUser(w, r)
	if !ok {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	if !IsAdmin(w, r) {
		sendErrorResponse(w, defs.ErrorRoleFaults)
		return
	}
	//登录状态中
	cid := p.ByName("cid")
	err := dbops.DeleteCircle(cid)

	if err != nil {
		log.Printf("Error in delete circle:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	s := &defs.Success{Success: true}
	if resp, err := json.Marshal(s); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

//设置圈子描述
func SetCircleDesc(
	w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, ok := ValidateUser(w, r)
	if !ok {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	if !IsAdmin(w, r) {
		sendErrorResponse(w, defs.ErrorRoleFaults)
		return
	}
	//登录状态中

	cid := p.ByName("cid")
	res, _ := ioutil.ReadAll(r.Body)
	body := &defs.CircleDesc{}
	if err := json.Unmarshal(res, body); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	err := dbops.SetCircleDesc(cid, body.Description)

	if err != nil {
		log.Printf("Error in set circle description:%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	s := &defs.Success{Success: true}
	if resp, err := json.Marshal(s); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

//上传
func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		log.Printf("error in ParseMultipartForm :File is too big")
		return
	}
	//file
	file, head, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	//获取文件后缀 确定为图片
	fileSuffix := strings.ToLower(path.Ext(head.Filename))
	if fileSuffix != ".bmp" && fileSuffix != ".jpg" &&
		fileSuffix != ".jpeg" && fileSuffix != ".png" && fileSuffix != ".gif" {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		log.Printf("file type error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, defs.ErrorInternalFaults)
	}
	t := time.Now().UnixNano()

	ts := strconv.FormatInt(t, 10)
	fileurl := PIC_DIR + ts + head.Filename
	err = ioutil.WriteFile(fileurl, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	w.WriteHeader(http.StatusCreated)
	url := &FilePath{}
	url.Url = fileurl
	js, _ := json.Marshal(url)
	//返回服务器位置
	sendNormalResponse(w, string(js), 200)
}

/*
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	websocket.Upgrade(w, r, responseHeader, readBufSize, writeBufSize)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	clients[ws] = true

	for {
		var msg Message // Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}
func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		log.Fatal(msg)
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
*/
