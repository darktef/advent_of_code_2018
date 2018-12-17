/**
After feeling like you've been falling for a few minutes, you look at the device's tiny screen. "Error: Device must be calibrated before first use. Frequency drift detected. Cannot maintain destination lock." Below the message, the device shows a sequence of changes in frequency (your puzzle input). A value like +6 means the current frequency increases by 6; a value like -3 means the current frequency decreases by 3.

For example, if the device displays frequency changes of +1, -2, +3, +1, then starting from a frequency of zero, the following changes would occur:

Current frequency  0, change of +1; resulting frequency  1.
Current frequency  1, change of -2; resulting frequency -1.
Current frequency -1, change of +3; resulting frequency  2.
Current frequency  2, change of +1; resulting frequency  3.
In this example, the resulting frequency is 3.

Here are other example situations:

+1, +1, +1 results in  3
+1, +1, -2 results in  0
-1, -2, -3 results in -6
Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?
**/

/**
1. Read in the file, line by line
2. Parse the line to integer, input could be +10 or -10
3. Add the input to result while parsing
4. Return the final result

Knowledge point:
1. Use go to parse file line by line
2. Write the function in go
**/

/**
--- Part Two ---
You notice that the device repeats the same frequency change list over and over. To calibrate the device, you need to find the first frequency it reaches twice.

For example, using the same list of changes above, the device would loop as follows:

Current frequency  0, change of +1; resulting frequency  1.
Current frequency  1, change of -2; resulting frequency -1.
Current frequency -1, change of +3; resulting frequency  2.
Current frequency  2, change of +1; resulting frequency  3.
(At this point, the device continues from the start of the list.)
Current frequency  3, change of +1; resulting frequency  4.
Current frequency  4, change of -2; resulting frequency  2, which has already been seen.
In this example, the first frequency reached twice is 2. Note that your device might need to repeat its list of frequency changes many times before a duplicate frequency is found, and that duplicates might be found while in the middle of processing the list.

Here are other examples:

+1, -1 first reaches 0 twice.
+3, +3, +4, -2, -4 first reaches 10 twice.
-6, +3, +8, +5, -6 first reaches 5 twice.
+7, +7, -2, -7, -4 first reaches 14 twice.
What is the first frequency your device reaches twice?
**/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	partTwo()
}

func partOne() {
	result := 0

	f, err := os.Open("day_1_input")
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t, err := strconv.Atoi(scanner.Text())
		check(err)
		result += t
	}

	err = scanner.Err()
	check(err)

	fmt.Println(result)
}

func partTwo() {
	freqs := getFreqList()
	dictionary := make(map[int]int)
	result := 0
	round := 0

	for n := 0; n <= len(freqs); n++ {
		if n == len(freqs) {
			n = 0
			round += 1
		}

		result += freqs[n]
		if dictionary[result] != 0 {
			fmt.Println(result)
			break
		} else {
			dictionary[result] = 1
		}
	}
}

func getFreqList() []int {
	var freq []int

	f, err := os.Open("day_1_input")
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		check(err)

		freq = append(freq, num)
	}

	err = scanner.Err()
	check(err)

	return freq
}
