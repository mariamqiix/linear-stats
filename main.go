package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var array []int
	for scanner.Scan() {
		currentY, _ := strconv.Atoi(scanner.Text())
		array = append(array, currentY)
		value1 := m(array)
		value2 := value1*float32(len(array)) + float32(currentY)

		s := " y = " +strconv.Itoa(int(value1)) + "x +" +strconv.Itoa(int(value2))
		fmt.Println("Linear Regression Line =" + s)
		Correlation:= strconv.Itoa(int(CorrelationCoefficent(array)))
		fmt.Println("Correlation Coefficent =" + Correlation)

	}
}

func m(i []int) float32 {
	theM := (float32(i[len(i)-1]) - float32(i[0])/float32(len(i)))
	return theM
}

func CorrelationCoefficent(i []int) float64 {
	theX := 0.0
	theY := 0.0
	theXY := 0
	for x := 0; x < len(i); x++ {
		theX += math.Pow(float64((0 - x)), 2)
		theY += math.Pow(float64((i[0] - i[x])), 2)
		theXY += (i[0] - i[x]) * (0 - x)
	}
	r := (float64((theX * theY)) / math.Sqrt(float64(theXY)))
	return r
}
