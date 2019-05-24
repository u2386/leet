# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):

    def printTree(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[str]]
        """
        if not root:
            return []

        depth = 0
        nodes = []
        q, p = [root], []
        while True:
            flag = False
            nodes.append([])
            for n in q:
                if n:
                    flag = True
                    nodes[-1].append(str(n.val))
                    p.extend([n.left, n.right])
                    continue
                nodes[-1].append('')
                p.extend([None, None])
            if not flag:
                nodes = nodes[:-1]
                break
            depth += 1
            q = p[:]
            p[:] = []

        ret = []
        for i in range(len(nodes)):
            ret.append([])
            padding = [''] * (2 ** (depth - i - 1) - 1)
            for j in range(len(nodes[i])):
                ret[-1].extend(padding)
                ret[-1].append(nodes[i][j])
                ret[-1].extend(padding)
                if j < len(nodes[i]) - 1:
                    ret[-1].append('')
        return ret
