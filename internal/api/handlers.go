package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/enquiry/internal/models"
	"github.com/tanush-128/openzo_backend/enquiry/internal/service"
)

type Handler struct {
	enquiryService service.EnquiryService
}

func NewHandler(enquiryService *service.EnquiryService) *Handler {
	return &Handler{enquiryService: *enquiryService}
}

func (h *Handler) CreateEnquiry(ctx *gin.Context) {
	var enquiry models.Enquiry
	if err := ctx.BindJSON(&enquiry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdEnquiry, err := h.enquiryService.CreateEnquiry(ctx, enquiry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdEnquiry)

}

func (h *Handler) GetEnquiryByID(ctx *gin.Context) {
	id := ctx.Param("id")

	enquiry, err := h.enquiryService.GetEnquiryByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, enquiry)
}

func (h *Handler) GetEnquirysByStoreID(ctx *gin.Context) {
	store_id := ctx.Param("store_id")

	enquirys, err := h.enquiryService.GetEnquirysByStoreID(ctx, store_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, enquirys)
}

func (h *Handler) ChangeOrderStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	status := ctx.Query("status")

	updatedEnquiry, err := h.enquiryService.ChangeEnquiryStatus(ctx, id, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedEnquiry)
}

func (h *Handler) UpdateEnquiry(ctx *gin.Context) {
	var enquiry models.Enquiry
	if err := ctx.BindJSON(&enquiry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEnquiry, err := h.enquiryService.UpdateEnquiry(ctx, enquiry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedEnquiry)
}

func (h *Handler) SendMessage(ctx *gin.Context) {
	var message models.Message
	if err := ctx.BindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sentMessage, err := h.enquiryService.SendMessage(ctx, message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sentMessage)
}
