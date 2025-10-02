package main

import (
	"fmt"
	"os"

	"github.com/MatthewTabatneck/Concurrent-Stock-Screener/internal/tickers"
)

func main() {
	file, err := os.Open("sp500_tickers.csv")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	tickers, err := tickers.LoadtickersCSV(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tickers)

}
