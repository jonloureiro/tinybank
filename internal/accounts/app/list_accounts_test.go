package app_test

import (
	"context"
	"testing"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/app"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories/test"
	"github.com/stretchr/testify/require"
)

func Test_ListUsecase(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name  string
		args  args
		setup func(context.Context) accounts.ListAccountsUsecase
		want  accounts.ListAccountsOutput
		err   error
	}{
		{
			name: "list accounts without error",
			args: args{
				ctx: context.Background(),
			},
			setup: func(ctx context.Context) accounts.ListAccountsUsecase {
				repo := repositories.NewAccountsRepositoryInMemory()
				_ = repo.Save(ctx, test.Account(1))
				_ = repo.Save(ctx, test.Account(2))
				return app.NewListAccountsUsecase(repo)
			},
			want: accounts.ListAccountsOutput{
				Accounts: []accounts.Account{
					test.Account(1),
					test.Account(2),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := tt.setup(tt.args.ctx)
			got, err := uc.Execute(tt.args.ctx)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, got)
		})
	}
}
