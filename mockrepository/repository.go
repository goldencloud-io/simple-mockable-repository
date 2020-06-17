package mockrepository

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/rs/zerolog/log"

	"pkg.goldencloud.dev/simple-mockable-repository/repository"
)

type mockRepository struct {
	DataFilePath string // absolute path
	shopProfiles []repository.ShopProfile
}

func NewMockRepository(filePath string) (*mockRepository, error) {
	if !filepath.IsAbs(filePath) {
		absFilepath, err := filepath.Abs(filePath)
		if err != nil {
			return nil, err
		}
		filePath = absFilepath
	}
	repo := &mockRepository{
		DataFilePath: filePath,
	}
	return repo, nil
}

func (repo *mockRepository) Open(ctx context.Context) error {
	fileContent, err := ioutil.ReadFile(repo.DataFilePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileContent, &repo.shopProfiles); err != nil {
		return err
	}
	return nil
}

func (repo *mockRepository) Close(ctx context.Context) error {
	log.Info().Msg("closed")
	return nil
}

func (repo *mockRepository) SaveAccount(ctx context.Context, shopID int64, account *repository.GoogleAnalyticAccount) error {
	return nil
}

func (repo *mockRepository) GetActivateAccount(ctx context.Context, shopID int64) (*repository.GoogleAnalyticAccount, error) {
	for _, profile := range repo.shopProfiles {
		if profile.ID != shopID {
			continue
		}
		if profile.ActivateAccount == nil {
			return nil, repository.ErrNoAccountFound
		}
		return profile.ActivateAccount, nil
	}
	return nil, repository.ErrShopNotFound
}
