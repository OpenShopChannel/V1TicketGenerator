package V1TicketGenerator

import (
	"github.com/wii-tools/wadlib"
)

type SectionType uint16

const Subscription SectionType = 2

type V1Ticket struct {
	// Ticket is common among all tickets on the Wii.
	Ticket                wadlib.Ticket
	V1TicketHeader        V1TicketHeader
	V1SubscriptionRecords []V1SubscriptionRecord
	V1SectionHeaders      []V1SectionHeader
}

// V1TicketHeader is the header for a v1 Ticket found immediately after the actual Ticket.
type V1TicketHeader struct {
	// Version will always be 1.
	Version uint16
	// HeaderSize will always be 20
	HeaderSize uint16
	// V1TicketSize is the size of the ticket minus V1Ticket.Ticket
	V1TicketSize             uint32
	SectionHeaderTableOffset uint32
	NumberOfSectionHeaders   uint16
	// SectionHeaderSize will always be 20
	SectionHeaderSize uint16
	_                 uint32
}

type V1SectionHeader struct {
	// RecordOffset is the offset to the subscription record.
	RecordOffset uint32
	// NumberOfRecords in our case will always be 1.
	NumberOfRecords uint32
	// RecordSize will always be 24 in our case
	RecordSize uint32
	// SectionSize will always be 20.
	SectionSize uint32
	SectionType SectionType
	_           uint16
}

type V1SubscriptionRecord struct {
	// ExpirationTime is when the subscription to the current content ends, in Unix timestamp
	ExpirationTime uint32
	ReferenceID    [16]byte
	_              uint32
}
