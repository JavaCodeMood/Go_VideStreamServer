package dbops

import(
	"time"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/myproject/video_server/api/defs"
	"github.com/myproject/video_server/api/utils"
)

// 用户数据操作
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


//  视频数据操作
func AddNewVideo(aid int , name string)(*defs.VideoInfo,error){
		
	vid, err := utils.NewUUID()
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")// 格式化时间，必须这么写
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info
		(id,author_id,name,display_ctime) VALUES(?,?,?,?)`)
	if err != nil{
		log.Printf("%s",err)
		return nil,err
	}
	_, err = stmtIns.Exec(vid,aid,name,ctime)
	if err != nil{
		return nil ,err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime:ctime}
	defer stmtIns.Close()
	return res,nil
}

func GetVideoInfo(vid string)(*defs.VideoInfo,error){
	stmtOut, err := dbConn.Prepare("SELECT author_id,name,display_ctime FROM video_info WHERE id=?")
	if err != nil{
		return nil,err
	}
	var aid int
	var dct string
	var name string
	err = stmtOut.QueryRow(vid).Scan(&aid,&name,&dct)
	if err != nil{
		return nil , err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime:dct}
	defer stmtOut.Close()
	return res , nil
}

func DeleteVideoInfo(vid string) error{
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil{
		return err 
	}
	_, err = stmtDel.Exec(vid);
	if err != nil{
		return err 
	}
	defer stmtDel.Close()
	return nil
}