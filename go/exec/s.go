package main

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"go/format"
	"io"
	"log"
	"net"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"text/template"
	"time"
)

const tpl = `
package {{.pkg}}

var messages = map[{{.type}}]string{
	{{range $key,$value := .comments}}
	{{$key}}:"{{$value}}",{{end}}

func GetErrMsg(code {{.type}}) string {
	if msg,ok :=message[code];ok{
		return msg
	}
		return ""	
}
`

func gen(cType interface{}, comments map[interface{}]string) ([]byte, error) {
	var buf = bytes.NewBufferString("")
	data := map[string]interface{}{
		"pkg":      "example",
		"type":     cType,
		"comments": comments,
	}
	fmt.Println(data)
	t, err := template.New("").Parse(tpl)
	fmt.Println(t)
	if err != nil {
		return nil, errors.New("snihaoqwqw")
	}
	t.Execute(buf, data)
	return format.Source(buf.Bytes())
}

func OrBySelect(channls ...<-chan interface{}) <-chan interface{} {
	switch len(channls) {
	case 0:
		return nil
	case 1:
		return channls[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		var cases []reflect.SelectCase
		for _, c := range channls {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		reflect.Select(cases)
	}()
	return orDone
}

func doSomething() {
	defer CountTime("dosomething")()
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}

func CountTime(msg string) func() {
	start := time.Now()
	fmt.Printf("run func: %s", msg)
	return func() {
		fmt.Printf("func name: %s run time: %f s \n", msg, time.Since(start).Seconds())
	}
}

func cal() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func() {
			for {
				if isok(ctx) {
					fmt.Println(1)
					break
				}
			}
		}()
	}
	cancel()

	time.Sleep(3 * time.Second)
}

func isok(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

//go:embed c.txt
var bs []byte

func main() {
	str := "排山倒海 qwqw"
	fmt.Println(str[0])

	//遍历每一个字符
	for i, j := range str {
		// 注意： i 不是递增的，ci 是根据 他的字节进行增加的
		log.Printf("第%d个字符：%s", i, string(j))
	}

	cc := strings.Fields(str)
	fmt.Printf("%v", cc)
	sdd := []struct {
		Name string
		Age  int
	}{{
		Name: "huyufan",
		Age:  12,
	},
		{
			Name: "ddd",
			Age:  3,
		},
		{
			Name: "dddd",
			Age:  8,
		},
	}
	sort.Slice(sdd, func(i, j int) bool {
		return sdd[i].Age > sdd[j].Age
	})

	fmt.Println(sdd)

	sysType := runtime.GOOS
	fmt.Println(sysType)

	cm, _ := exec.Command("cmd", "/c", "echo", "hello world").Output()
	fmt.Println(string(cm))
	r := regexp.MustCompile(`[A-Z]+`)

	strs := "sdhh收到sdsd"

	c := r.FindAllString(strs, -1)
	if len(c) > 0 {
		for _, v := range c {
			fmt.Println(v)
		}
	}
	d := strings.Join(c, ",")
	fmt.Println(d)
	// ssd := []struct {
	// 	Name  string `json:"name"`
	// 	Value int    `json:"value"`
	// }{
	// 	{Name: "huyufan", Value: 12},
	// 	{Name: "cccc", Value: 14},
	// }

	// c, _ := json.MarshalIndent(ssd, "", "  ")

	// fmt.Println(string(c))
	// http.HandleFunc()
	// http.Handler
	// http.HandlerFunc()
	// data := bytes.NewReader(json)
	// http.Post("http://www.baidu.com", "application/json", data)
	// ctx, cancel := context.WithCancel(context.Background())
	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		for {
	// 			if isok(ctx) {
	// 				fmt.Println(1)
	// 				break
	// 			}
	// 		}
	// 	}()
	// }
	// cancel()

	// time.Sleep(3 * time.Second)
	// listen, err := net.Listen("tcp", ":8082")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("tcp server is  running...")
	// for {
	// 	conn, err := listen.Accept()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	go process(conn)

	// }

	// ch1 := make(chan interface{})
	// ch2 := make(chan interface{})

	// go func() {
	// 	time.AfterFunc(2*time.Second, func() {
	// 		ch1 <- 1
	// 	})
	// }()

	// go func() {
	// 	time.AfterFunc(3*time.Second, func() {
	// 		ch2 <- 2
	// 	})
	// }()
	// fmt.Println(time.Now().Format(time.DateTime))
	// <-Ors(ch1, ch2)
	// fmt.Println(time.Now().Format(time.DateTime))
	// time.Sleep(5 * time.Second)

}
func Order(chs ...<-chan interface{}) <-chan interface{} {
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

func Ors(chs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var once sync.Once
		for _, c := range chs {
			go func(c <-chan interface{}) {
				select {
				case <-c:
					once.Do(func() {
						close(out) // 关闭out,提醒外部可以继续执行了
					})
					// case <-out:
				}
				fmt.Println(1)
			}(c)
		}
	}()
	return out
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var data [10]byte
		n, err := conn.Read(data[:])
		if err != nil && err != io.EOF {
			fmt.Printf("failed to read msg from client, err: %v \n", err)
			break
		}
		str := string(data[:n])
		if str == "exit" {
			fmt.Println("client exit...")
			break
		}
		fmt.Printf("read msg:%s\n", str)

		conn.Write([]byte(fmt.Sprintf("%s ok", str)))
	}
}
