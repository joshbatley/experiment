package clients

import "time"

const (
	APIKeyHeader = "x-api-key"
	APIKeyTemp   = "JBSWY3DPEHPK3PXP"
)

type VaultClient struct {
	service string
}

func NewVaultClient(service string) *VaultClient {
	return &VaultClient{
		service,
	}
}

type APIKey struct {
	ID          int
	Key         string
	Description string
	AccessLevel string
	CreatedAt   time.Time
	ExpiresAt   *time.Time
	Revoked     bool
	Owner       int // foreign key reference to the user or system record
}

func (key *APIKey) IsExpired() bool {
	if key.ExpiresAt == nil {
		return false
	}
	return key.ExpiresAt.Before(time.Now())
}

func (key *APIKey) IsActive() bool {
	return !key.Revoked && (key.ExpiresAt == nil || !key.IsExpired())
}

func (key *APIKey) Revoke() {
	key.Revoked = true
}

func (key *APIKey) IsValidAccessLevel(level string) bool {
	return key.AccessLevel == level
}

type APIKeyStore interface {
	CreateAPIKey(key *APIKey) error
	GetAPIKeyByID(id int) (*APIKey, error)
	GetAPIKeyByValue(value string) (*APIKey, error)
}
