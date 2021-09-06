package domain

type CustomerRepositoryStub struct{
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll()([]Customer, error){
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub{
	customers :=  []Customer{
		{"1001","Ashish","New Delhi","110011","2000-10-10","1"},
		{"1002","Rob","New Delhi","110011","2000-10-10","1"},
	}

	return CustomerRepositoryStub{customers: customers}
}