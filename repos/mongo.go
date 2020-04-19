package repos

import (
	mgo "github.com/caixos/tokit/clients/mongo"
	"github.com/caixos/tokit/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

////////// mongodb 操作
func Mongo(database ...string) *mongo.Database {
	config := configs.LoadMongoConfig()
	if database == nil {
		return mgo.GetMongo().Database(config.Database)

	} else {
		return mgo.GetMongo().Database(database[0])
	}
}
