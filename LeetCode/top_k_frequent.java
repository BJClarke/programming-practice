/*

Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

*/
class Solution {
    public List<Integer> topKFrequent(int[] nums, int k) {
        List<Integer> solutions = new ArrayList<>();
        HashMap<Integer, Integer> m = new HashMap<Integer, Integer>();
        
        for (int i = 0; i < nums.length; i++) {
            int count = m.getOrDefault(nums[i], 0);
            m.put(nums[i], count+1);
        }
        
        List[] buckets = new List[nums.length+1];
        for (Map.Entry<Integer, Integer> entry : m.entrySet()) {
            int key = entry.getKey();
            int count = entry.getValue();
            if (buckets[count] == null) {
                buckets[count] = new ArrayList<Integer>();
            }
            buckets[count].add(key);
        }
        
        for (int i = nums.length; i >= 0; i--) {
            if (buckets[i] != null) {
                for (int num : (List<Integer>)buckets[i]) {
                    solutions.add(num);
                    if (--k == 0) {
                        return solutions;
                    }
                }
            } 
        }
        
        return solutions;
    }
}