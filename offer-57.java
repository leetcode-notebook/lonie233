import java.lang.*;
class Solution{

	public static void main(String[] args){
		System.out.println("hello");
	}

	public int[] twoSum(int[] nums, int target) {
		int i = 0, j = nums.length-1;
		while(i<j) {
			int s = nums[i] + nums[j];
			if(s < target) {
				i++;
			}else if(s > target){
				j--;
			}else{
				return new int[] {nums[i], nums[j]};
			}
		}

		return new int[0];
    }

}
