package dtos

import (
	"time"
)

type CustomerResponse struct {
	ID              int64           `json:"id"`
	Firstname       string          `json:"firstname"`
	Lastname        string          `json:"lastname"`
	Genda           string          `json:"genda"`
	Email           string          `json:"email"`
	PhoneNumber     string          `json:"phone_number"`
	Status          string          `json:"status"`
	AddressResponse AddressResponse `json:"address"`
	CreatedAt       time.Time       `json:"created_date"`
	UpdatedAt       *time.Time      `json:"updated_date,omitempty"`
}
type AddressResponse struct {
	ID           int64   `json:"id"`
	StreetNumber *string `json:"street_number,omitempty"`
	StreetName   string  `json:"street_name"`
	ZipCode      string  `json:"zip_code"`
	City         string  `json:"city"`
	Region       string  `json:"region"`
	Country      string  `json:"country"`
	Complement   *string `json:"complement,omitempty"`
}

type LightCustomerResponse struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Genda       string `json:"genda"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}
