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

func Test_List(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name  string
		args  args
		setup func(context.Context) account.ListUsecase
		want  account.ListOutput
		err   error
	}{
		{
			name: "list accounts without error",
			args: args{
				ctx: context.Background(),
			},
			setup: func(ctx context.Context) account.ListUsecase {
				repo := inmemoryrepository.New()
				_ = repo.Save(ctx, test.Account(1))
				_ = repo.Save(ctx, test.Account(2))
				return usecase.NewList(repo)
			},
			want: account.ListOutput{
				Accounts: []account.Account{
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
