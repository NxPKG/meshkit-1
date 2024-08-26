package entity

import "github.com/khulnasoft/meshkit/database"

type EntityStatus string

const (
	Ignored   EntityStatus = "ignored"
	Enabled   EntityStatus = "enabled"
	Duplicate EntityStatus = "duplicate"
)

type Status interface {
	UpdateStatus(db *database.Handler, status EntityStatus) error
}
