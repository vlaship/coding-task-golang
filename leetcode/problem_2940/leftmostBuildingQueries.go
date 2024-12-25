package problem_2940

import "sort"

const infinity = 1 << 30

// BinaryIndexedTree represents a Fenwick Tree for range minimum queries
type BinaryIndexedTree struct {
    size int
    tree []int
}

// NewBIT creates a new Binary Indexed Tree
func NewBIT(size int) *BinaryIndexedTree {
    tree := make([]int, size+1)
    for i := range tree {
        tree[i] = infinity
    }
    return &BinaryIndexedTree{size: size, tree: tree}
}

// update updates the tree with value v at index x
func (bit *BinaryIndexedTree) update(x, v int) {
    for x <= bit.size {
        if v < bit.tree[x] {
            bit.tree[x] = v
        }
        x += x & -x // Add least significant bit
    }
}

// query returns the minimum value up to index x
func (bit *BinaryIndexedTree) query(x int) int {
    minimum := infinity
    for x > 0 {
        if bit.tree[x] < minimum {
            minimum = bit.tree[x]
        }
        x -= x & -x // Remove least significant bit
    }
    if minimum == infinity {
        return -1
    }
    return minimum
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
    n := len(heights)
    m := len(queries)
    
    // Normalize queries to ensure left < right
    for i := range queries {
        if queries[i][0] > queries[i][1] {
            queries[i][0], queries[i][1] = queries[i][1], queries[i][0]
        }
    }
    
    // Create indices array for sorting
    indices := make([]int, m)
    for i := range indices {
        indices[i] = i
    }
    
    // Sort indices based on right bound of queries in descending order
    sort.Slice(indices, func(i, j int) bool {
        return queries[indices[j]][1] < queries[indices[i]][1]
    })
    
    // Copy and sort heights
    sortedHeights := make([]int, n)
    copy(sortedHeights, heights)
    sort.Ints(sortedHeights)
    
    // Initialize result array
    result := make([]int, m)
    
    // Initialize Binary Indexed Tree
    tree := NewBIT(n)
    
    // Process queries from right to left
    j := n - 1
    for _, idx := range indices {
        left, right := queries[idx][0], queries[idx][1]
        
        // Update tree for buildings to the right
        for j > right {
            // Find position in sorted array
            pos := n - sort.SearchInts(sortedHeights, heights[j]) + 1
            tree.update(pos, j)
            j--
        }
        
        // Check if we can directly use right building
        if left == right || heights[left] < heights[right] {
            result[idx] = right
        } else {
            // Find position for minimum height needed
            k := n - sort.SearchInts(sortedHeights, heights[left])
            result[idx] = tree.query(k)
        }
    }
    
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
