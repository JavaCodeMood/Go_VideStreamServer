package dbops

import(
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/myproject/video_server/defs"
)


func AddUserCredential(loginName string , pwd string) error{
	//log.Printf("AddUserCredential %v",dbConn)
	stmtIns,err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil{
		return err
	}
	_, err = stmtIns.Exec(loginName,pwd)
	if(err != nil){
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)(string ,error){
	stmtOut,err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil{
		log.Printf("%s",err)
		return "",err
	}

	var pwd string 
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "",err
	}
	defer stmtOut.Close()
	return pwd,nil
}

func DeleteUser(loginName string,pwd string) error{
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? and pwd = ?")
	if err != nil{
		log.Printf("%s",err)
		return err
	}
	_,err = stmtDel.Exec(loginName,pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}



func AddNewVideo(aid int , name string)(*defs.VideoInfo,error){
	

}