package main

import "fmt"

func main() {
	n := 10
	fmt.Println(firstBadVersion(n))

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	min := 1
	max := n
	middle := (max - min) / 2

	for min <= max {
		switch isBadVersion(middle) {
		case true:
			max = middle - 1
		case false:
			min = middle + 1
		}
		middle = min + (max-min)/2

	}

	return middle
}

func isBadVersion(version int) bool {
	return version >= 1
}
