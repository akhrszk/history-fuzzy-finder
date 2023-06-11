package app

func reduce[T any, K any](src []T, initial K, fn func(T, K) K) K {
	var k = initial
	for _, t := range src {
		k = fn(t, k)
	}
	return k
}
