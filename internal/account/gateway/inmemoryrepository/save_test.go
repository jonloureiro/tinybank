package inmemoryrepository_test

import (
	"context"
	"testing"

	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/internal/account/gateway/inmemoryrepository"
	"github.com/jonloureiro/tinybank/internal/account/test"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.Parallel()

	repo := inmemoryrepository.New()

	require.NotNil(t, repo)
}

func Test_Save(t *testing.T) {
	t.Parallel()

	_account := test.Account(1)

	type args struct {
		ctx context.Context
		acc account.Account
	}

	tests := []struct {
		name  string
		args  args
		setup func(context.Context) account.SaveRepository
		err   error
	}{
		{
			name: "valid account",
			args: args{
				ctx: context.Background(),
				acc: _account,
			},
			setup: func(ctx context.Context) account.SaveRepository {
				return inmemoryrepository.New()
			},
		},
		{
			name: "invalid account (cpf)",
			args: args{
				ctx: context.Background(),
				acc: _account,
			},
			setup: func(ctx context.Context) account.SaveRepository {
				repo := inmemoryrepository.New()
				_ = repo.Save(ctx, _account)
				return repo
			},
			err: account.ErrAccountAlreadyExists,
		},
		{
			name: "invalid account (id)",
			args: args{
				acc: _account,
			},
			setup: func(ctx context.Context) account.SaveRepository {
				repo := inmemoryrepository.New()
				_ = repo.Save(ctx, _account)
				return repo
			},
			err: account.ErrAccountAlreadyExists,
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
