package potterkata

const (
	bookPrice = 8.0
)

func calculateDiscount(firstBookNum, secondBookNum, thirdBookNum, fourthBookNum, fifthBookNum int) float64 {
	const (
		twoBookDiscount   = 0.05
		threeBookDiscount = 0.1
	)

	var (
		bookCount int     = 0
		discount  float64 = 0.0
	)

	if firstBookNum > 0 {
		firstBookNum -= 1
		bookCount += 1
	}

	if secondBookNum > 0 {
		secondBookNum -= 1
		bookCount += 1
	}

	if thirdBookNum > 0 {
		thirdBookNum -= 1
		bookCount += 1
	}

	if fourthBookNum > 0 {
		fourthBookNum -= 1
		bookCount += 1
	}

	if fifthBookNum > 0 {
		fifthBookNum -= 1
		bookCount += 1
	}

	if bookCount >= 3 {
		// if contains 3 different books - (3 * bookPrice * threeBookDiscount)
		discount = (3 * bookPrice * threeBookDiscount)

	} else if bookCount == 2 {
		// if contains 2 different books - (2 * bookPrice * twoBookDiscount)
		discount = (2 * bookPrice * twoBookDiscount)
	}

	return discount
}

func CalculatePotterBookPrice(firstBookNum, secondBookNum, thirdBookNum, fourthBookNum, fifthBookNum int) float64 {
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
