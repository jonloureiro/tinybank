package usecase_test

import (
	"context"
	"testing"

	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/internal/account/app/usecase"
	"github.com/jonloureiro/tinybank/internal/account/gateway/inmemoryrepository"
	"github.com/jonloureiro/tinybank/internal/account/test"
	"github.com/stretchr/testify/require"
)

func Test_GetBalance(t *testing.T) {
	t.Parallel()

	validUUID := "c0916b90-fe5e-4f88-b5ae-248ebeeb5125"

	type args struct {
		ctx   context.Context
		input account.GetBalanceInput
	}

	tests := []struct {
		name  string
		args  args
		setup func(context.Context) account.GetBalanceUsecase
		want  account.GetBalanceOutput
		err   error
	}{
		{
			name: "get account balance without error",
			args: args{
				ctx: context.Background(),
				input: account.GetBalanceInput{
					AccountID: validUUID,
				},
			},
			setup: func(ctx context.Context) account.GetBalanceUsecase {
				repo := inmemoryrepository.New()
				acc := test.Account(111)
				acc.SetID(validUUID)
				acc.SetBalance(12345)
				_ = repo.Save(ctx, acc)
				return usecase.NewGetBalance(repo)
			},
			want: account.GetBalanceOutput{
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
