package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left, right, bad := 1, n, n

	for left <= right {
		current := (left + right) / 2

		if isBadVersion(current) {
			bad = current
			right = current - 1
		} else {
			left = current + 1
		}
	}

	return bad
}
