class Tree:
    # init the Tree Class as Root node left and right subnode
    def __init__(self, val=None):
        self.value = val
        if self.value:
            self.left = Tree()
            self.right= Tree()
        else:
            self.left = None
            self.right = None
    # If the value is None, then it is a empty tree
    def is_empty(self):
        return self.value == None

    def insert(self,data):
        # node with data is the root if the tree is empty
        if self.is_empty():
            self.value = data
            self.left = Tree()
            self.right = Tree()
        # insert to the left subtree if the node data is greater than the inserted data
        elif self.value > data:
            self.left.insert(data=data)
            return
        # insert to the right subtree if the node data is less than the inserted data
        elif self.value < data:
            self.right.insert(data=data)
            return
        # do not insert the data if node values is equal to the data; do nothing just return
        elif self.value == data:
            return

    def isleaf(self) -> bool:
        # if the left and right node is empty, then the node is a leaf node
        if self.left.is_empty() and self.right.is_empty():
            return True
        return False

    def find(self,v) -> bool:
        # not found if the node is empty
        if self.is_empty():
            print("{} is not found".format(v))
            return False
        # found if the node val matches
        if self.value == v:
            print("found")
            return True
        # go to left if val greater than the v
        if self.value > v:
            print("path ",self.value)
            return self.left.find(v)
        # go to right if val less than the v
        else:
            print("path ",self.value)
            return self.right.find(v)

    def delete(self, v):
        if self.is_empty():
            print("could not find the item to delete")
        if self.value == v :
            # if it is leaf node, then just delete it
            if self.isleaf():
                return Tree()
            else:
                if self.left.is_empty():
                    return  self.right
                elif self.right.is_empty():
                    return self.left
                else:
                    # get the right most node of left subtree
                    subnode = self.left
                    if subnode.right.is_empty():
                        subnode.right = self.right
                        return subnode
                    tmp = None
                    while not subnode.right.is_empty():
                        tmp = subnode
                        subnode = subnode.right
                    # point parent's right to childnode's left no matter it is empty or noe
                    tmp.right = subnode.left
                    subnode.left = self.left
                    subnode.right = self.right
                    return subnode
        elif self.value > v:
            self.left = self.left.delete(v)
        else:
            self.right = self.right.delete(v)
        return self


        
    def in_order(self) -> list:
        items = []
        self._in_order(items=items)
        return items
    def _in_order(self, items : list):
        # in_orde is the left node right
        if not self.is_empty():
            self.left._in_order(items)
            items.append(self.value)
            self.right._in_order(items)



t = Tree(15)
print(t.right)
print(t.right.right)
t.insert(34)
t.insert(24)
t.insert(14)
t.insert(35)
t.insert(38)
t.insert(3)
t.insert(21)
t.insert(26)
t.insert(25)
t.delete(25)
t.delete(34)
t.delete(26)
t.delete(35)
t.find(24)
# print(t.right)
# print(t.right.right)
items = t.in_order()
print(items)
for item in items:
    print(item)
