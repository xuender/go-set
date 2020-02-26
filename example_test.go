package set

import "fmt"

func ExampleDBSet() {
	s := NewDBSet().Add([]byte("str1"), []byte("str2"))
	defer s.Release()

	i := s.Iteration()
	defer i.Release()

	for i.Next() {
		fmt.Println("value:", i.Key())
	}
	// Output:
	// value: [115 116 114 49]
	// value: [115 116 114 50]
}
