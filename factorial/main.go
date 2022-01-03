package main

func FactorialFor(num int) (o int) {

	if num <= 1 {
		return 1
	}

	temp := 1
	for i := 1; i <= num; i++ {
		temp = temp * i
	}

	return temp
}

func FactorialReqursive(num int) (o int) {

	if num <= 1 {
		return 1
	}
	return num * FactorialReqursive(num-1)

}

func FactorialReqursiveTail(total int, num int) (o int) {

	if num <= 1 {
		return total
	}
	return FactorialReqursiveTail(total*num, num-1)

}
