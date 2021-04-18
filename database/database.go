package database

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const(
	dburi="mongodb://localhost:27017"
	dbname="yuthi"
	performance=100
)
var DB mongo.Database

func ConnectToDB(){
	ctx,cancel:=NewDBContext(10*time.Second)
	defer cancel()
	client,err:=mongo.Connect(ctx,options.Client().ApplyURI(dburi))
	if err!=nil{
		log.Fatal("Error Connecting to database",err.Error())
	}
	DB=*client.Database(dbname)
}
func NewDBContext(d time.Duration)(context.Context,context.CancelFunc){
	return context.WithTimeout(context.Background(),d*performance/100)
}
