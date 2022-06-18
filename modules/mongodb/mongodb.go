package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/danisbagus/simple-go-ddd/utils/config"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mutex            = &sync.Mutex{}
	mongoClient      *mongo.Client
	mongoDatabase    *mongo.Database
	defaultAtlasOpts = map[string]string{
		"retryWrites": "true",
		"w":           "majority",
	}
)

const (
	StandardPrefix    = "mongodb://"
	DNSSeedListPrefix = "mongodb+srv://"
)

func GetClient() (*mongo.Client, error) {
	if mongoClient == nil {
		return nil, errors.New("please init mongodb")
	}

	return mongoClient, nil
}

func GetDatabase() (*mongo.Database, error) {
	if mongoDatabase == nil {
		return nil, errors.New("please init mongodb")
	}

	return mongoDatabase, nil
}

func Init(cfg config.MongoDatabaseConfig) {

	mutex.Lock()
	defer mutex.Unlock()

	if mongoDatabase != nil {
		return
	}

	client, err := NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to connect MongoDB %v", err)
	}

	mongoClient = client
	mongoDatabase = client.Database(cfg.DBName)
}

func NewClient(cfg config.MongoDatabaseConfig) (*mongo.Client, error) {
	if mongoClient != nil {
		return mongoClient, nil
	}

	prefix := StandardPrefix
	opts := make(map[string]string)
	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	if cfg.UseSRV {
		prefix = DNSSeedListPrefix
		opts = defaultAtlasOpts
		address = cfg.Host
	}

	uri := fmt.Sprintf("%v%s", prefix, address)

	if cfg.Username != "" {
		uri = fmt.Sprintf("%v%s:%s@%s", prefix, cfg.Username, cfg.Password, address)
	}

	if len(opts) != 0 {
		counter := 0
		for key, val := range opts {
			if counter == 0 {
				uri = fmt.Sprintf("%s/?%s=%s", uri, key, val)
			} else {
				uri = fmt.Sprintf("%s&%s=%s", uri, key, val)
			}
			counter++
		}
	}

	// direct connect
	// uri := "mongodb://<username>:<password>@<host>:<port>?directConnection=true"

	// mongo atlas
	// uri := "mongodb+srv://<username>:<password>@<host>?retryWrites=true&w=majority"

	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Println(evt.Command)
		},
	}

	mongoOptions := options.Client().ApplyURI(uri).SetMonitor(monitor)

	client, err := mongo.NewClient(mongoOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		client.Disconnect(context.TODO())
		return nil, err
	}

	return client, nil

}
