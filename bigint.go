package bigint

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Bigint ...
type Bigint struct {
	Value string
}

var (
	// ErrorNotNumber used when input consists of not only numbers
	ErrorNotNumber = errors.New("input not a number")
)

func validation(num string) (string, bool) {
	// TODO
	allowed := "1234567890"
	var err bool

	if strings.HasPrefix(num, "+") {
		num = strings.Replace(num, "+", "", 1)
	}
	for strings.HasPrefix(num, "0") {
		num = strings.Replace(num, "0", "", 1)
	}
	arr := strings.Split(num, "")
	for _, v := range arr {
		if !strings.Contains(allowed, v) {
			err = true
		}
	}
	if strings.HasPrefix(num, "-") {
		err = false
	}

	return num, err
}

// NewInt returns new Bigint Value based on input
func NewInt(num string) (Bigint, error) {
	// TODO
	// step 1: validation
	// step 2: clean up 0000123 -> 123 | +123 -> 123 | +-+123 -> error
	// step 3: insert Value to the bigint struct and return   Bigint{Value: num}
	num, err := validation(num)
	if err {
		return Bigint{Value: num}, ErrorNotNumber
	} else {
		return Bigint{Value: num}, nil
	}
}

// Set methods updates Value
func (z *Bigint) Set(num string) error {
	// TODO
	// step 1: validation
	// step 2: clean up
	// step 3: set new Value
	num, err := validation(num)
	if err {
		return ErrorNotNumber
	}
	z.Value = num
	return nil
}

// Add ...
func Add(a, b Bigint) Bigint {
	// TODO
	// step 1:  cast - a.Value -> int32,  b.Value -> int32
	// step 2: add casted numbers
	// step 3: cast result to string
	// step 4: return result
	ans := add(a, b)
	if a.Value[0] != '-' && b.Value[0] == '-' {
		//b.Value = strings.Replace(b.Value, "-", "", 1)

		return sub(a, b)
	}
	// if a.Value[0] == '-' && b.Value[0] == '-' {
	// 	//b.Value = strings.Replace(b.Value, "-", "", 1)
	// 	ans := add(a, b)
	// 	return ans
	// }
	return ans
}
func add(a, b Bigint) Bigint {
	ans := Bigint{}
	first := strings.Split(a.Value, "")
	second := strings.Split(b.Value, "")
	isRest := false
	n1 := 0
	n2 := 0

	for i, j := len(first)-1, len(second)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		n1 = 0
		n2 = 0

		if i >= 0 {
			n1, _ = strconv.Atoi(first[i])
		}
		if j >= 0 {
			n2, _ = strconv.Atoi(second[j])
		}
		if isRest {
			n1 = n1 + 1
			isRest = false
		}
		if n1+n2 >= 10 {
			ans.Value = strconv.Itoa(n1+n2-10) + ans.Value
			isRest = true
		} else {
			ans.Value = strconv.Itoa(n1+n2) + ans.Value
			isRest = false
		}

	}
	if isRest {
		ans.Value = "1" + ans.Value
	}

	return ans
}

// Sub ...
func Sub(a, b Bigint) Bigint {
	// TODO
	// step 1:  cast - a.Value -> int32,  b.Value -> int32
	// step 2: sub casted numbers
	// step 3: cast result to string
	// step 4: return result

	ans := sub(a, b)
	if a.Value[0] != '-' && b.Value[0] == '-' {
		b.Value = strings.Replace(b.Value, "-", "", 1)
		return add(a, b)
	}
	return ans
}
func sub(a, b Bigint) Bigint {
	ans := Bigint{}
	first := strings.Split(a.Value, "")
	second := strings.Split(b.Value, "")
	isNeed := false
	n1 := 0
	n2 := 0

	for i, j := len(first)-1, len(second)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		n1 = 0
		n2 = 0

		if i >= 0 {
			n1, _ = strconv.Atoi(first[i])
		}
		if j >= 0 {
			n2, _ = strconv.Atoi(second[j])
		}
		if isNeed {
			n1 = n1 - 1
			isNeed = false
		}
		if n1-n2 >= 0 {
			ans.Value = strconv.Itoa(n1-n2) + ans.Value
			isNeed = false
		} else {
			ans.Value = strconv.Itoa(n1-n2+10) + ans.Value
			isNeed = true
		}

	}

	return ans
}

// Multiply ...
func Multiply(a, b Bigint) Bigint {
	// TODO
	// step 1:  cast - a.Value -> int32,  b.Value -> int32
	// step 2: multiply casted numbers
	// step 3: cast result to string
	// step 4: return result
	ans := Bigint{}
	n1, _ := strconv.Atoi(a.Value)
	n2, _ := strconv.Atoi(b.Value)
	mul := n1 * n2
	ans.Value = strconv.Itoa(mul)
	return ans
}

// Mod ...
func Mod(a, b Bigint) Bigint {
	// TODO
	// step 1:  cast - a.Value -> int32,  b.Value -> int32
	// step 2: mod casted numbers
	// step 3: cast result to string
	// step 4: return result
	ans := Bigint{}
	n1, _ := strconv.ParseFloat(a.Value, 32)
	n2, _ := strconv.ParseFloat(b.Value, 32)
	mul := int(math.Mod(n1, n2))
	ans.Value = strconv.Itoa(mul)
	return ans
}

// Abs ...
func (z Bigint) Abs() Bigint {
	// TODO
	// Abs
	ans := Bigint{}
	if strings.HasPrefix(z.Value, "-") {
		z.Value = strings.Replace(z.Value, "-", "", 1)
	}
	ans.Value = z.Value
	return ans
}
