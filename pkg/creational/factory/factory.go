package factory

import "errors"

import "fmt"

type PaymentMethod interface {
	Pay(amount float32) string
}

type PayMentType int

const (
	Card PayMentType = 1
	Cash             = 2
)

type CardpaymentMethod struct {
}

//Pay card pay method
func (cp *CardpaymentMethod) Pay(amount float32) string {

	return fmt.Sprintf("%0.2f paid using card...", amount)
}

type CashpaymentMethod struct {
}

//Pay cash pay method
func (cp *CashpaymentMethod) Pay(amount float32) string {

	return fmt.Sprintf("%0.2f paid using cash...", amount)
}

//GetPaymentMethod Factory method used to get payment
func GetPaymentMethod(paymentType PayMentType) (PaymentMethod, error) {
	switch paymentType {
	case Card:
		return new(CardpaymentMethod), nil
	case Cash:
		return new(CashpaymentMethod), nil
	default:
		return nil, errors.New("No payment type for this pay type")
	}

}
