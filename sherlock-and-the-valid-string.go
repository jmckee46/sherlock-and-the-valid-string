package main

import "fmt"

func isValid(s string) string {
	sCount := make([]int32, 26)
	sCountFrontLoaded := make([]int32, 26)

	for _, value := range s {
		sCount[int(value)-97]++
	}

	var j int32
	for _, value := range sCount {
		if value != 0 {
			sCountFrontLoaded[j] = value
			j++
		}
	}
	sCountFiltered := sCountFrontLoaded[:j]
	var matchFails int32
	var suspect1 int32
	var suspect2 int32
	var base int32

	// sCountFiltered: [2 2 1 1]

	fmt.Println("sCountFiltered:", sCountFiltered)
	for i := int32(1); i < j; i++ {
		fmt.Println("i:", i, "matchFails", matchFails)
		if matchFails == 0 {
			if sCountFiltered[i-1] != sCountFiltered[i] {
				fmt.Println("suspects collected")
				suspect1 = sCountFiltered[i-1]
				base = sCountFiltered[i-1]
				suspect2 = sCountFiltered[i]
				matchFails++
			}
		} else if matchFails == 1 {
			fmt.Println("suspect1:", suspect1, "suspect2", suspect2, "sCountFiltered[i]", sCountFiltered[i])
			if sCountFiltered[i] != suspect1 && sCountFiltered[i] != suspect2 {
				fmt.Println("nobody matches!")
				return "NO"
			} else if sCountFiltered[i] == suspect1 && sCountFiltered[i] != suspect2 {
				fmt.Println("suspect 2 is odd man")
				if suspect2-1 != sCountFiltered[i] {
					return "NO"
				}
			} else if sCountFiltered[i] != suspect1 && sCountFiltered[i] == suspect2 {
				fmt.Println("suspect 1 is odd man")
				if suspect1-1 != sCountFiltered[i] {
					return "NO"
				}
			}
			matchFails++
		} else if matchFails == 2 {
			if sCountFiltered[i-1] != sCountFiltered[i] {
				return "NO"
			}
		}

	}
	// fmt.Println(sCount)
	return "YES"
}

func main() {}
