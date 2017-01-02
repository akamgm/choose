package choose

// Choose takes a slice s of unique objects and a positive integer k 
// and returns all unique unordered subsets of s of length k. This can
// be a potentially enormous set - be careful.
func Choose(s []interface{}, k int) [][]interface{} {
	if k >= len(s) {
		return [][]interface{}{s}
	} else if k <= 0 {
		return [][]interface{}{}
	}

	// initialize index vector
	idx := make([]int, k+1)
	for i:=0; i < k; i++ {
		idx[i] = i
	}
	idx[k] = len(s)	// terminating condition

	var result [][]interface{}

	level := 0
	for level < k {
		if idx[level] < idx[level+1] {
			// emit elements
			chosen := make([]interface{}, k)
			for i, e := range idx[:len(idx)-1] {
				chosen[i] = s[e]
			}
			result = append(result, []interface{}(chosen))

			idx[level] += 1
		} else {
			idx[level] = level
			level += 1
			idx[level] += 1
			if (level < k) && (idx[level] < idx[level+1]) {
				level = 0
			}
		}
	}

	return result
}

// Convenience function to Choose that handles string conversion.
// As with the underlying Choose func, each rune in string s should be unique. 
func ChooseString(s string, k int) []string {
	runes := []rune(s)
	var interfaceSlice []interface{} = make([]interface{}, len(runes))
	for i, d := range runes {
	    interfaceSlice[i] = d
	}

	gen := Choose(interfaceSlice, k)	

	result := make([]string, len(gen))
	for i, d := range gen {
		r := make([]rune, len(d))
		for n, b := range d {
			r[n] = b.(rune)
		}
	    result[i] = string(r)
	}

	return result
}

