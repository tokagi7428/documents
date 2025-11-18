package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

type Store struct {
	sync.RWMutex
	docs map[int]*DocumentItem
}

func NewStore() *Store {
	s := &Store{
		docs: make(map[int]*DocumentItem),
	}
	// mock data
	s.docs[1] = &DocumentItem{ID: 1, Name: "รายการที่ 1", Reason: "xxxxx", Status: "รออนุมัติ", UpdatedAt: time.Now()}
	s.docs[2] = &DocumentItem{ID: 2, Name: "รายการที่ 2", Reason: "approve", Status: "อนุมัติ", UpdatedAt: time.Now()}
	s.docs[3] = &DocumentItem{ID: 3, Name: "รายการที่ 3", Reason: "not approve", Status: "รออนุมัติ", UpdatedAt: time.Now()}
	s.docs[4] = &DocumentItem{ID: 4, Name: "รายการที่ 4", Reason: "xxxxx", Status: "รออนุมัติ", UpdatedAt: time.Now()}
	s.docs[5] = &DocumentItem{ID: 5, Name: "รายการที่ 5", Reason: "xxxxx", Status: "ไม่อนุมัติ", UpdatedAt: time.Now()}
	s.docs[6] = &DocumentItem{ID: 6, Name: "รายการที่ 6", Reason: "xxxxx", Status: "รออนุมัติ", UpdatedAt: time.Now()}
	s.docs[7] = &DocumentItem{ID: 7, Name: "รายการที่ 7", Reason: "approve", Status: "อนุมัติ", UpdatedAt: time.Now()}
	s.docs[8] = &DocumentItem{ID: 8, Name: "รายการที่ 8", Reason: "not approve", Status: "ไม่อนุมัติ", UpdatedAt: time.Now()}
	s.docs[9] = &DocumentItem{ID: 9, Name: "รายการที่ 9", Reason: "xxxxx", Status: "รออนุมัติ", UpdatedAt: time.Now()}
	s.docs[10] = &DocumentItem{ID: 10, Name: "รายการที่ 10", Reason: "xxxxx", Status: "รออนุมัติ", UpdatedAt: time.Now()}
	return s
}

func (s *Store) ListDocuments() []*DocumentItem {
	s.RLock()
	defer s.RUnlock()

	out := make([]*DocumentItem, 0, len(s.docs))
	for _, v := range s.docs {
		out = append(out, v)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].ID < out[j].ID
	})

	return out
}

func (s *Store) Approve(ids []int, reason string) {
	s.Lock()
	defer s.Unlock()
	for _, id := range ids {
		fmt.Println(id)
		fmt.Println(s.docs[id])
		if doc, ok := s.docs[id]; ok {
			doc.Status = "อนุมัติ"
			doc.Reason = reason
			doc.UpdatedAt = time.Now()
		}
	}
}

func (s *Store) Reject(ids []int, reason string) {
	s.Lock()
	defer s.Unlock()
	for _, id := range ids {
		if doc, ok := s.docs[id]; ok {
			fmt.Println(id)
			fmt.Println(s.docs[id])
			doc.Status = "ไม่อนุมัติ"
			doc.Reason = reason
			doc.UpdatedAt = time.Now()
		}
	}
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "http://127.0.0.1:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	store := NewStore()

	api := r.Group("/api")
	{
		api.GET("/documents", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": store.ListDocuments()})
		})

		api.POST("/documents/approve", func(c *gin.Context) {
			var req ActionRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
				return
			}
			if len(req.IDs) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "no ids"})
				return
			}
			store.Approve(req.IDs, req.Reason)
			c.JSON(http.StatusOK, gin.H{"message": "approved"})
		})

		api.POST("/documents/reject", func(c *gin.Context) {
			var req ActionRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
				return
			}
			if len(req.IDs) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "no ids"})
				return
			}
			store.Reject(req.IDs, req.Reason)
			c.JSON(http.StatusOK, gin.H{"message": "rejected"})
		})
	}

	addr := ":8080"
	log.Printf("Starting server at %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
