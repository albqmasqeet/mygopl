package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/albqmasqeet/mygopl/src/distconv"
	"github.com/albqmasqeet/mygopl/src/tempconv"
	"github.com/albqmasqeet/mygopl/src/weightconv"
)

func printMeasurement(s string) {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("%v isn't a number.", s)
	}

	// convert number to temperatures
	f := tempconv.Fahrenheit(num)
	c := tempconv.Celsius(num)

	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c))

	// convert number to distances
	mtr := distconv.Meter(num)
	ft := distconv.Feet(num)

	fmt.Printf("%s = %s, %s = %s\n", mtr, distconv.MtoF(mtr), ft, distconv.FtoM(ft))

	// convert number to weights
	kgs := weightconv.KiloGram(num)
	lbs := weightconv.Pound(num)

	fmt.Printf("%s = %s, %s = %s\n", kgs, weightconv.KtoL(kgs), lbs, weightconv.LtoK(lbs))

}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			printMeasurement(arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			printMeasurement(scan.Text())
		}
	}
}
