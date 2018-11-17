# HOMEWORK 9: 
#
### code: complete 'insertionsort' and 'mergesort'.
### written: analyze 'mergesort' and 'binarysearch'.


# quicksort(xs):
#
# Sorts an array using the famous "divide-and-conquer"
# method.
#
def quicksort(xs):

    def choose_pivot(l,r):
        # Find the median of three values.
        c = (l + r) // 2
        if xs[c] < xs[l] < xs[r] or xs[r] < xs[l] < xs[c]:
            return l
        elif xs[l] < xs[c] < xs[r] or xs[r] < xs[c] < xs[l]:
            return c
        else:
            return r

    def partition(l,r,p):
        xs[r], xs[p] = xs[p], xs[r]
        i = l
        for j in range(l,r):
            if xs[j] < xs[r]:
                xs[i], xs[j] = xs[j], xs[i]
                i = i + 1
        xs[r], xs[i] = xs[i], xs[r]
        return i

    # The actual recursive (helper) function.
    def qsort(l, r):
        if l < r:
            p = choose_pivot(l,r)
            m = partition(l,r,p)
            qsort(l, m-1)
            qsort(m+1, r)

    # Sort the whole list.
    qsort(0,len(xs)-1)



# bubblesort(xs):
#
# Repeatedly makes passes over the list, swapping
# neighboring out-of-order pairs of values.
#
def bubblesort(xs):
 
    n = len(xs)

    # Make several passes.
    for p in range(1,n):

        # Scan the list, examining a 2-element portion of the
        # list, from left to right
        i = 0
        while i < n-p:
            # If the values at i and i+1 are out of order...
            if xs[i] > xs[i+1]:
                # ...then swap them.
                xs[i], xs[i+1] = xs[i+1], xs[i]

            # Move one position to the right. 
            i = i + 1

        # Note that this places the (n-p)-th element into
        # its final resting place.


# selectsort(xs):
#
# Repeatedly "selects" the minimum value of an unsorted
# region of the list, placing this value just to the 
# right of the sorted region, extending the sorted
# region.
#
def selectsort(xs):

    n = len(xs)

    # Extend the sorted region with each step of the 
    # loop below.
    #
    for i in range(n):

        # One step places the i-th smallest value into xs[i].

        # To so so, find the minimum of xs[i:n].
        min = xs[i]
        minj = i
        for j in range(i+1,n):
            # Compare xs[j] with the minimum value found so far.
            if xs[j] < min:
                # Update our current sense of the minimum.
                min = xs[j]
                minj = j    # Keep track of where it was found.

        # At this point, minj is the location of the 
        # minimum of xs[i:n].

        # Swap the values at index i and index minj.
        xs[i],xs[minj] = xs[minj],xs[i]



# EXERCISE 1
#
def insertionsort(xs):
    i=1
    while i<(len(xs)):
        x=0
        while x<=len(xs[:i]):
            if xs[i]<xs[x]:
                n=xs[i]
                del xs[i]
                xs.insert(x,n)
                x=len(xs)
            else:
                x+=1
        i+=1

            


# EXERCISE 2 (and 3)
#
# merge(ys,zs):
#
# Given two sorted lists 'ys' and 'zs', build a third sorted 
# list of their items.
#
def merge(ys,zs):
    xs = []
    lz=len(zs)
    ly=len(ys)
    z=0
    y=0
    while z<lz and y<ly:
        if zs[z]<=ys[y]:
            xs.append(zs[z])
            z+=1
        else:
            xs.append(ys[y])
            y+=1
    while y<ly:
        xs.append(ys[y])
        y+=1
    while z<lz:
        xs.append(zs[z])
        z+=1
    #
    # Your code for merging goes here...
    #
    # You'll append onto 'xs' the items from 'ys' and 'zs' 
    # in the correct order
    # 
    #
    return xs
#
#
def mergesort(xs):

    # Recursive helper function
    def ms(ls):
        if len(ls) < 2:
            return ls
        else:
            m = len(ls) // 2
            left = ms(ls[:m])
            right = ms(ls[m:])
            return merge(left,right)
    
    # Make a copy of the list.
    cp = xs.copy()

    # Sort the copy.
    sorted = ms(cp)

    # Place the values in the right order into 'xs'.
    i = 0
    for x in sorted:
        xs[i] = x
        i = i + 1
    

# binarysearch(k,xs):
#
# Determine whether 'k' is in the list 'xs'.
#
def binarysearch(k,xs):
    if len(xs) == 0:
        return False
    else:
        m = len(xs) // 2
        if xs[m] == k:
            return True
        else:
            if k < xs[m]:
                return binarysearch(k,xs[:m])
            else:
                return binarysearch(k,xs[m+1:])

# permutation(n):
#
# Produce a list of the values 0..n-1, randomly arranged.
#
def permutation(n):
    import random
    xs = [ i for i in range(n) ]
    for i in range(n-1):
        r = random.randint(i,n-1)
        xs[r], xs[i] = xs[i], xs[r]
    return xs


# Test the three sorts above.
def testsort(ls):
    msort = mergesort(ls)
    isort = insertionsort(ls)
    bsort = bubblesort(ls)
    ssort = selectsort(ls)
    qsort = quicksort(ls)
    check = ls.sort()
    if (msort != check):
    	print("mergesort incorrect")
    if (isort != check):
    	print("insertionsort incorrect")
    if (bsort != check):
    	print("bubblesort incorrect")
    if (ssort != check):
    	print("selectionsort incorrect")
    if (qsort != check):
    	print("quicksort incorrect")

i = 0
while i<100:
    ls = permutation(10)
    testsort(ls)
    i = i + 1
    


