package inmemory_test

import (
	"context"
	"testing"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/app"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories/inmemory"
	"github.com/jonloureiro/tiny-bank/internal/common"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.Parallel()

	repo := inmemory.NewAccountsRepositoryInMemory()

	require.NotNil(t, repo)
}

func Test_Save(t *testing.T) {
	t.Parallel()

	validAccount, _ := app.NewAccount("jonloureiro", "93105949186", "123456")

	type args struct {
		ctx context.Context
		acc accounts.Account
	}

	tests := []struct {
		name  string
		args  args
		setup func(context.Context) accounts.SaveAccountsRepository
		err   error
	}{
		{
			name: "valid account",
			args: args{
				ctx: context.Background(),
				acc: validAccount,
			},
			setup: func(ctx context.Context) accounts.SaveAccountsRepository {
				return inmemory.NewAccountsRepositoryInMemory()
			},
		},
		{
			name: "invalid account (cpf)",
			args: args{
				ctx: context.Background(),
				acc: validAccount,
			},
			setup: func(ctx context.Context) accounts.SaveAccountsRepository {
				repo := inmemory.NewAccountsRepositoryInMemory()
				_ = repo.Save(ctx, validAccount)
				return repo
			},
			err: common.ErrConflict,
		},
		{
			name: "invalid account (id)",
			args: args{
				acc: validAccount,
			},
			setup: func(ctx context.Context) accounts.SaveAccountsRepository {
				repo := inmemory.NewAccountsRepositoryInMemory()
				_ = repo.Save(ctx, validAccount)
				return repo
			},
			err: common.ErrConflict,
		},
	}

	for _, tt := range tests {
		tc := tt

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repo := tc.setup(tc.args.ctx)

			err := repo.Save(tc.args.ctx, tc.args.acc)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}

			require.NoError(t, err)
		})
	}
}
