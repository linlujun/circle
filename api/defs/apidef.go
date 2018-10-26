package defs

//reqeusts
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type Success struct {
	Success bool `json:"success"`
}

type SignedIn struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

// Data model
type UserInfo struct {
	Id           int    `json:"id"`
	LoginName    string `json:"login_name"`
	Role         int    `json:"role"`
	NickName     string `json:"nake_name"`
	Desc         string `json:"desc"`
	PicUrl       string `json:"pic_url"`
	TopicCount   int    `json:"topic_count"`
	CommentCount int    `json:"comment_count"`
}

//修改用户信息
type SetUserInfo struct {
	NickName string `json:"nake_name"`
	Desc     string `json:"desc"`
	PicUrl   string `json:"pic_url"`
}

type SetUserPwd struct {
	OrgPwd string `json:"org_pwd"`
	NewPwd string `json:"new_pwd"`
}
type JoninACircle struct {
	Cid string `json:"cid"`
}

//用于返回
type TopicInfo struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	AuthorLoginName string `json:"author_name"`
	CircleId        string `json:"circle_id"`
	Content         string `json:"content"`
	Ctime           string `json:"creat_time"`
	CommentCount    int    `json:"comment_count"`
}
type TopicInfoList struct {
	Topics []*TopicInfo `json:"topics"`
}

//用于创建
type NewTopic struct {
	Title    string `json:"title"`
	CircleId string `json:"circle_id"`
	Content  string `json:"content"`
}

type CircleInfo struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Master       string `json:"master"`
	UserCount    int    `json:"user_count"`
	TopicCount   int    `json:"topic_count"`
	CommentCount int    `json:"comment_count"`
}
type CircleInfoList struct {
	Circles []*CircleInfo `json:"circles"`
}
type NewCircle struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type CircleDesc struct {
	Description string `json:"description"`
}
type Comment struct {
	Id         string `json:"id"`
	TopicId    string `json:"topic_id"`
	Content    string `json:"content"`
	PicUrl     string `json:"pic_url"`
	AuthorName string `json:"author_name"`
	CircleId   string `json:"ciecle_id"`
	CommentId  string `json:"comment_id"`
	CreateTime string `json:"create_time"`
	Like       int    `json:"like"`
}
type CommentList struct {
	Comments []*Comment `json:"comments"`
}
type NewComment struct {
	Content   string `json:"content"`
	PicUrl    string `json:"pic_url"`
	CommentId string `json:"comment_id"`
	CircleId  string `json:"circle_id"`
}
type SimpleSession struct {
	Username string //login name
	TTL      int64
}
