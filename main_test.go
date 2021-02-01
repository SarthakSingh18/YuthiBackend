package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
	"yuthi.com/collab/google/protobuf"
)

func TestAuthServer_CreateCollab(t *testing.T) {
	server:=AuthServer{}
	res,err:=server.CreateCollab(context.Background(),&protobuf.CreateCollabRequest{
		Name:        "Hello",
		Description: "Hey this collab is for creating yuthi",
		AccessType:  0,
		IconUrl:     "asdasdasd",
	})
	fmt.Println(res)
	if err!=nil{
		t.Error("Error Occured",err.Error())
	}
}

func TestAuthServer_GetCollabDetailInfo(t *testing.T) {
	server:=AuthServer{}
	res,err:=server.GetCollabDetailInfo(context.Background(),&protobuf.CollabDetailRequest{CollabId: "60114a5cc35794e9e4dd78d4"})
	fmt.Println(res)
	if err!=nil{
		t.Error(err)
	}
}

func TestAuthServer_UpdateCollabInfo(t *testing.T) {
	server:=AuthServer{}
	res,err:=server.UpdateCollabInfo(context.Background(),&protobuf.UpdateCollabInfoRequest{
		CollabId:    "60114a5cc35794e9e4dd78d4",
		Name:        "Yuthi",
		Description: "This is yuthi",
		AccessType:  0,
		IconUrl:     "htsfdsf",
	})
	fmt.Println("printing res",res)
	if err!=nil{
		t.Error(err)
	}
}

func TestAuthServer_GetAllCollabs(t *testing.T) {
	server:=AuthServer{}
	res,err:=server.GetAllCollabs(context.Background(),&emptypb.Empty{})
	fmt.Println("printing res",res)
	if err!=nil{
		t.Error(err)
	}
}