/*Write a function that outputs the transpose of a matrix - a new matrix where the columns and rows of the original are swapped.

For example, the transpose of:

| 1 2 3 |
| 4 5 6 |
is

| 1 4 |
| 2 5 |
| 3 6 |
The input to your function will be an array of matrix rows.
 You can assume that each row has the same length, and that the height and width of the matrix are both positive.
*/
package kata


func Transpose(matrix [][]int) [][]int {
      // your code here
  
       result :=  [][]int{}
      for i:=0;i < len(matrix[0]);i++{
        row := []int{}
        for j:=0;j< len(matrix);j++{
          row = append(row,matrix[j][i])
        }
        result = append(result,row)
      }
  return result
}