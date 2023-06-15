package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	tab := []float64{}
	inf := 0
	sup := 0
	reader := bufio.NewScanner(os.Stdin)

	//*******************************************************************************

	for reader.Scan() {
		nbr, _ := strconv.Atoi(reader.Text())
		moy := int(math.Round(Moyenne(tab)))

		//*******************************

		if int(math.Abs(float64(nbr))) > moy+10000 {
			if len(tab) > 0 {
				nbr = moy
			}
		}
		tab = append(tab, float64(nbr))

		//*******************************

		if len(tab) <= 10 {
			inf = int(nbr) - 100
			sup = int(nbr) + 100
		} else {
			average := Moyenne(tab)
			std := Deviation(tab, average)
			inf = int(math.Round(average - 1.28*std))
			sup = int(math.Round(average + 1.28*std))
		}
		fmt.Printf("%d %d\n", inf, sup)
	}
}

func Variance(tab []float64, moyenne float64) int {
	var total float64
	for _, v := range tab {
		total += math.Pow(v-moyenne, 2)
	}
	return int(math.Round(total / float64(len(tab))))
}

func Deviation(tab []float64, moyenne float64) float64 {
	return math.Sqrt(float64(Variance(tab, moyenne)))
}

func Moyenne(tab []float64) float64 {
	var total float64
	for _, v := range tab {
		total += v
	}
	return total / float64(len(tab))
}
