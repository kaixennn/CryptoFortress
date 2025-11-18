package handlers

import (
	"net/http"

	"github.com/cryptofortress/backend/auth/internal/services"
	"github.com/gin-gonic/gin"
)

// RBACHandler handles role-based access control HTTP requests
type RBACHandler struct {
	rbacService services.RBACService
}

// NewRBACHandler creates a new RBAC handler
func NewRBACHandler(rbacService services.RBACService) *RBACHandler {
	return &RBACHandler{
		rbacService: rbacService,
	}
}

// CreateRoleRequest represents the create role request payload
type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// CreateRole handles creating a new role
func (h *RBACHandler) CreateRole(c *gin.Context) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create role
	err := h.rbacService.CreateRole(req.Name, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Role created successfully"})
}

// DeleteRoleRequest represents the delete role request payload
type DeleteRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

// DeleteRole handles deleting a role
func (h *RBACHandler) DeleteRole(c *gin.Context) {
	var req DeleteRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete role
	err := h.rbacService.DeleteRole(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete role"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}

// AssignRoleRequest represents the assign role request payload
type AssignRoleRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}

// AssignRole handles assigning a role to a user
func (h *RBACHandler) AssignRole(c *gin.Context) {
	var req AssignRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign role to user
	err := h.rbacService.AssignRoleToUser(req.UserID, req.RoleName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign role to user"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Role assigned to user successfully"})
}

// RemoveRoleRequest represents the remove role request payload
type RemoveRoleRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}

// RemoveRole handles removing a role from a user
func (h *RBACHandler) RemoveRole(c *gin.Context) {
	var req RemoveRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Remove role from user
	err := h.rbacService.RemoveRoleFromUser(req.UserID, req.RoleName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove role from user"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Role removed from user successfully"})
}

// CreatePermissionRequest represents the create permission request payload
type CreatePermissionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// CreatePermission handles creating a new permission
func (h *RBACHandler) CreatePermission(c *gin.Context) {
	var req CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create permission
	err := h.rbacService.CreatePermission(req.Name, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create permission"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Permission created successfully"})
}

// AssignPermissionRequest represents the assign permission request payload
type AssignPermissionRequest struct {
	RoleName       string `json:"role_name" binding:"required"`
	PermissionName string `json:"permission_name" binding:"required"`
}

// AssignPermission handles assigning a permission to a role
func (h *RBACHandler) AssignPermission(c *gin.Context) {
	var req AssignPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign permission to role
	err := h.rbacService.AssignPermissionToRole(req.RoleName, req.PermissionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign permission to role"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Permission assigned to role successfully"})
}

// RemovePermissionRequest represents the remove permission request payload
type RemovePermissionRequest struct {
	RoleName       string `json:"role_name" binding:"required"`
	PermissionName string `json:"permission_name" binding:"required"`
}

// RemovePermission handles removing a permission from a role
func (h *RBACHandler) RemovePermission(c *gin.Context) {
	var req RemovePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Remove permission from role
	err := h.rbacService.RemovePermissionFromRole(req.RoleName, req.PermissionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove permission from role"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Permission removed from role successfully"})
}

// CheckPermissionRequest represents the check permission request payload
type CheckPermissionRequest struct {
	UserID         string `json:"user_id" binding:"required"`
	PermissionName string `json:"permission_name" binding:"required"`
}

// CheckPermissionResponse represents the check permission response payload
type CheckPermissionResponse struct {
	HasPermission bool `json:"has_permission"`
}

// CheckPermission handles checking if a user has a specific permission
func (h *RBACHandler) CheckPermission(c *gin.Context) {
	var req CheckPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check permission
	hasPermission, err := h.rbacService.CheckPermission(req.UserID, req.PermissionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check permission"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, CheckPermissionResponse{
		HasPermission: hasPermission,
	})
}

// GetUserRolesRequest represents the get user roles request payload
type GetUserRolesRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

// GetUserRolesResponse represents the get user roles response payload
type GetUserRolesResponse struct {
	Roles []string `json:"roles"`
}

// GetUserRoles handles retrieving all roles assigned to a user
func (h *RBACHandler) GetUserRoles(c *gin.Context) {
	var req GetUserRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user roles
	roles, err := h.rbacService.GetUserRoles(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user roles"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, GetUserRolesResponse{
		Roles: roles,
	})
}

// GetRolePermissionsRequest represents the get role permissions request payload
type GetRolePermissionsRequest struct {
	RoleName string `json:"role_name" binding:"required"`
}

// GetRolePermissionsResponse represents the get role permissions response payload
type GetRolePermissionsResponse struct {
	Permissions []string `json:"permissions"`
}

// GetRolePermissions handles retrieving all permissions assigned to a role
func (h *RBACHandler) GetRolePermissions(c *gin.Context) {
	var req GetRolePermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get role permissions
	permissions, err := h.rbacService.GetRolePermissions(req.RoleName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get role permissions"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, GetRolePermissionsResponse{
		Permissions: permissions,
	})
}