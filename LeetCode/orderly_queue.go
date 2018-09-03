/*

A string S of lowercase letters is given.  Then, we may make any number of moves.

In each move, we choose one of the first K letters (starting from the left), remove it, and place it at the end of the string.

Return the lexicographically smallest string we could have after any number of moves.


Example 1:

Input: S = "cba", K = 1
Output: "acb"
Explanation: 
In the first move, we move the 1st character ("c") to the end, obtaining the string "bac".
In the second move, we move the 1st character ("b") to the end, obtaining the final result "acb".


Example 2:

Input: S = "baaca", K = 3
Output: "aaabc"
Explanation: 
In the first move, we move the 1st character ("b") to the end, obtaining the string "aacab".
In the second move, we move the 3rd character ("c") to the end, obtaining the final result "aaabc".


Note:

1 <= K <= S.length <= 1000
S consists of lowercase letters only.

*/

/*

The idea here is that if K = 1, then there are len(S) different permutations,
    so we can easily iterate over then and decide the solution

If K > 1, there always exists a way to move characters to the back such that all permuations of the characters of S are possible.
    As a result, we can simply alphabetically sort the collection of chars that make up S

*/

func orderlyQueue(S string, K int) string {
    if K == 1 {
        result := S
        for i := 1; i < len(S); i++ {
            temp := S[i:] + S[:i]
            if temp < result {
                result = temp
            }
        }
        return result
    } else {
        cs := strings.Split(S, "")
        sort.Strings(cs)
        return strings.Join(cs, "")
    }
}