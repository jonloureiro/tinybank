package usecases_test

import (
	"context"
	"testing"

	"github.com/jonloureiro/tiny-bank/internal"
	"github.com/jonloureiro/tiny-bank/internal/accounts"
	"github.com/jonloureiro/tiny-bank/internal/accounts/gateways/repositories"
	"github.com/jonloureiro/tiny-bank/internal/accounts/usecases"
	"github.com/stretchr/testify/require"
)

func Test_CreateAccount(t *testing.T) {
	t.Parallel()

	validName := "jonloureiro"
	validCPF := "93105949186"
	validSecret := "123456"

	type args struct {
		ctx   context.Context
		input usecases.CreateAccountInput
	}

	tests := []struct {
		name  string
		args  args
		setup func(context.Context) usecases.CreateAccountUC
		err   error
	}{
		{
			name: "create account without error",
			args: args{
				ctx: context.Background(),
				input: usecases.CreateAccountInput{
					Name:   validName,
					CPF:    validCPF,
					Secret: validSecret,
				},
			},
			setup: func(ctx context.Context) usecases.CreateAccountUC {
				return usecases.NewCreateAccountUC(
					repositories.NewAccountsRepositoryInMemory(),
				)
			},
		},
		{
			name: "create account error when input is invalid",
			args: args{
				ctx: context.Background(),
				input: usecases.CreateAccountInput{
					Name:   "f",
					CPF:    "1",
					Secret: "error",
				},
			},
			setup: func(ctx context.Context) usecases.CreateAccountUC {
				return usecases.NewCreateAccountUC(
					repositories.NewAccountsRepositoryInMemory(),
				)
			},
			err: internal.ErrFailedDependency,
		},
		{
			name: "create account error when cpf is already in use",
			args: args{
				ctx: context.Background(),
				input: usecases.CreateAccountInput{
					Name:   validName,
					CPF:    validCPF,
					Secret: validSecret,
				},
			},
			setup: func(ctx context.Context) usecases.CreateAccountUC {
				repo := repositories.NewAccountsRepositoryInMemory()
				acc, _ := accounts.New(validName, validCPF, validSecret)
				_ = repo.Save(ctx, acc)
				return usecases.NewCreateAccountUC(repo)
			},
			err: internal.ErrConflict,
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
