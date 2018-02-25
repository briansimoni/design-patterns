// When using the factory method design pattern, we gain an extra
// layer of encapsulation so that our program can grow in a
// controlled environment. With the Factory method, we delegate the
// creation of families of objects to a different package or object
// to abstract us frothe knowledge of the pool of possible objects we could
// se. Imagine that you want to organize your holidays using a trip agency
// You don't deal with hotels and traveling and you just tell the agency
// the destination you are interested in so that they provide you with
// everything you need. The trip agency represents a Factory of trips.

package factory

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(CreditCardPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}

type CreditCardPM struct{}

func (d *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card (new) implementation\n", amount)
}
