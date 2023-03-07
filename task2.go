package solution

/*
There is an int array of size n.
The value of each element is greater than or equal to 1 and less or equal n.
Some elements can be repeated. Find all elements of the set from 1 to n inclusive, which are not originally in the array.

Example
Input: [4,3,2,7,8,2,3,1]
Output: [5,6]
*/
func Task2(input []int) []int {
	var r = make([]int, len(input))
	var res = make([]int, 0)

	for _, value := range input {
		r[value-1] = 1
	}

	for i, v := range r {
		if v == 0 {
			res = append(res, i+1)
		}
	}

	return res
}
