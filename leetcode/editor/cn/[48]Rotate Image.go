//You are given an n x n 2D matrix representing an image, rotate the image by 90
// degrees (clockwise). 
//
// You have to rotate the image in-place, which means you have to modify the inp
//ut 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation. 
//
// 
// Example 1: 
//
// 
//Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
//Output: [[7,4,1],[8,5,2],[9,6,3]]
// 
//
// Example 2: 
//
// 
//Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
// 
//
// Example 3: 
//
// 
//Input: matrix = [[1]]
//Output: [[1]]
// 
//
// Example 4: 
//
// 
//Input: matrix = [[1,2],[3,4]]
//Output: [[3,1],[4,2]]
// 
//
// 
// Constraints: 
//
// 
// matrix.length == n 
// matrix[i].length == n 
// 1 <= n <= 20 
// -1000 <= matrix[i][j] <= 1000 
// 
// Related Topics 数组 
// 👍 645 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int)  {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		for j := 0; j < (length+1)/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[length-j-1][i]
			matrix[length-j-1][i] = matrix[length-i-1][length-j-1]
			matrix[length-i-1][length-j-1] = matrix[j][length-i-1]
			matrix[j][length-i-1] = temp
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
