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

	r := a
	for i := 0; i < int(d); i++ {
		first := r[0]
		r = r[1:]
		r = append(r, first)
	}

	return r
}
