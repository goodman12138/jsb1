package emailcodedispose

import (
	"fmt"
	"go/pkg/mylog"
	"go/conf/utils"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

var Code = make(map[string]int) //key为Email,value为随机码

var log = mylog.NewLog("Info")

// MailboxConf 邮箱配置
type MailboxConf struct {
	// 邮件标题
	Title string
	// 邮件内容
	Body string
	// 收件人列表
	RecipientList []string
	// 发件人账号
	Sender string
	// 发件人密码，QQ邮箱这里配置授权码
	SPassword string
	// SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPAddr string
	// SMTP端口 QQ邮箱是25
	SMTPPort int
}

func SendEmail(email string) string {
	var mailConf MailboxConf
	mailConf.Title = "验证来咯"

	mailConf.RecipientList = []string{email}
	mailConf.Sender = utils.GetEmailsender()

	mailConf.SPassword = utils.GetEmailspassword()

	mailConf.SMTPAddr = utils.GetEmailsmtpaddr()
	mailConf.SMTPPort = utils.GetEmailsmtpport()

	//产生六位数验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	//内容
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%s,为了保证账号安全，验证码有效期为60秒。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, vcode)

	m := gomail.NewMessage()

	// 第三个参数是我们发送者的名称，但是如果对方有发送者的好友，优先显示对方好友备注名
	m.SetHeader(`From`, mailConf.Sender, "记事本")
	m.SetHeader(`To`, mailConf.RecipientList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)

	err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
	if err != nil {
		log.Error(err.Error())
		return "40" 
	} else {
		log.Info("Send Email Success")
		return vcode
	}

}

func AdminSendEmail(email string, msg string) {

	var mailConf MailboxConf
	mailConf.Title = "温馨提示"

	mailConf.RecipientList = []string{email}
	mailConf.Sender = utils.GetEmailsender()

	mailConf.SPassword = utils.GetEmailspassword()

	mailConf.SMTPAddr = utils.GetEmailsmtpaddr()
	mailConf.SMTPPort = 587

	//内容
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>%s 感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, msg)

	m := gomail.NewMessage()

	// 第三个参数是我们发送者的名称，但是如果对方有发送者的好友，优先显示对方好友备注名
	m.SetHeader(`From`, mailConf.Sender, "记事本")
	m.SetHeader(`To`, mailConf.RecipientList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)

	err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("Send Email Success")
	}
}
