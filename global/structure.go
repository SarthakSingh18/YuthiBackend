package global

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Collab struct{
	ID primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Description string `bson:"description"`
	CollabAccessType string `bson:"collabAccessType"`
	IconUrl string `bson:"iconUrl"`
}
type UpdateCollab struct {
	Name string `bson:"name"`
	Description string `bson:"description"`
	CollabAccessType string `bson:"collabAccessType"`
	IconUrl string `bson:"iconUrl"`
}

