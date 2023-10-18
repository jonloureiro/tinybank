package usecase_test

import (
	"context"
	"testing"

	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/internal/account/app/domain"
	"github.com/jonloureiro/tinybank/internal/account/app/usecase"
	"github.com/jonloureiro/tinybank/internal/account/gateway/inmemoryrepository"
	"github.com/jonloureiro/tinybank/pkg/defaulterr"
	"github.com/stretchr/testify/require"
)

func Test_Create(t *testing.T) {
	t.Parallel()

	validName := "jonloureiro"
	validCPF := "93105949186"
	validSecret := "123456"

	type args struct {
		ctx   context.Context
		input account.CreateInput
	}

	tests := []struct {
		name  string
		args  args
		setup func(context.Context) account.CreateUsecase
		err   error
	}{
		{
			name: "create account without error",
			args: args{
				ctx: context.Background(),
				input: account.CreateInput{
					Name:   validName,
					CPF:    validCPF,
					Secret: validSecret,
				},
			},
			setup: func(ctx context.Context) account.CreateUsecase {
				return usecase.NewCreate(
					inmemoryrepository.New(),
				)
			},
		},
		{
			name: "create account error when input is invalid",
			args: args{
				ctx: context.Background(),
				input: account.CreateInput{
					Name:   "f",
					CPF:    "1",
					Secret: "error",
				},
			},
			setup: func(ctx context.Context) account.CreateUsecase {
				return usecase.NewCreate(
					inmemoryrepository.New(),
				)
			},
			err: defaulterr.FailedDependency(),
		},
		{
			name: "create account error when cpf is already in use",
			args: args{
				ctx: context.Background(),
				input: account.CreateInput{
					Name:   validName,
					CPF:    validCPF,
					Secret: validSecret,
				},
			},
			setup: func(ctx context.Context) account.CreateUsecase {
				repo := inmemoryrepository.New()
				acc, _ := domain.NewAccount(validName, validCPF, validSecret)
				_ = repo.Save(ctx, acc)
				return usecase.NewCreate(repo)
			},
			err: defaulterr.Conflict(),
		},
	}

	for _, tt := range tests {
		tc := tt

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			uc := tc.setup(tc.args.ctx)

			got, err := uc.Execute(tc.args.ctx, tc.args.input)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}

			require.NotEmpty(t, got.AccountID)
		})
	}
}
