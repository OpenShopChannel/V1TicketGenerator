package v1Ticket

import (
	"os"
	"testing"
	"time"
)

func TestCreateV1Ticket(t *testing.T) {
	// test.tik contains a v0 ticket that allows for ticket exporting. Although not checked by us,
	// IOS validates this on GetTicketFromView.
	buffer, err := os.ReadFile("test.tik")
	if err != nil {
		panic(err)
	}

	// Create a subscription record
	records := []V1SubscriptionRecord{
		{
			// 1 month (30 days) from now
			ExpirationTime: uint32(time.Now().AddDate(0, 1, 0).Unix()),
			ReferenceID:    [16]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0x01, 0x23, 0x45, 0x67, 0x89, 0x01, 0x23, 0x45, 0x67, 0x89, 0x12},
		},
		{
			ExpirationTime: uint32(time.Now().AddDate(0, 1, 0).Unix()),
			ReferenceID:    [16]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0x01, 0x23, 0x45, 0x67, 0x89, 0x01, 0x23, 0x45, 0x67, 0x89, 0x13},
		},
	}

	v1Ticket, err := CreateV1Ticket(buffer, records)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("ticket.tv1", v1Ticket, 0666)
	if err != nil {
		panic(err)
	}
}
