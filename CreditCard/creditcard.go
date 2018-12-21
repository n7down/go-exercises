package payment

type CreditCard struct {
	firstName       string
	lastName        string
	cardNumber      int
	expirationMonth int
	expirationYear  int
	balance         float64
}

func CreateCreditCard(firstName string, lastName string, cardNumber int, expirationMonth int, expirationYear int, balance float64) *CreditCard {
	return &CreditCard{
		firstName:       firstName,
		lastName:        lastName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		balance:         balance,
	}
}

func (c *CreditCard) SetFirstName(firstName string) {
	c.firstName = firstName
}

func (c CreditCard) FirstName() string {
	return c.firstName
}

func (c *CreditCard) SetLastName(lastName string) {
	c.lastName = lastName
}

func (c CreditCard) LastName() string {
	return c.lastName
}

func (c *CreditCard) SetCardNumber(cardNumber int) {
	c.cardNumber = cardNumber
}

func (c CreditCard) CardNumber() int {
	return c.cardNumberd
}

func (c *CreditCard) SetExpirationDate(month, year int) {
	c.expirationMonth = month
	c.expirationYear = year
}

func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

func (c *CreditCard) AddBalance(d float64) {
	c.balance = c.balance + d
}
