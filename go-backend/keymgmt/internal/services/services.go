package services

import (
	"time"

	"github.com/cryptofortress/backend/keymgmt/internal/config"
)

// Services holds references to all key management services
type Services struct {
	Key         KeyService
	Rotation    RotationService
	Shamir      ShamirService
	Replication ReplicationService
}

// KeyService defines the interface for key management operations
type KeyService interface {
	// Key generation
	GenerateKey(algorithm string, size int) (keyID string, key []byte, err error)
	GenerateKeyPair(algorithm string, size int) (keyID string, privateKey, publicKey []byte, err error)
	
	// Key storage and retrieval
	StoreKey(keyID string, key []byte, metadata KeyMetadata) error
	RetrieveKey(keyID string) ([]byte, *KeyMetadata, error)
	DeleteKey(keyID string) error
	
	// Key metadata
	ListKeys() ([]KeyInfo, error)
	GetKeyMetadata(keyID string) (*KeyMetadata, error)
	UpdateKeyMetadata(keyID string, metadata KeyMetadata) error
}

// RotationService defines the interface for key rotation operations
type RotationService interface {
	// Key rotation
	RotateKey(keyID string) (newKeyID string, err error)
	ScheduleRotation(keyID string, interval time.Duration) error
	CancelRotation(keyID string) error
	GetRotationSchedule(keyID string) (*RotationSchedule, error)
	
	// Automatic rotation
	EnableAutoRotation(keyID string, period time.Duration) error
	DisableAutoRotation(keyID string) error
	ListAutoRotations() ([]RotationInfo, error)
}

// ShamirService defines the interface for Shamir's Secret Sharing operations
type ShamirService interface {
	// Secret sharing
	SplitSecret(secret []byte, threshold, shares int) ([][]byte, error)
	CombineShares(shares [][]byte) ([]byte, error)
	
	// Key distribution
	DistributeKey(keyID string, threshold, shares int, recipients []string) error
	RecoverKey(shares [][]byte) (keyID string, key []byte, err error)
}

// ReplicationService defines the interface for key replication operations
type ReplicationService interface {
	// Key replication
	ReplicateKey(keyID string, regions []string) error
	EnableCrossRegionReplication(keyID string, regions []string) error
	DisableCrossRegionReplication(keyID string) error
	
	// Disaster recovery
	BackupKey(keyID string, destination string) error
	RestoreKey(backupID string, destination string) error
	ListBackups() ([]BackupInfo, error)
}

// KeyMetadata represents metadata associated with a key
type KeyMetadata struct {
	Algorithm   string            `json:"algorithm"`
	Size        int               `json:"size"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	ExpiresAt   *time.Time        `json:"expires_at,omitempty"`
	RevokedAt   *time.Time        `json:"revoked_at,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Description string            `json:"description,omitempty"`
	Version     int               `json:"version"`
}

// KeyInfo represents basic information about a key
type KeyInfo struct {
	KeyID     string    `json:"key_id"`
	Algorithm string    `json:"algorithm"`
	Size      int       `json:"size"`
	CreatedAt time.Time `json:"created_at"`
}

// RotationSchedule represents a key rotation schedule
type RotationSchedule struct {
	KeyID      string        `json:"key_id"`
	Interval   time.Duration `json:"interval"`
	NextRotate time.Time     `json:"next_rotate"`
	Enabled    bool          `json:"enabled"`
}

// RotationInfo represents information about an automatic rotation
type RotationInfo struct {
	KeyID    string        `json:"key_id"`
	Period   time.Duration `json:"period"`
	LastRun  time.Time     `json:"last_run"`
	NextRun  time.Time     `json:"next_run"`
	Enabled  bool          `json:"enabled"`
}

// BackupInfo represents information about a key backup
type BackupInfo struct {
	BackupID    string    `json:"backup_id"`
	KeyID       string    `json:"key_id"`
	CreatedAt   time.Time `json:"created_at"`
	Destination string    `json:"destination"`
	Size        int64     `json:"size"`
}