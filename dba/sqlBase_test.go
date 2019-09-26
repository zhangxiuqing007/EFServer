package dba

import (
	"EFServer/forum"
	"fmt"
	"math/rand"
	"testing"
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
	//创建随机个用户
	const testCount = 5
	users := rander.buildRandomUsers(testCount)
	//新增用户
	for _, v := range users {
		if sqlIns.AddUser(v) != nil {
			t.Error("x失败：新增用户" + v.Name)
			t.FailNow()
		}
	}
	t.Log("成功：新增用户")
	//通过id查询用户
	for _, v := range users {
		quser, err := sqlIns.QueryUserByID(v.ID)
		if err != nil || !rander.isTwoUserSame(v, quser) {
			t.Error("x失败：通过id查询用户")
			t.FailNow()
		}
	}
	t.Log("成功：通过id查询用户")
	//通过账户密码查询用户
	for _, v := range users {
		quser, err := sqlIns.QueryUserByAccountAndPwd(v.Account, v.PassWord)
		if err != nil || !rander.isTwoUserSame(v, quser) {
			t.Error("x失败：通过账户密码查询用户")
			t.FailNow()
		}
	}
	t.Log("成功：通过账户密码查询用户")
	//查询用户的统计信息
	/* 需要其他测试函数完成，在这里暂时没有条件测试 */
	//查询用户名是否存在
	//查询用户账号是否存在
	for _, v := range users {
		if !sqlIns.IsUserNameExist(v.Name) || !sqlIns.IsUserAccountExist(v.Account) {
			t.Error("x失败：查询是否用户名或账号已存在")
		}
	}
	t.Log("成功：查询是否用户名或账号已存在")
	//删除用户
	for _, v := range users {
		if sqlIns.DeleteUser(v.ID) != nil {
			t.Error("x失败：删除用户" + v.Name)
			t.FailNow()
		}
	}
	t.Log("成功：删除用户")
}

//帖子和评论增删改查操作	go test -v -run Test_PostAndCmt
func Test_PostAndCmt(t *testing.T) {
	rander := new(testResourceBuilder)
	rander.initRandomSeed()
	sqlIns := rander.buildCurrentTestSQLIns()
	defer sqlIns.Close()
	const testUserCount = 5
	const cmtCount = 50
	//创建临时主题
	tmIns := rander.buildRandomTheme(1)[0]
	sqlIns.AddTheme(tmIns)
	//创建临时用户
	users := rander.buildRandomUsers(testUserCount)
	for _, v := range users {
		sqlIns.AddUser(v)
	}
	//每个用户创建2个帖子以及主内容
	posts := make([]*forum.PostInDB, 0, testUserCount*2)
	for i := 0; i < testUserCount; i++ {
		posts = append(posts, rander.buildRandomPost(tmIns.ID, users[i].ID))
		posts = append(posts, rander.buildRandomPost(tmIns.ID, users[i].ID))
	}
	//逐个新增帖子
	for i := 0; i < testUserCount; i++ {
		if sqlIns.AddPost(posts[i]) != nil {
			t.Error("x失败：新增帖子")
			t.FailNow()
		}
	}
	//批量新增帖子
	if sqlIns.AddPosts(posts[testUserCount:]) != nil {
		t.Error("x失败：批量新增帖子")
		t.FailNow()
	}
	//生成对应的帖子主内容评论
	cmts := make([]*forum.CommentInDB, 0, testUserCount*2)
	for _, v := range posts {
		cmts = append(cmts, rander.buildRandomCmt(v.ID, v.UserID))
	}
	//逐个新增评论（其实是帖子的主内容）
	for _, v := range cmts {
		if sqlIns.AddComment(v) != nil {
			t.Error("x失败：新增帖子主内容")
			t.FailNow()
		}
	}
	//追加一定数量的评论
	cmts = make([]*forum.CommentInDB, 0, cmtCount)
	for i := 0; i < cmtCount; i++ {
		postID := posts[rand.Intn(len(posts))].ID
		userID := users[rand.Intn(len(users))].ID
		cmts = append(cmts, rander.buildRandomCmt(postID, userID))
	}
	//批量增加评论
	if sqlIns.AddComments(cmts) != nil {
		t.Error("x失败：批量增加评论")
		t.FailNow()
	}
	//查询用户的统计信息，属于用户类操作
	for _, v := range users {
		if _, err := sqlIns.QueryUserSaInfoByID(v.ID); err != nil {
			t.Error("x失败：统计用户信息")
			t.FailNow()
		}
	}
	//查询帖子
	for _, v := range posts {
		p, err := sqlIns.QueryPost(v.ID)
		if err != nil || !rander.isTwoPostSame(p, v) {
			t.Error("x失败：查询帖子失败或内容不一致")
			t.FailNow()
		}
	}
	//查询主题帖子总数量
	if count, err := sqlIns.QueryPostCountOfTheme(tmIns.ID); err != nil || count != testUserCount*2 {
		t.Error("x失败：查询主题帖子总量")
		t.FailNow()
	}
	//查询用户发帖总数量
	for _, v := range users {
		if count, err := sqlIns.QueryPostCountOfUser(v.ID); err != nil || count != 2 {
			t.Error("x失败：查询用户发帖总量")
			t.FailNow()
		}
	}
	//查询主题下的帖子列表
	if ps, err := sqlIns.QueryPostsOfTheme(tmIns.ID, testUserCount, testUserCount, 0); err != nil || len(ps) != testUserCount {
		t.Error("x失败：查询主题下的帖子列表，按发帖顺序排序")
		t.FailNow()
	}
	if ps, err := sqlIns.QueryPostsOfTheme(tmIns.ID, testUserCount, testUserCount, 1); err != nil || len(ps) != testUserCount {
		t.Error("x失败：查询主题下的帖子列表，按最后评论顺序排序")
		t.FailNow()
	}
	//查询用户发的帖子列表
	for _, v := range users {
		if ps, err := sqlIns.QueryPostsOfUser(v.ID, 1, 1); err != nil || len(ps) != 1 {
			t.Error("x失败：查询用户发的帖子列表")
			t.FailNow()
		}
	}
	for _, v := range posts {
		//查询帖子页内，帖子的展示内容
		if p, err := sqlIns.QueryPostOfPostPage(v.ID); err != nil || p.Title != v.Title || p.ThemeName != tmIns.Name {
			t.Error("x失败：查询帖子页内，帖子的展示内容")
			t.FailNow()
		}
		//查询DB评论
		cs, err := sqlIns.QueryComments(v.ID)
		if err != nil || cs[0].PostID != v.ID || cs[0].UserID != v.UserID {
			t.Error("x失败：查询DB评论")
			t.FailNow()
		}
		//统计帖子的评论数量
		count, err := sqlIns.QueryCommentsCountOfPost(v.ID)
		if err != nil || count != len(cs) {
			t.Error("x失败：统计帖子的评论数量")
			t.FailNow()
		}
		//查询帖子页内，评论的展示内容
		scs, err := sqlIns.QueryCommentsOfPostPage(v.ID, 10000, 0)
		if err != nil || len(scs) != count || scs[0].CmterID != v.UserID {
			t.Error("x失败：查询帖子页内，评论的展示内容")
			t.FailNow()
		}
	}
	//删除一半评论
	for i := 0; i < len(cmts)/2; i++ {
		if sqlIns.DeleteComment(cmts[i].ID) != nil {
			t.Error("x失败：删除评论")
			t.FailNow()
		}
	}
	//删除帖子（连同其剩余的评论）
	for _, v := range posts {
		if sqlIns.DeletePost(v.ID) != nil {
			t.Error("x失败：删除帖子（连同其剩余的评论）")
			t.FailNow()
		}
	}
	//删除用户
	for _, v := range users {
		sqlIns.DeleteUser(v.ID)
	}
	//删除主题
	sqlIns.DeleteTheme(tmIns.ID)
	t.Log("成功：帖子和评论增删改查操作")
}
