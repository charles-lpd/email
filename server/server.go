package main

import (
	"fmt"
	"net/http"
	"strings"

	email "github.com/charles-lpd/email"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/sendMail", sendMail)

	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}

func sendMail(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		ErrorResponse(c, Err_headers_content_type.Code, Err_headers_content_type.Msg, err.Error())
		return
	}
	var p email.MailParams
	if err := c.ShouldBind(&p); err != nil {
		ErrorResponse(c, Err_request_params.Code, Err_request_params.Msg, err.Error())
		return
	}
	// 前端传的数组会自动替换成字符串, 通过 , 进行分割
	p.To = strings.Split(p.To[0], ",")
	emailErr := email.SendEmailAPI(p, form.File, c)
	if emailErr != nil {
		fmt.Println(emailErr)
		ErrorResponse(c, Err_email_send_failed.Code, Err_email_send_failed.Msg, emailErr.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

func ErrorResponse(c *gin.Context, errorCode int, err, msg string) {
	var respCode int
	// 客服端 error
	if errorCode == Err_request_params.Code {
		respCode = http.StatusBadRequest
	} else {
		// 服务器端 error
		respCode = http.StatusInternalServerError
	}
	c.JSON(respCode, Response{
		Status: errorCode,
		Error:  err,
		Msg:    msg,
	})

}

type err struct {
	Code int
	Msg  string
}

var (
	Err_request_params       = err{4001, "err_request_params"}
	Err_headers_content_type = err{4002, "err_headers_content_type"}
	Err_email_send_failed    = err{4003, "err_email_send_failed"}
)
