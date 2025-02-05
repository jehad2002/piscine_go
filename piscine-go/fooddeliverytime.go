package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	time := food{}
	if order == "burger" {
		time.preptime = 15
	} else if order == "chips" {
		time.preptime = 10
	} else if order == "nuggets" {
		time.preptime = 12
	} else {
		time.preptime = 404
	}
	return time.preptime
}
