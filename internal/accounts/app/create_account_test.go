package app_test

import (
	"context"
	"testing"

	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/app"
	"github.com/jonloureiro/tiny-bank/internal/accounts/app/domain"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories/inmemory"
	"github.com/jonloureiro/tiny-bank/internal/common"
	"github.com/stretchr/testify/require"
)

func Test_CreateAccount(t *testing.T) {
	t.Parallel()

	validName := "jonloureiro"
	validCPF := "93105949186"
	validSecret := "123456"

	type args struct {
		ctx   context.Context
		input accounts.CreateAccountInput
	}

	tests := []struct {
		name  string
		args  args
		setup func(context.Context) accounts.CreateAccountUC
		err   error
	}{
		{
			name: "create account without error",
			args: args{
				ctx: context.Background(),
				input: accounts.CreateAccountInput{
					Name:   validName,
					CPF:    validCPF,
					Secret: validSecret,
				},
			},
			setup: func(ctx context.Context) accounts.CreateAccountUC {
				return app.NewCreateAccountUC(
					inmemory.NewAccountsRepositoryInMemory(),
				)
			},
		},
		{
			name: "create account error when input is invalid",
			args: args{
				ctx: context.Background(),
				input: accounts.CreateAccountInput{
					Name:   "f",
					CPF:    "1",
					Secret: "error",
				},
			},
			setup: func(ctx context.Context) accounts.CreateAccountUC {
				return app.NewCreateAccountUC(
					inmemory.NewAccountsRepositoryInMemory(),
				)
			},
			err: common.ErrFailedDependency,
		},
		{
			name: "create account error when cpf is already in use",
			args: args{
				ctx: context.Background(),
				input: accounts.CreateAccountInput{
					Name:   validName,
					CPF:    validCPF,
					Secret: validSecret,
				},
			},
			setup: func(ctx context.Context) accounts.CreateAccountUC {
				repo := inmemory.NewAccountsRepositoryInMemory()
				acc, _ := domain.NewAccount(validName, validCPF, validSecret)
				_ = repo.Save(ctx, acc)
				return app.NewCreateAccountUC(repo)
			},
			err: common.ErrConflict,
		},
	}

	for _, tt := range tests {
		tc := tt

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			uc := tc.setup(tc.args.ctx)

			got, err := uc.CreateAccount(tc.args.ctx, tc.args.input)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}

			require.NotEmpty(t, got.AccountID)
		})
	}
}
