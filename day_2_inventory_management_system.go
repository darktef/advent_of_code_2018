/**
Part One

To make sure you didn't miss any, you scan the likely candidate boxes again, counting the number that have an ID containing exactly two of any letter and then separately counting those with exactly three of any letter. You can multiply those two counts together to get a rudimentary checksum and compare it to what your device predicts.

For example, if you see the following box IDs:

abcdef contains no letters that appear exactly two or three times.
bababc contains two a and three b, so it counts for both.
abbcde contains two b, but no letter appears exactly three times.
abcccd contains three c, but no letter appears exactly two times.
aabcdd contains two a and two d, but it only counts once.
abcdee contains two e.
ababab contains three a and three b, but it only counts once.
Of these box IDs, four of them contain a letter which appears exactly twice, and three of them contain a letter which appears exactly three times. Multiplying these together produces a checksum of 4 * 3 = 12.

What is the checksum for your list of box IDs?
**/

/**
Part Two

Confident that your list of box IDs is complete, you're ready to find the boxes full of prototype fabric.

The boxes will have IDs which differ by exactly one character at the same position in both strings. For example, given the following box IDs:

abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz
The IDs abcde and axcye are close, but they differ by two characters (the second and fourth). However, the IDs fghij and fguij differ by exactly one character, the third (h and u). Those must be the correct boxes.

What letters are common between the two correct box IDs? (In the example above, this is found by removing the differing character from either ID, producing fgij.)
**/

package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  result := partTwo()

  fmt.Println("The result is %s", result)
}

func partOne() int {
  f, err := os.Open("day_2_input")
  check(err)

  scanner := bufio.NewScanner(f)
  globalTwos := 0
  globalThrees := 0

  for scanner.Scan() {
    t := scanner.Text()

    // Those sequences represent Unicode code points, called runes.
    // ref: https://blog.golang.org/strings
    freqs := make(map[rune]int)
    twos := 0
    threes := 0

    for _, runeValue := range t {
      // If the requested key doesn't exist, we get the value type's zero value.
      // ref: https://blog.golang.org/go-maps-in-action
      freqs[runeValue] += 1
    }

    for _, v := range freqs {
      if v == 2 {
        twos += 1
      } else if v == 3 {
        threes += 1
      }

      if twos >= 1 && threes >= 1 {
        break
      }
    }

    if twos >= 1 {
      globalTwos += 1
    }

    if threes >= 1 {
      globalThrees += 1
    }
  }

  return globalTwos * globalThrees
}

/**
Inspire by the checksum approach, if two strings are only differ by one rune, their checksum difference will be differ by no more than 26, assuming their checksum function is sum(rune * freq)

1. Iterate through the files, get checksum for each string
2. Find the pair which their checksum are only diff by 26
3. Manually iterate through the pair to confirm that is the case, a. mis-order b. 1+9=5+5 kind effect
**/
func partTwo() string {
  f, err := os.Open("day_2_input")
  check(err)

  scanner := bufio.NewScanner(f)

  codes := make([]string, 1)

  for scanner.Scan() {
    code := scanner.Text()
    codes = append(codes, code)
  }

  var result string
  for i := 0; i < len(codes); i++ {
    for j := i+1; j < len(codes); j++ {
      mismatchIndex := moreThanOneMismatch(codes[i], codes[j])
      if  mismatchIndex != -1 {
        var str1 = codes[i][:mismatchIndex]
        var str2 = codes[i][(mismatchIndex + 1):]

        result = str1 + str2
        break
      }
    }
  }

  return result
}

func moreThanOneMismatch(s1, s2 string) int {
  var mismatches int
  var mismatchPosition int

  for i := 0; i < len(s1); i++ {
    if s1[i] != s2[i] {
      mismatches++
      mismatchPosition = i
    }

    if mismatches > 1 {
      return -1
    }
  }

  if mismatches == 0 {
    return -1
  }

  return mismatchPosition
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}
