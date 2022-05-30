package acceptfunction

import (
	"encoding/json"
	"fmt"
	"go/BusinessLogicLayer/dispose"
	"go/BusinessLogicLayer/tokenoperation"
	"go/DataAccessLayer/redisdb"
	"go/pkg/mylog"
	"go/pkg/structpkg/data"
	"go/pkg/structpkg/res"
	"net/http"
	"strconv"
)

var log = mylog.NewLog("Warning")

const (
	SecretKey = "welcome to jishibeng"
)

func Send(w http.ResponseWriter, r interface{}) {
	var Str, _ = json.Marshal(r)
	w.Header().Set("Content-Length", strconv.Itoa(len(Str)))
	w.Write(Str)
}

func Sethead(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域

	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型

	w.Header().Set("content-type", "application/json") //返回数据格式是json
}

func Sendemailcode(w http.ResponseWriter, r *http.Request) { //发送邮箱验证码
	Sethead(w)
	// 解析参数
	var tempdata data.Userregister
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	fmt.Println(tempdata.Email)

	temp, statuscode, msg := dispose.DisposeSendemailcode(tempdata.Email)

	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	Send(w, b)
	if temp {
		log.Info("发送成功！")
	} else {
		log.Info("发送失败！")
	}

}

func Senduseremailcode(w http.ResponseWriter, r *http.Request) { //发送用户邮箱验证码
	Sethead(w)
	// 解析参数
	var tempdata data.Userregister
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	//fmt.Println(tempdata.Email)

	temp, statuscode, msg := dispose.DisposeSenduseremailcode(tempdata.Email)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	Send(w, b)
	if temp {
		log.Info("发送成功！")
	} else {
		log.Info("发送失败！")
	}

}

func Register(w http.ResponseWriter, r *http.Request) { //用户注册
	Sethead(w)
	// 解析参数
	var tempdata data.Userregister
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	temp, statuscode, msg := dispose.DisposeRegister(tempdata.Account, tempdata.Password, tempdata.Email, tempdata.Name, tempdata.Sex, tempdata.Emailcode)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	Send(w, b)
	if temp {
		log.Info("注册成功!")
	} else {
		log.Info("注册失败!")
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	Sethead(w)
	// 解析参数
	var tempdata data.Userregister
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	//	fmt.Println(tempdata.Account, tempdata.Email, tempdata.Password)
	temp, statuscode, msg := dispose.DisposeLogin(tempdata.Account, tempdata.Email, tempdata.Password)

	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if temp {
		log.Info("登录成功!")

		t1 := tokenoperation.Buildtoken(tempdata.Account, tempdata.Email)
		var Str, _ = json.Marshal(t1)
		//w.Header().Set("Content-Length", strconv.Itoa(len(Str)))
		w.Write(Str) //发送token
		//存储到redis

		_ = redisdb.Setkey(t1.Account, t1.Tokenstring, 15*24*60*60)

	} else {
		log.Info("登录失败!")
		//Send(w, b)
	}
	Send(w, b)

}

func ForgetPassword(w http.ResponseWriter, r *http.Request) {

	Sethead(w)
	// 解析参数
	var tempdata data.Userregister
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	temp, statuscode, msg := dispose.DisposeFogetPassword(tempdata.Email, tempdata.Password, tempdata.Emailcode)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	Send(w, b)
	if temp {
		log.Info("修改成功!")
	} else {
		log.Info("修改失败!")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

	Sethead(w)
	// 解析参数
	var tempdata data.Logoutuser
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeLogoutAccount(tempdata.Account, tempdata.Email, tempdata.Reason, tempdata.Emailcode)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("注销成功!")
		} else {
			log.Info("注销失败!")
		}
	}
}

func Userhome(w http.ResponseWriter, r *http.Request) {

	Sethead(w)
	// 解析参数
	var tempdata data.Userregister
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp := dispose.DisposeUserhome(tempdata.Account)
		Str, _ := json.Marshal(temp)
		w.Header().Set("Content-Length", strconv.Itoa(len(Str))) //发送查询到的用户信息
		w.Write(Str)
	}
}

func RedactUser(w http.ResponseWriter, r *http.Request) {
	Sethead(w)
	// 解析参数
	var tempdata data.Userregister
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeRedactUser(tempdata.Account, tempdata.Name, tempdata.Sex, tempdata.Headphoto)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("修改成功!")
		} else {
			log.Info("修改失败!")
		}
	}
}

func RedactUseremail(w http.ResponseWriter, r *http.Request) {
	Sethead(w)
	// 解析参数
	var tempdata data.Modifyuseremail
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)
	fmt.Println(tempdata)
	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeRedactUseremail(tempdata.Oldemail, tempdata.Newemail, tempdata.Reason, tempdata.Emailcode)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("修改成功!")
		} else {
			log.Info("修改失败!")
		}
	}
}

func RedactUserpassword(w http.ResponseWriter, r *http.Request) {
	Sethead(w)
	// 解析参数
	var tempdata data.Userregister
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)
	fmt.Println(tempdata)
	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeRedactUserpassword(tempdata.Email, tempdata.Emailcode, tempdata.Password)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("修改成功!")
		} else {
			log.Info("修改失败!")
		}
	}
}

func Showdairy(w http.ResponseWriter, r *http.Request) {
	Sethead(w)

	var tempdata data.Dairy
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)

	//获取tokenstring
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp := dispose.DisposeShowdairy(tempdata.Account)

		Str, _ := json.Marshal(temp)
		w.Header().Set("Content-Length", strconv.Itoa(len(Str))) //发送查询到的日记

		w.Write(Str)
	}

}

func Adddairy(w http.ResponseWriter, r *http.Request) {

	Sethead(w)
	var tempdata data.Dairy
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeAdddiary(tempdata.Account, tempdata.Details, tempdata.Title, tempdata.Classify, tempdata.Performance)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("添加成功！")
		} else {
			log.Info("添加失败！")
		}
	}
}

func Modifydairy(w http.ResponseWriter, r *http.Request) {

	Sethead(w)

	var tempdata data.Dairy
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)

	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeModifydairy(tempdata.Dairyid, tempdata.Details, tempdata.Title, tempdata.Classify, tempdata.Performance)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("修改成功！")
		} else {
			log.Info("修改失败！")
		}
	}
}

func Deletedairy(w http.ResponseWriter, r *http.Request) {

	Sethead(w)

	var tempdata data.Dairy
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeDeletedairy(tempdata.Dairyid)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("删除成功！")
		} else {
			log.Info("删除失败！")
		}
	}
}

func Emptydairy(w http.ResponseWriter, r *http.Request) {

	Sethead(w)

	var tempdata data.Dairy
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeEmptydairy(tempdata.Account)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("清空成功！")
		} else {
			log.Info("清空失败！")
		}
	}
}

func BackRegister(w http.ResponseWriter, r *http.Request) {
	Sethead(w)
	// 解析参数
	var tempdata data.Admin
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	//account string, password string, email string ,emailcode string
	temp, statuscode, msg := dispose.DisposeBackRegister(tempdata.Account, tempdata.Password, tempdata.Email, tempdata.Emailcode)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	Send(w, b)
	if temp {
		log.Info("注册成功!")
	} else {
		log.Info("注册失败!")
	}
}

func BackLogin(w http.ResponseWriter, r *http.Request) {
	Sethead(w)
	// 解析参数
	var tempdata data.Admin
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	//	fmt.Println(tempdata.Account, tempdata.Email, tempdata.Password)
	temp, statuscode, msg := dispose.DisposeBackLogin(tempdata.Email, tempdata.Password)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if temp {
		log.Info("登录成功!")

		t1 := tokenoperation.Buildadmintoken(tempdata.Email)
		var Str, _ = json.Marshal(t1)

		w.Write(Str) //发送token
		//存储到redis

		_ = redisdb.Setkey(t1.Account, t1.Tokenstring, 15*24*60*60)

	} else {
		log.Info("登录失败!")
		//Send(w, b)
	}
	Send(w, b)

}

func Showtempuser(w http.ResponseWriter, r *http.Request) {

	Sethead(w)

	temp := dispose.Disposetempuser()
	//temp, _ := dbfunction.ShowUserdiary(tempdata.Username)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		Str, _ := json.Marshal(temp)
		w.Header().Set("Content-Length", strconv.Itoa(len(Str)))
		w.Write(Str)
	}

}

func Passtempuser(w http.ResponseWriter, r *http.Request) {

	// 解析参数
	var tempdata data.Userregister
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	//	fmt.Println(tempdata.Account, tempdata.Email, tempdata.Password)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposePasstempuser(tempdata.Account, tempdata.Email)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("审核成功!")
		} else {
			log.Info("审核失败!")
		}
	}

}

func Turntempuser(w http.ResponseWriter, r *http.Request) {

	// 解析参数
	var tempdata data.Logoutuser
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	//	fmt.Println(tempdata.Account, tempdata.Email, tempdata.Password)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeTurntempuser(tempdata.Account, tempdata.Email, tempdata.Reason)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("驳回成功!")
		} else {
			log.Info("驳回失败!")
		}
	}
}

func Showlogoutuser(w http.ResponseWriter, r *http.Request) {

	Sethead(w)

	//var tempdata tempdata.Userregister
	//Recipientuser(&tempdata, r) //读取前端发送的数据
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {

		temp := dispose.DisposeShowlogoutuser()
		//temp, _ := dbfunction.ShowUserdiary(tempdata.Username)

		Str, _ := json.Marshal(temp)
		w.Header().Set("Content-Length", strconv.Itoa(len(Str)))
		w.Write(Str)
	}

}

func Passlogoutuser(w http.ResponseWriter, r *http.Request) {

	// 解析参数
	var tempdata data.Logoutuser
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)

	fmt.Println(tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		//	fmt.Println(tempdata.Account, tempdata.Email, tempdata.Password)
		temp, statuscode, msg := dispose.DisposePasslogoutuser(tempdata.Account, tempdata.Email)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("注销成功!")
		} else {
			log.Info("注销失败!")
		}
	}

}

func Turnlogoutuser(w http.ResponseWriter, r *http.Request) {

	// 解析参数
	var tempdata data.Logoutuser
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	//	fmt.Println(tempdata.Account, tempdata.Email, tempdata.Password)

	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeTurnlogoutuser(tempdata.Account, tempdata.Email, tempdata.Reason)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("驳回成功!")
		} else {
			log.Info("驳回失败!")
		}
	}

}

func Showuser(w http.ResponseWriter, r *http.Request) {

	Sethead(w)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp := dispose.Disposeshowuser()

		Str, _ := json.Marshal(temp)
		w.Header().Set("Content-Length", strconv.Itoa(len(Str)))
		w.Write(Str)
	}

}

func AdminRedactUser(w http.ResponseWriter, r *http.Request) {
	Sethead(w)
	// 解析参数
	var tempdata data.Userregister
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		//account string, email string, name string, sex string, headphoto string
		_, statuscode, msg := dispose.DisposeAdminRedactUser(tempdata.Account, tempdata.Email, tempdata.Name, tempdata.Sex, tempdata.Headphoto)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
	}

}

func ShowModifyemail(w http.ResponseWriter, r *http.Request) {

	Sethead(w)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {

		temp := dispose.DisposeShowModifyemail()

		Str, _ := json.Marshal(temp)
		w.Header().Set("Content-Length", strconv.Itoa(len(Str)))
		w.Write(Str)
	}

}

func PassModifyemail(w http.ResponseWriter, r *http.Request) {

	// 解析参数
	var tempdata data.Modifyuseremail
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		//	fmt.Println(tempdata.Account, tempdata.Email, tempdata.Password)
		temp, statuscode, msg := dispose.DisposePassModifyemail(tempdata.Oldemail, tempdata.Newemail)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("审核成功!")
		} else {
			log.Info("审核失败!")
		}
	}

}

func TurnModifyemail(w http.ResponseWriter, r *http.Request) {

	// 解析参数
	var tempdata data.Modifyuseremail
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	account, tokenstring := tokenoperation.GetUsertoken(r)

	t, statuscode, msg := dispose.Judgeusertoken(account, tokenstring)
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg

	if !t {
		Send(w, b)
	} else {
		temp, statuscode, msg := dispose.DisposeTurnModifyemail(tempdata.Oldemail, tempdata.Reason)
		var b res.Res
		b.StatusCode = statuscode
		b.Msg = msg

		Send(w, b)
		if temp {
			log.Info("驳回成功!")
		} else {
			log.Info("驳回失败!")
		}
	}
}

func Sendadminemailcode(w http.ResponseWriter, r *http.Request) { //发送管理员邮箱验证码
	Sethead(w)
	// 解析参数
	var tempdata data.Admin
	decoder := json.NewDecoder(r.Body)
	// 解析参数
	decoder.Decode(&tempdata)
	//fmt.Println(tempdata.Email)

	temp, statuscode, msg := dispose.DisposeSendadminemailcode(tempdata.Email) // 发送验证码 存入redis3
	var b res.Res
	b.StatusCode = statuscode
	b.Msg = msg
	Send(w, b)
	if temp {
		log.Info("发送成功！")
	} else {
		log.Info("发送失败！")
	}

}
