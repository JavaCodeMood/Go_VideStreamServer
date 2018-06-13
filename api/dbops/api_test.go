package dbops

import(
	//"log"
	"testing"
	"strconv"
	"time"
	"fmt"
)

//init(dblogin, truncate tables)

var tempvid string

func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M){
	clearTables()
	m.Run()
	clearTables();
}

// 测试用户例

func TestUserWorkFlow(t *testing.T){
	//log.Printf("%v",t)
	t.Run("Add",testAddUser)
	t.Run("Get",testGetUser)
	t.Run("Del",testDeleteUser)
	t.Run("Reget",testRegetUser)
}

func testAddUser(t *testing.T){
	err := AddUserCredential("jiang","123")
	if err != nil{
		t.Errorf("Error of AddUser: %v",err)
	}
}

func testGetUser(t *testing.T){
	pwd,err := GetUserCredential("jiang")
	if pwd != "123" || err != nil{
		t.Errorf("Error of getUser")
	}
}

func testDeleteUser(t *testing.T){
	err := DeleteUser("jiang","123")
	if err != nil{
		t.Errorf("Error of delUser: %v",err)
	}
}

func testRegetUser(t *testing.T){
	pwd ,err := GetUserCredential("jiang")
	if err != nil{
		t.Errorf("Error of regetUser: %v",err)
	}

	if pwd != ""{
		t.Errorf("deleting user test failed!")
	}
}

// 测试视频增删改查
func TestVideoWorkFlow(t *testing.T){
	//log.Printf("%v",t)
	t.Run("PrepareUser",testAddUser)
	t.Run("AddVideo",testAddVideoInfo)
	t.Run("GetVideo",testGetVideoInfo)
	t.Run("DelVideo",testDeleteVideoInfo)
}

func testAddVideoInfo(t *testing.T){
	vi , err := AddNewVideo(1,"my-video")
	if err != nil{
		t.Errorf("Error of AddViewInfo :v" , err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T){
	_ , err := GetVideoInfo(tempvid)
	if err != nil{
		t.Errorf("Error of GetVideoInfo : %v" , err)
	}
}

func testDeleteVideoInfo(t *testing.T){
	err := DeleteVideoInfo(tempvid)
	if err != nil{
		t.Errorf("Error of DeleteVideoInfo : %v", err)
	}
}


// 测试评论增删改查

func TestCommentWorkFlow(t *testing.T){
	clearTables()
	t.Run("PrepareUser",testAddUser)
	t.Run("AddComment",testAddComment)
	t.Run("ListComment",testListComment)
	clearTables()
}

func testAddComment(t *testing.T){
	vid := "12345"
	aid :=1
	content := " i like it"
	err := AddNewComment(vid , aid, content)
	if err != nil{
		t.Errorf("Error of AddComment: %v",err)
	}
}

func testListComment(t *testing.T){
	vid := "12345"
	from := 1514764800
	to , _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000,10))
	res , err := ListComments(vid,from,to)
	if err != nil{
		t.Errorf("Error of ListComment: %v",err)
	}
	for i , ele := range res{
		fmt.Printf("Comment: %d , %v, \n",i ,ele)
	}
}

