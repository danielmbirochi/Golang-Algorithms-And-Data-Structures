<div class="html">
<p>
    Given a binary tree, write a function to test if the tree is a binary search
    tree.
</p>
<p>You can assume that there will only be integers value.</p>
<p>
  Each <span>BST</span> node has an integer <span>value</span>, a
  <span>left</span> child node, and a <span>right</span> child node. A node is
  said to be a valid <span>BST</span> node if and only if it satisfies the BST
  property: its <span>value</span> is strictly greater than the values of every
  node to its left; its <span>value</span> is less than or equal to the values
  of every node to its right; and its children nodes are either valid
  <span>BST</span> nodes themselves or <span>None</span> / <span>null</span>.
</p>
<h3>Sample Input</h3>
<pre><span class="CodeEditor-promptParameter">tree</span> =   5
       /     \
      2       7
    /   \   /   \
   1     3 6     8
</pre>
<h3>Sample Output</h3>
<pre>true</pre>
</div>