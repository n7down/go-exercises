package potterkata

const (
	bookPrice = 8.0
)

// TODO: potterkata - http://codingdojo.org/kata/Potter/
func calculateDiscount(firstBookNum int, secondBookNum int, thirdBookNum int, fourthBookNum int, fifthBookNum int) float64 {
	return 0.0
}

func CalculatePotterBookPrice(firstBookNum int, secondBookNum int, thirdBookNum int, fourthBookNum int, fifthBookNum int) float64 {
	var (
		firstBookPrice  float64 = 0.0
		secondBookPrice float64 = 0.0
		thirdBookPrice  float64 = 0.0
		fourthBookPrice float64 = 0.0
		fifthBookPrice  float64 = 0.0
	)

	if firstBookNum > 0 {
		firstBookPrice = bookPrice * float64(firstBookNum)
	}

	if secondBookNum > 0 {
		secondBookPrice = bookPrice * float64(secondBookNum)
	}

	if thirdBookNum > 0 {
		thirdBookPrice = bookPrice * float64(thirdBookNum)
	}

	if fourthBookNum > 0 {
		fourthBookPrice = bookPrice * float64(fourthBookNum)
	}

	if fifthBookNum > 0 {
		fifthBookPrice = bookPrice * float64(fifthBookNum)
	}

	discount := calculateDiscount(firstBookNum, secondBookNum, thirdBookNum, fourthBookNum, fifthBookNum)

	total := firstBookPrice + secondBookPrice + thirdBookPrice + fourthBookPrice + fifthBookPrice - discount

	return total
}
