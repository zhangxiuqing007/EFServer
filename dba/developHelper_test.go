package dba

import (
	"EFServer/forum"
	"EFServer/tool"
	"math/rand"
	"testing"
	"time"
)

//清空数据库
func Test_ClearCurrentDb(t *testing.T) {
	iotool := SqliteIns{}
	iotool.Open("../ef.db")
	defer iotool.Close()
	if iotool.Clear() != nil {
		t.Error("清空失败")
		t.FailNow()
	}
}

//增加标准主题
func Test_HelpAddStandardThemes(t *testing.T) {
	iotool := SqliteIns{}
	iotool.Open("../ef.db")
	defer iotool.Close()
	iotool.AddTheme("要闻")
	iotool.AddTheme("国内")
	iotool.AddTheme("国际")
	iotool.AddTheme("社会")
	iotool.AddTheme("军事")
	iotool.AddTheme("娱乐")
	iotool.AddTheme("体育")
	iotool.AddTheme("汽车")
	iotool.AddTheme("科技")
}

func initRandomNameData() {
	spe := []rune{' ', '\r', '\n'}
	tool.InitNameWords(
		tool.SplitText(tool.MustStr(tool.ReadAllTextUtf8("../config/中文姓氏.txt")), spe),
		tool.SplitText(tool.MustStr(tool.ReadAllTextUtf8("../config/中文名字.txt")), spe))
}

//增加一些用户，其中包括二把刀
func Test_HelpAddSomeUsers(t *testing.T) {
	initRandomNameData()
	const addCount = 10
	iotool := new(SqliteIns)
	iotool.Open("../ef.db")
	defer iotool.Close()
	user := buildRandomUser()
	user.Name = "二把刀"
	user.Account = "erbadao"
	user.PassWord = "erbadao"
	iotool.AddUser(user)
	count := 0
	for {
		user := buildRandomUser()
		if iotool.IsUserNameExist(user.Name) {
			continue
		}
		if iotool.IsUserAccountExist(user.Account) {
			continue
		}
		if iotool.AddUser(user) != nil {
			t.Error("添加随机用户 失败")
			t.FailNow()
		} else {
			count++
		}
		if count >= addCount {
			break
		}
	}
}

//增加一些帖子和评论
func Test_HelpAddSomePostAndCmts(t *testing.T) {
	//随机种子1
	rand.Seed(time.Now().UnixNano())
	//指定用户，分别在指定主题，发5-10个测试帖子，然后给予几轮评论
	const userCount = 11
	const themeCount = 9
	//评论总数
	const cmtMaxCountOneUser = 150000

	userIDs := [userCount]int64{}
	for i := 0; i < userCount; i++ {
		userIDs[i] = int64(150 + i)
	}
	themeIDs := [themeCount]int64{}
	for i := 0; i < themeCount; i++ {
		themeIDs[i] = int64(1764 + i)
	}
	iotool := new(SqliteIns)
	iotool.Open("../ef.db")
	defer iotool.Close()

	posts := make([]*forum.PostInDB, 0, 200)
	cmts := make([]*forum.CommentInDB, 0, 2000)
	//针对每一个主题
	for _, tm := range themeIDs {
		//确定发帖轮数
		postCount := 30
		//开始发帖
		for p := 0; p < postCount; p++ {
			//每一个用户
			for _, ur := range userIDs {
				//发帖
				randPost := buildRandomPost(tm, ur)
				posts = append(posts, randPost)
				//帖子主体内容（第0条评论）
				randCmt := buildRandomCmt(randPost.ID, ur)
				cmts = append(cmts, randCmt)
			}
		}
	}
	if iotool.AddPosts(posts) != nil {
		t.Error("批量新增帖子失败3")
		t.FailNow()
	}
	if iotool.AddComments(cmts) != nil {
		t.Error("添加评论失败")
		t.FailNow()
	}
	cmts = make([]*forum.CommentInDB, 0, 2000)
	//多轮评论
	for cmti := 0; cmti < cmtMaxCountOneUser; cmti++ {
		cmts = append(cmts, buildRandomCmt(posts[rand.Intn(len(posts))].ID, userIDs[rand.Intn(userCount)]))
		if len(cmts) >= 2000 {
			if iotool.AddComments(cmts) != nil {
				t.Error("添加评论失败")
				t.FailNow()
			}
			cmts = make([]*forum.CommentInDB, 0, 2000)
		}
	}
}
