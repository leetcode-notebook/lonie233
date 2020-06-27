import java.lang.*;
class Solution{
	class TreeNode{
		int val;
		TreeNode left;
		TreeNode right;
		TreeNode(int x){val = x;}
	}
	

	public int maxDepth(TreeNode root){
	
		if(null == root){
			return 0;
		}

		return Math.max(1+maxDepth(root.left),1+maxDepth(root.right));

	}

}
