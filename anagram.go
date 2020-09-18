package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

var dic = make(map[string]string)

func TimeTracker(start time.Time, arg string) {
	FindAnagram(dic, arg)
	fmt.Printf(">> It took %v to find the anagrams.", time.Since(start))

}

func main() {
	arg := os.Args[1]
	defer TimeTracker(time.Now(), arg)
	file, err := os.Open("/Users/Aidar/Desktop/words.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value := scanner.Text()
		key := SortStr(value)
		dic[key] += value + "\n"
	}

}

func SortStr(arg string) string {

	arr := strings.Split(arg, "")
	sort.Strings(arr)
	return strings.Join(arr, "")

}
func FindAnagram(dic map[string]string, arg string) {
	fmt.Printf("Find anagrams of %v:\n", arg)
	fmt.Print(dic[SortStr(arg)])

}
