package services

import (
	"context"
	"fmt"
	"testing"
	"yuthi.com/mailService/proto"
)

func TestAuthServer_SendMail(t *testing.T) {
	server:=AuthServer{}
	_,err:=server.SendMail(context.Background(),&proto.Mail{
		ToEmail: "mrrobotyup@gmail.com",
		Subject: "This is regarding test email",
		Body:    "This is the mail body",
	})
	if err!=nil{
		fmt.Println(err)
	}
}
