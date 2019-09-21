package dba

import (
	"EFServer/forum"
	"EFServer/tool"
	"math/rand"
	"strings"
	"testing"
	"time"
)

//用户增删改查操作
func Test_UserOperations(t *testing.T) {
	initRandomNameData()
	const testCount = 5
	ioTool := &SqliteIns{}
	ioTool.Open("../ef.db")
	defer ioTool.Close()
	//创建testCount个用户
	testUsers := make([]*forum.UserInDB, 0, testCount)
	for {
		newUser := buildRandomUser()
		if ioTool.IsUserNameExist(newUser.Name) {
			continue
		}
		if ioTool.IsUserAccountExist(newUser.Account) {
			continue
		}
		if ioTool.AddUser(newUser) != nil {
			t.Error("新增用户失败")
			t.FailNow()
		}
		testUsers = append(testUsers, newUser)
		if len(testUsers) >= testCount {
			break
		}
	}
	//查、校验
	for _, u := range testUsers {
		if rand.Float64() > 0.5 {
			tempUser, err := ioTool.QueryUserByID(u.ID)
			if err != nil {
				t.Error("查询失败")
				t.FailNow()
			}
			if !isTwoUserSame(u, tempUser) {
				t.Error("校验失败")
				t.FailNow()
			}
		} else {
			tempUser, err := ioTool.QueryUserByAccountAndPwd(u.Account, u.PassWord)
			if err != nil {
				t.Error("查询失败")
				t.FailNow()
			}
			if !isTwoUserSame(u, tempUser) {
				c := 0
				c++
				t.Error("校验失败")
				t.FailNow()
			}
		}
	}
	//删
	for _, u := range testUsers {
		if ioTool.DeleteUser(u.ID) != nil {
			t.FailNow()
		}
	}
}

func buildRandomUser() *forum.UserInDB {
	randUser := new(forum.UserInDB)
	randUser.Name = tool.RandomChineseName()
	randUser.Account = tool.NewUUID()
	randUser.PassWord = tool.NewUUID()
	randUser.SignUpTime = time.Now().UnixNano()
	randUser.UserState = forum.UserStateNormal
	randUser.UserType = forum.UserTypeNormalUser
	return randUser
}

func isTwoUserSame(user1, user2 *forum.UserInDB) bool {
	return user1.ID == user2.ID &&
		user1.Name == user2.Name &&
		user1.Account == user2.Account &&
		user1.PassWord == user2.PassWord &&
		user1.SignUpTime == user2.SignUpTime &&
		user1.UserState == user2.UserState &&
		user1.UserType == user2.UserType
}

//主题增删改查操作
func Test_ThemeOperations(t *testing.T) {
	initRandomNameData()
	const testCount = 5
	ioTool := &SqliteIns{}
	ioTool.Open("../ef.db")
	defer ioTool.Close()
	//增
	for i := 0; i < testCount; i++ {
		tmName := buildRandomThemeName()
		if _, err := ioTool.AddTheme(tmName); err != nil {
			t.Error("增加")
			t.FailNow()
		}
	}
	//查
	tms, err := ioTool.QueryThemes()
	if err != nil {
		t.Error("查")
		t.FailNow()
	}
	//改
	for _, tm := range tms {
		tm.Name = buildRandomThemeName()
		if ioTool.UpdateTheme(tm) != nil {
			t.Error("改")
			t.FailNow()
		}
	}
	//删
	for _, tm := range tms {
		if ioTool.DeleteTheme(tm.ID) != nil {
			t.Error("删")
			t.FailNow()
		}
	}
}

func buildRandomTheme() *forum.ThemeInDB {
	return &forum.ThemeInDB{Name: buildRandomThemeName()}
}

func buildRandomThemeName() string {
	uuidStr := tool.NewUUID()
	rs := []rune(uuidStr)
	return string(rs[0 : rand.Int()%32+5])
}

//帖子和评论增删改查操作
func Test_PostAndCmt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	initRandomNameData()
	const TestUserCount = 11
	const CmtCountOfOneUser = 10
	ioTool := new(SqliteIns)
	ioTool.Open("../ef.db")
	defer ioTool.Close()
	//创建主题
	tmIns, err := ioTool.AddTheme(buildRandomThemeName())
	if err != nil {
		t.Error("新增主题失败")
		t.FailNow()
	}
	//创建用户 n 个
	users := make([]*forum.UserInDB, 0, TestUserCount)
	for {
		//创建用户，直到够数
		newUser := buildRandomUser()
		if ioTool.IsUserNameExist(newUser.Name) {
			continue
		}
		if ioTool.IsUserAccountExist(newUser.Account) {
			continue
		}
		if ioTool.AddUser(newUser) != nil {
			t.Error("新增用户失败")
			t.FailNow()
		} else {
			users = append(users, newUser)
		}
		if len(users) >= TestUserCount {
			break
		}
	}
	//由第一个用户 创建一个帖子
	post := buildRandomPost(tmIns.ID, users[0].ID)
	//增加帖子
	if ioTool.AddPost(post) != nil {
		t.Error("增加帖子失败")
		t.FailNow()
	}
	//制造评论
	cmts := make([]*forum.CommentInDB, 0, 30)
	for i := 0; i < CmtCountOfOneUser; i++ {
		for j := 0; j < TestUserCount; j++ {
			cmts = append(cmts, buildRandomCmt(post.ID, users[j].ID))
		}
	}
	//增加评论
	for _, v := range cmts {
		if ioTool.AddComment(v) != nil {
			t.Error("增加评论失败")
			t.FailNow()
		}
	}
	//查询帖子，
	qPost, err := ioTool.QueryPost(post.ID)
	if err != nil {
		t.Error("查询帖子失败")
		t.FailNow()
	}
	//帖子内容比较
	if !isTwoPostSame(post, qPost) {
		t.Error("帖子内容不一致")
		t.FailNow()
	}
	//删除一部分评论
	deleteCmtCount := TestUserCount * CmtCountOfOneUser / 2
	for i := 0; i < deleteCmtCount; i++ {
		if ioTool.DeleteComment(cmts[i].ID) != nil {
			t.Error("删除部分评论失败")
			t.FailNow()
		}
	}
	//删除帖子
	if ioTool.DeletePost(post.ID) != nil {
		t.Error("删除帖子失败")
		t.FailNow()
	}
	//删除所有用户
	for _, v := range users {
		if ioTool.DeleteUser(v.ID) != nil {
			t.Error("删除用户失败")
			t.FailNow()
		}
	}
	//删除所有主题
	if ioTool.DeleteTheme(tmIns.ID) != nil {
		t.Error("删除主题失败")
		t.FailNow()
	}
}

func buildRandomPost(themeID, userID int64) *forum.PostInDB {
	post := new(forum.PostInDB)
	post.ID = 0
	post.ThemeID = themeID
	post.UserID = userID
	post.Title = buildPostTitle()
	post.State = forum.PostStateNormal
	return post
}

func buildRandomCmt(postID, userID int64) *forum.CommentInDB {
	cmt := new(forum.CommentInDB)
	cmt.ID = 0
	cmt.UserID = userID
	cmt.State = forum.CmtStateNormal
	cmt.PostID = postID
	cmt.Content = buildPostContent()
	cmt.CreatedTime = time.Now().UnixNano()
	cmt.LastEditTime = cmt.CreatedTime
	cmt.EditTimes = rand.Int()%5 + 1
	cmt.PraiseTimes = rand.Int() % 200
	cmt.BelittleTimes = rand.Int() % 1000
	return cmt
}

func buildPostTitle() string {
	return combineUuids(rand.Int()%6 + 1)
}

func buildPostContent() string {
	return combineUuids(rand.Int()%100 + 1)
}

func combineUuids(count int) string {
	var uids = make([]string, 0, count)
	for i := 0; i < count; i++ {
		uids = append(uids, tool.NewUUID())
	}
	return strings.Join(uids, " ")
}

func isTwoPostSame(post1, post2 *forum.PostInDB) bool {
	return post1.ID == post2.ID &&
		post1.ThemeID == post2.ThemeID &&
		post1.UserID == post2.UserID &&
		post1.Title == post2.Title &&
		post1.State == post2.State
}
