/*

Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.

*/

class Solution {
    
    public boolean exist(char[][] board, String word) {
        if (word.length() == 0) {
            return true;
        }
        char ch = word.charAt(0);
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (board[i][j] == ch) {
                    if (findWord(board, i, j, word)) {
                        return true;
                    }
                }
            }
        }
        return false;
    }
    
    private boolean findWord(char[][] board, int i, int j, String word) {
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length || board[i][j] == '@') {
            return false;
        }
        char ch = word.charAt(0);
        if (board[i][j] == ch) {
            if (word.length() == 1) {
                return true;
            }
            board[i][j] = '@';
            boolean isWordFound = findWord(board, i+1, j, word.substring(1)) ||
                findWord(board, i-1, j, word.substring(1)) ||
                findWord(board, i, j+1, word.substring(1)) ||
                findWord(board, i, j-1, word.substring(1));
            board[i][j] = ch;
            if (isWordFound) {
                return true;
            }
        }
        return false;
    }
}