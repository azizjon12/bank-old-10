package types

// Money represents amount of money in minimum units e.g. dirams, cents, pennies
type Money int64

// Category represents the category of the payment made e.g. restaurants, auto, pharmacy
type Category string

// Payment represents payment information
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
