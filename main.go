package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"log"
)

func main() {
	if len(os.Args) == 2 {
		array := returnFileinArray(os.Args[1])
		value1,Correlation := CorrelationCoefficent(array)
		value2 := a(os.Args[1])
		fmt.Printf("Linear Regression Line: y = %.6fx + %.6f", value1, value2)
		fmt.Printf("\nCorrelation Coefficent: %.10f\n", Correlation)
	}
}

func a(fileName string) float64 {
	var arr []int
	sum := 0.0
	xsum := 0.0
	XYNum := 0.0
	XXsum := 0.0
	YYsum := 0.0
	readFile, err := os.Open(fileName)
	defer readFile.Close()
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		ha, e := strconv.Atoi(fileScanner.Text())
		if e != nil {
			fmt.Printf("%T \n %v", ha, ha)
		}
		arr = append(arr, ha)
		sum = sum + float64(ha)
		xsum = xsum + float64(len(arr)-1)
		XYNum = XYNum + (float64(ha) * float64(len(arr)-1))
		XXsum = XXsum + (float64(len(arr)-1) * float64(len(arr)-1))
		YYsum = YYsum + (float64(ha) * float64(ha))
	}
	d := math.Round((((sum*XXsum)-(xsum*XYNum))/((float64(len(arr))*XXsum)-(xsum*xsum)))*1000000) / 1000000
	return d
}

func returnFileinArray(fileName string) []int {
	ReadFile, error := os.Open(fileName)
	if error != (nil) {
		printError(error)
	}
	FileScanner := bufio.NewScanner(ReadFile)
	var numbers []int
	for FileScanner.Scan() {
		if s, err := strconv.Atoi(FileScanner.Text()); err == nil {
			numbers = append(numbers, s)
		}
	}
	ReadFile.Close()
	return numbers
}




func CorrelationCoefficent(i []int) (float64,float64) {
	theX := 0.0
	theY := 0.0
	theXY := 0.0
	theXX := 0.0
	theYY := 0.0
	SXY := 0.0
	SXX := 0.0
	for x := 0; x < len(i); x++ {
		theXY += float64((x + 1) * i[x])
		theX += float64(x + 1)
		theY += float64(i[x])
		theXX += float64((x + 1) * (x + 1))
		theYY += float64(i[x] * i[x])
	}
	r := ((float64(len(i)) * theXY) - float64(theX*theY)) / math.Sqrt(((float64(len(i))*theXX)-(theX*theX))*((float64(len(i))*theYY)-(theY*theY)))
	SXY = theXY - ((theY * theX) / float64(len(i)))
	SXX = theXX - ((theX * theX) / float64(len(i)))
	a := SXY / SXX
	return a,r
}

func printError(err error) {
	fmt.Println("ERROR: " + err.Error())
	os.Exit(1)
}
