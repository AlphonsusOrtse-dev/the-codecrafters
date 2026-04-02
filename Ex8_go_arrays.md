# Go Array
An array is a fixed size list of items that are of  the same data type. this means that when the size of an array is created, it can not be changed.
In Go the length of an arrary is either defined or inferred.

### For example:
var numbers = [5]int{1, 2, 3, 4, 5} 
* Here, this means that then length of the array is defined with a value = 5.

var numbers = [...]int{1, 2, 3, 4, 5} 
* when you don't want to manually count the numbers, you just type [...] and Go look at the length and count it for you. this is done to avoid making mistate.


var numbers = [5]int{1, 2, 3}
* here, even though only (3) values are provided Go automatically prints 5 values adding (00)to the 3 values provided because the size of an array is fix and can not be changed.