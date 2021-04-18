package services

import (
	"context"
	"fmt"
	"github.com/myuser/myrepo/database"
	createOrganisationService "github.com/myuser/myrepo/proto/helloworld"
	"github.com/myuser/myrepo/versions"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"time"
)

type Server2 struct{
	createOrganisationService.UnimplementedOrganisationServer
}

type newOrganisationStruct struct{
	Version string `bson:"version"`
	Name string `bson:"name"`
	Desc string `bson:"desc"`
	Email string `bson:"email"`
	Contact string `bson:"contact"`
	Image string `bson:"image"`
	Video string `bson:"video"`

}
func(s *Server2) createOrganisation(ctx context.Context,request *createOrganisationService.OrganisationRequest)(*httpbody.HttpBody,error){
	name,desc,email,contact,image,video:=request.GetName(),request.GetDesc(),request.GetEmail(),request.GetContact(),request.GetImage(),request.GetVideo()
	if name == " " || desc ==" " || email == " "{
		return &httpbody.HttpBody{
			ContentType: "text/html",
			Data: []byte("error"),
		}, nil
	}
	newOrganisation:=newOrganisationStruct{
		Version: versions.Version,
		Name:    name,
		Desc:    desc,
		Email:   email,
		Contact: contact,
		Image:   image,
		Video:   video,
	}
	ctx, cancel := database.NewDBContext(5 * time.Second)
	defer cancel()
	_,err := database.DB.Collection("admin").InsertOne(ctx,newOrganisation)
	if err!=nil{
		fmt.Print("HERE??")
		fmt.Print(err)
		return &httpbody.HttpBody{
			ContentType: "text/html",
			Data: []byte("error"),
		}, nil
	}
	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:  []byte("Success"),
	},nil
}