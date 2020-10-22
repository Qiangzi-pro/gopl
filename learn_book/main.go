package main

import (
	"fmt"
	"sync"
)


var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	//fmt.Println(runtime.NumCPU())
	//fmt.Println(strconv.Itoa())
	//reader := bufio.NewReader(os.Stdin)

	//temDirs := []string{}
	//for _, dir := range temDirs {
	//	dir := dir
	//	fmt.Println(dir)
	//}
	//
	//const day = 24 * time.Hour
	//
	//p := geometry.Point{1, 2}
	//pptr := &p
	//pptr.ScaleBy(2)
	//fmt.Println(pptr)

	//var initial [64]byte
	//var buf []byte
	//fmt.Println(initial)
	//fmt.Println(buf==nil, len(buf), cap(buf))
	//buf = initial[:1]
	//fmt.Println(buf, len(buf), cap(buf))
	//day := 24 * time.Hour
	//fmt.Println(day.Seconds())
	//
	//var err error = syscall.Errno(2)
	//e2 := syscall.Errno(3)
	//fmt.Println(err)
	//fmt.Println(e2)
	//
	//fmt.Println(math.Pow(2, 3))
	//
	//var z float64
	//fmt.Println(z, -z, 1/z, -1/z, z/z)
	//math.NaN()
	//
	//var w io.Writer
	// w = os.Stdout
	// f, ok := w.(*os.File)
	// fmt.Println(ok)
	// fmt.Printf("f=%[1]v, %[1]T\n", f)
	// fmt.Printf("w=%[1]v, %[1]T\n", w)

	 sqlQuote("fascinating")
	sqlQuote(11)
	 fmt.Println()

}

func sqlQuote(x interface{}) {

	switch x := x.(type) {
	case string:
		fmt.Printf("x=%[1]v, %[1]T, %p\n", x, &x)
		//f := x.(int)
		//fmt.Println(f)
	case int, uint:

		fmt.Printf("x=%[1]v, %[1]T, %p\n", x, &x)
		x = 2
		f := x.(int)
		fmt.Println(f)
	default:
		panic("err")
	}
	fmt.Printf("x=%[1]v, %[1]T, %p\n", x, &x)
}