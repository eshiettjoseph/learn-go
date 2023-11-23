package main 

import(
	"fmt"
	"flag"
	"os"
)
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	
}