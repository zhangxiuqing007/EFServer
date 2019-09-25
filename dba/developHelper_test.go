package dba

import (
	"EFServer/forum"
	"EFServer/tool"
	"math/rand"
	"testing"
)

//清空数据库	go test -v -run Test_ClearCurrentDb
func Test_ClearCurrentDb(t *testing.T) {
	iotool := new(testResourceBuilder).buildCurrentTestSQLIns()
	defer iotool.Close()
	if iotool.Clear() != nil {
		t.Error("清空失败1")
		t.FailNow()
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
	if iotool.AddUser(users[0]) != nil {
		t.Error("x失败：添加测试用户")
		t.FailNow()
	} else {
		t.Log("成功：添加测试用户")
	}
	initRandomNameData()
	for i := 1; i < addCount; i++ {
		user := users[i]
		for {
			user.Name = tool.RandomChineseName()
			if iotool.IsUserNameExist(user.Name) {
				continue
			}
		}
		if iotool.AddUser(user) != nil {
			t.Error("x失败：添加随机用户")
			t.FailNow()
		} else {
			t.Log("成功：添加测试用户")
		}
	}
}

func initRandomNameData() {
	spe := []rune{' ', '\r', '\n'}
	tool.InitNameWords(
		tool.SplitText(tool.MustStr(tool.ReadAllTextUtf8("../config/中文姓氏.txt")), spe),
		tool.SplitText(tool.MustStr(tool.ReadAllTextUtf8("../config/中文名字.txt")), spe))
}

//增加一些帖子和评论	go test -v -run Test_HelpAddSomePostAndCmts
func Test_HelpAddSomePostAndCmts(t *testing.T) {
	const userCount = 11
	const themeCount = 9
	//帖子总数 100W
	const postMaxCount = 1000000
	//评论总数 100W
	const cmtMaxCount = 1000000

	rander := new(testResourceBuilder)
	rander.initRandomSeed()
	iotool := rander.buildCurrentTestSQLIns()
	defer iotool.Close()
	userIDs := [userCount]int64{}
	for i := 0; i < userCount; i++ {
		userIDs[i] = int64(183 + i)
	}
	themeIDs := [themeCount]int64{}
	for i := 0; i < themeCount; i++ {
		themeIDs[i] = int64(1792 + i)
	}
	posts := make([]*forum.PostInDB, 0, postMaxCount)
	for i := 0; i < postMaxCount; i++ {
		posts = append(posts, rander.buildRandomPost(themeIDs[rand.Intn(themeCount)], userIDs[rand.Intn(userCount)]))
	}
	if iotool.AddPosts(posts) != nil {
		t.Error("x失败：插入巨量测试帖子")
		t.FailNow()
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
		t.Error("x失败：插入楼主巨量测试评论")
		t.FailNow()
	} else {
		t.Log("成功：插入楼主巨量测试评论")
	}
	//前500个帖子追加评论
	posts = posts[0:501]
	cmts = cmts[0:0]
	//多轮评论
	for cmti := 0; cmti < cmtMaxCount; cmti++ {
		cmts = append(cmts, rander.buildRandomCmt(posts[rand.Intn(len(posts))].ID, userIDs[rand.Intn(userCount)]))
		if cmti == cmtMaxCount-1 || len(cmts) >= 50000 {
			if iotool.AddComments(cmts) != nil {
				t.Error("x失败：插入数万条随机评论")
				t.FailNow()
			} else {
				t.Log("成功：插入数万条随机评论")
			}
			cmts = cmts[0:0]
		}
	}
}
