package potterkata

const (
	bookPrice = 8.0
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func calculateDiscount(firstBookNum, secondBookNum, thirdBookNum, fourthBookNum, fifthBookNum int) float64 {
	const (
		twoBookDiscount   = 0.05
		threeBookDiscount = 0.1
		fourBookDiscount  = 0.2
		fiveBookDiscount  = 0.25
	)

	var (
		bookCount     int     = 0
		totalDiscount float64 = 0.0
	)

	// FIXME: get the max number from each of the num
	// FIXME: iterate over that number of times to get the discount value
	maxNum := firstBookNum
	maxNum = max(maxNum, secondBookNum)
	maxNum = max(maxNum, thirdBookNum)
	maxNum = max(maxNum, fourthBookNum)
	maxNum = max(maxNum, fifthBookNum)

	for i := 0; i < maxNum; i++ {
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

		var discount float64 = 0.0
		if bookCount == 5 {
			discount = (5 * bookPrice * fiveBookDiscount)
		} else if bookCount == 4 {
			discount = (5 * bookPrice * fourBookDiscount)
		} else if bookCount == 3 {
			discount = (3 * bookPrice * threeBookDiscount)
		} else if bookCount == 2 {
			discount = (2 * bookPrice * twoBookDiscount)
		}

		totalDiscount += discount
	}

	return totalDiscount
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
