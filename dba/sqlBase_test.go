package dba

import (
	"EFServer/forum"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//测试主题表相关操作	go test -v -run TestThemeTableOperations
func TestThemeTableOperations(t *testing.T) {
	rander := new(testResourceBuilder)
	rander.initRandomSeed()
	sqlIns := rander.buildCurrentTestSQLIns()
	defer sqlIns.Close()
	const testCount = 5
	//逐个新增主题
	tms := rander.buildRandomTheme(testCount)
	for i := 0; i < testCount; i++ {
		err := sqlIns.AddTheme(tms[i])
		if err != nil {
			t.Error("x失败：新增主题")
			t.FailNow()
		}
		tms = append(tms, tms[i])
		t.Logf("成功：新增主题，ID:%d,Name:%s", tms[i].ID, tms[i].Name)
	}
	//逐个更新主题名称
	for i, v := range tms {
		v.Name = fmt.Sprintf("主题改名%d", i)
		if sqlIns.UpdateTheme(v) != nil {
			t.Error("x失败：修改主题名称")
			t.FailNow()
		} else {
			t.Log("成功：修改主题名")
		}
	}
	//逐个查询主题并对比信息
	for _, v := range tms {
		qtm, err := sqlIns.QueryTheme(v.ID)
		if err != nil || qtm.ID != v.ID || qtm.Name != v.Name {
			t.Error("x失败：查询主题")
			t.FailNow()
		} else {
			t.Log("成功：查询主题，一致")
		}
	}
	//查询所有主题
	_, err := sqlIns.QueryAllThemes()
	if err != nil {
		t.Error("x失败：查询所有主题失败")
		t.FailNow()
	} else {
		t.Log("成功：查询所有主题")
	}
	//删除刚才新增的主题
	for _, v := range tms {
		if sqlIns.DeleteTheme(v.ID) != nil {
			t.Error("x失败：删除主题")
			t.FailNow()
		} else {
			t.Logf("成功：删除主题，ID:%d,Name:%s", v.ID, v.Name)
		}
	}
}

//测试用户表相关操作	go test -v -run TestUserTableOperations
func TestUserTableOperations(t *testing.T) {
	rander := new(testResourceBuilder)
	rander.initRandomSeed()
	sqlIns := rander.buildCurrentTestSQLIns()
	defer sqlIns.Close()
	//创建随机用户
	const testCount = 10
	users := rander.buildRandomUsers(testCount)
	//新增用户
	for _, v := range users {
		if sqlIns.AddUser(v) != nil {
			t.Error("x失败：新增用户" + v.Name)
			t.FailNow()
		} else {
			t.Log("成功：新增用户" + v.Name)
		}
	}
	//通过id查询用户
	//通过账号密码查询用户
	//查询用户名是否存在
	//查询账号是否存在
	//查询用户的统计信息，需要其他模块配合，暂时没思路进行单元测试
	//删除用户
	for _, v := range users {
		if sqlIns.DeleteUser(v.ID) != nil {
			t.Error("x失败：删除用户" + v.Name)
			t.FailNow()
		} else {
			t.Log("成功：删除用户" + v.Name)
		}
	}
}

//用户增删查操作	go test -v -run Test_UserOperations
func Test_UserOperations(t *testing.T) {
	initRandomNameData()
	const testCount = 200

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

//主题增删改查操作	go test -v -run Test_ThemeOperations
func Test_ThemeOperations(t *testing.T) {
	initRandomNameData()
	const testCount = 5
	ioTool := &SqliteIns{}
	ioTool.Open("../ef.db")
	defer ioTool.Close()
	//增
	for i := 0; i < testCount; i++ {
		if err := ioTool.AddTheme(buildRandomTheme()); err != nil {
			t.Error("增加")
			t.FailNow()
		}
	}
	//查
	tms, err := ioTool.QueryAllThemes()
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

//帖子和评论增删改查操作	go test -v -run Test_PostAndCmt
func Test_PostAndCmt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	initRandomNameData()
	const TestUserCount = 11
	const CmtCountOfOneUser = 10
	ioTool := new(SqliteIns)
	ioTool.Open("../ef.db")
	defer ioTool.Close()
	//创建主题
	tmIns := buildRandomTheme()
	err := ioTool.AddTheme(tmIns)
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
