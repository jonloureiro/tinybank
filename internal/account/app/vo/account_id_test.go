package vo_test

import (
	"testing"

	"github.com/jonloureiro/tinybank/internal/account"
	"github.com/jonloureiro/tinybank/internal/account/app/vo"
	"github.com/stretchr/testify/assert"
)

func TestParseAccountID(t *testing.T) {
	t.Parallel()

	validUUID := "d939fd6e-596f-429e-88c1-a32d17416c6d"

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "valid account id",
			args: args{
				s: validUUID,
			},
			want: validUUID,
		},
		{
			name: "invalid account id",
			args: args{
				s: "1",
			},
			err: account.ErrInvalidAccountID,
		},
	}

	for _, tt := range tests {
		tc := tt

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := vo.ParseAccountID(tc.args.s)
			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err)
				return
			}

			assert.Equal(t, tc.want, got.Value())
		})
	}
}
