package domain

import "time"

type Role string

const (
	COOK      Role = "COOK"
	ATTENDANT Role = "ATTENDANT"
	MANAGER   Role = "MANAGER"
)

type Staff struct {
	ID        uint64
	Name      string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
