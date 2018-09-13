/*

Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note: 
You may assume k is always valid, 1 ≤ k ≤ array's length.

*/
class Solution {
    public int findKthLargest(int[] nums, int k) {
        int min = nums[0];
        int max = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < min) {
                min = nums[i];
            }
            if (nums[i] > max) {
                max = nums[i];
            }
        }
        
        int[] buckets = new int[max-min+1];
        for (int i = 0; i < nums.length; i++) {
            buckets[nums[i]-min]+=1;
        }
        
        for (int i = buckets.length-1; i >= 0; i--) {
            k -= buckets[i];
            if (k <= 0) {
                return i+min;
            }
        }
        
        return -1;
    }
}