package main

//75. Sort Colors
// https://leetcode.com/problems/sort-colors/

import "fmt"

func sortColors(num []int) {
    for i:= 0 ; i < len(num); i++{
        for j:= i; j < len(num)-1; j++{
            if num[i] > num[j]{
                temp := num[i]
                num[i] = num[j]
                num[j] = temp
            }
        }
    }
}

func main() {
    n:=0
    fmt.Println("Размерность:")
    fmt.Scanln(&n)
    num := make([]int, n)
    for i:=0; i < n; i++{
        fmt.Scanln(&num[i])
    }
    fmt.Println("Массив:")
    fmt.Println(num)
    fmt.Println("Сортированный массив:")
    sortColors(num)
    fmt.Println(num)
}
