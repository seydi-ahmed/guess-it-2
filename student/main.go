package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	tab := []int{}
	inf := 0
	sup := 0
	reader := bufio.NewScanner(os.Stdin)

	//*******************************************************************************

	x := 1
	for reader.Scan() {
		nbr, _ := strconv.Atoi(reader.Text())
		moy := int(math.Round(Moyenne(tab)))

		//*******************************

		if int(math.Abs(float64(nbr))) > moy+10000 {
			if len(tab) > 0 {
				nbr = moy
			}
		}
		tab = append(tab, nbr)

		//*******************************

		//if (calculatePearsonCoefficient(tab) < 1 && calculatePearsonCoefficient(tab) > 0.5) || (calculatePearsonCoefficient(tab) > -1 && calculatePearsonCoefficient(tab) < -0.5) {
		if len(tab) <= 10 {
			inf = int(nbr) - 100
			sup = int(nbr) + 100
		} else {
			a, b := calculateLinearRegressionLine(tab)
			inf = int(math.Round((a * float64(x)) - b))
			sup = int(math.Round((a * float64(x)) + b))
		}
		//}

		fmt.Printf("%d %d\n", inf, sup)

		x++
	}
}

func calculateLinearRegressionLine(data []int) (float64, float64) {
	n := len(data)
	sumX, sumY, sumXY, sumXSquare := 0, 0, 0, 0

	for i := 0; i < n; i++ {
		sumX += i
		sumY += data[i]
		sumXY += (i) * data[i]
		sumXSquare += (i) * (i)
	}

	slope := float64(n*sumXY-sumX*sumY) / float64(n*sumXSquare-sumX*sumX)
	// intercept := float64(sumY)/float64(n) - slope*float64(sumX)/float64(n)
	intercept := float64((float64(sumY) - slope*float64(sumX)) / float64(n))
	return slope, intercept
}

func calculatePearsonCoefficient(data []int) float64 {
	n := len(data)
	sumX, sumY, sumXY, sumXSquare, sumYSquare := 0, 0, 0, 0, 0

	for i := 0; i < n; i++ {
		sumX += i
		sumY += data[i]
		sumXY += (i) * data[i]
		sumXSquare += (i) * (i)
		sumYSquare += data[i] * data[i]
	}

	numerator := float64(n*sumXY - sumX*sumY)
	denominator := math.Sqrt(float64(n*sumXSquare-sumX*sumX) * float64(n*sumYSquare-sumY*sumY))

	return numerator / denominator
}

func Moyenne(tab []int) float64 {
	var total int
	for _, v := range tab {
		total += v
	}
	return float64(float64(total) / float64(len(tab)))
}
