package main

import (
	"encoding/json"
	"fmt"
)

type rupiah int

var (
	exchangeRupiah = []rupiah{
		100000,
		50000,
		20000,
		10000,
		5000,
		2000,
		1000,
		500,
		200,
		100,
	}
	mapKey = map[rupiah]string{
		100000: "100000",
		50000:  "50000",
		20000:  "20000",
		10000:  "10000",
		5000:   "5000",
		2000:   "2000",
		1000:   "1000",
		500:    "500",
		200:    "200",
		100:    "100",
	}
)

func moneyExchange(currency rupiah) ([]byte, error) {
	result := map[string]int{}
	var key string
	for _, v := range exchangeRupiah {
		if currency-v >= 0 {
			key = mapKey[v]
			for {
				if currency-v >= 0 {
					currency -= v
					result[fmt.Sprintf("RP. %s", key)] = result[fmt.Sprintf("RP. %s", key)] + 1
				} else {
					break
				}
			}
		}
		if currency > 0 && currency < 100 {
			result["RP. 100"] = result["RP. 100"] + 1
			break
		}
	}

	return json.Marshal(result)
}

func main() {
	result, err := moneyExchange(450000)
	if err != nil {
		fmt.Println("Opps, something wrong")
	}
	fmt.Println(string(result))

}
