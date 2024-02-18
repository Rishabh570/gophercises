package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Hello from challenge 12")

	subDirToSkip := "skip"
	re := regexp.MustCompile(`(.+\/file)(\d+)(\..*)$`)

	err := filepath.Walk("./sample", func(path string, info fs.FileInfo, err error) error {
		// fmt.Println("path:", path)
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}

		fmt.Printf("visited file or dir: %q\n", path)

		// regex check
		patternMatched := re.FindStringSubmatch(path)
		fmt.Println("patternMatched:", patternMatched)
		if patternMatched == nil {
			fmt.Println("No number found in the file name\n")
			return nil
		}

		// Get the number and convert it to its word counterpart
		num, _ := strconv.Atoi(patternMatched[2])
		fmt.Println("num:", num)
		wordNum := numToWord(num)
		fmt.Println("wordNum:", wordNum)

		// Construct the new file name
		newFileName := fmt.Sprintf("%s-%s%s", patternMatched[1], wordNum, patternMatched[3])
		fmt.Println("newFileName:", newFileName)

		renameErr := os.Rename(path, newFileName)
		if renameErr != nil {
			fmt.Printf("Error renaming file:%s\n", renameErr)
			return nil
		}

		fmt.Printf("renamed file path: %s with this: %s\n", path, newFileName)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return
	}
}

func numToWord(num int) string {
	words := map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
	}

	// Return the word if it exists in the mapping, otherwise return the number as a string
	if word, ok := words[num]; ok {
		return word
	}
	return strconv.Itoa(num)
}
