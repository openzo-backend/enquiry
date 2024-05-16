package repository

import (
	"github.com/google/uuid"
	"github.com/tanush-128/openzo_backend/enquiry/internal/models"

	"gorm.io/gorm"
)

type EnquiryRepository interface {
	CreateEnquiry(Enquiry models.Enquiry) (models.Enquiry, error)
	GetEnquiryByID(id string) (models.Enquiry, error)
	GetEnquirysByStoreID(store_id string) ([]models.Enquiry, error)
	ChangeOrderStatus(id string, status string) (models.Enquiry, error)
	UpdateEnquiry(Enquiry models.Enquiry) (models.Enquiry, error)
	// Add more methods for other Enquiry operations (GetEnquiryByEmail, UpdateEnquiry, etc.)

}

type enquiryRepository struct {
	db *gorm.DB
}

func NewEnquiryRepository(db *gorm.DB) EnquiryRepository {

	return &enquiryRepository{db: db}
}

func (r *enquiryRepository) CreateEnquiry(Enquiry models.Enquiry) (models.Enquiry, error) {
	Enquiry.ID = uuid.New().String()
	tx := r.db.Create(&Enquiry)

	if tx.Error != nil {
		return models.Enquiry{}, tx.Error
	}

	return Enquiry, nil
}

func (r *enquiryRepository) GetEnquiryByID(id string) (models.Enquiry, error) {
	var Enquiry models.Enquiry
	tx := r.db.Where("id = ?", id).First(&Enquiry)
	if tx.Error != nil {
		return models.Enquiry{}, tx.Error
	}

	return Enquiry, nil
}

func (r *enquiryRepository) GetEnquirysByStoreID(store_id string) ([]models.Enquiry, error) {
	var Enquirys []models.Enquiry
	tx := r.db.Where("store_id = ?", store_id).Find(&Enquirys)
	if tx.Error != nil {
		return []models.Enquiry{}, tx.Error

	}

	return Enquirys, nil
}


func (r *enquiryRepository) ChangeOrderStatus(id string, status string) (models.Enquiry, error) {
	var Enquiry models.Enquiry
	tx := r.db.Where("id = ?", id).First(&Enquiry)
	if tx.Error != nil {
		return models.Enquiry{}, tx.Error
	}
	
	// if status == "accepted" {
	// 	Enquiry.OrderStatus = models.OrderAccepted
	// } else if status == "rejected" {
	// 	Enquiry.OrderStatus = models.OrderRejected
	// } else if status == "out_for_delivery" {
	// 	Enquiry.OrderStatus = models.OrderOutForDel
	// } else if status == "cancelled" {
	// 	Enquiry.OrderStatus = models.OrderCancelled
	// } else if status == "delivered" {
	// 	Enquiry.OrderStatus = models.OrderDelivered
	// }

	tx = r.db.Save(&Enquiry)
	if tx.Error != nil {
		return models.Enquiry{}, tx.Error
	}

	return Enquiry, nil
}

func (r *enquiryRepository) UpdateEnquiry(Enquiry models.Enquiry) (models.Enquiry, error) {
	tx := r.db.Save(&Enquiry)
	if tx.Error != nil {
		return models.Enquiry{}, tx.Error
	}

	return Enquiry, nil
}

// Implement other repository methods (GetEnquiryByID, GetEnquiryByEmail, UpdateEnquiry, etc.) with proper error handling
