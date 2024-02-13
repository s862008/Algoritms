package main

import "fmt"

func cocktailSort(arr []int) {
left := 0
right := len(arr) - 1
swapped := true

for left < right && swapped {
swapped = false

// Проход слева направо
for i := left; i < right; i++ {
if arr[i] > arr[i+1] {
arr[i], arr[i+1] = arr[i+1], arr[i]
swapped = true
}
}
right--

if !swapped {
break
}

// Проход справа налево
for i := right; i > left; i— {
if arr[i] < arr[i-1] {
arr[i], arr[i-1] = arr[i-1], arr[i]
swapped = true
}
}
left++
}
}

func main() {
arr := []int{64, 34, 25, 12, 22, 11, 90}
fmt.Println("Исходный массив:", arr)

cocktailSort(arr)
fmt.Println("Отсортированный массив:", arr)
}
