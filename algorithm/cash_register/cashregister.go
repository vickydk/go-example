package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	m := make(map[string]float64)
	m["PENNY"] = 0.01
	m["NICKEL"] = 0.05
	m["DIME"] = 0.10
	m["QUARTER"] = 0.25
	m["HALF DOLLAR"] = 0.50
	m["ONE"] = 1.00
	m["TWO"] = 2.00
	m["FIVE"] = 5.00
	m["TEN"] = 20.00
	m["TWENTY"] = 20.00
	m["FIFTY"] = 50.00
	m["ONE HONDRED"] = 100.00
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		pp, _ := strconv.ParseFloat(s[0], 64)
		ch, _ := strconv.ParseFloat(s[1], 64)
		if pp == ch {
			fmt.Println("ZERO")
		} else if ch < pp {
			fmt.Println("ERROR")
		} else {
			value := math.Round((ch-pp)*100) / 100
			var res string
			for {
				var tempvalue float64
				var tempName string
				for k, v := range m {
					if value >= v && tempvalue < v {
						tempvalue = v
						tempName = k
					}
				}
				value = math.Round((value-tempvalue)*100) / 100
				res = res + tempName
				if math.Round(value*100)/100 <= 0 {
					break
				}
				res = res + ","
			}
			fmt.Println(res)
		}
	}
}
