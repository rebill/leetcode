package leetcode

func hammingDistance(x int, y int) int {
    dist := 0
    val := x ^ y

    // Count the number of bits set
    for (val != 0) {
        // A bit is set, so increment the count and clear the bit
        dist++
        val &= val - 1
    }

    // Return the number of differing bits
    return dist;
}
