package logs

import (
	"context"
	"github.com/kainonly/collection-service/src/facade"
	"github.com/mongodb/mongo-go-driver/bson"
)

var err error

type Base struct {
	Database string
}

func (m *Base) ValidateWhitelist(key string, value string) bool {
	collection := facade.Db[m.Database].Collection("whitelist")
	var someone map[string]interface{}
	result := collection.FindOne(context.Background(), bson.D{{key, value}})
	return result.Decode(&someone) == nil
}
