package dba

import "testing"
import "math/rand"
import "strings"
import "EFServer/forum"
import "EFServer/tool"
import "time"

func Test_UserOperations(t *testing.T) {
	const testCount = 5
	ioTool := &SqliteIns{}
	ioTool.Open("../ef.db")
	defer ioTool.Close()
	//创建testCount个用户
	testUsers := make([]*forum.User, 0, testCount)
	for i := 0; i < testCount; i++ {
		testUsers = append(testUsers, buildRandomUser())
	}
	//增
	for _, u := range testUsers {
		ioTool.AddUser(u)
	}
	//改
	for _, u := range testUsers {
		value := rand.Float64()
		if value < 0.333 {
			u.Name = tool.NewUUID()
		} else if value < 0.666 {
			u.PassWord = tool.NewUUID()
		} else {
			if rand.Float64() > 0.5 {
				u.UserType = forum.UserTypeAdministrator
				u.UserState = forum.UserStateNormal
			} else {
				u.UserType = forum.UserTypeNormalUser
				u.UserState = forum.UserStateLock
			}
		}
		ioTool.UpdateUser(u)
	}
	//查、校验
	for _, u := range testUsers {
		if rand.Float64() > 0.5 {
			tempUser, err := ioTool.QueryUserByID(u.ID)
			if err != nil {
				t.FailNow()
			}
			if !isTwoUserSame(u, tempUser) {
				t.FailNow()
			}
		} else {
			tempUser, err := ioTool.QueryUserByAccountAndPwd(u.Account, u.PassWord)
			if err != nil {
				t.FailNow()
			}
			if !isTwoUserSame(u, tempUser) {
				t.FailNow()
			}
		}
	}
	//删
	for _, u := range testUsers {
		ioTool.DeleteUser(u.ID)
	}
}

func buildRandomUser() *forum.User {
	randUser := new(forum.User)
	randUser.Name = tool.NewUUID()
	randUser.Account = tool.NewUUID()
	randUser.PassWord = tool.NewUUID()
	randUser.SignUpTime = time.Now().UnixNano()
	randUser.UserState = forum.UserStateNormal
	randUser.UserType = forum.UserTypeNormalUser
	return randUser
}

func isTwoUserSame(user1, user2 *forum.User) bool {
	return user1.ID == user2.ID &&
		user1.Name == user2.Name &&
		user1.Account == user2.Account &&
		user1.PassWord == user2.PassWord &&
		user1.SignUpTime == user2.SignUpTime &&
		user1.UserState == user2.UserState &&
		user1.UserType == user2.UserType
}

func Test_ThemeOperations(t *testing.T) {
	const testCount = 5
	ioTool := &SqliteIns{}
	ioTool.Open("../ef.db")
	defer ioTool.Close()
	//增
	for i := 0; i < testCount; i++ {
		tmName := buildRandomThemeName()
		if ioTool.AddTheme(tmName) != nil {
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

func buildRandomTheme() *forum.Theme {
	return &forum.Theme{Name: buildRandomThemeName()}
}

func buildRandomThemeName() string {
	uuidStr := tool.NewUUID()
	rs := []rune(uuidStr)
	return string(rs[0 : rand.Int()%32+5])
}

func Test_PostAndCmt(t *testing.T) {
	const TestUserCount = 11
	const CmtCountOfOneUser = 10
	ioTool := new(SqliteIns)
	ioTool.Open("../ef.db")
	defer ioTool.Close()
	//创建主题
	themeName := buildRandomThemeName()
	if ioTool.AddTheme(themeName) != nil {
		t.Error("新增主题失败")
		t.FailNow()
	}
	//查询主题
	tmIns, err := ioTool.QueryTheme(themeName)
	if err != nil {
		t.Error("主题查询失败")
		t.FailNow()
	}
	//创建用户 n 个
	users := make([]*forum.User, 0, TestUserCount)
	for i := 0; i < TestUserCount; i++ {
		users = append(users, buildRandomUser())
	}
	//新增用户
	for _, v := range users {
		if ioTool.AddUser(v) != nil {
			t.Error("新增用户失败")
			t.FailNow()
		}
	}
	//由第一个用户 创建一个帖子
	post := buildRandomPost(tmIns.ID, users[0].ID)
	//增加帖子
	if ioTool.AddPost(post) != nil {
		t.Error("增加帖子失败")
		t.FailNow()
	}
	//修改帖子
	post.Title = buildPostTitle()
	post.Content = buildPostContent()
	post.State = forum.PostStateNormal
	post.LastEditTime = time.Now().UnixNano()
	post.EditTimes = 1
	post.PraiseTimes = 100
	post.BelittleTimes = 1000
	if ioTool.UpdatePost(post) != nil {
		t.Error("修改帖子失败")
		t.FailNow()
	}
	//制造评论
	cmts := make([]*forum.Comment, 0, 30)
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
	//更新评论
	for _, v := range cmts {
		v.Content = buildPostContent()
		v.State = forum.CmtStateNormal
		v.LastEditTime = time.Now().UnixNano()
		v.EditTimes = 1
		v.PraiseTimes = 100
		v.BelittleTimes = 1000
		if ioTool.UpdateComment(v) != nil {
			t.Error("更新评论失败")
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
	//评论内容比较
	for i, v := range cmts {
		if !isTwoCmtSame(qPost.Comments[i], v) {
			t.Error("评论内容不一致")
			t.FailNow()
		}
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

func buildRandomPost(themeID, userID int64) *forum.Post {
	post := new(forum.Post)
	post.ID = 0
	post.UserID = userID
	post.ThemeID = themeID
	post.State = forum.PostStateNormal
	post.Title = buildPostTitle()
	post.Content = buildPostContent()
	post.CreatedTime = time.Now().UnixNano()
	post.LastEditTime = post.CreatedTime
	post.EditTimes = rand.Int()%5 + 1
	post.PraiseTimes = rand.Int() % 200
	post.BelittleTimes = rand.Int() % 1000
	return post
}

func buildRandomCmt(postID, userID int64) *forum.Comment {
	cmt := new(forum.Comment)
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

func isTwoPostSame(post1, post2 *forum.Post) bool {
	return post1.ID == post2.ID &&
		post1.UserID == post2.UserID &&
		post1.ThemeID == post2.ThemeID &&
		post1.State == post2.State &&
		post1.Title == post2.Title &&
		post1.Content == post2.Content &&
		post1.CreatedTime == post2.CreatedTime &&
		post1.LastEditTime == post2.LastEditTime &&
		post1.EditTimes == post2.EditTimes &&
		post1.PraiseTimes == post2.PraiseTimes &&
		post1.BelittleTimes == post2.BelittleTimes
}

func isTwoCmtSame(cmt1, cmt2 *forum.Comment) bool {
	return cmt1.ID == cmt2.ID &&
		cmt1.UserID == cmt2.UserID &&
		cmt1.State == cmt2.State &&
		cmt1.PostID == cmt2.PostID &&
		cmt1.Content == cmt2.Content &&
		cmt1.CreatedTime == cmt2.CreatedTime &&
		cmt1.LastEditTime == cmt2.LastEditTime &&
		cmt1.EditTimes == cmt2.EditTimes &&
		cmt1.PraiseTimes == cmt2.PraiseTimes &&
		cmt1.BelittleTimes == cmt2.BelittleTimes
}
