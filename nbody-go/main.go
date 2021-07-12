package main

import (
	"github.com/teodoranedic/ntp2021/parallel"
	"github.com/teodoranedic/ntp2021/sequential"
)

func main() {
	parallel.Parallel()
	sequential.Sequential()
}
