# Homework 10
#
# Binary search tree search and traversal.
#
# Complete the three methods/function definitions
# below:
#
# Ex 1. method 'min'.
# Ex 2. method 'depth'.
# Ex 3. function 'as_string', a helper function for method 'toString'.
class BiNode:

    def __init__(self,k):
        self.key = k
        self.left = None
        self.right = None


# class BSTree:
#
# A type that defines binary search tree objects.
#
class BSTree:
    

    # t = BSTree():
    #
    # Constructs a new empty binary search tree instance.
    #
    def __init__(self):
        self.root = None


    # t.isEmpty():
    #
    # Returns whether binary search tree 't' is empty.
    #
    def isEmpty(self):
        return (self.root == None)


    # b = t.contains(k):
    #
    # Returns whether key 'k' is in binary search tree 't'.
    #
    def contains(self, k):
        # Walk down the tree, moving left or 
        # right, looking for k.
        n = self.root
        while n != None:
            if n.key == k:
                return True
            elif n.key > k:
                n = n.left
            else:
                n = n.right
        return False

    # t.insert(k):
    #
    # Adds key 'k' to binary search tree 't' if it
    # is not yet in 't'.
    #
    def insert(self, k):
        p = None
        n = self.root

        # Walk through the search path for
        # k until we "fall off" while
        # maintaining a follower p.
        while n != None:
            p = n
            if n.key == k:
                return
            elif n.key > k:
                n = n.left
            else:
                n = n.right

        # Node p is now the parent of
        # the node we'll create.
        e = BiNode(k)
        if p == None:
            # No parent so tree must be empty.
            self.root = e
        elif p.key > k:
            # Hang the new node left of p.
            p.left = e
        else:
            # Hang the new node right of p.
            p.right = e


    # m = t.min():
    #
    # Returns the smallest key held in binary search tree 't'.
    # Returns 'None' if 't' is empty.
    #
    def min(self):
        if self.isEmpty():
            return None
        else:
            n=self.root
            p=None
            while n!=None:
                p=n
                n=n.left
            # Your code goes below here.
            return p.key


    # d = t.depth(k):
    #
    # Returns an integer describing the depth at
    # which key 'k' appears within tree 't'.  It
    # returns 'None' if 'k' is not in 't'.
    #
    # The key at the root node is at depth 0.
    # A key in a child node of the root is at depth 1.
    # A key in a root's granchild is at depth 2.
    # Etc.
    #
    def depth(self,k):
        if not self.contains(k):
            return None
        else:
            n=self.root
            depth=0
            while n.key!=k:
                if k<n.key:
                    depth+=1
                    n=n.left
                else:
                    depth+=1
                    n=n.right
            return depth
    # ks = t.asList():
    #
    # Returns a list of the keys in 't' in sorted order.
    #
    def asList(self):

        # list_of(n)
        #
        # Given a BiNode 'n', computes the list of the
        # keys held at and below 'n' in sorted order.
        #
        def list_of(n):

            if n == None:
                # Our subtree is empty, return an empty
                # list.
                #
                return []

            else:
                # Get the list of keys from the left node,
                # and the list of keys from the right node.
                # Append them together with this node's
                # key in the middle. Return that result.
                #
                ks = []
                ks += list_of(n.left)  # Traverse the left,
                ks += [n.key]          # then add our key,
                ks += list_of(n.right) # and then traverse the right.
                return ks

        # Use the list_of helper function to do all the work.
        return list_of(self.root)


    # s = t.asString():
    #
    # Returns a string representing the structure of 't'.
    #
    # An empty tree is represented by "."
    #
    # A balanced tree with three keys 1, 2, and 3 is represented by
    # "[2: [1: . .] [3: . .]]"
    #
    # A totally unbalanced three-node tree with children only to the
    # left, containing 3 at the root, 2 as its one child, and 1 as
    # its one grandchild is represented by
    # "[3: [2: [1: . .] .] .]"
    #
    # In general, a non-empty subtree rooted at node n has the string
    # representation suggested by "[n.key: n.left n.right]".
    #
    def asString(self):
        
        def string_of(n):
            if n == None:
                return "."
            else:
                return "["+str(n.key)+": "+str(string_of(n.left))+" "+ str(string_of(n.right))+ "]"   
        return string_of(self.root)

    __str__ = asString
    __repr__ = asString
    