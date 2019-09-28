package dba

import (
	"EFServer/forum"
	"math/rand"
	"testing"
)

//目前mysql不能通过sql清空
// SET foreign_key_checks = 0;
// truncate cmt;
// truncate post;
// truncate user;
// truncate theme;
// SET foreign_key_checks = 1;

//清空数据库	go test -v -run Test_ClearCurrentDb
func Test_ClearCurrentDb(t *testing.T) {
	iotool := new(testResourceBuilder).buildCurrentTestSQLIns()
	defer iotool.Close()
	err := iotool.Clear()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

//增加标准主题		go test -v -run Test_HelpAddStandardThemes
func Test_HelpAddStandardThemes(t *testing.T) {
	iotool := new(testResourceBuilder).buildCurrentTestSQLIns()
	defer iotool.Close()
	iotool.AddTheme(&forum.ThemeInDB{ID: 0, Name: "要闻"})
	iotool.AddTheme(&forum.ThemeInDB{ID: 0, Name: "国内"})
	iotool.AddTheme(&forum.ThemeInDB{ID: 0, Name: "国际"})
	iotool.AddTheme(&forum.ThemeInDB{ID: 0, Name: "社会"})
	iotool.AddTheme(&forum.ThemeInDB{ID: 0, Name: "军事"})
	iotool.AddTheme(&forum.ThemeInDB{ID: 0, Name: "娱乐"})
	iotool.AddTheme(&forum.ThemeInDB{ID: 0, Name: "体育"})
	iotool.AddTheme(&forum.ThemeInDB{ID: 0, Name: "汽车"})
	iotool.AddTheme(&forum.ThemeInDB{ID: 0, Name: "科技"})
}

//增加一些用户，其中包括二把刀	go test -v -run Test_HelpAddSomeUsers
func Test_HelpAddSomeUsers(t *testing.T) {
	const addCount = 11
	rander := new(testResourceBuilder)
	rander.initRandomSeed()
	iotool := rander.buildCurrentTestSQLIns()
	defer iotool.Close()
	users := rander.buildRandomUsers(addCount)
	users[0].Name = "二把刀"
	users[0].Account = "erbadao"
	users[0].PassWord = "erbadao"
	if err := iotool.AddUser(users[0]); err != nil {
		t.Fatalf("x失败：添加测试用户：" + err.Error())
	} else {
		t.Log("成功：添加测试用户")
	}
	for i := 1; i < addCount; i++ {
		user := users[i]
		for {
			user.Name = rander.buildRandomChineseName()
			if iotool.IsUserNameExist(user.Name) {
				continue
			}
			break
		}
		if iotool.AddUser(user) != nil {
			t.Fatalf("x失败：添加随机用户")
		} else {
			t.Log("成功：添加测试用户")
		}
	}
}

//增加一些帖子和评论	go test -v -run Test_HelpAddSomePostAndCmts
func Test_HelpAddSomePostAndCmts(t *testing.T) {
	const userCount = 11
	const themeCount = 9
	//帖子总数 100W
	const postMaxCount = 10000
	//评论总数 100W
	const cmtMaxCount = 200000

	rander := new(testResourceBuilder)
	rander.initRandomSeed()
	iotool := rander.buildCurrentTestSQLIns()
	defer iotool.Close()
	userIDs := [userCount]int64{}
	for i := 0; i < userCount; i++ {
		userIDs[i] = int64(1 + i)
	}
	themeIDs := [themeCount]int64{}
	for i := 0; i < themeCount; i++ {
		themeIDs[i] = int64(1 + i)
	}
	posts := make([]*forum.PostInDB, 0, postMaxCount)
	for i := 0; i < postMaxCount; i++ {
		posts = append(posts, rander.buildRandomPost(themeIDs[rand.Intn(themeCount)], userIDs[rand.Intn(userCount)]))
	}
	if err := iotool.AddPosts(posts); err != nil {
		t.Fatalf("x失败：插入巨量测试帖子，" + err.Error())
	} else {
		t.Log("成功：插入巨量测试帖子")
	}

	cmts := make([]*forum.CommentInDB, 0, cmtMaxCount)
	for _, v := range posts {
		//帖子主体内容（第0条评论）
		randCmt := rander.buildRandomCmt(v.ID, v.UserID)
		cmts = append(cmts, randCmt)
	}
	if iotool.AddComments(cmts) != nil {
		t.Fatalf("x失败：插入楼主巨量测试评论")
	} else {
		t.Log("成功：插入楼主巨量测试评论")
	}
	//最后500个帖子追加评论
	posts = posts[len(posts)-500:]
	cmts = cmts[0:0]
	//多轮评论
	for cmti := 0; cmti < cmtMaxCount; cmti++ {
		cmts = append(cmts, rander.buildRandomCmt(posts[rand.Intn(len(posts))].ID, userIDs[rand.Intn(userCount)]))
		if cmti == cmtMaxCount-1 || len(cmts) >= 50000 {
			if iotool.AddComments(cmts) != nil {
				t.Fatalf("x失败：插入数万条随机评论")
			} else {
				t.Log("成功：插入数万条随机评论")
			}
			cmts = cmts[0:0]
		}
	}
}
