# 2. Progression (coding task)

Write a program that finds the `n`th term (`n` is a non-negative integer) of
the progression

<img src="https://render.githubusercontent.com/render/math?math=\begin{cases}a_0 = 1\\a_1 = 2\\a_i = 4a_{i-2} %2b a_{i-1}\end{cases}">
   
1. Create the program in the directory `progression`.
1. Implement the logic in a function declared as `func nthTerm(n uint) int`.
1. Implement `main()` function that
   1. reads `n` from the keyboard
   1. calls `nthTerm` function
   1. prints a value returned from the function.
1. Create a file with tests for the program. The tests must cover all
   representative inputs.
   
## Example 1
   
```
Enter n: 1
The nth term is 2.
```
   
**Explanation**:

`nthTerm(1)` must return `2`, because <img src="https://render.githubusercontent.com/render/math?math=a_1 = 2">.
 
## Example 2
 
```
Enter n: 4
The nth term is 38.
```
    
`nthTerm(4)` must return `38`, because
    
<img src="https://render.githubusercontent.com/render/math?math=\begin{align*}a_0 %26= 1\\a_1 %26= 2\\ a_2 %26 = 4a_0 %2b a_1 = 4 \cdot 1 %2b 2 = 6\\a_3 %26 = 4a_1 %2b a_2 = 4 \cdot 2 %2b 6 = 14\\a_4 %26 = 4a_2 %2b a_3 = 4 \cdot 6 %2b 14 = 38\end{align*}">

# 3. Filter (coding task)

Write a program that removes all integers divisible by `5` from a slice of
integers.

1. Write the program in the directory `filter`.
1. Implement the logic in a function declared as `func filter(s []int) []int`.
1. Implement `main()` function that
   1. reads a slice from the keyboard (a user provides the number of elements
      in a slice and the element values)
   1. calls `filter` function
   1. prints the filtered slice.
1. Create a file with tests for the program. The tests must cover all
   representative inputs.

## Example 1

```
Enter the number of elements in a slice: 5
Enter the elements: 10 -15
The filtered slice: []
```

`filter([]int{-10, -15})` returns `[]int{}`, because all integer in the slice
are divisible by `3`, and are removed from the slice.

## Example 2

```
Enter the number of elements in a slice: 5
Enter the elements: -1 0 -3 9 4
The filtered slice: [-1 4]
```

`filter([]int{-1, 0, -3, 9, 4})` returns `[]int{-1, 4}`. `0`, `-3` and `9` are
integers divisible by `3` and are removed from the slice.
