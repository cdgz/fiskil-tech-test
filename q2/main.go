package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type reqStore struct {
	data map[string]bool
	mu   sync.RWMutex
}

func (store *reqStore) Set(key string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.data[key] = true
}

func (store *reqStore) Check(key string) bool {
	store.mu.RLock()
	defer store.mu.RUnlock()
	_, exists := store.data[key]
	return exists
}

func generateID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	id := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	return id, nil
}

func main() {
	rs := &reqStore{
		data: make(map[string]bool),
	}
	r := gin.Default()

	r.GET("/api/req", func(c *gin.Context) {
		id, err := generateID()
		if err != nil {
			fmt.Println("Error generating ID:", err)
			return
		}
		if exists := rs.Check(id); !exists {
			rs.Set(id)
		} else {
			for {
				id, err = generateID()
				if err != nil {
					fmt.Println("Error generating ID:", err)
					return
				}

				if exists := rs.Check(id); !exists {
					rs.Set(id)
					break
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{"generated_id": id})
	})

	r.GET("/api/list_ids", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ids": rs.data})
	})

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
