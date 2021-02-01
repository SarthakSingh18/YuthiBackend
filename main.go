package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
	"yuthi.com/collab/global"
	"yuthi.com/collab/google/protobuf"
)

type AuthServer struct{}

//func (s AuthServer) GetAllCollabs(_ context.Context, _ *emptypb.Empty) (*protobuf.CollabsListingResponse, error) {
//	ctx,cancel:=global.NewDBContext(5*time.Second)
//	defer cancel()
//	cursor,err:=global.DB.Collection("collabs").Find(ctx,bson.M{})
//	if err!=nil{
//		log.Fatal(err)
//	}
//	collab:=&global.Collab{}
//	defer cursor.Close(context.Background())
//	for cursor.Next(context.Background()){
//		err:=cursor.Decode(collab)
//		if err!=nil{
//			return &protobuf.CollabsListingResponse{},nil
//		}
//		proto
//	}
//
//	return &protobuf.CollabsListingResponse{}, nil
//}

func (s AuthServer) UpdateCollabInfo(_ context.Context, request *protobuf.UpdateCollabInfoRequest) (*protobuf.Info, error) {
	collabId, name, description, iconUrl := request.GetCollabId(), request.GetName(), request.GetDescription(), request.IconUrl
	collabAccessType := request.GetAccessType()
	if collabId == "" || name == "" || description == "" || iconUrl == "" || collabAccessType.String() == "" {
		return &protobuf.Info{
			Status: 300,
			Error:  "BadRequest",
		}, nil
	}
	collabIdHex, err := primitive.ObjectIDFromHex(collabId)
	if err != nil {
		return &protobuf.Info{}, nil
	}
	filter := bson.M{
		"_id": collabIdHex,
	}
	updateCollab := bson.M{
		"$set": bson.M{
			"name":             name,
			"description":      description,
			"collabAccessType": collabAccessType.String(),
			"iconUrl":          iconUrl,
		},
	}
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	global.DB.Collection("collabs").FindOneAndUpdate(ctx, filter, updateCollab)
	return &protobuf.Info{
		Status: 200,
		Error:  "",
	}, nil
}

func (s AuthServer) GetCollabDetailInfo(_ context.Context, request *protobuf.CollabDetailRequest) (*protobuf.CollabDetailResponse, error) {
	collab_id := request.GetCollabId()
	if collab_id == "" {
		return &protobuf.CollabDetailResponse{}, nil
	}
	fmt.Println(collab_id)
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	var collab global.Collab
	objectId, err := primitive.ObjectIDFromHex(collab_id)
	if err != nil {
		return &protobuf.CollabDetailResponse{}, nil
	}
	global.DB.Collection("collabs").FindOne(ctx, bson.M{"_id": objectId}).Decode(&collab)
	if collab.Name == "" {
		return &protobuf.CollabDetailResponse{}, nil
	}
	collabAcessType := protobuf.CollabAccessType_value[collab.CollabAccessType]
	var cat protobuf.CollabAccessType
	cat = protobuf.CollabAccessType(collabAcessType)
	fmt.Println("prinintng cat", cat)
	collabDetailInfo := &protobuf.CollabDetailInfo{
		Id:           collab.ID.String(),
		Name:         collab.Name,
		Description:  collab.Description,
		AccessType:   cat,
		IconUrl:      collab.IconUrl,
		Participants: nil,
		SubCollabs:   nil,
	}
	return &protobuf.CollabDetailResponse{Collab: collabDetailInfo}, nil

}

func (s AuthServer) CreateCollab(_ context.Context, request *protobuf.CreateCollabRequest) (*protobuf.Info, error) {
	name, description, iconUrl := request.GetName(), request.GetDescription(), request.GetIconUrl()
	collabAccess := request.GetAccessType()
	if name != "" && description != "" && iconUrl != "" {
		newCollab := global.Collab{
			ID:               primitive.NewObjectID(),
			Name:             name,
			Description:      description,
			CollabAccessType: collabAccess.String(),
			IconUrl:          iconUrl,
		}
		ctx, cancel := global.NewDBContext(5 * time.Second)
		defer cancel()
		_, err := global.DB.Collection("collabs").InsertOne(ctx, newCollab)
		if err != nil {
			log.Println("Error Inserting Collab", err.Error())
		}
		return &protobuf.Info{
			Status: 200,
			Error:  "",
		}, nil
	}
	return &protobuf.Info{
		Status: 300,
		Error:  "Bad Request",
	}, nil
}

func main() {
	server := grpc.NewServer()
	protobuf.RegisterCollabServiceServer(server, AuthServer{})
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal("Error Creating listener", err.Error())
	}
	server.Serve(listener)

}
