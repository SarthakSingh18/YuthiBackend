package services

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/smtp"
	"yuthi.com/mailService/constants"
	"yuthi.com/mailService/proto"
)

type AuthServer struct{

}

func (a AuthServer) mustEmbedUnimplementedMailServiceServer() {
	panic("implement me")
}

func (a AuthServer) SendMail(ctx context.Context, mail *proto.Mail) (*emptypb.Empty, error) {
	toEmail,subject,body:= mail.GetToEmail(),mail.GetSubject(),mail.GetBody()
	if toEmail== " " || subject == " " || body == " "{
		fmt.Println("Mail Does not contain proper values")
	}
	toEmailString:=[]string{
		toEmail,
	}
	smtpHost:="smtp.gmail.com"
	smtpPort:="587"
	message:=[]byte(`this is the message`+body)
	auth:=smtp.PlainAuth("",constants.Email,constants.Password,smtpHost)
	err:=smtp.SendMail(smtpHost+":"+smtpPort,auth,constants.Email,toEmailString,message)
	if err!=nil{
		fmt.Println(err)
		return nil, err
	}
	return nil,nil
}
