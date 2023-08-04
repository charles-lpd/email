# sendEmail
发送邮件

## go

```golang
import (
  "github.com/charles-lpd/email"
)

func main() {

	p := SendParams{
		From:         "",           // 发送者 gmail 地址
		Password:     "",           // Google 账户的 应用专用密码
		To:           []string{""}, // 接收着 gmail 地址
		Title:        "",           // email 标题
		ContentType:  "",           // email 编码格式 text/html or text/plain
		EmailContent: "",           // email 主体内容
		Attachments: []Attachment{{ // 附件
			FileName: "文件",          // 配置文件名， 可选参数
			FilePath: "./README.md", // 文件地址
		}}, // 可配置多个附件
	}
	err := SendEmail(p)
	if err != nil {
		fmt.Println(err)
	}

}
```
