package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("00000000000000000", "John Doe", 12, 2024, 123)
	assert.Equal(t, "Invalid credit card number", err.Error())

	_, err = NewCreditCard("4193523830170205", "John Doe", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "John Doe", 13, 2024, 123)
	assert.Equal(t, "Invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "John Doe", 00, 2024, 123)
	assert.Equal(t, "Invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "John Doe", 11, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)

	_, err := NewCreditCard("4193523830170205", "John Doe", 11, lastYear.Year(), 123)
	assert.Equal(t, "Invalid expiration year", err.Error())
}
