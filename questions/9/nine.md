Сколько существует способов задать переменную типа slice или map?

* map: 
    - m := map[int]int{0: 0, 1: 10}
    - m := map[int]int{}
    - m := make(map[int]int, 0)

* slice:
    - s := []int{}
    - s := []int{0, 1, 3, 7}
    - s := make([]int, 0)