package main

func twoSum(numbers []int, target int) []int {
	//создаем мапу бля индексов массива и их значений
	seen := make(map[int]int)
	for i, n := range numbers {
		remind := target - n           //считаем разницу текущего элемента и таргета
		if _, ok := seen[remind]; ok { //есть ли индекс этого значения в мапе
			// если есть то возвращаем индексы элементов
			if i > seen[remind] {
				return []int{seen[remind], i}
			} else {
				return []int{i, seen[remind]}
			}
		} else {
			seen[n] = i //если в мапе нет разницы, добавляем в мапу индекс
		}
	}
	return []int{}
}
