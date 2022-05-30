package route

import (
	"go/App/acceptfunction"
	"net/http"
)

func Init() {
	// //前台功能
	http.HandleFunc("/sendemailcode", acceptfunction.Sendemailcode)         //发送邮箱验证码(注册)
	http.HandleFunc("/register", acceptfunction.Register)                   //用户注册
	http.HandleFunc("/login", acceptfunction.Login)                         //用户登录
	http.HandleFunc("/senduseremailcode", acceptfunction.Senduseremailcode) //发送邮箱验证码
	http.HandleFunc("/forgetpassword", acceptfunction.ForgetPassword)       //忘记密码

	http.HandleFunc("/userhome", acceptfunction.Userhome)               //用户主页
	http.HandleFunc("/logout", acceptfunction.Logout)                   //用户注销
	http.HandleFunc("/redactuser", acceptfunction.RedactUser)           //修改用户资料
	http.HandleFunc("/redactuseremail", acceptfunction.RedactUseremail) //修改用户邮箱
	http.HandleFunc("/redactuserpassword", acceptfunction.RedactUserpassword)

	http.HandleFunc("/showdairy", acceptfunction.Showdairy)     //显示用户日记
	http.HandleFunc("/adddairy", acceptfunction.Adddairy)       //添加日志
	http.HandleFunc("/modifydairy", acceptfunction.Modifydairy) //修改日志
	http.HandleFunc("/deletedairy", acceptfunction.Deletedairy) //删除日志
	http.HandleFunc("/emptydairy", acceptfunction.Emptydairy)   //清空日志

	//后台功能
	http.HandleFunc("/sendadminemailcode", acceptfunction.Sendadminemailcode) //发送管理员注册邮箱验证码
	http.HandleFunc("/backregister", acceptfunction.BackRegister)             //后台注册
	http.HandleFunc("/backlogin", acceptfunction.BackLogin)                   //后台登录

	http.HandleFunc("/showtempuser", acceptfunction.Showtempuser) //显示注册待审核用户
	http.HandleFunc("/passtempuser", acceptfunction.Passtempuser) //通过用户注册
	http.HandleFunc("/turntempuser", acceptfunction.Turntempuser) //驳回用户注册

	http.HandleFunc("/showlogoutuser", acceptfunction.Showlogoutuser) //显示待注销用户
	http.HandleFunc("/passlogoutuser", acceptfunction.Passlogoutuser) //通过用户注销
	http.HandleFunc("/turnlogoutuser", acceptfunction.Turnlogoutuser) //驳回用户注销

	http.HandleFunc("/showuser", acceptfunction.Showuser) //显示所有用户

	http.HandleFunc("/adminredactuser", acceptfunction.AdminRedactUser) //管理员修改用户信息
	http.HandleFunc("/showmodifyemail", acceptfunction.ShowModifyemail) //显示用户修改邮箱信息
	http.HandleFunc("/passmodifyemail", acceptfunction.PassModifyemail) //通过用户邮箱修改
	http.HandleFunc("/turnmodifyemail", acceptfunction.TurnModifyemail) //驳回用户邮箱修改

}
