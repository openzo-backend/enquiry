package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/enquiry/internal/models"
	"github.com/tanush-128/openzo_backend/enquiry/internal/pb"
	"github.com/tanush-128/openzo_backend/enquiry/internal/repository"
)

type EnquiryService interface {

	//CRUD
	CreateEnquiry(ctx *gin.Context, req models.Enquiry) (models.Enquiry, error)
	GetEnquiryByID(ctx *gin.Context, id string) (models.Enquiry, error)
	GetEnquirysByStoreID(ctx *gin.Context, store_id string) ([]models.Enquiry, error)
	ChangeEnquiryStatus(ctx *gin.Context, id string, status string) (models.Enquiry, error)
	UpdateEnquiry(ctx *gin.Context, req models.Enquiry) (models.Enquiry, error)
	SendMessage(ctx *gin.Context, req models.Message) (models.Message, error)
}

type enquiryService struct {
	enquiryRepository   repository.EnquiryRepository
	notificationService pb.NotificationServiceClient
	storeService        pb.StoreServiceClient
}

func NewEnquiryService(enquiryRepository repository.EnquiryRepository,
	notificationService pb.NotificationServiceClient, storeService pb.StoreServiceClient,
) EnquiryService {
	return &enquiryService{enquiryRepository: enquiryRepository, notificationService: notificationService, storeService: storeService}
}

func (s *enquiryService) CreateEnquiry(ctx *gin.Context, req models.Enquiry) (models.Enquiry, error) {

	token, err := s.storeService.GetFCMToken(ctx, &pb.StoreId{Id: string(req.StoreID)})
	if err != nil {
		return models.Enquiry{}, err
	}

	createdEnquiry, err := s.enquiryRepository.CreateEnquiry(req)
	if err != nil {
		return models.Enquiry{}, err // Propagate error
	}

	_, err = s.notificationService.SendData(ctx, &pb.Data{
		Data:  createdEnquiry.ToMap(),
		Token: token.Token,
	})
	if err != nil {
		return models.Enquiry{}, err
	}

	return createdEnquiry, nil
}

func (s *enquiryService) ChangeEnquiryStatus(ctx *gin.Context, id string, status string) (models.Enquiry, error) {

	token := ""

	if status == "accepted" {
		acceptedNotification.Token = token
		_, err := s.notificationService.SendNotification(ctx, &acceptedNotification)
		if err != nil {
			return models.Enquiry{}, err
		}
	} else if status == "rejected" {
		rejectedNotification.Token = token
		_, err := s.notificationService.SendNotification(ctx, &rejectedNotification)
		if err != nil {
			return models.Enquiry{}, err
		}

	} else if status == "out_for_delivery" {
		outForDeliveryNotification.Token = token
		_, err := s.notificationService.SendNotification(ctx, &outForDeliveryNotification)
		if err != nil {
			return models.Enquiry{}, err
		}

	} else if status == "cancelled" {
		cancelledNotification.Token = token
		_, err := s.notificationService.SendNotification(ctx, &cancelledNotification)
		if err != nil {
			return models.Enquiry{}, err
		}

	} else if status == "delivered" {
		deliveredNotification.Token = token
		_, err := s.notificationService.SendNotification(ctx, &deliveredNotification)
		if err != nil {
			return models.Enquiry{}, err
		}
	}

	changedEnquiry, err := s.enquiryRepository.ChangeOrderStatus(id, status)
	if err != nil {
		return models.Enquiry{}, err
	}

	return changedEnquiry, nil
}

func (s *enquiryService) UpdateEnquiry(ctx *gin.Context, req models.Enquiry) (models.Enquiry, error) {
	updatedEnquiry, err := s.enquiryRepository.UpdateEnquiry(req)
	if err != nil {
		return models.Enquiry{}, err
	}

	return updatedEnquiry, nil
}

func (s *enquiryService) SendMessage(ctx *gin.Context, req models.Message) (models.Message, error) {
	// token := ""
	// if req.Sender == models.SenderCustomer {
	var enquiry models.Enquiry
	enquiry, err := s.enquiryRepository.GetEnquiryByID(req.EnquiryID)
	if err != nil {
		return models.Message{}, err
	}

	token, err := s.storeService.GetFCMToken(ctx, &pb.StoreId{Id: string(enquiry.StoreID)})
	if err != nil {
		return models.Message{}, err
	}

	_, err = s.notificationService.SendData(ctx, &pb.Data{
		Data:  req.ToMap(),
		Token: token.Token,
	})
	if err != nil {
		return models.Message{}, err
	}

	return req, nil
}
