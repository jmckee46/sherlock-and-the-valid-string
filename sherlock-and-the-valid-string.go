package main

import "fmt"

func isValid(s string) string {
	sCount := make([]int32, 26)
	sCountFrontLoaded := make([]int32, 26)

	// count frequencies
	for _, value := range s {
		sCount[int(value)-97]++
	}
	fmt.Println("sCount:", sCount)

	// filter out 0's
	var j int32
	for _, value := range sCount {
		if value != 0 {
			sCountFrontLoaded[j] = value
			j++
		}
	}
	sCountFiltered := sCountFrontLoaded[:j]
	fmt.Println("sCountFiltered:", sCountFiltered)

	// determin if valid
	var matchFails int32
	var base = sCountFiltered[0]
	var baseConfirmed bool

	for i := int32(1); i < j; i++ {

		fmt.Println("")
		fmt.Println("baseConfirmed:", baseConfirmed)
		fmt.Println("base:", base)
		fmt.Println("sCountFiltered[i]:", sCountFiltered[i])
		fmt.Println("matchFails:", matchFails)

		if base != sCountFiltered[i] {
			if absValue(base-sCountFiltered[i]) > 1 {
				if sCountFiltered[i] == int32(1) {
					fmt.Println("special case = 1")
					matchFails++
				} else {
					fmt.Println("more than 1 apart")
					matchFails = 2
				}
				break
			} else if !baseConfirmed && base-1 == sCountFiltered[i] {
				base = sCountFiltered[i]
				matchFails++
			} else if base == sCountFiltered[i]-1 {
				fmt.Println("decrement one scenario")
				matchFails++
			} else {
				fmt.Println("incrementing matchFails")
				matchFails++
			}
		} else {
			fmt.Println("equal and base confirmed")
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
