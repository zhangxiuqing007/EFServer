package dba

import (
	"EFServer/forum"
	"EFServer/tool"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//测试资源创建者，放置包内有太多的全局函数
type testResourceBuilder struct {
}

//随机种子值
func (t *testResourceBuilder) initRandomSeed() {
	rand.Seed(time.Now().UnixNano())
}

//制造当前测试类型的sql对象
func (t *testResourceBuilder) buildCurrentTestSQLIns() *sqlBase {
	const testNowIsMysql = true
	if testNowIsMysql {
		db := MySQLIns{}
		db.Open("mysql5856")
		return &db.sqlBase
	}
	db := &SqliteIns{}
	db.Open("../ef.db")
	return &db.sqlBase
}

//生成随机主题
func (t *testResourceBuilder) buildRandomTheme(count int) []*forum.ThemeInDB {
	tms := make([]*forum.ThemeInDB, 0, count)
	for i := 0; i < count; i++ {
		newTheme := new(forum.ThemeInDB)
		newTheme.Name = fmt.Sprintf("随机主题：%d", i+1)
		tms = append(tms, newTheme)
	}
	return tms
}

//生成随机用户
func (t *testResourceBuilder) buildRandomUsers(count int) []*forum.UserInDB {
	users := make([]*forum.UserInDB, 0, count)
	for i := 0; i < count; i++ {
		newUser := new(forum.UserInDB)
		newUser.Account = tool.NewUUID()
		newUser.PassWord = tool.NewUUID()
		newUser.Name = tool.NewUUID()
		if rand.Intn(2) == 1 {
			newUser.UserType = forum.UserTypeAdministrator
		} else {
			newUser.UserType = forum.UserTypeNormalUser
		}
		newUser.UserState = forum.UserStateNormal
		newUser.SignUpTime = time.Now().UnixNano()
		users = append(users, newUser)
	}
	return users
}

//生成随机帖子
func (t *testResourceBuilder) buildRandomPost(themeID, userID int64) *forum.PostInDB {
	post := new(forum.PostInDB)
	post.ThemeID = themeID
	post.UserID = userID
	post.Title = t.buildRandomPostTitle()
	post.State = forum.PostStateNormal
	return post
}

//生成随机评论
func (t *testResourceBuilder) buildRandomCmt(postID, userID int64) *forum.CommentInDB {
	cmt := new(forum.CommentInDB)
	cmt.PostID = postID
	cmt.UserID = userID
	cmt.Content = t.buildRandomPostContent()
	cmt.State = forum.CmtStateNormal
	cmt.CreatedTime = time.Now().UnixNano()
	cmt.LastEditTime = cmt.CreatedTime
	cmt.EditTimes = rand.Int()%5 + 1
	cmt.PraiseTimes = rand.Int() % 10
	cmt.BelittleTimes = rand.Int() % 20
	return cmt
}

//生成随机标题
func (t *testResourceBuilder) buildRandomPostTitle() string {
	return t.combineUuids(rand.Int()%6 + 1)
}

//生成随机内容
func (t *testResourceBuilder) buildRandomPostContent() string {
	return t.combineUuids(rand.Int()%20 + 1)
}

//合并uuid
func (t *testResourceBuilder) combineUuids(count int) string {
	var uids = make([]string, 0, count)
	for i := 0; i < count; i++ {
		uids = append(uids, tool.NewUUID())
	}
	return strings.Join(uids, "#")
}

//判断两个帖子内容是否相同
func (t *testResourceBuilder) isTwoPostSame(post1, post2 *forum.PostInDB) bool {
	return post1.ID == post2.ID &&
		post1.ThemeID == post2.ThemeID &&
		post1.UserID == post2.UserID &&
		post1.Title == post2.Title &&
		post1.State == post2.State
}

//判断两个用户是否相同
func (t *testResourceBuilder) isTwoUserSame(user1, user2 *forum.UserInDB) bool {
	return user1.ID == user2.ID &&
		user1.Account == user2.Account &&
		user1.PassWord == user2.PassWord &&
		user1.Name == user2.Name &&
		user1.UserType == user2.UserType &&
		user1.UserState == user2.UserState &&
		user1.SignUpTime == user2.SignUpTime
}
