package main

import (
	"fmt"
	"time"
)

const SOME_CONST string = "constant"

func main() {
	fmt.Println("hello world");
	fmt.Println("and &&", true && false);
	fmt.Println("or ||", true || false);

	// variables
	var b, c int = 1, 2
	fmt.Println(b, c)
	f := "apple"
	fmt.Println(f)
	fmt.Println("constant", SOME_CONST)

	// for loop
	for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
		}
		if n==5 {
			break
		}
        fmt.Println(n)
	}

	// switch case
	t := time.Now()
    switch {
    case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() >= 20:
		fmt.Println("It's night")
    default:
        fmt.Println("It's after noon")
	}

	// array
	var d [5]int
	d[4] = 100
	fmt.Println("emp:", d, d[2:5], len(d))

	e := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", e)
	
	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
	fmt.Println("2d: ", twoD)
	
	// slices
	letters := []string{"a", "b", "c", "d"}
	fmt.Println("letters:", letters, len(letters))
	le := []string{}
	fmt.Println("le:", le, len(le))
	sl := e[:] // a slice referencing the storage of x
	fmt.Println("sl:", sl, len(sl))
	s := make([]string, 3)
	s[0] = "a"
	s[2] = "c"
	s = append(s, "b")
	fmt.Println("s:", s, len(s))

	s1 := make([]string, len(s))
    copy(s1, s)
	fmt.Println("s1:", s1)
	s2 := s[2:5]
	fmt.Println("s2:", s2)

	s3 := make([]int, 0)
	s3 = append(s3, 1)
	fmt.Println("s3:", s3)

	s4 := []int{}
	s4 = append(s4, 2)
	fmt.Println("s4:", s4)

	// map
	m := make(map[string]int)
	m["yo"] = 1
	m["k"] = 2
	delete(m, "yo")
	fmt.Println("m", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	
	// range
	sum := 0
    for _, num := range e {
        sum += num
    }
	fmt.Println("sum:", sum)
	
	for k, v := range n {
        fmt.Printf("%s -> %d\n", k, v)
    }
}
