/*Write a simple parser that will parse and run Deadfish.

Deadfish has 4 commands, each 1 character long:

i increments the value (initially 0)
d decrements the value
s squares the value
o outputs the value into the return array
Invalid characters should be ignored.

Parse("iiisdoso") == []int{8, 64}
*/
package kata

func Parse(data string) []int{
  //TODO: write your code here
  result := []int{}
  deadfish :=0
  for _,character := range data {
    switch character {
      case 'i': deadfish++
      case 'd': deadfish--
      case 's': deadfish*=deadfish
      case 'o': result = append(result,deadfish)
    }
  }
  return result
}