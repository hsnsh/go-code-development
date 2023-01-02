package calc

type Calculator struct {
	balance int
}

func (c *Calculator) Add(a int, b int) int {
	c.balance = a + b
	return c.balance
}

func (c *Calculator) Sub(a int, b int) int {
	c.balance = a - b
	return c.balance
}
