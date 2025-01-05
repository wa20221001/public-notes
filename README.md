### Function Styles

`go mod tidy` Downloads packages required. Packages located in

`go.mod` located at root of module

`go.sum`

```golang
func <function name>(<declare inputs>) (<declare outputs>)
```

```go
func nameoffunction(a int, b int) (resulta int, result b int)
```

This is a special function/ entry point for the program

```golang
func main()
```

packages are installed in `GOPATH`

**Error Handling**

if functions, they should be  returning error (type error) with value of error to return

```golang
package main

import (
    "errors"
    "fmt"
)

func main() {
    var numerator int = 11
    var denominator int = 2
    result, reminder, err := intDivision(numerator, denominator)
    switch {
        case err!=nil:
           fmt.Printf(err.Error())
        default:
            fmt.Printf("result is %v and reminder is %v, result,reminder)
    }
}

func intDivision(numerator int, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, errors.New("denominator can not be zero")
    }
    result := numerator / denominator
    reminder := numerator % denominator
    return result, reminder, nil
}
```

#### 1. **Arrays . memory efficient,**   Arrays have a fixed length, so they’re not flexible.   Passing arrays to functions involves copying the entire array unless you use pointers.

- **Create**
    - Empty (fixed size): `var arr [5]int`
    - With values: `arr := [5]int{1, 2, 3, 4, 5}`
    - `[...]int32{1,2,3,4}`
- **Modify**
    - `arr[0] = 10`
- **Add Values**
    - Not possible. Arrays have a fixed size. Use slices instead.
- **Remove Values**
    - Not possible. Arrays are fixed-size. Use slices for dynamic behavior.
- **Search**
    - No inbuilt search. Use a loop to iterate and check values.
- **Sort**
    - Convert to a slice and use: `sort.Ints(arr[:])` (requires `import "sort"`)

#### 2. **Slices . wrappers around arrays with better functionality**

- **Create**
    - Empty: `var slice []int`
    - With values: `slice := []int{1, 2, 3}`
    - Using `make()`: `slice := make([]int, 5)` (creates a slice of length 5 with default values)
    - Using `make()` with capacity: `slice := make([]int, 5, 10)` (creates a slice of length 5 with capacity 10)
- **Modify**
    - `slice[0] = 10`
- **Add Values**
    - Use `append()`: `slice = append(slice, 4, 5)`
    - you can append slices into slices i.e.  `slice = append(slice,slice2)`
- **Remove Values**
    - Remove by index: `slice = append(slice[:index], slice[index+1:]...)`
- **Search**
    - No inbuilt search. Use a loop or custom function.
- **Sort**
    - Use `sort.Ints(slice)` (requires `import "sort"`)

#### 3. **Strings**

- **Create**
    - Empty: `var str string`
    - With value: `str := "hello"`
- **Modify**
    - Strings are immutable. Convert to `[]rune`: `runes := []rune(str); runes[0] = 'H'; str = string(runes)`
- **Add Values**
    - Concatenate: `str = str + " world"`
- **Remove Values**
    - Create a new string by slicing: `str = str[:index] + str[index+1:]`
- **Search**
    - Use `strings.Contains(str, "world")` (requires `import "strings"`)
- **Sort**
    - Convert to `[]rune`, then use `sort.Slice()`: `runes := []rune(str); sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] }); str = string(runes)`

#### 4. **Runes**

- **Create**
    - Single rune: `var r rune = 'a'`
- **Modify**
    - Convert string to runes: `runes := []rune("hello"); runes[1] = 'a'`
- **Add Values**
    - Append to a slice of runes: `runes = append(runes, 'c')`
- **Remove Values**
    - Remove by index: `runes = append(runes[:index], runes[index+1:]...)`
- **Search**
    - No inbuilt search. Use a loop to iterate.
- **Sort**
    - Use `sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })`

#### 5. **Maps , pretty much dictonary**

- **Create**
    
    - Empty map: `var m map[string]int` or `m := make(map[string]int)`
    - With values: `m := map[string]int{"one": 1, "two": 2}`
- **Modify**
    
    - Add or update value: `m["three"] = 3`
- **Add Values**
    
    - Simply assign a value to a new key: `m["four"] = 4`
- **Remove Values**
    
    - Use `delete()`: `delete(m, "two")`
- **Search**
    
    - Check if a key exists:
        
        ```go
        val, exists := m["one"]
        if exists {
            // key exists, use val
        }
        ```
        
- **Sort**
    
    - Maps are inherently unordered. To sort, extract keys to a slice, sort the slice, and then access map values in order.
        
        ```go
        keys := make([]string, 0, len(m))
        for k := range m {
            keys = append(keys, k)
        }
        sort.Strings(keys) // requires `import "sort"`
        for _, k := range keys {
            fmt.Println(k, m[k])
        }
        ```
        

`switch`  
evaluated once if x, y is equal to whatever expression

```
switch expression {
case x:

case y:

default:
}
```

```go
if condition1 {
   // code to be executed if condition1 is true
} else if condition2 {
   // code to be executed if condition1 is false and condition2 is true
} else {
   // code to be executed if condition1 and condition2 are both false
} 
```

Approaches:

**strings and arrays**:  
two pointers: already sorted input, PAIRWISE comparison , looking for specific pair/subarray  
sliding window: extension of above

do not use: unsorted input, non linear data

### Techniques Used in Previous Questions

Here are some of the common techniques used in the previous questions, along with explanations of when and why to use them.

#### 1. **Two Pointers Technique**

- **Use Case**: Suitable for problems involving sorted arrays or linked lists, especially when looking for a pairwise comparison, finding subarrays, or managing two ends of a sequence.
- **Logic**: Use two pointers, typically starting from either end of the array or one starting at the beginning and the other at the end. The pointers move towards each other based on certain conditions.
- **Example**: The **Trapping Rain Water** problem uses two pointers to efficiently track the maximum heights on the left and right, allowing the calculation of trapped water based on the lower boundary.

#### 2. **Hash Set for Tracking Elements**

- **Use Case**: Useful when checking for existence, uniqueness, or identifying elements that have already been seen, especially when order doesn't matter.
    
- **Logic**: Store elements in a hash set (map) for quick lookup and to prevent duplication.
    
- **Pseudocode**:
    
    ```
    function findConsecutiveSequence(nums):
        set = new HashSet()
        for num in nums:
            set.add(num)
    
        maxLength = 0
        for num in set:
            if num - 1 not in set:
                currentLength = 1
                currentNum = num
                while currentNum + 1 in set:
                    currentNum += 1
                    currentLength += 1
                maxLength = max(maxLength, currentLength)
        return maxLength
    ```
    
- **Example**: The **Longest Consecutive Sequence** problem uses a set to quickly determine if the current number is the start of a sequence, allowing an efficient way to track consecutive numbers.
    

#### 3. **Breadth-First Search (BFS)**

- **Use Case**: Applicable for shortest path problems in graphs or any scenario where you need to explore level by level (e.g., word transformations, tree traversals).
    
- **Logic**: Use a queue to explore nodes level by level. Each level's nodes are explored before moving on to the next level.
    
- **Pseudocode**:
    
    ```
    function bfs(startNode):
        queue = new Queue()
        queue.enqueue(startNode)
        visited = new Set()
        visited.add(startNode)
    
        while not queue.isEmpty():
            currentNode = queue.dequeue()
            for neighbor in currentNode.neighbors:
                if neighbor not in visited:
                    queue.enqueue(neighbor)
                    visited.add(neighbor)
    ```
    
- **Example**: The **Word Ladder** problem uses BFS to find the shortest transformation sequence between two words, utilizing a queue to manage the level-wise traversal of possible transformations.
    

#### 4. **Sliding Window Technique**

- **Use Case**: Useful for finding subarrays or subranges within an array/string where a certain condition needs to be met (e.g., sum, distinct characters, fixed size window).
    
- **Logic**: Use a window that slides across the array, expanding or shrinking based on conditions to maintain an optimal solution.
    
- **Pseudocode**:
    
    ```
    function slidingWindowSum(nums, k):
        currentSum = 0
        maxSum = 0
    
        # Initialize the first window
        for i in range(0, k):
            currentSum += nums[i]
    
        maxSum = currentSum
    
        # Slide the window
        for i in range(k, len(nums)):
            currentSum += nums[i] - nums[i - k]
            maxSum = max(maxSum, currentSum)
    
        return maxSum
    ```
    
- **Example**: The **Subarray Sum Equals K** problem uses a sliding window with a map that tracks cumulative sums to efficiently count the number of subarrays summing to `k`.
    

#### 5. **Sorting for k-th Largest or Smallest Elements**

- **Use Case**: When finding the k-th largest or smallest element, sorting is a straightforward approach, particularly for small datasets.
    
- **Logic**: Sort the entire array, then directly index to find the desired element.
    
- **Pseudocode**:
    
    ```
    function findKthLargest(nums, k):
        sort(nums, descending)
        return nums[k - 1]
    ```
    
- **Example**: The **Kth Largest Element in an Array** problem sorts the array in descending order and returns the `k`-th element.
    

#### 6. **Prefix Sum with Hash Map**

- **Use Case**: Applicable when needing to efficiently find subarray sums or track cumulative sums to find a target sum.
    
- **Logic**: Calculate a running sum and use a hash map to track occurrences of cumulative sums. This allows checking if a specific cumulative sum has occurred before.
    
- **Pseudocode**:
    
    ```
    function subarraySum(nums, k):
        cumulativeSum = 0
        sumMap = {0: 1}
        count = 0
    
        for num in nums:
            cumulativeSum += num
            if cumulativeSum - k in sumMap:
                count += sumMap[cumulativeSum - k]
            if cumulativeSum in sumMap:
                sumMap[cumulativeSum] += 1
            else:
                sumMap[cumulativeSum] = 1
    
        return count
    ```
    
- **Example**: The **Subarray Sum Equals K** problem uses a map to store cumulative sums and check if a particular difference exists, allowing for constant time lookup.
    

#### 7. **Dynamic Update of Boundaries**

- **Use Case**: Useful when dealing with subproblems that require dynamically updating constraints, such as maximum or minimum values encountered so far.
    
- **Logic**: Track current boundaries (e.g., `leftMax` and `rightMax`) and update them as you iterate through the data.
    
- **Pseudocode**:
    
    ```go
    function trapRainWater(heights):
        left = 0
        right = len(heights) - 1
        leftMax = 0
        rightMax = 0
        water = 0
    
        while left < right:
            if heights[left] < heights[right]:
                if heights[left] >= leftMax:
                    leftMax = heights[left]
                else:
                    water += leftMax - heights[left]
                left += 1
            else:
                if heights[right] >= rightMax:
                    rightMax = heights[right]
                else:
                    water += rightMax - heights[right]
                right -= 1
    
        return water
    ```
    
- **Example**: The **Trapping Rain Water** problem maintains `leftMax` and `rightMax` boundaries to efficiently calculate the volume of trapped water.
    

### Summary

- **Two Pointers**: Ideal for sorted data and pairwise comparisons.
- **Hash Set**: Fast lookups to track uniqueness or check existence.
- **BFS**: Level-by-level traversal to find shortest paths or explore all possible states.
- **Sliding Window**: Optimal for subarray/substring problems, expanding and shrinking based on conditions.
- **Sorting**: Straightforward solution for finding ranked elements, especially when complexity is not a concern.
- **Prefix Sum with Hash Map**: Efficient subarray sum tracking using cumulative sums and a hash map.
- **Dynamic Boundaries Update**: Suitable for problems involving maintaining or updating limits dynamically.

### Advanced Algorithmic Techniques with Detailed Examples and Explanations

Below are some advanced techniques commonly used to solve algorithmic problems, along with detailed explanations, pseudocode, and when to apply them.

#### 1. **Binary Search Technique**

- **Use Case**: Ideal for searching in sorted datasets. Efficient for problems that involve finding an element or determining the position of an element in a sorted array.
    
- **Logic**: Repeatedly divide the search range in half. If the target is less than the middle element, search the left half; otherwise, search the right half.
    
- **Pseudocode**:
    
    ```
    function binarySearch(arr, target):
        left = 0
        right = len(arr) - 1
    
        while left <= right:
            mid = (left + right) // 2
            if arr[mid] == target:
                return mid
            elif arr[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        return -1
    ```
    
- **Example**: Searching for a specific element in a sorted list of numbers or finding the insertion point for an element.
    

**Additional Details**:

- **Complexity**: O(log n) due to halving the search space each time. This makes it highly efficient for large datasets.
- **Applications**: Searching for elements in dictionaries, index lookups, and finding boundaries in problems involving ranges.

#### 2. **Dynamic Programming (DP)**

- **Use Case**: Best for problems involving optimization, such as finding the maximum or minimum result, or problems that have overlapping subproblems.
    
- **Logic**: Break down a problem into smaller subproblems, solve each subproblem once, and store the result to avoid redundant calculations.
    
- **Pseudocode**:
    
    ```
    function fibonacci(n):
        dp = [0] * (n + 1)
        dp[1] = 1
        for i in range(2, n + 1):
            dp[i] = dp[i - 1] + dp[i - 2]
        return dp[n]
    ```
    
- **Example**: The **Fibonacci Sequence** problem is solved efficiently using dynamic programming to avoid recalculating intermediate values.
    

**Additional Details**:

- **Complexity**: O(n) for time and O(n) for space when using an array. Space can be optimized to O(1) by using just two variables for Fibonacci.
- **Applications**: Problems involving sequences (Fibonacci, longest common subsequence), knapsack problems, and optimal pathfinding.

#### 3. **Backtracking**

- **Use Case**: Ideal for constraint satisfaction problems, such as generating permutations, combinations, or solving puzzles like Sudoku.
    
- **Logic**: Incrementally build candidates for the solution, and backtrack as soon as you determine that the current candidate cannot lead to a valid solution.
    
- **Pseudocode**:
    
    ```
    function solveSudoku(board):
        if board is complete:
            return true
        for each empty cell in board:
            for each possible value:
                if placing value is valid:
                    place value
                    if solveSudoku(board):
                        return true
                    remove value (backtrack)
        return false
    ```
    
- **Example**: The **N-Queens Problem**, where you need to place queens on a chessboard so that no two queens threaten each other, can be solved with backtracking.
    

**Additional Details**:

- **Complexity**: O(N!) in the worst case, due to exploring all possible configurations.
- **Applications**: Puzzles (Sudoku, N-Queens), generating all possible subsets or permutations, and solving constraint satisfaction problems.
- **Approach**: Backtracking is often combined with pruning to cut off invalid branches early, making it more efficient.

#### 4. **Greedy Algorithm**

- **Use Case**: Useful for optimization problems where local decisions can lead to the global optimum, such as finding the minimum number of coins for change.
    
- **Logic**: Make the optimal choice at each step, hoping that these local decisions will lead to the best overall solution.
    
- **Pseudocode**:
    
    ```
    function coinChange(coins, amount):
        coins.sort(descending)
        count = 0
        for coin in coins:
            while amount >= coin:
                amount -= coin
                count += 1
        return count if amount == 0 else -1
    ```
    
- **Example**: The **Coin Change Problem** can be approached with a greedy algorithm if the denominations are such that local optimization leads to the best global solution.
    

**Additional Details**:

- **Complexity**: O(n log n) for sorting plus O(n) for iteration.
- **Applications**: Minimum spanning trees (Kruskal's algorithm), Huffman coding, activity selection problems.
- **Limitations**: Not all problems can be solved optimally with a greedy approach; it only works when local optima are guaranteed to lead to a global optimum.

#### 5. **Depth-First Search (DFS)**

- **Use Case**: Suitable for tree or graph traversal, especially when you need to explore all possible paths, such as maze-solving or generating combinations.
    
- **Logic**: Start at the root and explore as far as possible along each branch before backtracking.
    
- **Pseudocode**:
    
    ```
    function dfs(node, visited):
        if node is null or node in visited:
            return
        visit(node)
        visited.add(node)
        for neighbor in node.neighbors:
            dfs(neighbor, visited)
    ```
    
- **Example**: **Maze Solving** can use DFS to explore all possible routes until it finds the correct path to the destination.
    

**Additional Details**:

- **Complexity**: O(V + E) where V is the number of vertices and E is the number of edges.
- **Applications**: Finding connected components in a graph, topological sorting, and solving puzzles like mazes.
- **Approach**: Often implemented recursively or with a stack for non-recursive traversal.

#### 6. **Heap/Priority Queue**

- **Use Case**: Ideal for problems involving dynamically fetching the minimum or maximum element, such as Dijkstra's algorithm or finding the k-largest elements.
    
- **Logic**: Use a heap to efficiently maintain the top elements in a dataset.
    
- **Pseudocode**:
    
    ```
    function findKSmallestElements(nums, k):
        heap = new MinHeap()
        for num in nums:
            heap.add(num)
            if heap.size() > k:
                heap.remove()
        return heap.elements()
    ```
    
- **Example**: The **Kth Largest Element** problem can be efficiently solved by maintaining a heap of the top k elements.
    

**Additional Details**:

- **Complexity**: O(n log k) for maintaining the heap of size k for n elements.
- **Applications**: Scheduling tasks, implementing priority queues, and algorithms like Prim's and Dijkstra's for shortest path.
- **Approach**: In Go, you can use the `container/heap` package to implement custom heaps for different use cases.

#### 7. **Union-Find (Disjoint Set Union - DSU)**

- **Use Case**: Best for problems that involve grouping elements and determining whether they are connected, such as finding connected components in a graph.
    
- **Logic**: Use a parent array to track the representative (root) of each element, and perform union operations to merge groups.
    
- **Pseudocode**:
    
    ```
    function find(x, parent):
        if parent[x] != x:
            parent[x] = find(parent[x], parent) # Path compression
        return parent[x]
    
    function union(x, y, parent, rank):
        rootX = find(x, parent)
        rootY = find(y, parent)
        if rootX != rootY:
            if rank[rootX] > rank[rootY]:
                parent[rootY] = rootX
            elif rank[rootX] < rank[rootY]:
                parent[rootX] = rootY
            else:
                parent[rootY] = rootX
                rank[rootX] += 1
    ```
    
- **Example**: The **Connected Components in a Graph** problem can be solved using the Union-Find algorithm to efficiently determine if nodes are in the same component.
    

**Additional Details**:

- **Complexity**: Nearly O(1) per operation with path compression and union by rank.
- **Applications**: Network connectivity, detecting cycles in undirected graphs, and Kruskal's algorithm for finding minimum spanning trees.
- **Approach**: The combination of **path compression** and **union by rank** ensures that each operation is extremely efficient.

### Summary

- **Binary Search**: Efficient for searching in sorted datasets by repeatedly dividing the search space.
- **Dynamic Programming**: Best for optimization problems involving overlapping subproblems, storing intermediate results for reuse.
- **Backtracking**: Used in problems involving constraints, where incremental decisions need to be verified, and invalid solutions are undone.
- **Greedy Algorithm**: Make the locally optimal choice at each step, useful when local decisions yield the global optimum.
- **DFS**: Suitable for deep exploration of trees or graphs, often used in scenarios where all paths must be explored.
- **Heap/Priority Queue**: Maintains a dynamic set of top elements, useful in scenarios involving efficient retrieval of minimum or maximum elements.
- **Union-Find**: Efficiently handles grouping and connectivity queries, best for graph problems involving connected components.
