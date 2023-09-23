package vos_test

import (
	"testing"

	"github.com/jonloureiro/tiny-bank/internal/accounts/app/vos"
	"github.com/jonloureiro/tiny-bank/internal/common"
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
			err: common.ErrFailedDependency,
		},
	}

	for _, tt := range tests {
		tc := tt

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := vos.NewCPF(tc.args.cpf)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}

			require.Equal(t, tc.want, got.Value())
		})
	}
}
