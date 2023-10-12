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

func Test_GetAccountBalance(t *testing.T) {
	t.Parallel()

	validUUID := "c0916b90-fe5e-4f88-b5ae-248ebeeb5125"

	type args struct {
		ctx   context.Context
		input accounts.GetAccountBalanceInput
	}

	tests := []struct {
		name  string
		args  args
		setup func(context.Context) accounts.GetAccountBalanceUsecase
		want  accounts.GetAccountBalanceOutput
		err   error
	}{
		{
			name: "get account balance without error",
			args: args{
				ctx: context.Background(),
				input: accounts.GetAccountBalanceInput{
					AccountID: validUUID,
				},
			},
			setup: func(ctx context.Context) accounts.GetAccountBalanceUsecase {
				repo := repositories.NewAccountsRepositoryInMemory()
				acc := test.Account(111)
				acc.SetID(validUUID)
				acc.SetBalance(12345)
				_ = repo.Save(ctx, acc)
				return app.NewGetAccountBalanceUsecase(repo)
			},
			want: accounts.GetAccountBalanceOutput{
				Balance: 12345,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := tt.setup(tt.args.ctx)
			got, err := uc.Execute(tt.args.ctx, tt.args.input)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, got)
		})
	}
}
