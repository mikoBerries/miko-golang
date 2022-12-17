package main

import "github.com/MikoBerries/miko-golang/13_testing_benchmarking/02_testing/acdc"

func main() {
	acdc.Sum(10, 10, 10, 10)
	acdc.Sum(20, 10, 10, 10)
	acdc.Sum(30, 40, 10, 10)
}
