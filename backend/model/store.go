package model

import (
	"time"
)

// DocumentItem represents a document record
type DocumentItem struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Reason    string    `json:"reason"`
	Status    string    `json:"status"` // "รออนุมัติ" | "อนุมัติ" | "ไม่อนุมัติ"
	UpdatedAt time.Time `json:"updatedAt"`
}

type ActionRequest struct {
	IDs    []int  `json:"ids"`
	Reason string `json:"reason"`
}
