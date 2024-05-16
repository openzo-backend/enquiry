package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/enquiry/internal/models"
)

func (s *enquiryService) GetEnquiryByID(ctx *gin.Context, id string) (models.Enquiry, error) {
	enquiry, err := s.enquiryRepository.GetEnquiryByID(id)
	if err != nil {
		return models.Enquiry{}, err
	}

	return enquiry, nil
}

func (s *enquiryService) GetEnquirysByStoreID(ctx *gin.Context, store_id string) ([]models.Enquiry, error) {
	enquirys, err := s.enquiryRepository.GetEnquirysByStoreID(store_id)
	if err != nil {
		return []models.Enquiry{}, err
	}

	return enquirys, nil
}

