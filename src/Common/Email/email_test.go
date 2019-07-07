package Email

import (
	"Common/Conf"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"testing"
)

func TestGoMail(t *testing.T) {
	mailConf := Conf.GetMailConf("mail.toml")
	t.Log("current log mail ", mailConf.Mail.Host)
	m := gomail.NewMessage()
	m.SetAddressHeader("From", mailConf.Mail.From, mailConf.Mail.From)
	m.SetAddressHeader("To", "leeforrest@126.com", "leeforrest")
	m.SetHeader("Subject", "gomail-邮件测试9")
	m.SetBody("text/html", "<h1>hello world tests sample</h1>")

	d := gomail.NewDialer(mailConf.Mail.Host, int(mailConf.Mail.Port), mailConf.Mail.Mailer, mailConf.Mail.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("*** faulte %s\n", err.Error())
		t.Error("failure send mail: ", err)
	}
}

func TestSendMail(t *testing.T) {
	m := MailContent{
		Address: "leeforrest@126.com",
		Name:    "leeforrest",
		Title:   "gomail-发送测试0",
		Type:    "text/html",
		Body:    "<h1>hello world tests sample</h1>",
	}
	if err := SendMail(&m); err != nil {
		t.Error("failure send mail ", err)
	}
}
