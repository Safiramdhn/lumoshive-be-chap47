package service

import (
	"golang-chap47/models"
	"golang-chap47/repository"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
}

type productService struct {
	repo repository.Repository
}

func NewProductService(repo repository.Repository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.Product.GetAll()
}
