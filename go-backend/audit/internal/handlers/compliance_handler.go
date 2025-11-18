package handlers

import (
	"net/http"
	"time"

	"github.com/cryptofortress/backend/audit/internal/services"
	"github.com/gin-gonic/gin"
)

// ComplianceHandler handles compliance HTTP requests
type ComplianceHandler struct {
	complianceService services.ComplianceService
}

// NewComplianceHandler creates a new compliance handler
func NewComplianceHandler(complianceService services.ComplianceService) *ComplianceHandler {
	return &ComplianceHandler{
		complianceService: complianceService,
	}
}

// GenerateReportRequest represents the compliance report generation request payload
type GenerateReportRequest struct {
	Standard string    `json:"standard" binding:"required"`
	Period   time.Time `json:"period" binding:"required"`
}

// GenerateReportResponse represents the compliance report generation response payload
type GenerateReportResponse struct {
	ReportID  string `json:"report_id"`
	Standard  string `json:"standard"`
	Generated string `json:"generated"`
	Status    string `json:"status"`
}

// GenerateReport handles compliance report generation requests
func (h *ComplianceHandler) GenerateReport(c *gin.Context) {
	var req GenerateReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate report
	report, err := h.complianceService.GenerateReport(req.Standard, req.Period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate compliance report"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, GenerateReportResponse{
		ReportID:  report.ID,
		Standard:  report.Standard,
		Generated: report.Generated.Format(time.RFC3339),
		Status:    report.Status,
	})
}

// GetComplianceStatusRequest represents the compliance status retrieval request payload
type GetComplianceStatusRequest struct {
	Standard string `json:"standard" binding:"required"`
}

// GetComplianceStatusResponse represents the compliance status retrieval response payload
type GetComplianceStatusResponse struct {
	Standard    string   `json:"standard"`
	Compliant   bool     `json:"compliant"`
	LastChecked string   `json:"last_checked"`
	Issues      []string `json:"issues,omitempty"`
	NextAudit   string   `json:"next_audit"`
}

// GetComplianceStatus handles compliance status retrieval requests
func (h *ComplianceHandler) GetComplianceStatus(c *gin.Context) {
	var req GetComplianceStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get compliance status
	status, err := h.complianceService.GetComplianceStatus(req.Standard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get compliance status"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, GetComplianceStatusResponse{
		Standard:    status.Standard,
		Compliant:   status.Compliant,
		LastChecked: status.LastChecked.Format(time.RFC3339),
		Issues:      status.Issues,
		NextAudit:   status.NextAudit.Format(time.RFC3339),
	})
}

// ListComplianceStandardsResponse represents the compliance standards listing response payload
type ListComplianceStandardsResponse struct {
	Standards []string `json:"standards"`
}

// ListComplianceStandards handles compliance standards listing requests
func (h *ComplianceHandler) ListComplianceStandards(c *gin.Context) {
	// List compliance standards
	standards, err := h.complianceService.ListComplianceStandards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list compliance standards"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, ListComplianceStandardsResponse{
		Standards: standards,
	})
}

// HandleDataSubjectRequestRequest represents the data subject request handling request payload
type HandleDataSubjectRequestRequest struct {
	RequestType string `json:"request_type" binding:"required"`
	RequesterID string `json:"requester_id" binding:"required"`
	Data        string `json:"data,omitempty"`
}

// HandleDataSubjectRequestResponse represents the data subject request handling response payload
type HandleDataSubjectRequestResponse struct {
	RequestID   string `json:"request_id"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	ProcessedAt string `json:"processed_at,omitempty"`
}

// HandleDataSubjectRequest handles data subject request handling requests
func (h *ComplianceHandler) HandleDataSubjectRequest(c *gin.Context) {
	var req HandleDataSubjectRequestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create data subject request
	dsRequest := services.DataSubjectRequest{
		RequestType: req.RequestType,
		RequesterID: req.RequesterID,
		RequestedAt: time.Now(),
		Status:      "pending",
		Data:        req.Data,
	}

	// Handle data subject request
	err := h.complianceService.HandleDataSubjectRequest(dsRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to handle data subject request"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, HandleDataSubjectRequestResponse{
		RequestID:   dsRequest.ID,
		Message:     "Data subject request handled successfully",
		Status:      dsRequest.Status,
		ProcessedAt: time.Now().Format(time.RFC3339),
	})
}

// GenerateDataInventoryResponse represents the data inventory generation response payload
type GenerateDataInventoryResponse struct {
	InventoryID string `json:"inventory_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Message     string `json:"message"`
}

// GenerateDataInventory handles data inventory generation requests
func (h *ComplianceHandler) GenerateDataInventory(c *gin.Context) {
	// Generate data inventory
	inventory, err := h.complianceService.GenerateDataInventory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate data inventory"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, GenerateDataInventoryResponse{
		InventoryID: inventory.ID,
		CreatedAt:   inventory.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   inventory.UpdatedAt.Format(time.RFC3339),
		Message:     "Data inventory generated successfully",
	})
}