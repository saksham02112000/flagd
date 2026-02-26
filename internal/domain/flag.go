package domain

import (
	"errors"
	"time"
)

type Environment struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`  // used in API calls e.g. "production"
	Color     string    `json:"color"` // hex color for UI badge e.g. "#ef4444"
	CreatedAt time.Time `json:"created_at"`
}


type Flag struct {
	ID          string     `json:"id"`
	Key         string     `json:"key"`         // used in SDK e.g. "new-checkout"
	Name        string     `json:"name"`        // display name e.g. "New Checkout Flow"
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	ArchivedAt  *time.Time `json:"archived_at,omitempty"` // nil = active

	// Populated when fetching flag detail — not always present
	Environments []FlagEnvironment `json:"environments,omitempty"`
	Rules        []Rule            `json:"rules,omitempty"`
}

// IsActive returns true if the flag has not been archived
func (f *Flag) IsActive() bool {
	return f.ArchivedAt == nil
}

type FlagEnvironment struct {
	FlagID          string    `json:"flag_id"`
	EnvironmentID   string    `json:"environment_id"`
	EnvironmentSlug string    `json:"environment_slug"` // joined from environments table
	Enabled         bool      `json:"enabled"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Rule struct {
	ID            string     `json:"id"`
	FlagID        string     `json:"flag_id"`
	EnvironmentID string     `json:"environment_id"`
	Name          string     `json:"name"`
	Type          RuleType   `json:"type"`
	Priority      int        `json:"priority"`
	ServeValue    bool       `json:"serve_value"` // value to return when this rule matches
	Config        RuleConfig `json:"config"`
	Enabled       bool       `json:"enabled"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// RuleType defines what kind of targeting this rule does
type RuleType string

const (
	RuleTypePercentage RuleType = "percentage" // serve to X% of users
	RuleTypeUserIDs    RuleType = "user_ids"   // serve to specific user IDs
	RuleTypeAttribute  RuleType = "attribute"  // serve based on user attribute
	RuleTypeDefault    RuleType = "default"    // catch-all, always last
)

type RuleConfig struct {
	Percentage int `json:"percentage,omitempty"`

	IDs []string `json:"ids,omitempty"`

	Key      string       `json:"key,omitempty"`
	Operator RuleOperator `json:"operator,omitempty"`
	Value    string       `json:"value,omitempty"`
}

// RuleOperator defines how to compare attribute values
type RuleOperator string

const (
	OperatorEq       RuleOperator = "eq"       // equals
	OperatorNeq      RuleOperator = "neq"      // not equals
	OperatorContains RuleOperator = "contains"  // string contains
	OperatorIn       RuleOperator = "in"        // value in list
	OperatorNotIn    RuleOperator = "not_in"    // value not in list
	OperatorGt       RuleOperator = "gt"        // greater than (numeric string)
	OperatorLt       RuleOperator = "lt"        // less than (numeric string)
)

type APIKey struct {
	ID            string     `json:"id"`
	EnvironmentID string     `json:"environment_id"`
	Name          string     `json:"name"`
	KeyPrefix     string     `json:"key_prefix"` // first 12 chars shown in UI
	KeyHash       string     `json:"-"`          // never serialized to JSON
	CreatedAt     time.Time  `json:"created_at"`
	LastUsedAt    *time.Time `json:"last_used_at,omitempty"`
	DeletedAt     *time.Time `json:"-"`
}


type UserContext struct {
	ID         string            `json:"id"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

var (
	ErrFlagNotFound    = errors.New("flag not found")
	ErrFlagKeyExists   = errors.New("flag key already exists")
	ErrFlagArchived    = errors.New("flag is archived")
	ErrEnvNotFound     = errors.New("environment not found")
	ErrRuleNotFound    = errors.New("rule not found")
	ErrInvalidRuleType = errors.New("invalid rule type")
)