package v1Ticket

import (
	"bytes"
	"encoding/binary"
	"github.com/wii-tools/wadlib"
	"io"
)

const (
	V1TicketHeaderSize       uint32 = 20
	V1SectionHeaderSize      uint32 = 20
	V1SubscriptionRecordSize uint32 = 24
)

func CreateV1Ticket(ticket []byte, subscriptionRecords []V1SubscriptionRecord) ([]byte, error) {
	var tempWad wadlib.WAD
	// Check if passed ticket is actually a ticket
	err := tempWad.LoadTicket(ticket)
	if err != nil {
		return nil, err
	}

	// Init v1 Ticket
	v1Ticket := V1Ticket{
		Ticket: tempWad.Ticket,
		V1TicketHeader: V1TicketHeader{
			Version:                  1,
			HeaderSize:               uint16(V1TicketHeaderSize),
			V1TicketSize:             V1TicketHeaderSize + (V1SubscriptionRecordSize * uint32(len(subscriptionRecords))) + (V1SectionHeaderSize * uint32(len(subscriptionRecords))),
			SectionHeaderTableOffset: V1TicketHeaderSize + (V1SubscriptionRecordSize * uint32(len(subscriptionRecords))),
			NumberOfSectionHeaders:   uint16(len(subscriptionRecords)),
			SectionHeaderSize:        uint16(V1SectionHeaderSize),
		},
		V1SubscriptionRecords: subscriptionRecords,
		V1SectionHeaders:      nil,
	}

	// Now create the section headers from the subscription records
	var sectionHeaders []V1SectionHeader
	for i, _ := range subscriptionRecords {
		sectionHeader := V1SectionHeader{
			RecordOffset:    V1TicketHeaderSize + (V1SubscriptionRecordSize * uint32(i)),
			NumberOfRecords: 1,
			RecordSize:      V1SubscriptionRecordSize,
			SectionSize:     V1SubscriptionRecordSize,
			SectionType:     Subscription,
		}
		sectionHeaders = append(sectionHeaders, sectionHeader)
	}

	v1Ticket.V1SectionHeaders = sectionHeaders

	// Now write to a buffer and return the bytes
	buffer := bytes.NewBuffer(nil)
	err = Write(buffer, v1Ticket.Ticket)
	if err != nil {
		return nil, err
	}

	// If the previous write passed we can assume the rest won't fail.
	Write(buffer, v1Ticket.V1TicketHeader)
	Write(buffer, v1Ticket.V1SubscriptionRecords)
	Write(buffer, v1Ticket.V1SectionHeaders)

	return buffer.Bytes(), nil
}

// Write writes the passed data to an io.Writer method.
func Write(writer io.Writer, data any) error {
	return binary.Write(writer, binary.BigEndian, data)
}
