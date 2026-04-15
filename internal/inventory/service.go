package inventory

import (
	"errors"
	"strings"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List() ([]Item, error) {
	return s.repo.List()
}

func (s *Service) Find(id int64) (*Item, error) {
	return s.repo.Find(id)
}

func (s *Service) Create(payload Payload) (*Item, map[string]string, error) {
	if validationErrors := validate(payload); len(validationErrors) > 0 {
		return nil, validationErrors, nil
	}

	item, err := s.repo.Create(normalize(payload))
	return item, nil, err
}

func (s *Service) Update(id int64, payload Payload) (*Item, map[string]string, error) {
	if id <= 0 {
		return nil, nil, errors.New("invalid id")
	}
	if validationErrors := validate(payload); len(validationErrors) > 0 {
		return nil, validationErrors, nil
	}

	item, err := s.repo.Update(id, normalize(payload))
	return item, nil, err
}

func (s *Service) Delete(id int64) (bool, error) {
	if id <= 0 {
		return false, errors.New("invalid id")
	}

	return s.repo.Delete(id)
}

func validate(payload Payload) map[string]string {
	errors := map[string]string{}

	if strings.TrimSpace(payload.SKU) == "" {
		errors["sku"] = "field is required"
	}
	if strings.TrimSpace(payload.Name) == "" {
		errors["name"] = "field is required"
	}
	if strings.TrimSpace(payload.Category) == "" {
		errors["category"] = "field is required"
	}
	if strings.TrimSpace(payload.SupplierName) == "" {
		errors["supplier_name"] = "field is required"
	}
	if payload.PurchasePrice < 0 {
		errors["purchase_price"] = "must be greater than or equal to 0"
	}
	if payload.SellingPrice < 0 {
		errors["selling_price"] = "must be greater than or equal to 0"
	}
	if payload.Stock < 0 {
		errors["stock"] = "must be greater than or equal to 0"
	}
	if payload.ReorderLevel < 0 {
		errors["reorder_level"] = "must be greater than or equal to 0"
	}

	return errors
}

func normalize(payload Payload) Payload {
	payload.SKU = strings.TrimSpace(payload.SKU)
	payload.Name = strings.TrimSpace(payload.Name)
	payload.Category = strings.TrimSpace(payload.Category)
	payload.SupplierName = strings.TrimSpace(payload.SupplierName)

	return payload
}
