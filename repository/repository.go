package repository

import (
	"context"
	"time"
)

type AccessToken struct {
	AccessToken  string    `bson:"access_token,omitempty"`
	RefreshToken string    `bson:"refresh_token,omitempty"`
	Expiry       time.Time `bson:"expiry,omitempty"`
}

type GoogleAnalyticAccount struct {
	ID      string      `bson:"id"`
	Picture string      `bson:"picture"`
	Token   AccessToken `bson:"token" bson:"token"`
}

type ShopProfile struct {
	ID              int64                             `json:"id" bson:"id"`
	ActivateAccount *GoogleAnalyticAccount            `json:"activate_account,omitempty" bson:"activate_account,omitempty"`
	Accounts        map[string]*GoogleAnalyticAccount `json:"accounts,omitempty" bson:"accounts,omitempty"`
}

type Repository interface {
	Open(ctx context.Context) error
	Close(ctx context.Context) error
	SaveAccount(ctx context.Context, shopID int64, account *GoogleAnalyticAccount) error
	GetActivateAccount(ctx context.Context, shopID int64) (*GoogleAnalyticAccount, error)
}
