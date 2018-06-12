package dbops

import(
	//"log"
	"testing"
)

//init(dblogin, truncate tables)


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
