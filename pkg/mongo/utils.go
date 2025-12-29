package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectIDFromHexOrNil returns an ObjectID from the provided hex representation.
func ObjectIDFromHexOrNil(id string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(id)
	return objID
}

func ObjectIDsFromHexOrNil(ids []string) []primitive.ObjectID {
	objIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objIDs[i] = ObjectIDFromHexOrNil(id)
	}
	return objIDs
}

func HexFromObjectID(id primitive.ObjectID) string {
	return id.Hex()
}

func HexFromObjectIDs(ids []primitive.ObjectID) []string {
	hexs := make([]string, len(ids))
	for i, id := range ids {
		hexs[i] = HexFromObjectID(id)
	}
	return hexs
}

func BuildQueryWithSoftDelete(query bson.M) bson.M {
	query["deleted_at"] = nil
	return query
}

func GetMongoDateTimeNow() primitive.DateTime {
	return primitive.NewDateTimeFromTime(time.Now())
}
