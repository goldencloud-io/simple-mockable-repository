package main

import (
	"context"
	"testing"

	"pkg.goldencloud.dev/simple-mockable-repository/mockrepository"
	_ "pkg.goldencloud.dev/simple-mockable-repository/mockrepository"
	"pkg.goldencloud.dev/simple-mockable-repository/repository"
)

func Test_getActivateAccountID(t *testing.T) {
	repo, err := mockrepository.NewMockRepository("../mockrepository/case1.json")
	if err != nil {
		t.Fatal(err)
	}

	if err := repo.Open(context.Background()); err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx    context.Context
		shopID int64
		repo   repository.Repository
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "mock",
			args: args{
				ctx:    context.Background(),
				shopID: 1,
				repo:   repo,
			},
			want:    "1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getActivateAccountID(tt.args.ctx, tt.args.shopID, tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("getActivateAccountID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getActivateAccountID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
