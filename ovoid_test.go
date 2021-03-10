package ovoid

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedBalanceResponse RespBalance = RespBalance{
	Balance: Balance{
		OvoPoint: OvoPoint{
			CardBalance:   0.0,
			CardNo:        "8009574905404600",
			PaymentMethod: "OVO",
		},
		OvoCash: OvoCash{
			CardBalance:   0,
			CardNo:        "8009574905404001",
			PaymentMethod: "OVO Cash",
		},
		OvoMain: OvoMain{
			CardBalance:   0,
			CardNo:        "800XXXXXXXXXX000",
			PaymentMethod: "MAIN",
		},
	},
}

func TestGetAllBalance(t *testing.T) {
	ovo, err := NewClient(os.Getenv("OVOTOKEN"), 0)
	if err != nil {
		panic(err)
	}

	b, err := ovo.GetAllBalances(context.Background())
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedBalanceResponse, *b, "balance not the same as expected")
}

func TestGetAllNotification(t *testing.T) {
	ovo, err := NewClient(os.Getenv("OVOTOKEN"), 0)
	if err != nil {
		panic(err)
	}

	b, err := ovo.GetAllNotifications(context.Background())
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 0, len(b), "notification error")
}

func TestGetCountUnreadNotification(t *testing.T) {
	ovo, err := NewClient(os.Getenv("OVOTOKEN"), 0)
	if err != nil {
		panic(err)
	}

	b, err := ovo.GetCountUnreadNotifications(context.Background())
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 0, *b, "total not same with the expected")
}
