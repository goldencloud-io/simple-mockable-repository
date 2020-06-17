package main

import (
	"context"
	"log"
	"pkg.goldencloud.dev/simple-mockable-repository/mongorepository"
	"pkg.goldencloud.dev/simple-mockable-repository/repository"
)

func main() {
	ctx := context.Background()
	repo, err := mongorepository.NewMongoRepository("db", "uri")
	if err != nil {
		panic(err)
	}

	if err := repo.Open(ctx); err != nil {
		panic(err)
	}
	defer repo.Close(ctx)

	var shopID int64 = 111111

	accountID, err := getActivateAccountID(ctx, shopID, repo)
	if err != nil {
		panic(err)
	}
	log.Println(accountID)
}

func getActivateAccountID(ctx context.Context, shopID int64, repo repository.Repository) (string, error) {
	account, err := repo.GetActivateAccount(ctx, shopID)
	if err != nil {
		return "", err
	}
	return account.ID, nil
}
