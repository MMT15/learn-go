package utils

type Numeric interface{
	~int | ~float64 | ~string
}

func Sum[T Numeric](numbers ...T) T {
	var sum T
	for _,i:=range numbers {
		sum+=i
	}
	return sum
}

