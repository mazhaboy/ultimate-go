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
	scan := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a word: ")
	scan.Scan()
	arg := scan.Text()

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
		if value!=arg {
		dic[key] += value + "\n"
		}
	}

}

func SortStr(arg string) string {

	arr := strings.Split(arg, "")
	sort.Strings(arr)
	return strings.Join(arr, "")

}
func FindAnagram(dic map[string]string, arg string) {

	a := strings.Split(dic[SortStr(arg)], "\n")
	fmt.Println("The word", arg, "has the following", len(a)-1, "anagrams:")
	if val, ok := dic[SortStr(arg)]; ok {
		fmt.Print(val)
	}

}
