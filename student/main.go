package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	var arr []float64
	ysum := 0.0
	xsum := 0.0
	XYNum := 0.0
	XXsum := 0.0
	YYsum := 0.0
	index := 0.0
	avr := 0.0
	sd := 0.0
	numOne := 0.0
	fileScanner := bufio.NewScanner(os.Stdin)
	for fileScanner.Scan() {
		num, e := strconv.ParseFloat(fileScanner.Text(), 64)
		if e != nil {
			fmt.Printf("%T \n %v", num, num)
		}
		if index == 0 {
			numOne = num
		}
		xsum = xsum + index
		XXsum = XXsum + (index * index)
		if num < 100 || num > 200 {
			arr = append(arr, ysum/(index+1))
			ysum = ysum + ysum/(index+1)
			XYNum = XYNum + (ysum/(index+1))*index
			YYsum = YYsum + (ysum/(index+1))*(ysum/(index+1))
		} else {
			arr = append(arr, num)
			ysum = ysum + num
			XYNum = XYNum + (num * index)
			YYsum = YYsum + (num * num)
		}
		index++
		b := ((XYNum - ((ysum * xsum) / index)) / (XXsum - ((xsum * xsum) / index)))
		a := (((ysum * XXsum) - (xsum * XYNum)) / ((index * XXsum) - (xsum * xsum)))
		rrr := 0.0

		calVal1 := b*index + a + index
		if index <= 3 {
			calVal1 = arr[int(index)-1]
		} else {
			rrr = arr[int(index)-1] - arr[int(index)-2]
		}
		calVal := num + rrr
		sort.Float64s(arr)
		avr = math.Round(ysum / float64(len(arr)))
		for j := 0; j < len(arr); j++ {
			sd += math.Pow((float64(arr[j]) - avr), 2)
		}
		sd = math.Round(math.Sqrt(sd / float64(len(arr))))
		used := 0.0
		low := 0
		up := 0
		used = (calVal + calVal1) / 2
		if numOne > 169 && numOne < 180 {
			low = int(math.Round(used - sd))
			up = int(math.Round(used + (sd)))
		} else if numOne > 159.0 {
			low = int(math.Round(used - 5))
			up = int(math.Round(used + (6)))
		} else if numOne > 120.0 {
			low = int(math.Round(used - 1))
			up = int(math.Round(used + (1)))
		} else {
			low = int(math.Round(calVal1 - 10))
			up = int(math.Round(calVal1 + (9)))
		}

		fmt.Printf("%d %d\n", low, up)
	}
}
