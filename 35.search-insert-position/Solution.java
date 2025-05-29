class Solution {
    public int searchInsert(int[] nums, int target) {
        int s = 0;
        int l = nums.length-1;
        while(s<=l){
            int mid = s + (l-s)/2;
            if(nums[mid]==target) return mid;
            else if(nums[mid]>target) l=mid-1;
            else s=mid+1;
        }
        return s;
    }
}