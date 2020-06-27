import java.lang.*;
class Solution{
	class TreeNode{
		int val;
		TreeNode left;
		TreeNode right;
		TreeNode(int x){val = x;}
	}
	
	//offer 55
	public int maxDepth(TreeNode root){
	
		if(null == root){
			return 0;
		}

		return Math.max(1+maxDepth(root.left),1+maxDepth(root.right));

	}

	//offer 55-2
	public boolean isBalanced(TreeNode root) {
		if(root == null || (root.left == null && root.right == null)){
			return true;
		}

		int leftHeight = maxDepth(root.left);
		int rightHeight = maxDepth(root.right);
		
		int diff = leftHeight - rightHeight;

		if(diff > 1 || diff < -1){
			return false;
		}

		return isBalanced(root.left) && isBalanced(root.right);
    }

}
