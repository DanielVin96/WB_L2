package main

import "fmt"

/*Паттерн стратегия позволяет определить семейство алгоритмов, инкапсулировать каждый из них и сделать их взаимозаменяемыми.
Это позволяет изменять алгоритмы независимо от клиентов, которые используют эти алгоритмы.

Пример:
У нас есть система оплаты, которая должна работать с различными способами оплаты: кредитной картой,
PayPal, банковским переводом и т.д. Для этого мы можем использовать паттерн стратегия.

 Сначала определим интерфейс "PaymentStrategy", который будет содержать метод "Pay":*/

type PaymentStrategy interface {
	Pay(amount float64) error
}

//Затем мы можем создать несколько типов данных, которые реализуют этот интерфейс, например:
type CreditCardStrategy struct {
	CardNumber string
	ExpiryDate string
	CVV        string
}

func (c *CreditCardStrategy) Pay(amount float64) error {
	// реализация оплаты кредитной картой
	return nil
}

type PayPalStrategy struct {
	Email    string
	Password string
}

func (p *PayPalStrategy) Pay(amount float64) error {
	// реализация оплаты через PayPal
	return nil
}

//Теперь мы можем создать клиента, который будет использовать эти стратегии оплаты:

type PaymentClient struct {
	Strategy PaymentStrategy
}

func (p *PaymentClient) ProcessPayment(amount float64) error {
	return p.Strategy.Pay(amount)
}

//Клиент может использовать любую стратегию оплаты, реализующую интерфейс "PaymentStrategy". Например:

func main() {
	creditCard := &CreditCardStrategy{
		CardNumber: "1234 5678 9012 3456",
		ExpiryDate: "12/22",
		CVV:        "123",
	}

	paymentClient := &PaymentClient{
		Strategy: creditCard,
	}

	err := paymentClient.ProcessPayment(100.0)
	if err != nil {
		fmt.Println(err)
	}

	payPal := &PayPalStrategy{
		Email:    "example@example.com",
		Password: "password",
	}

	paymentClient.Strategy = payPal

	err = paymentClient.ProcessPayment(50.0)
	if err != nil {
		fmt.Println(err)
	}
}
