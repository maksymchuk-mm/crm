package models

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func TestCard(t *testing.T) {

	tests := []struct {
		name string
		args Card
		want string
	}{
		{
			name: "UAH",
			args: Card{
				ID:           10,
				Name:         "Test card UAH",
				CurrencyCode: "UAH",
				Balance:      152364,
				UpdatedAt:    time.Now(),
				Closed:       false,
			},
			want: "₴ 1 523,64",
		},
		{
			name: "USD",
			args: Card{
				ID:           11,
				Name:         "Test card USD",
				CurrencyCode: "USD",
				Balance:      152364,
				UpdatedAt:    time.Now(),
				Closed:       false,
			},
			want: "$ 1,523.64",
		},
		{
			name: "EUR",
			args: Card{
				ID:           12,
				Name:         "Test card EUR",
				CurrencyCode: "EUR",
				Balance:      152364,
				UpdatedAt:    time.Now(),
				Closed:       false,
			},
			want: "€ 1.523,64",
		},
		{
			name: "GBR",
			args: Card{
				ID:           13,
				Name:         "Test card GBR",
				CurrencyCode: "GBR",
				Balance:      152364,
				UpdatedAt:    time.Now(),
				Closed:       false,
			},
			want: "₴ 1 523,64",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.args.FormatBalance(), test.want)
		})
	}
}
