Most compiled languages (is there any?) doesn't have `insert` as a method for arrays, that is, in fact, one of its core components: fixed size. 

What I found interesting is how dynamic array-like data structures deal with this problem at runtime. To ilustrate this, when a list needs resize: when a python list is created, memory is allocated dynamically by CPython to store the list's elements. This memory allocation happens at runtime. The size of the initial memory block allocated for the list can vary depending on factors such as the initial size of the list.

Truly inserting a new element in a array in compiled languages would mean to shift the position of all the right side elements of the index of the element being inserted. This means losing one or more element of an array. For example: if I had an array of [3]int{1, 2, 3} and wanted to insert "5" at index "1", the final array would be: [3]int{1, 5, 2}. Three would be gone. 

While we can't insert, we can update. Updating is not inserting. 

So, arrays seem boring to implement, but they are really interesting conceptually. 
