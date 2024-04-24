package diff

import (
	"bufio"
	"fmt"
	"os"
)

func Diff(path1 string, path2 string ){
	//1st file
	file1, err := os.Open(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file1.Close()

	//2nd file
	file2, err := os.Open(path2)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer file2.Close()

	//Построчный сканер bufio.Scanner создан для построчного чтения данных.
	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	lineNumber := 1
	t1 := scanner1.Scan()
	t2 := scanner2.Scan()
	for t1 && t2 {
		line1 := scanner1.Text()
		line2 := scanner2.Text()
		// fmt.Println("COMPARE [", line1,"AND", line2, "]")
		if line1 != line2 {
			fmt.Printf("Difference found at line %d:\n", lineNumber)
			fmt.Printf("--- %s\n+++ %s\n", line1, line2)
		}
		lineNumber ++
		t1 = scanner1.Scan()
		t2 = scanner2.Scan()
	}
	
	for t1 { 
		fmt.Printf("Difference found at line %d:\n", lineNumber)
		fmt.Printf("--- %s\n", scanner1.Text())
		lineNumber++
		t1 = scanner1.Scan()
	   }
	

	for t2 {
		fmt.Printf("Difference found at line %d:\n", lineNumber)
		fmt.Printf("+++ %s\n", scanner2.Text())
		lineNumber++
		t2 = scanner2.Scan()
	   }
}