# sendEmail

发送邮件

## 使用条件

1. 打开对应 gmail 邮箱 找到 ` 转发和 POP/IMAP ` 并打开 ` IMAP 访问 `。
2. `Password` 参数，需在对应 Google 账户中 => 安全性 => 开启 `2FA` 并创建 `应用专用密码`。

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
		Attachments: []Attachment{{ // 配置本地 附件
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
