package repository

import "assignment-2/domain"

type OrderRepository interface {
	Create(data domain.Order) error
	Delete(orderId int) error
	GetAll() ([]domain.Order, error)
	Update(orderId int, data domain.Order) error
}
