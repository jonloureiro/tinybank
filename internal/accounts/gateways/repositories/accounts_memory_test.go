package repositories_test

import (
	"context"
	"testing"

	"github.com/jonloureiro/tiny-bank/internal"
	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories"
	"github.com/jonloureiro/tiny-bank/internal/accounts/usecases"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.Parallel()

	repo := repositories.NewAccountsRepositoryInMemory()

	require.NotNil(t, repo)
}

func Test_Save(t *testing.T) {
	t.Parallel()

	validAccount, _ := accounts.New("jonloureiro", "93105949186", "123456")

	type args struct {
		ctx context.Context
		acc accounts.Account
	}

	tests := []struct {
		name  string
		args  args
		setup func(context.Context) usecases.AccountsRepository
		err   error
	}{
		{
			name: "valid account",
			args: args{
				ctx: context.Background(),
				acc: validAccount,
			},
			setup: func(ctx context.Context) usecases.AccountsRepository {
				return repositories.NewAccountsRepositoryInMemory()
			},
		},
		{
			name: "invalid account (cpf)",
			args: args{
				ctx: context.Background(),
				acc: validAccount,
			},
			setup: func(ctx context.Context) usecases.AccountsRepository {
				repo := repositories.NewAccountsRepositoryInMemory()
				_ = repo.Save(ctx, validAccount)
				return repo
			},
			err: internal.ErrConflict,
		},
		{
			name: "invalid account (id)",
			args: args{
				acc: validAccount,
			},
			setup: func(ctx context.Context) usecases.AccountsRepository {
				repo := repositories.NewAccountsRepositoryInMemory()
				_ = repo.Save(ctx, validAccount)
				return repo
			},
			err: internal.ErrConflict,
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
