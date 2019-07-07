package Email

import (
	"Common/Conf"
	"crypto/tls"
	"github.com/astaxie/beego/logs"
	"gopkg.in/gomail.v2"
	_ "gopkg.in/gomail.v2"
)

type MailContent struct {
	Address string
	Name    string
	Title   string
	Type    string
	Body    string
}

var (
	mailConf   *Conf.TomlInfo
	mailDialer *gomail.Dialer
)

func init() {
	mailConf = Conf.GetMailConf("mail.toml")
	mailDialer = gomail.NewDialer(mailConf.Mail.Host,
		int(mailConf.Mail.Port),
		mailConf.Mail.Mailer,
		mailConf.Mail.Password)
	mailDialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

func SendMail(content *MailContent) error {

	m := gomail.NewMessage()
	m.SetAddressHeader("From", mailConf.Mail.From, mailConf.Mail.From)
	m.SetAddressHeader("To", content.Address, content.Name)
	m.SetHeader("Subject", content.Title)
	m.SetBody(content.Type, content.Body)

	if err := mailDialer.DialAndSend(m); err != nil {
		logs.Error("failure send mail to ", content.Address, err)
		return err
	}
	return nil
}
