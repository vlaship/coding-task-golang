package problem_0706

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMyHashMap(t *testing.T) {
	testCases := []struct {
		operations []string
		inputs     [][]int
		expected   []interface{}
	}{
		{
			operations: []string{"MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"},
			inputs:     [][]int{{}, {1, 1}, {2, 2}, {1}, {3}, {2, 1}, {2}, {2}, {2}},
			expected:   []interface{}{nil, nil, nil, 1, -1, nil, 1, nil, -1},
		},
		{
			operations: []string{"MyHashMap", "get", "put", "put", "put", "remove", "put", "put", "put", "get", "put", "put", "put", "put", "get", "put", "get", "put", "put", "put", "put", "remove", "put", "put", "put", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put", "put", "put", "put", "put", "put", "put", "remove", "put", "remove", "put", "remove", "put", "remove", "put", "put", "put", "remove", "put", "put", "put", "put", "get", "put", "put", "put", "get", "remove", "put", "put", "put", "put", "remove", "put", "put", "put", "get", "put", "put", "get", "get", "put", "put", "put", "put", "put", "put", "put", "put", "get", "put", "put", "put", "get", "get", "remove", "remove", "put", "get", "get", "put", "get", "put", "put", "get"},
			inputs:     [][]int{{}, {79}, {72, 7}, {77, 1}, {10, 21}, {26}, {94, 5}, {53, 35}, {34, 9}, {94}, {96, 8}, {73, 79}, {7, 60}, {84, 79}, {94}, {18, 13}, {18}, {69, 34}, {21, 82}, {57, 64}, {23, 60}, {0}, {12, 97}, {56, 90}, {44, 57}, {30, 12}, {17, 10}, {42, 13}, {62, 6}, {34}, {70, 16}, {51, 39}, {22, 98}, {82, 42}, {84, 7}, {29, 32}, {96, 54}, {57, 36}, {85, 82}, {49, 33}, {22, 14}, {63, 8}, {56, 8}, {94}, {78, 77}, {51}, {20, 89}, {51}, {9, 38}, {20}, {29, 64}, {92, 69}, {72, 25}, {73}, {6, 90}, {1, 67}, {70, 83}, {58, 49}, {79}, {73, 2}, {56, 16}, {58, 26}, {53}, {7}, {27, 17}, {55, 40}, {55, 13}, {89, 32}, {49}, {75, 75}, {64, 52}, {94, 74}, {81}, {39, 82}, {47, 36}, {57}, {66}, {3, 7}, {54, 34}, {56, 46}, {58, 64}, {22, 81}, {3, 1}, {21, 96}, {6, 19}, {77}, {60, 66}, {48, 85}, {77, 16}, {78}, {23}, {72}, {27}, {20, 80}, {30}, {94}, {74, 85}, {49}, {79, 59}, {15, 15}, {26}},
			expected:   []interface{}{nil, -1, nil, nil, nil, nil, nil, nil, nil, 5, nil, nil, nil, nil, 5, nil, 13, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 9, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, -1, nil, nil, nil, 35, nil, nil, nil, nil, nil, nil, nil, nil, nil, -1, nil, nil, 36, -1, nil, nil, nil, nil, nil, nil, nil, nil, 1, nil, nil, nil, 77, 60, nil, nil, nil, 12, 74, nil, -1, nil, nil, -1},
		},
		// You can add more test cases here if needed
	}

	for _, tc := range testCases {
		var myHashMap MyHashMap
		result := []interface{}{}

		for i, op := range tc.operations {
			var key int
			if i != 0 {
				key = tc.inputs[i][0]
			}
			switch op {
			case "MyHashMap":
				myHashMap = Constructor()
				result = append(result, nil) // Constructor returns nil
			case "put":
				myHashMap.Put(key, tc.inputs[i][1])
				result = append(result, nil) // Put returns nil
			case "get":
				res := myHashMap.Get(key)
				result = append(result, res)
			case "remove":
				myHashMap.Remove(key)
				result = append(result, nil) // Remove returns nil
			}
		}

		// Assert that the result matches the expected output using require.Equal
		require.Equal(t, tc.expected, result)
	}
}
