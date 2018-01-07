package main

import (
	"fmt"
)

func main() {
	data, err := DataRead()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d\n", SumTotal(data))
}

func SumTotal(list []int64) int64 {
	var totalSum int64 = 0

	for _, v := range list {
		totalSum += v
	}

	return totalSum
}

func DataRead() ([]int64, error) {
	var length int64

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	data := make([]int64, length)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}
