package logs

import (
	"context"
	"github.com/kainonly/collection-service/src/common"
	"github.com/kainonly/collection-service/src/facade"
	"github.com/mongodb/mongo-go-driver/bson"
)

var err error

type Base struct {

}

func CheckAllowDomain(domain string) bool {
	collection := facade.MGODb[common.Config.SystemDatabase].Collection("whitelist")
	var someone map[string]interface{}
	result := collection.FindOne(context.Background(), bson.D{{"domain", domain}})
	return result.Decode(&someone) == nil
}
