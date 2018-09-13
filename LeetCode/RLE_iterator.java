/*

Write an iterator that iterates through a run-length encoded sequence.

The iterator is initialized by RLEIterator(int[] A), where A is a run-length encoding of some sequence.  More specifically, for all even i, A[i] tells us the number of times that the non-negative integer value A[i+1] is repeated in the sequence.

The iterator supports one function: next(int n), which exhausts the next n elements (n >= 1) and returns the last element exhausted in this way.  If there is no element left to exhaust, next returns -1 instead.

For example, we start with A = [3,8,0,9,2,5], which is a run-length encoding of the sequence [8,8,8,5,5].  This is because the sequence can be read as "three eights, zero nines, two fives".

 

Example 1:

Input: ["RLEIterator","next","next","next","next"], [[[3,8,0,9,2,5]],[2],[1],[1],[2]]
Output: [null,8,8,5,-1]
Explanation: 
RLEIterator is initialized with RLEIterator([3,8,0,9,2,5]).
This maps to the sequence [8,8,8,5,5].
RLEIterator.next is then called 4 times:

.next(2) exhausts 2 terms of the sequence, returning 8.  The remaining sequence is now [8, 5, 5].

.next(1) exhausts 1 term of the sequence, returning 8.  The remaining sequence is now [5, 5].

.next(1) exhausts 1 term of the sequence, returning 5.  The remaining sequence is now [5].

.next(2) exhausts 2 terms, returning -1.  This is because the first term exhausted was 5,
but the second term did not exist.  Since the last term exhausted does not exist, we return -1.


*/
class RLEIterator {
    
    // index of the number we are pointing at
    int currIndex;

    // offset within that number
    // so if the currIndex points to 1, and there are twenty 1's,
    //  this number will be less than 20
    int offset;

    // the RLE array
    int[] arr;

    public RLEIterator(int[] A) {
        this.arr = A;
        this.currIndex = 0;
        this.offset = -1;
    }
    
    public int next(int n) {
        // we've already reached the end of the array
        if (this.currIndex >= this.arr.length) {
            return -1;
        }

        // get new index to query
        int index = this.offset + n;
        int num = 0;

        // if this new offset is less than the count of the current number, 
        //  use it and increase the offset
        if (index < this.arr[currIndex]) {
            num = this.arr[this.currIndex+1];
            this.offset += n;
        } else {
            // else find the next number by subtracting the count of each successive number from the new offset
            while (this.currIndex < this.arr.length && index >= arr[this.currIndex]) {
                index -= arr[this.currIndex];
                this.currIndex+=2;
            }
            if (this.currIndex >= this.arr.length) {
                return -1;
            }
            num = this.arr[this.currIndex+1];
            this.offset = index;
        }
        return num;
    }
}