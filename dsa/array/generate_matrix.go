package array

// GenerateMatrix generates a n x n matrix which is filled with elements from 1 to n^2 in spiral order.
//
// Example 1:
// Input: n = 3, Output: [[1,2,3],[8,9,4],[7,6,5]]
// Example 2:
// Input: n = 1, Output: [[1]]
func GenerateMatrix(n int) [][]int {
	var rowBegin, rowEnd, colBegin, colEnd = 0, n - 1, 0, n - 1
	var counter = 1

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for rowBegin <= rowEnd && colBegin <= colEnd {
		// Direction 1, fixed row, traverse from first column to last column.
		for i := colBegin; i <= colEnd; i++ {
			matrix[rowBegin][i] = counter
			counter++
		}
		// Begin from the next row.
		rowBegin++
		// Direction 2, fixed column, traverse from next row to last row.
		for i := rowBegin; i <= rowEnd; i++ {
			matrix[i][colEnd] = counter
			counter++
		}
		// Begin from the n-1 col.
		colEnd--
		// Direction 3, fixed row, traverse from n-1 column to first column.
		for i := colEnd; i >= colBegin; i-- {
			matrix[rowEnd][i] = counter
			counter++
		}
		// Begin from n-1 row.
		rowEnd--
		// Direction 4, fixed column, traverse from n-1 row to second row.
		for i := rowEnd; i >= rowBegin; i-- {
			matrix[i][colBegin] = counter
			counter++
		}
		// Move begin position of col for next traversal.
		colBegin++
	}
	return matrix
}
