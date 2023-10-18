package vo_test

import (
	"testing"

	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/internal/account/app/vo"
	"github.com/stretchr/testify/require"
)

func TestNewCPF(t *testing.T) {
	t.Parallel()

	fakeCPF := "93105949186"

	type args struct {
		cpf string
	}

	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "valid cpf",
			args: args{
				cpf: fakeCPF,
			},
			want: fakeCPF,
		},
		{
			name: "invalid cpf",
			args: args{
				cpf: "1",
			},
			err: account.ErrIncompatibleCPFLength,
		},
	}

	for _, tt := range tests {
		tc := tt

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := vo.ParseCPF(tc.args.cpf)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}

			require.Equal(t, tc.want, got.Value())
		})
	}
}
