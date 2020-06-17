package mongorepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"pkg.goldencloud.dev/simple-mockable-repository/repository"
)

const collectionName = "google-analytic-accounts"

func NewMongoRepository(databaseName, URI string) (*mongoRepository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}
	return &mongoRepository{
		MongoClient:  client,
		DatabaseName: databaseName,
	}, nil
}

type mongoRepository struct {
	MongoClient  *mongo.Client
	Database     *mongo.Database
	DatabaseName string
}

func (repo *mongoRepository) Open(ctx context.Context) error {
	if err := repo.MongoClient.Connect(ctx); err != nil {
		return err
	}

	if err := repo.MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	repo.Database = repo.MongoClient.Database(repo.DatabaseName)

	return nil
}

func (repo *mongoRepository) Close(ctx context.Context) error {
	return repo.MongoClient.Disconnect(ctx)
}

func (repo *mongoRepository) SaveAccount(ctx context.Context, shopID int64, account *repository.GoogleAnalyticAccount) error {
	accountCollection := repo.Database.Collection(collectionName)

	filter := bson.M{
		"id": shopID,
	}
	update := bson.M{
		"$set": bson.M{
			"activate_account":                               account,
			"accounts." + account.ID + ".id":                 account.ID,
			"accounts." + account.ID + ".picture":            account.Picture,
			"accounts." + account.ID + ".picture":            account.Picture,
			"accounts." + account.ID + ".token.access_token": account.Token.AccessToken,
			"accounts." + account.ID + ".token.expiry":       account.Token.Expiry,
		},
	}
	upsert := true
	option := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := accountCollection.UpdateOne(ctx, filter, update, &option)
	if err != nil {
		return err
	}

	return nil
}

func (repo *mongoRepository) GetActivateAccount(ctx context.Context, shopID int64) (*repository.GoogleAnalyticAccount, error) {
	return nil, nil
}
