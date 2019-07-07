package Conf

import (
	"testing"
)

func TestMailConf(t *testing.T) {
	//filePath :="D:\\PHOENIX\\Documents\\WORKSPACE\\Go\\Daphne\\src\\Common\\Conf\\mail.toml"

	path := "mail.toml"
	conf := GetMailConf(path)

	t.Log("title: ", conf.Title)
	t.Log("mail host", conf.Mail.Host)
	t.Log("mail port", conf.Mail.Port)
	t.Log("mail mail", conf.Mail.Mailer)
	t.Log("mail pass", conf.Mail.Password)
	t.Log("mail tls", conf.Mail.UseTls)
	t.Log("mail from", conf.Mail.From)

}
