package models

import (
	"time"
)

// type Enquiry struct {
// 	ID string `json:"id" gorm:"primaryKey"` // The ID field is also the UUID for the store

// }

type EnquiryStatus string

const (
	EnquiryStatusPending  EnquiryStatus = "pending"
	EnquiryStatusAccepted EnquiryStatus = "accepted"
	EnquiryStatusRejected EnquiryStatus = "rejected"
)

type Enquiry struct {
	ID          string        `json:"id" gorm:"primaryKey"`
	StoreID     string        `json:"store_id"`
	CustomerID  string        `json:"customer_id"`
	EnquiryTime time.Time     `json:"enquiry_time" gorm:"autoCreateTime"`
	Query       string        `json:"query"`
	Status      EnquiryStatus `json:"status" gorm:"default:'pending'"`
}

type Sender string

const (
	SenderCustomer Sender = "customer"
	SenderStore    Sender = "store"
)

type Message struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	EnquiryID string    `json:"enquiry_id"`
	Sender    Sender    `json:"sender"`
	Message   string    `json:"message"`
}

func (enquiry *Enquiry) ToMap() map[string]string {
	return map[string]string{
		"type":         "enquiry",
		"id":           enquiry.ID,
		"store_id":     enquiry.StoreID,
		"customer_id":  enquiry.CustomerID,
		"enquiry_time": enquiry.EnquiryTime.String(),
		"query":        enquiry.Query,
		"status":       string(enquiry.Status),
	}
}

func (message *Message) ToMap() map[string]string {
	return map[string]string{
		"type":       "message",
		"id":         message.ID,
		"created_at": message.CreatedAt.String(),
		"enquiry_id": message.EnquiryID,
		"sender":     string(message.Sender),
		"message":    message.Message,
	}
}
