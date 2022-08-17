package main

import "fmt"

func ex(inpA, inpB string) int {
	rang := len(inpB)
	i := 0
	fmt.Println(inpA, inpB)
	for {
		if rang == i {
			return 3
		}
		if inpA[i] != inpB[i] {
			new := fmt.Sprintf("%s%s", string(inpB[:i]), string(inpA[i:]))
			if new == inpA {
				return 1
			} else {
				return 2
			}
		}

		i++
	}
}

func editText(inpA, inpB string) bool {
	if inpA == inpB {
		return true
	}

	lengA := len(inpA)
	lengB := len(inpB)
	var total int

	// check 
	if lengA < lengB {
		total = ex(inpB, inpA)
	} else {
		total = ex(inpA, inpB)
	}

	if total == 3 && lengA%lengB == 1 {
		return true
	}

	if total > 1 {
		return false
	}

	return true
}

func main() {
	resp := editText("telkom", "telkm")

	fmt.Println(resp)
}
