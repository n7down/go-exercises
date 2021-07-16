package leftrotation

/*
 * Complete the 'rotLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER d
 */

func RotateLeft(a []int32, d int32) []int32 {
	if len(a) == 0 {
		return a
	}

	if len(a) == 1 {
		return a
	}

	// var r int32
	// for i := 0; i < int(d); i++ {
	first := a[0]
	rest := a[1:]

	rest = append(rest, first)
	// }

	return rest
}
