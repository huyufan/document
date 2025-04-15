package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"sync"
	"time"
)

type task struct {
	id  int
	str string
}

func Or(chs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var once sync.Once
		for _, c := range chs {
			go func(c <-chan interface{}) {
				select {
				case <-c:
					once.Do(func() {
						close(out)
					})
				case <-out:
				}
			}(c)
		}
	}()
	return out
}

const (
	x = iota
	_
	y
	z = "pi"
	k
	p = iota
	q
)

func inc(p *int) int {
	*p++
	return *p
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

type Person struct {
	age int
}
type S struct {
}

func m(x interface{}) {
}

func g(x *interface{}) {
}

type People interface {
	Speaks(string) string
}

type Students struct{}

func (stu *Students) Speaks(think string) string {
	fmt.Println(think)
	return think
}

const (
	a = iota
	b = iota
)
const (
	name = "name"
	c    = iota
	d    = iota
)

//var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	//fmt.Println(*p)
}

func f46(n int) (r int) {
	defer func() {
		fmt.Println(r)
		r += n
		recover()
	}()

	var f func()

	defer f()
	f = func() {
		fmt.Println(r)
		r += 2
	}
	fmt.Println(n)
	return n + 1
}

type data struct {
	sync.Locker
}

func (d *data) check() {
	d.Lock()
	defer d.Unlock()
	fmt.Println(1)
	time.Sleep(1 * time.Second)
}

func alwaysFalse() bool {
	return false
}

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c.Daemon)
}

type User struct {
	Name string
}

func (u *User) SetName(name string) {
	u.Name = name
	fmt.Println(u.Name)
}

type Employee User

type Persons struct {
	Age      int
	Name     string
	Birthday time.Time
}

type PersonSwapper struct {
	p  []Persons
	by func(c, q *Persons) bool
}

func (pw PersonSwapper) Swap(i, j int) {
	pw.p[i], pw.p[j] = pw.p[j], pw.p[i]
}

func (pw PersonSwapper) Len() int {
	return len(pw.p)
}

func (pw PersonSwapper) Less(i, j int) bool {
	return pw.by(&pw.p[i], &pw.p[j])
}

func let(it *[]int) {
	for i, m := range *it {
		(*it)[i] = m * m
	}
}

type ast struct {
	Name string `json:"name"`
}

var pools = sync.Pool{
	New: func() any {
		return new(ast)
	},
}
var buf, _ = json.Marshal(ast{Name: "huyufan"})

type config struct {
	name string
	ip   int
}

var (
	once sync.Once
	conf *config
)

type ParseOp int

const (
	Second ParseOp = iota
	SecondOptional
	Minute
	Hour
)

func testOnce() {
	fmt.Println("once")
	conf = &config{
		name: "huyufan",
		ip:   33,
	}
}

const (
	starBit = 1 << 63
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Fatal("sdsd")
	fmt.Errorf("sdsdsdsds")
	fmt.Errorf("fffff")

	strs := "sdpohsdh说的话 asas"
	cc := strings.Fields(strs)
	for c, v := range cc {
		fmt.Println(c, v)
	}
	fmt.Println(cc)
	fmt.Printf("%v", cc)
	t := time.Now()
	fmt.Println(uint(t.Month()))
	// for 1<<(3&3) == 0 {
	// 	fmt.Println("huyf")
	// }
	fmt.Println(1 << 3)

	fmt.Printf("%b\n", (1 << 3))
	fmt.Println(1 << 3 & 7)
	fmt.Println(2 & 2)
	return
	// 加载北京时间（Asia/Shanghai）
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	currentTime := time.Now().In(location)
	fmt.Println("Current time:", currentTime)
	fmt.Println("Location:", currentTime.Location()) // 打印当前时间的时区
	sdt := " djhj dsdj sdj \n"
	fmt.Println(strings.TrimSpace(sdt))
	sdsd := []int{1, 2, 3, 4, 55}
	fmt.Println(len(sdsd))
	i := float64(5)
	c := float64(3)
	dd := i / c
	qqc := math.Floor(dd*100) / 100
	fmt.Println(qqc)
	fmt.Printf("%.2f\n", dd)
	fmt.Println(dd)
	return
	fmt.Println(Second, SecondOptional, Minute, Hour)
	run, _ := os.Hostname()
	fmt.Println(run)
	for i := 0; i < 10; i++ {
		once.Do(func() {
			testOnce()
		})
	}

	fmt.Println(conf)

	fmt.Println(buf)

	st := pools.Get().(*ast)
	pools.Put(st)
	_ = json.Unmarshal(buf, st)

	d := ast{}

	json.Unmarshal(buf, &d)

	fmt.Println(d)

	fmt.Println(st)

	hd := make([]int, 0, 8)
	hd = append(hd, []int{1, 2, 3}...)
	fmt.Println(hd)

	le := []int{1, 2, 3, 4, 5, 6, 7}

	let(&le)

	fmt.Println(le)

	file, err := os.OpenFile("c.csv", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.Write([]string{"1212", "12123,44", "怕送"})
	w.Flush()

	f, err := os.Open("c.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	r := csv.NewReader(f)
	str, _ := r.Read()
	for i := range str {
		fmt.Println(str[i])
	}

}

type T struct {
	n int
}

func hTest() {
	ch := make(chan int)
	cd := make(chan int)

	go func() {
		for {
			select {
			case c := <-ch:
				fmt.Println(c)
			case c := <-cd:
				fmt.Println(c)
			}
		}
		print("end")
	}()
	go func() {
		for i := 0; i < 100; i++ {
			cd <- i
			cd <- i * i
		}
	}()

	time.Sleep(2 * time.Second)
}

func fTest() {
	c := make(chan interface{})
	d := make(chan interface{})
	ot := Or(c, d)
	go func() {
		time.Sleep(1 * time.Second)
		close(d)
	}()

	<-ot
	fmt.Println("Finished main")
}

func gtest() {
	ch := make(chan struct{})
	go func() {
		for {
			select {
			case s := <-ch:
				fmt.Print(s)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- struct{}{}
	}
}

func cTest() {
	taskQueue := make(chan chan task)

	for i := 0; i < 3; i++ {
		go func(w int) {
			for {
				tasChan := <-taskQueue
				fmt.Println(1)
				tk := <-tasChan
				fmt.Sprint(tk)
				//fmt.Println("c", tk.id)
				//fmt.Println(w, tk.id, tk.str)
			}
		}(i)
	}

	for c := 0; c < 10; c++ {
		time.Sleep(500 * time.Microsecond)
		taskCh := make(chan task)

		taskQueue <- taskCh

		taskCh <- task{id: c, str: "huyufan"}

		//fmt.Println(c)
	}

	//time.Sleep(10 * time.Second)
}
