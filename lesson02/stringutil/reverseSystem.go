package stringutil

func reverseSystem(str string) string {

	rslt := []rune(str)

	for i, j := 0, len(rslt)-1; i < len(rslt)/2; i, j = i+1, j-1 {

		rslt[i], rslt[j] = rslt[j], rslt[i]
	}

	return string(rslt)

}
