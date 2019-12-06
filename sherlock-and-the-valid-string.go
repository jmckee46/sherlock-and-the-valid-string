package main

func isValid(s string) string {
	sCount := make([]int32, 26)
	sCountFrontLoaded := make([]int32, 26)

	// count frequencies
	for _, value := range s {
		sCount[int(value)-97]++
	}

	// filter out 0's
	var j int32
	for _, value := range sCount {
		if value != 0 {
			sCountFrontLoaded[j] = value
			j++
		}
	}
	sCountFiltered := sCountFrontLoaded[:j]

	// determin if valid
	var matchFails int32
	var base = sCountFiltered[0]
	var baseConfirmed bool

	for i := int32(1); i < j; i++ {

		if base != sCountFiltered[i] {
			if absValue(base-sCountFiltered[i]) > 1 {
				if sCountFiltered[i] == int32(1) {
					matchFails++
				} else {
					matchFails = 2
				}
			} else if !baseConfirmed && base-1 == sCountFiltered[i] {
				base = sCountFiltered[i]
				matchFails++
			} else if base == sCountFiltered[i]-1 {
				matchFails++
			} else {
				matchFails++
			}
		} else {
			baseConfirmed = true
		}

	}
	if matchFails > 1 {
		return "NO"
	}

	return "YES"
}

func absValue(x int32) int32 {
	if x > 0 {
		return x
	}

	return -x
}

func main() {}
