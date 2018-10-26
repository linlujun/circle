package dbops

import (
	"circleTest/api/defs"
	"circleTest/api/utils"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

/**

*/
//users ops
func AddUserCredential(loginName string, pwd string, role int) error {
	stmtIns, err := dbConn.Prepare(
		"INSERT INTO users (login_name, pwd,role) VALUES (?, ?,?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd, role)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare(
		"SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare(
		"DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func SetUserInfo(loginName, nick, desc, picUrl string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE users SET nickname =?,description=?,userpic=? WHERE login_name=?")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(nick, desc, picUrl, loginName)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
func SetUserPwd(login_name string, pwd string, new_pwd string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE users SET pwd =? WHERE login_name=? AND pwd=?")
	if err != nil {
		return err
	}

	res, err := stmtIns.Exec(new_pwd, login_name, pwd)

	if err != nil {
		return err
	} else if num, _ := res.RowsAffected(); num == 0 {
		return errors.New("modify user password fail")
	}
	defer stmtIns.Close()
	return nil
}

func GetUserInfo(loginName string) (*defs.UserInfo, error) {
	stmtOut, err := dbConn.Prepare(
		`SELECT id, login_name,role,nickname,userpic,
		description,topic_count,comment_count
		 FROM users WHERE login_name = ?`)

	var (
		Id           int
		LoginName    string
		Role         int
		NickName     string
		Desc         string
		PicUrl       string
		TopicCount   int
		CommentCount int
	)

	err = stmtOut.QueryRow(loginName).Scan(
		&Id, &LoginName,
		&Role, &NickName,
		&PicUrl, &Desc,
		&TopicCount, &CommentCount)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.UserInfo{
		Id:           Id,
		LoginName:    LoginName,
		Role:         Role,
		NickName:     NickName,
		Desc:         Desc,
		PicUrl:       PicUrl,
		TopicCount:   TopicCount,
		CommentCount: CommentCount,
	}

	return res, nil
}
func GetUserRole(loginName string) (int, error) {
	stmtOut, err := dbConn.Prepare(
		`SELECT role
		 FROM users WHERE login_name = ?`)

	var Role int

	err = stmtOut.QueryRow(loginName).Scan(&Role)

	if err != nil && err != sql.ErrNoRows {
		return Role, err
	}

	if err == sql.ErrNoRows {
		return Role, nil
	}

	defer stmtOut.Close()

	return Role, nil
}
func SetUserRole(loginName string, role int) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE users SET role =? WHERE login_name=?")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(role, loginName)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func SetUserNickname(loginName string, nickname string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE users SET nickname =? WHERE login_name=?")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(nickname, loginName)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func SetUserDesc(loginName string, description string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE users SET description =? WHERE login_name=?")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(description, loginName)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func SetUserpic(loginName string, pic_url string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE users SET userpic =? WHERE login_name=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(pic_url, loginName)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func incUserTopicCount(loginName string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE users SET topic_count =topic_count+1 WHERE login_name=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
func decUserTopicCount(loginName string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE users SET topic_count =topic_count-1 WHERE login_name=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func incUserCommentCount(loginName string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE users SET comment_count =comment_count+1 WHERE login_name=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func decUserCommentCount(loginName string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE users SET comment_count =comment_count-1 WHERE login_name=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

/****************users dbops end*********************************/

//circles ops
func CreateCircle(cname, master_name, description string) (*defs.CircleInfo, error) {
	stmtIns, err := dbConn.Prepare(
		"INSERT INTO circles (id,cname, master,description) VALUES (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	id, _ := utils.NewUUID()
	_, err = stmtIns.Exec(id, cname, master_name, description)
	if err != nil {
		return nil, err
	}
	ci := &defs.CircleInfo{
		Id:           id,
		Name:         cname,
		Description:  description,
		Master:       master_name,
		UserCount:    0,
		TopicCount:   0,
		CommentCount: 0,
	}
	defer stmtIns.Close()
	return ci, nil
}

func SetCircleDesc(cid string, description string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circles SET description =? WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(description, cid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func incCircleUserCount(cid string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circles SET user_count =user_count+1 WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func decCircleUserCount(cid string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circles SET user_count =user_count-1 WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func incCircleTopicCount(cid string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circles SET topic_count =topic_count+1 WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func decCircleTopicCount(cid string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circles SET topic_count =topic_count-1 WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func incCircleCommentCount(cid string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circles SET Comment_count =Comment_count+1 WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
func decCircleCommentCount(cid string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circles SET Comment_count =Comment_count-1 WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
func DeleteCircle(cid string) error {
	stmtDel, err := dbConn.Prepare(
		"DELETE FROM circles WHERE id=?")
	if err != nil {
		log.Printf("DeleteCircle error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(cid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}
func GetCircleInfo(cid string) (*defs.CircleInfo, error) {
	stmtOut, err := dbConn.Prepare(
		`SELECT id, cname,description,masterid,user_count,
		topic_count,comment_count
		 FROM circles WHERE id = ?`)

	var (
		Cid           string
		Cname         string
		Cdescription  string
		Cmasterid     string
		CuserCcount   int
		CtopicCount   int
		CcommentCount int
	)

	err = stmtOut.QueryRow(cid).Scan(
		&Cid, &Cname,
		&Cdescription,
		&Cmasterid, &CuserCcount,
		&CtopicCount, &CcommentCount)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.CircleInfo{
		Id:           Cid,
		Name:         Cname,
		Description:  Cdescription,
		Master:       Cmasterid,
		UserCount:    CuserCcount,
		TopicCount:   CtopicCount,
		CommentCount: CcommentCount,
	}

	return res, nil
}

func ListCircleInfo() (*defs.CircleInfoList, error) {
	stmtOut, err := dbConn.Prepare(
		`SELECT id, cname,description,master,user_count,
		topic_count,comment_count
		 FROM circles`)

	res := &defs.CircleInfoList{}
	rows, err := stmtOut.Query()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var (
			Cid           string
			Cname         string
			Cdescription  string
			Cmaster       string
			CuserCcount   int
			CtopicCount   int
			CcommentCount int
		)
		if err := rows.Scan(
			&Cid, &Cname,
			&Cdescription,
			&Cmaster, &CuserCcount,
			&CtopicCount, &CcommentCount); err != nil {
			return res, err
		}

		c := &defs.CircleInfo{
			Id:           Cid,
			Name:         Cname,
			Description:  Cdescription,
			Master:       Cmaster,
			UserCount:    CuserCcount,
			TopicCount:   CtopicCount,
			CommentCount: CcommentCount,
		}
		res.Circles = append(res.Circles, c)
	}
	defer stmtOut.Close()
	return res, nil
}

/////////////////topic ops/////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////
func CreateTopic(
	title string, content string, uname, cid string) (*defs.TopicInfo, error) {
	stmtIns, err := dbConn.Prepare(
		"INSERT INTO topics (id,title, content,author,cid,create_time) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}

	id, _ := utils.NewUUID()
	t := time.Now()
	_, err = stmtIns.Exec(id, title, content, uname, cid, t)
	if err != nil {
		return nil, err
	}

	_ = incUserTopicCount(uname)
	_ = incCircleTopicCount(cid)
	_ = incTopicCount(cid, uname)

	ti := &defs.TopicInfo{
		Id:              id,
		Title:           title,
		AuthorLoginName: uname,
		CircleId:        cid,
		Ctime:           t.Format("2006-01-02 15:04:05"),
		Content:         content,
		CommentCount:    0,
	}
	defer stmtIns.Close()
	return ti, nil
}

func ModifyTopic(tid, uname string) {

}

func DeleteTopic(tid, aname, cid string) error {
	//评论是否连带删除
	stmtDel, err := dbConn.Prepare(
		"DELETE FROM topics WHERE id=? AND author=?")
	if err != nil {
		log.Printf("DeleteTopic error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(tid, aname)
	if err != nil {
		return err
	}

	decTopicCount(cid, aname)
	decUserTopicCount(aname)
	decCircleTopicCount(cid)

	defer stmtDel.Close()
	return nil
}

func incTopicCommentCount(tid string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE topics SET Comment_count =Comment_count+1 WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(tid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
func decTopicCommentCount(tid string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE topics SET Comment_count =Comment_count-1 WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(tid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetTopicInfo(tid string) (*defs.TopicInfo, error) {
	stmtOut, err := dbConn.Prepare(
		`SELECT topics.id, topics.title,users.login_name,
		 topics.cid,
		 topics.create_time,topics.content,topics.comment_count
		 FROM topics 
		 INNER JOIN users ON topics.author = users.login_name
		 WHERE topics.id = ?
		 ORDER By topics.create_time DESC`)

	var (
		Id              string
		Title           string
		AuthorLoginName string
		CircleId        string
		Ctime           string
		Content         string
		CommentCount    int
	)

	err = stmtOut.QueryRow(tid).Scan(
		&Id, &Title,
		&AuthorLoginName,
		&CircleId, &Ctime,
		&Content, &CommentCount)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.TopicInfo{
		Id:              Id,
		Title:           Title,
		AuthorLoginName: AuthorLoginName,
		CircleId:        CircleId,
		Ctime:           Ctime,
		Content:         Content,
		CommentCount:    CommentCount,
	}

	return res, nil
}
func ListCircleTopics(circle_id string) (*defs.TopicInfoList, error) {
	stmtOut, err := dbConn.Prepare(`
		 SELECT topics.id, topics.title,users.login_name,
		 topics.cid,
		 topics.create_time,topics.content,topics.comment_count
		 FROM topics 
		 INNER JOIN users ON topics.author = users.login_name
		 WHERE topics.cid = ?
		 ORDER By topics.create_time DESC`)
	res := &defs.TopicInfoList{}
	rows, err := stmtOut.Query(circle_id)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var (
			Id              string
			Title           string
			AuthorLoginName string
			CircleId        string
			Ctime           string
			Content         string
			CommentCount    int
		)
		if err := rows.Scan(
			&Id, &Title, &AuthorLoginName,
			&CircleId, &Ctime, &Content,
			&CommentCount); err != nil {
			return res, err
		}

		c := &defs.TopicInfo{
			Id:              Id,
			Title:           Title,
			AuthorLoginName: AuthorLoginName,
			CircleId:        CircleId,
			Ctime:           Ctime,
			Content:         Content,
			CommentCount:    CommentCount,
		}
		res.Topics = append(res.Topics, c)
	}
	defer stmtOut.Close()
	return res, nil
}
func ListUserTopics(uname string) (*defs.TopicInfoList, error) {
	stmtOut, err := dbConn.Prepare(`
		 SELECT topics.id, topics.title,users.login_name,
		 topics.author,topics.cid,
		 topics.create_time,topics.content,topics.comment_count
		 FROM topics 
		 INNER JOIN users ON topics.author = users.login_name
		 WHERE users.login_name = ?
		 ORDER By topics.create_time DESC`)
	res := &defs.TopicInfoList{}
	rows, err := stmtOut.Query(uname)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var (
			Id              string
			Title           string
			AuthorLoginName string
			CircleId        string
			Ctime           string
			Content         string
			CommentCount    int
		)
		if err := rows.Scan(
			&Id, &Title, &AuthorLoginName,
			&CircleId, &Ctime, &Content,
			&CommentCount); err != nil {
			return res, err
		}

		c := &defs.TopicInfo{
			Id:              Id,
			Title:           Title,
			AuthorLoginName: AuthorLoginName,
			CircleId:        CircleId,
			Ctime:           Ctime,
			Content:         Content,
			CommentCount:    CommentCount,
		}
		res.Topics = append(res.Topics, c)
	}
	defer stmtOut.Close()
	return res, nil
}

/**topic dbops api end**/
////////////////////////////////////////////////////////////////////////////
////////////////////////commment dbops//////////////////////////////////////////

func AddNewComments(
	content, picurl, tid, aname, cid, ccid string) (*defs.Comment, error) {
	stmtIns, err := dbConn.Prepare(
		`INSERT INTO comments (id,content, picurl,topic_id,
			author_name,circle_id,comment_id,create_time) VALUES (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return nil, err
	}

	t := time.Now()
	id, _ := utils.NewUUID()
	_, err = stmtIns.Exec(id, content, picurl, tid, aname, cid, ccid, t)
	if err != nil {
		return nil, err
	}
	//circle_user
	incCommentCount(cid, aname)
	//circle
	incCircleCommentCount(cid)
	//topics
	incTopicCommentCount(tid)
	//user
	incUserCommentCount(aname)

	res := &defs.Comment{
		Id:         id,
		TopicId:    tid,
		Content:    content,
		PicUrl:     picurl,
		AuthorName: aname,
		CircleId:   cid,
		CommentId:  ccid,
		Like:       0,
	}
	defer stmtIns.Close()
	return res, nil
}

func ListComments(tid string) (*defs.CommentList, error) {
	stmtOut, err := dbConn.Prepare(`
		 SELECT id, comment_id,content,
		 picurl,topic_id,author_name,
		 circle_id ,create_time,like_count
		 FROM comments
		 WHERE topic_id = ?
		 ORDER BY like_count DESC
		`)
	res := &defs.CommentList{}
	rows, err := stmtOut.Query(tid)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var (
			Id         string
			TopicId    string
			Content    string
			PicUrl     string
			AuthorName string
			CircleId   string
			CommentId  string
			CreateTime string
			Like       int
		)
		if err := rows.Scan(
			&Id, &CommentId, &Content,
			&PicUrl, &TopicId, &AuthorName,
			&CircleId, &CreateTime, &Like); err != nil {
			return res, err
		}

		c := &defs.Comment{
			Id:         Id,
			TopicId:    TopicId,
			Content:    Content,
			PicUrl:     PicUrl,
			AuthorName: AuthorName,
			CircleId:   CircleId,
			CommentId:  CommentId,
			CreateTime: CreateTime,
			Like:       Like,
		}
		res.Comments = append(res.Comments, c)
	}
	defer stmtOut.Close()
	return res, nil
}

func DeleteComment(comment_id string) error {
	stmtDel, err := dbConn.Prepare(
		"DELETE FROM comments WHERE id=?")
	if err != nil {
		log.Printf("Delete Comment error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(comment_id)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

//对评论点赞时调用
func IncCommentLike(cid string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE comments SET like_count =like_count+1 WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// func ListComments(topic_id string, from, to int) ([]*defs.Comment, error) {

// 	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name,
// 		comments.content,comments.picurl,comment_id,like_count FROM comments
// 		INNER JOIN users ON comments.author_name = users.login_name
// 		WHERE comments.topic_id = ?
// 		AND comments.create_time > FROM_UNIXTIME(?)
// 		AND comments.create_time <= FROM_UNIXTIME(?)
// 		ORDER By comments.create_time DESC`)
// 	var res []*defs.Comment

// 	rows, err := stmtOut.Query(topic_id, from, to)
// 	if err != nil {
// 		return res, err
// 	}

// 	for rows.Next() {
// 		var (
// 			id         string
// 			uname      string
// 			content    string
// 			picurl     string
// 			comment_id string
// 			like       int
// 		)
// 		if err := rows.Scan(
// 			&id, &uname, &content, &picurl, &comment_id, &like); err != nil {
// 			return res, err
// 		}

// 		c := &defs.Comment{
// 			Id:         id,
// 			Content:    content,
// 			PicUrl:     picurl,
// 			AuthorName: uname,
// 			CommentId:  comment_id,
// 			Like:       like,
// 		}
// 		res = append(res, c)
// 	}
// 	defer stmtOut.Close()

// 	return res, nil
// }

//////////////////////////////////////////////
///////////////masters dbops//////////////////
func AddNewMaster(mname, cid string) error {
	stmtIns, err := dbConn.Prepare(
		`INSERT INTO masters (mname,cid,create_time) VALUES (?,?,?)`)
	if err != nil {
		return err
	}

	t := time.Now()
	_, err = stmtIns.Exec(mname, cid, t)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func IsMaster(username string) (string, bool, error) {
	stmOut, err := dbConn.Prepare(`SELECT cid FROM masters WHERE mname=?`)
	if err != nil {
		return "", false, err
	}
	var cid string
	err = stmOut.QueryRow(username).Scan(&cid)
	if err != nil {
		return "", false, err
	}
	defer stmOut.Close()
	return cid, true, nil
}

/////////////////////////////////////////////////
////////////////////circle_user map/////////////
func JoinCircle(uname, cid string) error {
	stmtIns, err := dbConn.Prepare(
		`INSERT INTO circle_user (uname,cid,joined) VALUES (?,?,?)`)
	if err != nil {
		return err
	}

	t := time.Now()
	_, err = stmtIns.Exec(uname, cid, t)
	if err != nil {
		return err
	}
	//相应圈子人数加 1
	err = incCircleUserCount(cid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func incTopicCount(cid, uname string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circle_user SET topic_count =topic_count+1 WHERE cid=? AND uname=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
func decTopicCount(cid, uname string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circle_user SET topic_count =topic_count-1 WHERE cid=? AND uname=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
func incCommentCount(cid, uname string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circle_user SET comment_count =comment_count+1 WHERE cid=? AND uname=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
func decCommentCount(cid, uname string) error {
	stmtIns, err := dbConn.Prepare(
		"UPDATE circle_user SET comment_count =comment_count-1 WHERE cid=? AND uname=?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
