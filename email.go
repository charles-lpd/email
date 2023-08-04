package email

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/gomail.v2"
)

type Attachment struct {
	FileName string
	FilePath string
}
type SendParams struct {
	From         string       // 发送者
	Password     string       // 发送者密码
	To           []string     // 接收者地址
	Title        string       // 邮件主题
	ContentType  string       // 邮件内容格式
	EmailContent string       // 邮件内容
	Attachments  []Attachment // 附件
}

func SendEmail(p SendParams) error {

	n := gomail.NewDialer("smtp.gmail.com", 587, p.From, p.Password)

	msg := gomail.NewMessage()
	// 发送者
	msg.SetHeader("From", p.From)
	// 接受者
	msg.SetHeader("To", p.To...)
	//邮件主题，标题
	msg.SetHeader("Subject", p.Title)
	// 邮件主体内容
	msg.SetBody(p.ContentType, p.EmailContent)
	// 添加附件，只能添加 本地附件

	if len(p.Attachments) != 0 {
		for _, v := range p.Attachments {
			if !FileExists(v.FilePath) {
				err := errors.New(v.FilePath + " File does not exist")
				return err
			}
			filename := v.FileName
			if v.FileName != "" {
				filename += filepath.Ext(v.FilePath)
			} else {
				filename = v.FilePath
			}
			msg.Attach(v.FilePath, gomail.Rename(filename))
		}
	}

	// Send the email
	err := n.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}
