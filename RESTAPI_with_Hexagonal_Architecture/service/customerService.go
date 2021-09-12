package service

import (
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/domain"
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/dto"
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/errs"
)

type CustomerService interface{
	GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer()([]dto.CustomerResponse, *errs.AppError){
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}

	return response, nil
	
	// return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string)(*dto.CustomerResponse, *errs.AppError){
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repository}
}