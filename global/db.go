package global

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)
import "context"

var DB mongo.Database

func connectToDB(){
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

