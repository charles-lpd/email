package email

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmailTo(t *testing.T) {
	htmlBody := `
	<html lang="en">
	  <body>
	    <h1>Hello!</h1>
	  </body>
	</html>
	`
	p := SendParams{
		From:         "",
		Password:     "",
		To:           []string{""},
		Title:        "测试标题",
		ContentType:  "text/html", // text/html or text/plain
		EmailContent: htmlBody,
		Attachments: []Attachment{{
			FileName: "文件",          // 配置文件名， 可选参数
			FilePath: "./README.md", // 文件地址
		}}, // 可配置多个附件
	}
	err := SendEmail(p)
	assert.NoError(t, err)
	t.Log("Test SendEmail no error")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("发送成功")
	}
}
