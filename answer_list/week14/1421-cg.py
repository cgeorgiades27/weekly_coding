class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def serialize(node):

    def helper(node, arr):
        if node is None:
            return
        helper(node.right, arr)
        arr.append(node.val)
        helper(node.left, arr)
        return arr
    return ','.join(helper(node, []))


# def deserialize(node):
#    return true;
node = Node('root', Node('left', Node('left.left')), Node('right'))
print(serialize(node))

# assert deserialize(serialize(node)).left.left.val == 'left.left'
