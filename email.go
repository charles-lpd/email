package email

import (
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type Attachment struct {
	FileName string
	FilePath string
}

type MailParams struct {
	From         string   `form:"from" binding:"required"`         // 发送者
	Password     string   `form:"password" binding:"required"`     // 发送者密码
	To           []string `form:"to" binding:"required"`           // 接收者地址
	Title        string   `form:"title" binding:"required"`        // 邮件主题
	ContentType  string   `form:"contentType" binding:"required"`  // 邮件内容格式
	EmailContent string   `form:"emailContent" binding:"required"` // 邮件内容
}

func SendEmail(p MailParams, attachments []Attachment) error {

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

	if len(attachments) != 0 {
		for _, v := range attachments {
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

func SendEmailAPI(p MailParams, files map[string][]*multipart.FileHeader, c *gin.Context) error {

	n := gomail.NewDialer("smtp.gmail.com", 587, p.From, p.Password)

	msg := gomail.NewMessage()
	// 发送者
	msg.SetHeader("From", p.From)
	// 接受者
	msg.SetHeader("To", p.To...)
	// 邮件主题，标题
	msg.SetHeader("Subject", p.Title)
	// 邮件主体内容
	msg.SetBody(p.ContentType, p.EmailContent)

	// 遍历文件列表，逐个处理
	for _, file := range files {
		// 打开上传的文件
		f, err := file[0].Open()
		if err != nil {
			return err
		}
		// 处理上传的文件
		data, readErr := ioutil.ReadAll(f)
		if readErr != nil {
			return readErr
		}
		msg.Attach(file[0].Filename, gomail.SetCopyFunc(func(w io.Writer) error {
			_, writeErr := w.Write(data)
			return writeErr
		}))

		defer f.Close()
	}
	// Send the email
	err := n.DialAndSend(msg)
	if err != nil {
		return err
	}
	return nil
}
