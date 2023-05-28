package config

import (
	"github.com/google/uuid"
)
func GenerateId() (id string){
	id = uuid.New().String()
	return id
}