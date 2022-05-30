package dbfunction

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"

	//	"go/DataAccessLayer/db"
	"go/conf/utils"
	"go/pkg/mylog"
	"go/pkg/structpkg/data"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var log = mylog.NewLog("Info")

var (
	Db  *sql.DB
	err error
)

func init() {
	//var c utils.Conf
	//t := c.GetConf()
	tempsql := utils.GetMysql()
//	tempsql := "root:" + utils.GetMysqlpwd() + "@tcp(" + utils.GetMsqlhost() + ")/" + utils.GetMysqldbname()
	//tempsql := "root:" + t.Mysqlpwd  + "@tcp(" + t.Mysqlhost + ")/" + t.Mysqldbname
	// Db, err = sql.Open("mysql", "root:666666@tcp(localhost:3306)/data")
	Db, err = sql.Open("mysql", tempsql)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func SelectUserAccount(email string) string { //传入账号查询邮箱
	//写sql语句
	sqlStr := "select account from auser where email = ?"
	//执行
	row := Db.QueryRow(sqlStr, email)
	user := &data.Userregister{}
	row.Scan(&user.Account)
	if user.Account == "" {
		log.Info("找不到用户！")
		return ""
	} else {
		log.Info("用户存在")
		return user.Account
	}
}

//SaveTemperUser 向数据库中插入未审核的用户信息
func SaveTempUser(account string, password string, email string, name string, sex string) error {
	//写sql语句
	sqlStr := "insert into Auser(account,password,email,state,name,sex,time) values(?,?,?,?,?,?,?)"
	//执行
	temppassword := MD5(password)

	//timeObj := time.Now().Format("2006-01-02 15:04:05")
	var cstSH, _ = time.LoadLocation("Asia/Shanghai")
	timeObj := time.Now().In(cstSH).Format("2006-01-02 15:04:05")

	//	fmt.Println(tempdata)
	_, err := Db.Exec(sqlStr, account, temppassword, email, 0, name, sex, timeObj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func CheckUserAccount(account string) (bool, error) {
	//写sql语句
	sqlStr := "select account,password,email from tempuser where account = ?"
	//执行
	row := Db.QueryRow(sqlStr, account)
	user := &data.Userregister{}
	row.Scan(&user.Account, &user.Password, &user.Email)
	if user.Account == "" {
		log.Info("找不到用户！")
		return false, nil
	} else {
		log.Info("用户存在")
		return true, nil
	}
}

// func Checkemail(){//查询邮箱是否存在

// }
//-------------------------------------------------------------------------------
func CheckAuserEmail(email string) (bool, int) { //检查用户邮箱状态
	//写sql语句
	sqlStr := "select account, state from Auser where email = ?"
	//执行
	row := Db.QueryRow(sqlStr, email)
	user := &data.Userregister{}
	row.Scan(&user.Account, &user.State)
	//fmt.Println(user)
	if user.Account == "" {
		log.Info("找不到邮箱！")
		return false, -1
	} else {
		log.Info("邮箱存在")
		return true, user.State
	}
}

func CheckAuserAccount(account string) (bool, int) { //检查用户邮箱状态
	//写sql语句
	sqlStr := "select account, state from Auser where account = ?"
	//执行
	row := Db.QueryRow(sqlStr, account)
	user := &data.Userregister{}
	row.Scan(&user.Account, &user.State)
	if user.Account == "" {
		log.Info("找不到用户！")
		return false, user.State
	} else {
		log.Info("用户存在")
		return true, user.State
	}
}

//-------------------------------------------------------------------------------

func CheckUserEmail(email string) bool {
	//写sql语句
	sqlStr := "select account,password,email from Auser where email = ? and state = 1 "
	//执行
	row := Db.QueryRow(sqlStr, email)
	user := &data.Userregister{}
	row.Scan(&user.Account, &user.Password, &user.Email)
	if user.Account == "" {
		log.Info("找不到邮箱！")
		return false
	} else {
		log.Info("邮箱存在")
		return true
	}
}

func CheckAccountAndPassword(account string, password string) (bool, int) {
	//sql语句
	sqlStr := "select account,password,email,state from auser where account = ? and password = ?"
	//执行
	temppassword := MD5(password)
	row := Db.QueryRow(sqlStr, account, temppassword)

	//fmt.Println(account, temppassword)
	user := &data.Userregister{}
	row.Scan(&user.Account, &user.Password, &user.Email, &user.State)
	if user.Account == "" || user.State == 0 || user.State == 2 || user.State == 50 {
		log.Info("找不到用户！")
		return false, user.State
	} else {
		return true, user.State
	}
}

func CheckEmailAndPassword(email string, password string) (bool, int) {
	//sql语句
	sqlStr := "select account,password,email,state from auser where email = ? and password = ?"
	//执行
	temppassword := MD5(password)
	row := Db.QueryRow(sqlStr, email, temppassword)
	user := &data.Userregister{}
	//fmt.Println(email, temppassword)
	row.Scan(&user.Account, &user.Password, &user.Email, &user.State)
	if user.Account == "" || user.State == 0 || user.State == 2 || user.State == 50 {
		log.Info("找不到用户！")
		return false, user.State
	} else {
		return true, user.State
	}
}

func ModifyPassword(email string, password string) bool {
	//CheckEmailAndPassword(email,password)
	temppassword := MD5(password)
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}

	//写sql语句
	sql := "update auser set password = ? where email = ? "
	//执行
	Db.Exec(sql, temppassword, email)
	return true
}

func AddLogoutAccount(account string, email string, reason string) error {
	//写sql语句
	sqlStr := "update auser set state = 2  where email = ?"
	//执行
	_, err := Db.Exec(sqlStr, email)

	if err != nil {
		log.Error(err.Error())
		return err
	}
	sqlStr = "insert into logoutuser(account, email, reason) values(?,?,?)"
	//执行
	_, err = Db.Exec(sqlStr, account, email, reason)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func CheckUser(account string) *data.Userregister { //显示用户主页信息
	//sql语句
	sqlStr := "select account, email, state, name, sex, time, headphoto from auser where account = ? and state = 1 "
	//执行
	row := Db.QueryRow(sqlStr, account)
	user := &data.Userregister{}
	row.Scan(&user.Account, &user.Email, &user.State, &user.Name, &user.Sex, &user.Time, &user.Headphoto)
	if user.Account == "" {
		log.Error("找不到用户！")
		return nil
	} else {
		return user
	}
}

func ModifyUsername(account string, name string) bool {
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}

	//写sql语句
	sql := "update auser set name = ? where account = ? "
	//执行
	_, err := Db.Exec(sql, name, account)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}
func ModifyUsersex(account string, sex string) bool {
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}

	//写sql语句
	sql := "update auser set sex = ? where account = ? "
	//执行
	_, err := Db.Exec(sql, sex, account)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}
func ModifyUserheadphoto(account string, headphoto string) bool {
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}

	//写sql语句
	sql := "update auser set headphoto = ? where account = ? "
	//执行
	_, err := Db.Exec(sql, []byte(headphoto), account)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}

func AddModifyUseremail(oldemail string, newemail string, reason string) bool {
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}
	// fmt.Println(tempdata)
	// fmt.Println("!@#123")
	//写sql语句
	sql := "insert into modifyemail(oldemail,newemail,reason) values(?,?,?)"
	//执行
	_, err := Db.Exec(sql, oldemail, newemail, reason)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}

func SeekAlldairy(account string) ([]data.Dairy, error) {
	//写sql语句
	sqlStr := "select dairyid, account, time, details, title, classify, performance from dairy where account = ? "
	//执行
	rows, err := Db.Query(sqlStr, account)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	var diarys []data.Dairy
	for rows.Next() {
		var diary data.Dairy
		rows.Scan(&diary.Dairyid, &diary.Account, &diary.Time, &diary.Details, &diary.Title, &diary.Classify, &diary.Performance)
		diarys = append(diarys, diary)
	}
	return diarys, nil

}

func AddUserdiary(account string, details string, title string, classify string, performance string) error {

	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return temperr
	}

	//timeObj := time.Now().Format("2006-01-02 15:04:05")
	var cstSH, _ = time.LoadLocation("Asia/Shanghai")
	timeObj := time.Now().In(cstSH).Format("2006-01-02 15:04:05")

	//写sql语句
	sql := "insert into dairy(account, time, details, title, classify, performance) values(?,?,?,?,?,?)"
	//执行
	_, err := Db.Exec(sql, account, timeObj, details, title, classify, performance)
	if err != nil {
		log.Error(err.Error())
		_ = tx.Rollback()
		return err
	}
	_ = tx.Commit()
	return nil

}

func SeekUserdiaryid(dairyid int) *data.Dairy {
	//写sql语句
	sql := "select account from dairy where dairyid = ?"
	//执行
	row := Db.QueryRow(sql, dairyid)
	diary := &data.Dairy{}
	row.Scan(&diary.Account)
	if diary.Account == "" {
		log.Info("找不到日记")
		return nil
	} else {
		return diary
	}
}

func ModifyUSerdairy(dairyid int, details string, title string, classify string, performance string) bool {
	//tempdata.Dairyid, tempdata.Details, tempdata.Title, tempdata.Classify, tempdata.Performance
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}
	if SeekUserdiaryid(dairyid) == nil {
		return false
	}
	//timeObj := time.Now().Format("2006-01-02 15:04:05")
	var cstSH, _ = time.LoadLocation("Asia/Shanghai")
	timeObj := time.Now().In(cstSH).Format("2006-01-02 15:04:05")
	//写sql语句
	sql := "update dairy set time = ?,  details = ?, title = ?, classify = ?, performance = ? where dairyid = ? "
	//执行
	Db.Exec(sql, timeObj, details, title, classify, performance, dairyid)
	return true
}

func DeleteUserdairy(dairyid int) bool {

	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}

	if SeekUserdiaryid(dairyid) == nil {
		return false
	}
	//写sql语句
	sql := "delete from dairy where dairyid = ? "
	//执行
	Db.Exec(sql, dairyid)

	return true
}

func EmptyUserdairy(account string) bool {

	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}

	//写sql语句
	sql := "delete from dairy where account = ? "
	//执行
	Db.Exec(sql, account)

	return true
}

func CheckAdminEmail(email string) bool {
	//写sql语句
	sqlStr := "select account,email,password from admin where email = ?"
	//执行
	row := Db.QueryRow(sqlStr, email)
	user := &data.Admin{}
	row.Scan(&user.Account, &user.Email, &user.Password)
	if user.Account == "" {
		log.Info("找不到管理员！")
		return false
	} else {
		log.Info("管理员存在")
		return true
	}
}

func SaveAdmin(account string, email string, password string) error {
	//写sql语句
	sqlStr := "insert into auser(account,email,password, state ,time ) values(?,?,?,?,?)"
	//执行
	temppassword := MD5(password)

	//timeObj := time.Now().Format("2006-01-02 15:04:05") //获取当前时间，类型是Go的时间类型Time
	var cstSH, _ = time.LoadLocation("Asia/Shanghai")
	timeObj := time.Now().In(cstSH).Format("2006-01-02 15:04:05")
	//fmt.Println(tempdata.Account)
	_, err := Db.Exec(sqlStr, account, email, temppassword, 50, timeObj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func CheckAdminEmailAndPassword(email string, password string) (bool, error) {
	//sql语句
	sqlStr := "select account,email,password from auser where email = ? and password = ? and state =50"
	//执行
	temppassword := MD5(password)
	row := Db.QueryRow(sqlStr, email, temppassword)
	user := &data.Admin{}
	//fmt.Println(email, temppassword)
	row.Scan(&user.Account, &user.Email, &user.Password)
	if user.Account == "" {
		log.Info("找不到管理员！")
		return false, nil
	} else {
		return true, nil
	}
}

func Showtempuser() ([]data.Userregister, error) {
	//写sql语句
	sqlStr := "select account, email, name, sex, time from auser where state = 0"
	//执行
	rows, err := Db.Query(sqlStr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	var users []data.Userregister
	for rows.Next() {
		var user data.Userregister
		rows.Scan(&user.Account, &user.Email, &user.Name, &user.Sex, &user.Time)
		users = append(users, user)

		//	users = append(users, user)
	}
	return users, nil
}

func Deletetempuser(account string) bool {
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}
	//写sql语句
	sql := "delete from tempuser where account = ?"

	//执行
	_, err := Db.Exec(sql, account)

	if err != nil {
		log.Error(err.Error())
		_ = tx.Rollback()
		return false
	}
	_ = tx.Commit()

	return true
}

//------------------------------------------------

func Deleteauser(account string) bool {
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}
	//写sql语句
	sql := "delete from auser where account = ?"

	//执行
	_, err := Db.Exec(sql, account)

	if err != nil {
		log.Error(err.Error())
		_ = tx.Rollback()
		return false
	}
	_ = tx.Commit()

	return true
}

//------------------------------------------------

func Passuser(account string) bool {

	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}
	//写sql语句
	sql := "update auser set state = 1 where account = ?"

	//sql := "insert into user(account,password,email,state,time) select account,password,email,1,time from tempuser where account = ?"

	//执行
	_, err := Db.Exec(sql, account)
	if err != nil {
		log.Error(err.Error())
		_ = tx.Rollback()
		return false
	}
	_ = tx.Commit()
	//Deletetempuser(account)

	return true
}

func Showlogoutuser() ([]data.Logoutuser, error) {
	//写sql语句
	sqlStr := "select * from logoutuser"
	//执行
	rows, err := Db.Query(sqlStr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	var users []data.Logoutuser
	for rows.Next() {
		var user data.Logoutuser
		rows.Scan(&user.Account, &user.Email, &user.Reason)
		users = append(users, user)
	}
	return users, nil
}

func Passlogoutuser(account string) bool {

	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}
	//写sql语句
	sql1 := "delete from logoutuser where account = ?"

	//执行
	_, err1 := Db.Exec(sql1, account)
	sql2 := "delete from auser where account = ?"

	//执行
	_, err2 := Db.Exec(sql2, account)
	sql3 := "delete from dairy where account = ?"

	//执行
	_, err3 := Db.Exec(sql3, account)
	if err1 != nil || err2 != nil || err3 != nil {
		log.Error(err1.Error())
		log.Error(err2.Error())
		log.Error(err3.Error())
		_ = tx.Rollback()
		return false
	}
	_ = tx.Commit()

	return true
}

func Deletelogoutuser(account string) bool {
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}
	//写sql语句
	sql := "delete from logoutuser where account = ?"

	//执行
	_, err := Db.Exec(sql, account)

	if err != nil {
		log.Error(err.Error())
		_ = tx.Rollback()
		return false
	}
	_ = tx.Commit()

	return true
}

func ModifyUserstate(account string) bool {
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}

	//写sql语句
	sql := "update auser set state = 1 where account = ?"
	//执行
	Db.Exec(sql, account)
	return true
}

//func Showuser()
func Showuser() ([]data.Userregister, error) {
	//写sql语句
	sqlStr := "select account, email, state, name, sex, time, headphoto from auser where state != 50"
	//执行
	rows, err := Db.Query(sqlStr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	var users []data.Userregister
	for rows.Next() {
		var user data.Userregister
		rows.Scan(&user.Account, &user.Email, &user.State, &user.Name, &user.Sex, &user.Time, &user.Headphoto)
		fmt.Println(user)
		users = append(users, user)
	}
	return users, nil
}

func AdminModifyUserEmail(account string, email string) bool {
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}

	//写sql语句
	sql := "update auser set email = ? where account = ? "
	//执行

	_, err := Db.Exec(sql, email, account)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}

func Showmodifyemail() ([]data.Modifyuseremail, error) {
	//写sql语句
	sqlStr := "select * from modifyemail"
	//执行
	rows, err := Db.Query(sqlStr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	var users []data.Modifyuseremail
	for rows.Next() {
		var user data.Modifyuseremail
		rows.Scan(&user.Oldemail, &user.Newemail, &user.Reason)
		fmt.Println(user)
		users = append(users, user)
	}
	return users, nil
}

func PassModifyemail(newemail string, oldemail string) bool {

	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}
	//写sql语句

	sql := "update auser set email = ? where email = ? "
	//执行
	_, err := Db.Exec(sql, newemail, oldemail)
	if err != nil {
		log.Error(err.Error())
		_ = tx.Rollback()
		return false
	}
	_ = tx.Commit()
	DeleteModifyemail(oldemail)

	return true

}

func DeleteModifyemail(oldemail string) bool {
	tx, temperr := Db.Begin()
	if temperr != nil {
		log.Error(temperr.Error())
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Error(temperr.Error())
		return false
	}
	//写sql语句
	sql := "delete from modifyemail where oldemail = ?"

	//执行
	_, err := Db.Exec(sql, oldemail)

	if err != nil {
		log.Error(err.Error())
		_ = tx.Rollback()
		return false
	}
	_ = tx.Commit()

	return true
}
