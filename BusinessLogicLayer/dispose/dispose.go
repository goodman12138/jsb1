package dispose

import (
	"fmt"
	//	"time"

	"go/DataAccessLayer/dbfunction"
	"go/DataAccessLayer/emailcodedispose"
	"go/DataAccessLayer/redisdb"

	"go/pkg/structpkg/data"
)

const (
	SecretKey = "welcome to jishibeng"
)

//var log = mylog.NewLog("Warning")

func DisposeSendemailcode(email string) (bool, int, string) { //发验证码

	t, state := dbfunction.CheckAuserEmail(email)
	if t {
		if state == 0 { //如果邮箱状态为0

			return false, 404, "发送失败，该邮箱正等待审核!"
		} else if state == 1 || state == 50 { //如果邮箱状态为1
	
			return false, 403, "发送失败，该邮箱已被注册!"
		} else if state == 2 {

			return false, 411, "发送失败,用户正在等待注销！"
		}
	}

	emailcode := emailcodedispose.SendEmail(email)

	fmt.Println(emailcode)

	if emailcode == "40" { //验证码发送失败
	
		return false, 400, "发送失败"
	} else { //验证码发送成功

		temp := redisdb.Setkey(email, emailcode, 60) //存入redis 设置60秒过期
		if temp {
			return true, 200, "发送成功"
		} else {
			return false, 400, "发送失败"
		}
	}
}

func DisposeSenduseremailcode(email string) (bool, int, string) { //发验证码

	t, state := dbfunction.CheckAuserEmail(email)
	if t {
		if state == 2 {
		
			return false, 411, "发送失败,用户正在等待注销！"
		}
	} else {

		return false, 410, "发送失败，该邮箱不存在用户!"
	}

	emailcode := emailcodedispose.SendEmail(email)

	fmt.Println(emailcode)

	if emailcode == "40" { //验证码发送失败

		return false, 400, "发送失败"
	} else { //验证码发送成功

		temp := redisdb.Setkey(email, emailcode, 60) //存入redis 设置60秒过期
		if temp {
			return true, 200, "发送成功"
		} else {
			return false, 400, "发送失败"
		}
	}

}

func DisposeRegister(account string, password string, email string, name string, sex string, emailcode string) (bool, int, string) {

	val := redisdb.Getvalue(email)

	if val == "" {

		return false, 406, "注册失败，邮箱验证码过期，请重新发送！"
	}
	if val != string(emailcode) {

		return false, 405, "注册失败，邮箱验证码错误！"
	}

	t, state := dbfunction.CheckAuserAccount(account)
	if t {
		if state == 0 { //如果邮箱状态为0

			return false, 408, "注册失败，该用户名正等待审核！"
		} else if state == 1 || state == 2 || state == 50 { //如果邮箱状态为1

			return false, 407, "注册失败，该用户名已被使用！"
		}
	}

	t, state = dbfunction.CheckAuserEmail(email)

	if t {
		if state == 0 { //如果邮箱状态为0

			return false, 404, "注册失败，该邮箱正等待审核!"
		} else if state == 1 || state == 2 || state == 50 { //如果邮箱状态为1

			return false, 403, "注册失败，该邮箱已被注册!"
		}
	}

	dbfunction.SaveTempUser(account, password, email, name, sex) //包含用户注册信息
	
	return true, 200, "注册成功"

}

// func DisposeGetaccount(account string, email string ) string {
// 	var account1 string
// 	if account == "" {
// 		account1 = dbfunction.SelectUserAccount(email)
// 	} else {
// 		account1 = account
// 	}

// 	return account1
// }

func DisposeLogin(account string, email string, password string) (bool, int, string) {
	temp1, statu1 := dbfunction.CheckAccountAndPassword(account, password)
	tmep2, statu2 := dbfunction.CheckEmailAndPassword(email, password)

	if temp1 || tmep2 {

		return true, 200, "登录成功"
	} else if statu1 == 2 || statu2 == 2 {

		return false, 411, "登录失败,用户正在等待注销！"
	} else {

		return false, 400, "登录失败,账号或密码错误"
	}

}

func DisposeFogetPassword(email string, password string, emailcode string) (bool, int, string) {


	val := redisdb.Getvalue(email)

	if val == "" {

		return false, 406, "修改失败，邮箱验证码过期，请重新发送！"
	}
	if val != string(emailcode) {

		return false, 405, "修改失败，邮箱验证码错误！"
	}
	if !dbfunction.CheckUserEmail(email) {

		return false, 409, "修改失败，该邮箱不存在！"
	}

	dbfunction.ModifyPassword(email, password)


	return true, 200, "修改成功"

}

func DisposeRedactUserpassword(email string, emailcode string, password string) (bool, int, string) {

	val := redisdb.Getvalue(email)
	if val == "" {
		return false, 406, "修改失败，邮箱验证码过期，请重新发送！"
	}
	if val != string(emailcode) {
		return false, 405, "修改失败，邮箱验证码错误！"
	}
	dbfunction.ModifyPassword(email, password)

	return true, 200, "修改成功"

}

func DisposeLogoutAccount(account string, email string, reason string, emailcode string) (bool, int, string) {

	val := redisdb.Getvalue(email)
	// 判断查询是否出错

	if val == "" {
	
		return false, 406, "注销失败，邮箱验证码过期，请重新发送！"
	}
	if val != string(emailcode) {
		
		return false, 405, "注销失败，邮箱验证码错误！"
	}

	_ = dbfunction.AddLogoutAccount(account, email, reason)
	
	return true, 200, "注销成功,请等待审核！"

}

func DisposeUserhome(account string) *data.Userregister {

	temp := dbfunction.CheckUser(account)
	return temp
}

func DisposeRedactUser(account string, name string, sex string, headphoto string) (bool, int, string) {

	if headphoto != "" {
		t := dbfunction.ModifyUserheadphoto(account, headphoto)
		if !t {
			return false, 400, "修改失败"
		}
	}
	if name != "" {
		t := dbfunction.ModifyUsername(account, name)
		if !t {
			return false, 400, "修改失败"
		}
	}
	if sex != "" {
		t := dbfunction.ModifyUsersex(account, sex)
		if !t {
			return false, 400, "修改失败"
		}
	}
	return true, 200, "修改成功"

}

func DisposeRedactUseremail(oldemail string, newemail string, reason string, emailcode string) (bool, int, string) {

	val := redisdb.Getvalue(oldemail)

	if val == "" {
		
		return false, 406, "修改失败，邮箱验证码过期，请重新发送！"
	}
	if val != string(emailcode) {
	
		return false, 405, "修改失败，邮箱验证码错误！"
	}

	t, state := dbfunction.CheckAuserEmail(newemail)
	if t {
		if state == 0 { //如果邮箱状态为0
		
			return false, 404, "修改失败，该邮箱正等待审核!"
		} else if state == 1 || state == 2 || state == 50 { //如果邮箱状态为1
	
			return false, 403, "修改失败，该邮箱已被注册!"
		}
	}

	_ = dbfunction.AddModifyUseremail(oldemail, newemail, reason)

	return true, 200, "修改成功,请等待审核！"

}

func DisposeShowdairy(account string) []data.Dairy {

	temp, _ := dbfunction.SeekAlldairy(account)
	return temp
}

func DisposeAdddiary(account string, details string, title string, classify string, performance string) (bool, int, string) {
	error := dbfunction.AddUserdiary(account, details, title, classify, performance)

	if error != nil {
		
		return false, 400, "添加失败"

	} else {
	
		return true, 200, "添加成功"

	}
}

func DisposeModifydairy(dairyid int, details string, title string, classify string, performance string ) (bool, int, string) {

	temp := dbfunction.ModifyUSerdairy(dairyid, details, title, classify, performance)
	if !temp {
	
		return false, 400, "修改失败"
	} else {

		return true, 200, "修改成功"
	}

}

func DisposeDeletedairy(dairyid int) (bool, int, string) {
	temp := dbfunction.DeleteUserdairy(dairyid)
	if !temp {

		return false, 400, "删除失败"

	} else {
	
		return true, 200, "删除成功"

	}
}

func DisposeEmptydairy(account string) (bool, int, string) {
	temp := dbfunction.EmptyUserdairy(account)
	if !temp {

		return false, 400, "删除失败"

	} else {

		return true, 200, "删除成功"

	}
}

func DisposeBackRegister(account string, password string, email string ,emailcode string ) (bool, int, string) {

	val := redisdb.Getvalue(email)

	if val == "" {

		return false, 406, "注册失败，邮箱验证码过期，请重新发送！"
	}
	if val != string(emailcode) {

		return false, 405, "注册失败，邮箱验证码错误！"
	}

	t, _ := dbfunction.CheckAuserAccount(account)
	if t {


		return false, 407, "注册失败，该用户名已被使用！"
	}

	dbfunction.SaveAdmin(account, password, email)

	return true, 200, "注册成功"

}

func DisposeBackLogin(email string, password string) (bool, int, string) {

	temp, _ := dbfunction.CheckAdminEmailAndPassword(email, password)
	if temp {

		return true, 200, "登录成功"
	} else {

		return false, 400, "登录失败,账号或密码错误！"
	}

}

func Disposetempuser() []data.Userregister {
	temp, _ := dbfunction.Showtempuser()
	return temp
}

func DisposePasstempuser(account string, email string) (bool, int, string) {

	temp := dbfunction.Passuser(account)
	emailcodedispose.AdminSendEmail(email, "您的账号注册申请已审核通过，可以登录网站！")

	if temp {

		return true, 200, "通过成功"

	} else {
	
		return false, 400, "通过失败"

	}
}

func DisposeTurntempuser(account string, email string, reason string) (bool, int, string) {

	temp := dbfunction.Deleteauser(account)
	emailcodedispose.AdminSendEmail(email, "您的用户注册申请未通过审核，已被驳回，未通过原因 为："+reason)

	if temp {

		return true, 200, "驳回成功"

	} else {
	
		return false, 400, "驳回失败"

	}
}

func DisposeShowlogoutuser() []data.Logoutuser {
	temp, _ := dbfunction.Showlogoutuser()
	return temp
}

func DisposePasslogoutuser(account string, email string) (bool, int, string) {

	temp := dbfunction.Passlogoutuser(account)
	emailcodedispose.AdminSendEmail(email, "注销申请通过，账户信息已清空！")

	if temp {
	
		return true, 200, "注销成功"

	} else {

		return false, 400, "注销失败"

	}
}

func DisposeTurnlogoutuser(account string, email string, reason string) (bool, int, string) {

	temp1 := dbfunction.Deletelogoutuser(account)
	temp2 := dbfunction.ModifyUserstate(account)
	emailcodedispose.AdminSendEmail(email, "您的用户注销申请未通过审核，已被驳回，未通过原因 为："+reason)

	if temp1 && temp2 {
	
		return true, 200, "驳回成功"

	} else {

		return false, 400, "驳回失败"

	}
}

func Disposeshowuser() []data.Userregister {
	temp, _ := dbfunction.Showuser()
	return temp
}

func DisposeAdminRedactUser(account string, email string, name string, sex string, headphoto string) (bool, int, string) {

	t, state := dbfunction.CheckAuserEmail(email)
	if t {
		if state == 0 { //如果邮箱状态为0

			return false, 404, "修改失败，该邮箱正等待审核!"
		} else if state == 1 || state == 2 || state == 50 { //如果邮箱状态为1
	
			return false, 403, "修改失败，该邮箱已被注册!"
		}
	}

	if headphoto != "" {
		t := dbfunction.ModifyUserheadphoto(account, headphoto)
		if !t {

			return false, 400, "修改失败"
		}
	}
	if name != "" {
		t := dbfunction.ModifyUsername(account, name)
		if !t {

			return false, 400, "修改失败"
		}
	}
	if sex != "" {
		t := dbfunction.ModifyUsersex(account, sex)
		if !t {

			return false, 400, "修改失败"
		}
	}
	if email != "" {
		t := dbfunction.AdminModifyUserEmail(account, email)
		if !t {

			return false, 400, "修改失败"
		}
	}

	return true, 200, "修改成功"

}

func DisposeShowModifyemail() []data.Modifyuseremail {
	temp, _ := dbfunction.Showmodifyemail()
	return temp
}

func DisposePassModifyemail(oldemail string, newemail string) (bool, int, string) {

	temp := dbfunction.PassModifyemail(newemail, oldemail)
	emailcodedispose.AdminSendEmail(newemail, "您的邮箱修改审核已通过，可以使用新邮箱登录网站！")
	if temp {

		return true, 200, "通过成功"

	} else {

		return false, 400, "通过失败"

	}
}

func DisposeTurnModifyemail(oldemail string, reason string) (bool, int, string) {

	temp := dbfunction.DeleteModifyemail(oldemail)
	emailcodedispose.AdminSendEmail(oldemail, "您的用户邮箱修改申请未通过审核，已被驳回，未通过原因 为："+reason)
	if temp {

		return true, 200, "驳回成功"

	} else {

		return false, 400, "驳回失败"

	}
}

func Judgeusertoken(account string, tokenstring string) (bool, int, string) {

	//从redis获取token
	val := redisdb.Getvalue(account)

	if val == "" {

		return false, 402, "身份信息过期"
	}
	if tokenstring != val {

		return false, 401, "身份信息不符"
	} else {

		return true, 200, "登录成功"
	}

}

func DisposeSendadminemailcode(email string) (bool, int, string) { //发验证码

	t, _ := dbfunction.CheckAuserEmail(email)
	if t {

		return false, 410, "发送失败，该邮箱已被注册!"
	} else {

		emailcode := emailcodedispose.SendEmail(email)

		fmt.Println(emailcode)

		if emailcode == "40" { //验证码发送失败

			return false, 400, "发送失败"

		} else { //验证码发送成功

			//验证码存入redis

			temp := redisdb.Setkey(email, emailcode, 60) //存入redis 设置60秒过期
			if temp {
				return true, 200, "发送成功"

			}
			return false, 200, "发送成功"
		}
	}

}
