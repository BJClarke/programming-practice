class Solution {
    public int calculate(String s) {
        // remove whitespace
        s = s.replaceAll("\\s+","");

        // begin at back because recursion will then cause the first operation to happen first
        //  and work out like a stack
        return calc(s, s.length()-1, 0);
    }
    
    // index represents the index of the next character to consider
    // this will initially be equal to s.length()-1 and will decrease except
    // it will be set to s.length()-1 when calculating groupings
    public int calc(String s, int index, int result) {
        if (index < 0) {
            return result;
        }
        
        char c = s.charAt(index);
        switch (c) {
            case ')':
                // get offset of matching beginning parenthesis
                int offsetOfBeginParenthesis = findOffsetOfBeginParenthesis(s, index);

                // calculate result of inner grouping and pass that back to calc as the result of an operation
                int parenthesisResult = calc(s.substring(index - offsetOfBeginParenthesis + 1, index),
                                             offsetOfBeginParenthesis-2, result);

                return calc(s, index - offsetOfBeginParenthesis - 1, parenthesisResult);
            case '+':
                return result+calc(s, index-1, result);
            case '-':
                return calc(s, index-1, result)-result;
            default:
                // get indices of number, may be more than a single digit
                int i = index;
                while (i > 0) {
                    if (s.charAt(i-1) == '(' ||
                        s.charAt(i-1) == ')' ||
                        s.charAt(i-1) == '+' ||
                        s.charAt(i-1) == '-') {
                        break;
                    }
                    i--;
                }
                return calc(s, i-1, Integer.parseInt(s.substring(i, index+1)));
        }
    }
    
    // don't forget to consider that there may be inner parenthesis
    private int findOffsetOfBeginParenthesis(String s, int index) {
        int numClosed = 0;
        for (int i = index-1; i >= 0; i--) {
            if (s.charAt(i) == '(') {
                if (numClosed == 0) {
                    return index-i;
                }
                numClosed--;
            } else if (s.charAt(i) == ')') {
                numClosed++;
            }
        }
        return -1;
    }
}