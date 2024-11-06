package repository

import (
	"context"
	c "desa-temp-svc/config"
	"desa-temp-svc/util"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(ctx context.Context, username string, password string) (bson.M, error) {
	// logic

	client, err := c.ConnectMongoDB()
	if err != nil {
		return bson.M{}, err
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("Failed to disconnect from MongoDB")
		}
	}()
	var result bson.M
	fmt.Println("user : ", username)
	fmt.Println("pass : ", password)
	password = util.EncryptStringToSHA256(password)
	coll := client.Database("sidesa").Collection("account")
	err = coll.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&result)
	if err != nil {
		log.Println(err)
		return bson.M{}, err
	}

	return result, nil
}
