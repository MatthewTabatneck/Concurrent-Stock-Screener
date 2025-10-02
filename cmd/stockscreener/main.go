package main

import (
	"fmt"
)

func main() {
	tickers, err := loadtickersCSV("sp500_tickers.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tickers)

}
