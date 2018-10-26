package dbops

import (
	"circleTest/api/defs"
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*以下所有测试均已通过，后期api已发生改动，再次测试前注意更改测试方法
topics、comments、circles id 类型改变
topic contanturl->contant
*/
func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate circles")
	dbConn.Exec("truncate topics")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("SetRole", testSetRole)
	t.Run("SetNick", testSetNick)
	t.Run("SetDesc", testSetDesc)
	t.Run("SetPic", testSetPic)
	t.Run("IncTopic", testIncTopic)
	t.Run("DecTopic", testDecTopic)
	t.Run("IncComment", testIncComment)
	t.Run("DecComment", testDecComment)
	t.Run("GetInfo", testGetUserInfo)
	t.Run("SetPwd", testSetPwd)
	t.Run("Get", testGetUser)

	t.Run("DelUer", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func TestCircleWorkFlow(t *testing.T) {
	t.Run("Add", testAddCircle)
	t.Run("SetDesc", testSetCircleDesc)
	t.Run("incUserCount", testIncCircleUserCount)
	t.Run("incUserCount", testIncCircleUserCount)
	t.Run("decUserCount", testDecCircleUserCount)
	t.Run("incTopicCount", testIncCircleTopicCount)
	t.Run("incTopicCount", testIncCircleTopicCount)
	t.Run("decTopicCount", testDecCircleTopicCount)
	t.Run("incCommentCount", testIncCircleCommentCount)
	t.Run("incCommentCount", testIncCircleCommentCount)
	t.Run("decCommentCount", testDecCircleCommentCount)
	t.Run("GetInfo", testGetCircleInfo)
	t.Run("Del", testDelCircle)
	//t.Run("Reget", testRegetCircle)
}

func TestTopicWorkFlow(t *testing.T) {
	t.Run("CreateTopic", testCreateTopic)
	t.Run("incTopicCommentCount", testIncTopicCommentCount)
	t.Run("decTopicCommentCount", testDecTopicCommentCount)
	t.Run("GetTopicInfo", testGetTopicInfo)
	t.Run("DeleteTopic", testDeleteTopic)
	//t.Run("Reget", testRegetTopic)
}

///////////////////user test flow/////////////////////////////

func testAddUser(t *testing.T) {
	err := AddUserCredential("20161004054", "lljllj", 1222)
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("20161004054")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("20161004054", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("20161004054")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func testSetRole(t *testing.T) {
	err := SetUserRole("20161004054", 2)
	if err != nil {
		t.Errorf("Error of SetRole: %v", err)
	}
}

func testSetNick(t *testing.T) {
	err := SetUserNickname("20161004054", "jj")
	if err != nil {
		t.Errorf("Error of SetNick: %v", err)
	}
}

func testSetDesc(t *testing.T) {
	err := SetUserDesc("20161004054", "my name is llj")
	if err != nil {
		t.Errorf("Error of SetDesc: %v", err)
	}
}

func testSetPwd(t *testing.T) {
	err := SetUserPwd("20161004054", "lljllj", "123")
	if err != nil {
		t.Errorf("Error of SetPwd: %v", err)
	}
}

func testSetPic(t *testing.T) {
	err := SetUserpic("20161004054", "./userpic/user/20161004054/pic001.jpg")
	if err != nil {
		t.Errorf("Error of SetPic: %v", err)
	}
}

func testIncTopic(t *testing.T) {
	err := incUserTopicCount("20161004054")
	if err != nil {
		t.Errorf("Error of IncTopic: %v", err)
	}

	info, err := GetUserInfo("20161004054")
	num := info.TopicCount
	if err != nil || num != 1 {
		t.Errorf("IncTopic error")
	}
}
func testIncComment(t *testing.T) {
	err := incUserCommentCount("20161004054")
	if err != nil {
		t.Errorf("Error of IncComment: %v", err)
	}
	info, err := GetUserInfo("20161004054")
	num := info.CommentCount
	if err != nil || num != 1 {
		t.Errorf("IncTopic error")
	}
}
func testDecTopic(t *testing.T) {
	err := decUserTopicCount("20161004054")
	if err != nil {
		t.Errorf("Error of DecTopic: %v", err)
	}
	info, err := GetUserInfo("20161004054")
	num := info.TopicCount
	if err != nil || num != 0 {
		t.Errorf("DecTopic error")
	}
}
func testDecComment(t *testing.T) {
	err := decUserCommentCount("20161004054")
	if err != nil {
		t.Errorf("Error of DecComment: %v", err)
	}
	info, err := GetUserInfo("20161004054")
	num := info.CommentCount
	if err != nil || num != 0 {
		t.Errorf("IncTopic error")
	}
}

func testGetUserInfo(t *testing.T) {
	info := &defs.UserInfo{}
	info, err := GetUserInfo("20161004054")
	if err != nil {
		t.Errorf("Get user info error")
	}
	if info.LoginName != "20161004054" || info.NickName != "jj" ||
		info.Role != 2 || info.Desc != "my name is llj" ||
		info.PicUrl != "./userpic/user/20161004054/pic001.jpg" ||
		info.CommentCount != 0 || info.TopicCount != 0 {
		t.Errorf("user info error")

	}
}

////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////
////////////////////////circles test flow///////////////////////
func testAddCircle(t *testing.T) {
	err := CreateCircle("足球圈", 1)
	if err != nil {
		t.Errorf("Error of AddCircle: %v", err)
	}
}

func testSetCircleDesc(t *testing.T) {

	err := SetCircleDesc(1, "足球爱好者聚集地")
	if err != nil {
		t.Errorf("SetCircleDesc Error: %v", err)
	}
}

func testIncCircleUserCount(t *testing.T) {
	err := incCircleUserCount(1)
	if err != nil {
		t.Errorf("IncCircleUserCount Error: %v", err)
	}
}

func testDecCircleUserCount(t *testing.T) {
	err := decCircleUserCount(1)
	if err != nil {
		t.Errorf("decCircleUserCount Error: %v", err)
	}
}

func testIncCircleTopicCount(t *testing.T) {
	err := incCircleTopicCount(1)
	if err != nil {
		t.Errorf("incCircleTopicCount Error: %v", err)
	}
}

func testDecCircleTopicCount(t *testing.T) {
	err := decCircleTopicCount(1)
	if err != nil {
		t.Errorf("decCircleTopicCount Error: %v", err)
	}
}
func testIncCircleCommentCount(t *testing.T) {
	err := incCircleCommentCount(1)
	if err != nil {
		t.Errorf("incCircleCommentCount Error: %v", err)
	}
}
func testDecCircleCommentCount(t *testing.T) {
	err := decCircleCommentCount(1)
	if err != nil {
		t.Errorf("decCircleCommentCount Error: %v", err)
	}
}

func testGetCircleInfo(t *testing.T) {
	c := &defs.CircleInfo{}
	c, err := GetCircleInfo(1)
	if err != nil {
		t.Errorf("GetCircleInfo error")
	}
	if c.Id != 1 || c.MasterId != 1 || c.Description != "足球爱好者聚集地" ||
		c.Name != "足球圈" || c.TopicCount != 1 ||
		c.UserCount != 1 || c.CommentCount != 1 {
		t.Errorf("circle infomation error")
	}
}

func testDelCircle(t *testing.T) {
	err := DeleteCircle(1)
	if err != nil {
		t.Errorf("DeleteCircle Error: %v", err)
	}
}

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
//////////////////////////topic test flow//////////////////////

func testCreateTopic(t *testing.T) {
	err := CreateTopic("first topic", "/topics", 1, 1)
	if err != nil {
		t.Errorf("Add topic Error: %v", err)
	}
}

func testIncTopicCommentCount(t *testing.T) {
	if err != nil {
		t.Errorf("DeleteCircle Error: %v", err)
	}
}

func testDecTopicCommentCount(t *testing.T) {
	if err != nil {
		t.Errorf("DeleteCircle Error: %v", err)
	}
}

func testGetTopicInfo(t *testing.T) {
	info, err := GetTopicInfo(1)
	if err != nil || info.Title != "first topic" || info.AuthorId != 1 || info.CircleId != 1 {
		t.Errorf("topic information Error: %v", err)
	}
}

func testDeleteTopic(t *testing.T) {
	if err != nil {
		t.Errorf("DeleteCircle Error: %v", err)
	}
}

///////////////////////////////////////////////////////////
///////////////////////comment test flow///////////////////

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddCommnets", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	err := AddNewComments("this is comment content", "/comment/pic", 1, 1, 1, 0)
	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}
func testListComments(t *testing.T) {
	tid := 1
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(tid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i+1, ele)
	}
}
