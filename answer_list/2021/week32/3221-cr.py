from collections import deque


class Node:

    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


def get_level(root, level):
    """
    Does a BFS of the tree using a list of queues [queue].
    When the count reaches the desired level/height
    the current node queue is returned.
    """
    node_qu = deque()
    queue = [deque([root])]
    count = 0

    while queue:
        cur_queue = queue.pop()
        count += 1
        if count == level:
            return [val.value for val in cur_queue]

        while cur_queue:
            cur_node = cur_queue.popleft()

            if cur_node.left is not None:
                node_qu.append(cur_node.left)

            if cur_node.right is not None:
                node_qu.append(cur_node.right)

        queue.append(node_qu.copy())
        node_qu.clear()

    return None


a = Node(1)
a.left = Node(2)
a.right = Node(3)
a.left.left = Node(4)
a.left.right = Node(5)
a.right.right = Node(7)

print(get_level(a,3))
