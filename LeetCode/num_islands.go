/*

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.


Example 1:

Input:
11110
11010
11000
00000

Output: 1


Example 2:

Input:
11000
11000
00100
00011

Output: 3

*/

package main

import (
    "bufio"
    "os"
    "strings"
    "fmt"
)

var moves = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    grid := make([][]string, 0)

    for scanner.Scan() {
        line := scanner.Text()

        if line == "DONE" {
            break
        }

        arr := make([]string, 0)

        nums := strings.Split(line, "")

        for _, num := range nums { 
            arr = append(arr, num)
        }

        grid = append(grid, arr)

        if len(arr) != len(grid[0]) {
            fmt.Println("Error has occurred: The length of all rows should be equal")
            return
        }
    }

    num := numIslands(grid)

    if num == 1 {
        fmt.Println("\nThere is 1 island in the grid")
    } else {
        fmt.Println(fmt.Sprintf("\nThere are %d islands in the grid.", num))
    }
}

func numIslands(grid [][]string) int {
    if len(grid) == 0 {
        return 0
    }
    
    visited := make([][]bool, len(grid))
    for i := 0; i < len(grid); i++ {
        visited[i] = make([]bool, len(grid[0]))
    }
    
    numIslands := 0
    
    for i, row := range grid {
        for j := range row {
            if grid[i][j] == "1" && !visited[i][j] {
                setVisited(grid, i, j, visited)
                numIslands++
            }
        }
    }
    
    return numIslands
}

func setVisited(grid [][]string, i, j int, visited[][]bool) {
    queue := make([][]int, 0)
    queue = append(queue, []int{i, j})
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            loc := queue[0]
            queue = queue[1:]
            
            for _, move := range moves {
                a := loc[0] + move[0]
                b := loc[1] + move[1]
                
                if a < 0 || a >= len(visited) || b < 0 || b >= len(visited[0]) || grid[a][b] == "0" || visited[a][b] {
                    continue
                }
          
                queue = append(queue, []int{a, b})
                visited[a][b] = true
            }
        }
    }
}