package main


import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"bytes"
)

type Banker struct {
	N 			int 
	M 			int
	resource 	[]int
	available 	[]int 
	max 		[][]int
	allocation 	[][]int
	need 		[][]int
}

func handleError(error error) {
	if error != nil {
		log.Fatal(error)
		os.Exit(2)
	}
}

func get_input(splitLines [][]string) (Banker, []int, int) {
	
	N,err := strconv.Atoi(splitLines[0][0])
	handleError(err)

	M,err := strconv.Atoi(splitLines[0][1])
	handleError(err)
	
	resource := make([]int, M)
	for i := range splitLines[1] {
		resource[i], err = strconv.Atoi(splitLines[1][i])
		handleError(err)
	}

	available := make([]int, M)
	for i := range splitLines[2] {
		available[i], err = strconv.Atoi(splitLines[2][i])
		handleError(err)
	}

	max := make([][]int, N)
	allocation := make([][]int, N)
	need := make([][]int, N)
	for i := range max {
		max[i] = make([]int, M)
		allocation[i] = make([]int, M)
		need[i] = make([]int, M)
		for j := range max[i] {
			max[i][j], err = strconv.Atoi(splitLines[3 + i][j])
			handleError(err)
			allocation[i][j], err = strconv.Atoi(splitLines[3 + i + N][j])
			handleError(err)
			need[i][j] = max[i][j] - allocation[i][j]
		}
	}


	request := make([]int, M)
	lastThing := splitLines[len(splitLines)-1][0]
	initial, err := strconv.Atoi(strings.Split(lastThing, ":")[0])
	handleError(err)
	request[0], err = strconv.Atoi(strings.Split(lastThing, ":")[1])
	handleError(err)
	for i := range request {
		if i != 0 {
			request[i], err = strconv.Atoi(splitLines[3 + N * 2][i])
			handleError(err)
		}
	}

	return Banker{N, M, resource, available, max, allocation, need}, request , initial
}

func getHeader(m int) (string) {
	var buffer bytes.Buffer
	for i := 0; i < m; i++ {
		buffer.WriteString(string('A' + i) + " ")
	}
	return buffer.String()
}



func printBanker(banker Banker) {
	fmt.Println("The Resource Vector is...")
	fmt.Println(getHeader(banker.M))
	for i := 0; i < banker.M; i++ {
		fmt.Printf("%d ", banker.resource[i])
	}
	fmt.Println("\n")

	fmt.Println("The Available Vector is...")
	fmt.Println(getHeader(banker.M))
	for i := 0; i < banker.M; i++ {
		fmt.Printf("%d ", banker.available[i])
	}
	fmt.Println("\n")

	fmt.Println("The Max Matrix is...")
	fmt.Println("   " + getHeader(banker.M))
	for i := 0; i < banker.N; i++ {
		fmt.Printf("%d:", i)
		for j := 0; j < banker.M; j++ {
			fmt.Printf(" %d", banker.max[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("The Allocation Matrix is...")
	fmt.Println("   " + getHeader(banker.M))
	for i := 0; i < banker.N; i++ {
		fmt.Printf("%d:", i)
		for j := 0; j < banker.M; j++ {
			fmt.Printf(" %d", banker.allocation[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("The Need Matrix is...")
	fmt.Println("   " + getHeader(banker.M))
	for i := 0; i < banker.N; i++ {
		fmt.Printf("%d:", i)
		for j := 0; j < banker.M; j++ {
			fmt.Printf(" %d", banker.need[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func addVector(change, keep []int) {
	for i := range change {
		change[i] = change[i] + keep[i]
	}
}

func subtractVector(change, keep []int) {
	for i := range change {
		change[i] = change[i] - keep[i]
	}
}

func isLessThan(A, B []int) bool {
	for i := range A {
		if A[i] > B[i] {
			return false
		}
	}
	return true
}

func findProcess(banker Banker, work []int, finish []bool) int {
	for i := range finish {
		fmt.Println(i, banker.need[i], work, finish[i])
		if !finish[i] && isLessThan(banker.need[i], work) {
			return i
		}
	}
	return -1
}

func isInSafeState(banker Banker) bool {

	work := make([]int, banker.M)
	copy (work, banker.available)
	finish := make([]bool, banker.N)
	notFinish := true
	for notFinish {
		if index := findProcess(banker, work, finish); index >= 0 {
			addVector(work, banker.allocation[index])
			finish[index] = true
			fmt.Println(work)
		} else {
			notFinish = false
		}
	}

	for i := range finish {
		if !finish[i] {
			return false
		}
	}

	return true
}

func isAllocatable(banker Banker, request []int, initial int) bool {
	
	return true
}

func main() {
	if len(os.Args) <= 1 || len(os.Args) > 2 {
		fmt.Println("usage: bankerAlgorithm <resoureFile>")
		os.Exit(2)
	}
	file, err := os.Open(os.Args[1])
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var splitLines [][]string
	for scanner.Scan() {
		if line := scanner.Text(); len(line) != 0 {
			splitLines = append(splitLines, strings.Split(line, " "))
		}
	}
	handleError(scanner.Err())

	banker, request, initial := get_input(splitLines)
	fmt.Printf("There are %d processes and %d resources types in the system.\n\n", banker.N, banker.M)
	printBanker(banker)
	if isInSafeState(banker) {

		fmt.Println("THE SYSTEM IS IN A SAFE STATE\n")

		fmt.Println("The Request Vector is...")
		fmt.Println("   " + getHeader(banker.M))
		fmt.Printf("%d: ", initial)
		for i := 0; i < banker.M; i++ {
			fmt.Printf("%d ", request[i])
		}
		fmt.Println("\n")

		if isAllocatable(banker, request, initial) {
			fmt.Println("THE REQUEST CAN BE GRANTED: NEW STATE FOLLOWS\n")
			printBanker(banker)
		} else {
			fmt.Println("THE REQUEST CANNOT BE GRANTED\n")
		}
	} else {
		fmt.Println("THE SYSTEM IS NOT IN A SAFE STATE")
	}

}
