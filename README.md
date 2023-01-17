# GOLang 基础笔记

## 基础语法

运行

```powershell
go run xxx
```

生成二进制文件

```
go build xxx
```



#### hello.go

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

```

#### main.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {

	var a = "initial"

	var b, c int = 1, 2

	var d = true

	var e float64

	f := float32(e)

	g := a + "foo"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple

	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}

```

```
initial 1 2 true 0 0
initialfoo
constant 500000000 6e+11 -0.28470407323754404 0.7591864109375384
```

#### for.go

```go
package main

import "fmt"

func main() {

	i := 1
	for {
		fmt.Println("loop")
		break
	}
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}

```

```
loop
7
8
1
3
1
2
3
```

#### if.go

```go
package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

```

```
7 is odd
8 is divisible by 4
9 has 1 digit
```

#### switch.go

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}

```

```
two
It's before noon
```

#### array.go

```go
package main

import "fmt"

func main() {

	var a [5]int
	a[4] = 100
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

```

```
get: 0
len: 5
[1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]
```

#### slice.go

```go
package main

import "fmt"

func main() {

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) // [a b c d e f]

	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]
}

```

```
get: c
len: 3
[a b c d e f]
[a b c d e f]
[c d e]
[a b c d e]
[c d e f]
[g o o d]
```

#### map.go

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)           // map[one:1 two:2]
	fmt.Println(len(m))      // 2
	fmt.Println(m["one"])    // 1
	fmt.Println(m["unknow"]) // 0

	r, ok := m["unknow"]
	fmt.Println(r, ok) // 0 false

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)
}

```

```go
map[one:1 two:2]
2
1
0
0 false
map[one:1 two:2] map[one:1 two:2]
```

#### range.go

```go
package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 9

	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v) // b 8; a A
	}
	for k := range m {
		fmt.Println("key", k) // key a; key b
	}
}

```

```
index: 0 num: 2
9
a A
b B
key a
key b
```

func.go

```go
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True
}

```

```
3
A true
```

#### point.go

```go
package main

import "fmt"

func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

func main() {
	n := 5
	add2(n)
	fmt.Println(n) // 5
	add2ptr(&n)
	fmt.Println(n) // 7
}

```

```
5
7
```

#### struct.go

```go
package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	a := user{name: "wang", password: "1024"}
	b := user{"wang", "1024"}
	c := user{name: "wang"}
	c.password = "1024"
	var d user
	d.name = "wang"
	d.password = "1024"

	fmt.Println(a, b, c, d)                 // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
	fmt.Println(checkPassword(a, "haha"))   // false
	fmt.Println(checkPassword2(&a, "haha")) // false
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}

```

```
{wang 1024} {wang 1024} {wang 1024} {wang 1024}
false
false
```

#### struct-method.go

```go
package main

import "fmt"

type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user{name: "wang", password: "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048")) // true
}

```

```
true
```

#### error.go

```go
package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	u, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name) // wang

	if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err) // not found
		return
	} else {
		fmt.Println(u.name)
	}
}

```

```
wang
not found
```

#### string.go

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))                // true
	fmt.Println(strings.Count(a, "l"))                    // 2
	fmt.Println(strings.HasPrefix(a, "he"))               // true
	fmt.Println(strings.HasSuffix(a, "llo"))              // true
	fmt.Println(strings.Index(a, "ll"))                   // 2
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) // he-llo
	fmt.Println(strings.Repeat(a, 2))                     // hellohello
	fmt.Println(strings.Replace(a, "e", "E", -1))         // hEllo
	fmt.Println(strings.Split("a-b-c", "-"))              // [a b c]
	fmt.Println(strings.ToLower(a))                       // hello
	fmt.Println(strings.ToUpper(a))                       // HELLO
	fmt.Println(len(a))                                   // 5
	b := "你好"
	fmt.Println(len(b)) // 6
}

```

#### fmt.go

```go
package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n) // hello 123
	fmt.Println(p)    // {1 2}

	fmt.Printf("s=%v\n", s)  // s=hello
	fmt.Printf("n=%v\n", n)  // n=123
	fmt.Printf("p=%v\n", p)  // p={1 2}
	fmt.Printf("p=%+v\n", p) // p={x:1 y:2}
	fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2}

	f := 3.141592653
	fmt.Println(f)          // 3.141592653
	fmt.Printf("%.2f\n", f) // 3.14
}

```

#### json.go

```go
package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func main() {
	a := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)         // [123 34 78 97...]
	fmt.Println(string(buf)) // {"Name":"wang","age":18,"Hobby":["Golang","TypeScript"]}

	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b) // main.userInfo{Name:"wang", Age:18, Hobby:[]string{"Golang", "TypeScript"}}
}

```

```
[123 34 78 97 109 101 34 58 34 119 97 110 103 34 44 34 97 103 101 34 58 49 56 44 34 72 111 98 98 121 34 58 91 34 71 111 108 97 110 103 34 44 34 84 121 112 101 83 99 114 105 112 116 34 93 125]
{"Name":"wang","age":18,"Hobby":["Golang","TypeScript"]}
{
        "Name": "wang",
        "age": 18,
        "Hobby": [
                "Golang",
                "TypeScript"
        ]
}
main.userInfo{Name:"wang", Age:18, Hobby:[]string{"Golang", "TypeScript"}}
```

#### time.go

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // 2022-03-27 18:04:59.433297 +0800 CST m=+0.000087933
	t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)
	fmt.Println(t)                                                  // 2022-03-27 01:25:36 +0000 UTC
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()) // 2022 March 27 1 25
	fmt.Println(t.Format("2006-01-02 15:04:05"))                    // 2022-03-27 01:25:36
	diff := t2.Sub(t)
	fmt.Println(diff)                           // 1h5m0s
	fmt.Println(diff.Minutes(), diff.Seconds()) // 65 3900
	t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t)    // true
	fmt.Println(now.Unix()) // 1648738080
}

```

```
2023-01-13 08:28:58.822490392 +0000 UTC m=+0.000034401
2022-03-27 01:25:36 +0000 UTC
2022 March 27 1 25
2022-03-27 01:25:36
1h5m0s
65 3900
true
1673598538
```

#### strconv.go

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // 1.234

	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n) // 111

	n, _ = strconv.ParseInt("0x1000", 0, 64)
	fmt.Println(n) // 4096

	n2, _ := strconv.Atoi("123")
	fmt.Println(n2) // 123

	n2, err := strconv.Atoi("AAA")
	fmt.Println(n2, err) // 0 strconv.Atoi: parsing "AAA": invalid syntax
}

```

#### env.go

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// go run example/20-env/main.go a b c d
	fmt.Println(os.Args)           // [/var/folders/8p/n34xxfnx38dg8bv_x8l62t_m0000gn/T/go-build3406981276/b001/exe/main a b c d]
	fmt.Println(os.Getenv("PATH")) // /usr/local/go/bin...
	fmt.Println(os.Setenv("AA", "BB"))

	buf, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf)) // 127.0.0.1       localhost
}

```

#### guess.go

```go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		input = strings.Trim(input, "\r\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}

```

#### 翻译

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query(word string) {
	client := &http.Client{}
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}

```

####   简易的服务器

```go
package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const socks5Ver = 0x05
const cmdBind = 0x01
const atypIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	err := auth(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
	err = connect(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
}

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}

	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}
	return nil
}

func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 版本号，socks5的值为0x05
	// CMD 0x01表示CONNECT请求
	// RSV 保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名
	// DST.ADDR 一个可变长度的值
	// DST.PORT 目标端口，固定2个字节

	buf := make([]byte, 4)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("read header failed:%w", err)
	}
	ver, cmd, atyp := buf[0], buf[1], buf[3]
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not supported cmd:%v", ver)
	}
	addr := ""
	switch atyp {
	case atypIPV4:
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read atyp failed:%w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case atypeHOST:
		hostSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read hostSize failed:%w", err)
		}
		host := make([]byte, hostSize)
		_, err = io.ReadFull(reader, host)
		if err != nil {
			return fmt.Errorf("read host failed:%w", err)
		}
		addr = string(host)
	case atypeIPV6:
		return errors.New("IPv6: no supported yet")
	default:
		return errors.New("invalid atyp")
	}
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed:%w", err)
	}
	port := binary.BigEndian.Uint16(buf[:2])

	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		return fmt.Errorf("dial dst failed:%w", err)
	}
	defer dest.Close()
	log.Println("dial", addr, port)

	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER socks版本，这里为0x05
	// REP Relay field,内容取值如下 X’00’ succeeded
	// RSV 保留字段
	// ATYPE 地址类型
	// BND.ADDR 服务绑定的地址
	// BND.PORT 服务绑定的端口DST.PORT
	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		_, _ = io.Copy(dest, reader)
		cancel()
	}()
	go func() {
		_, _ = io.Copy(conn, dest)
		cancel()
	}()

	<-ctx.Done()
	return nil
}

```

#### 课后作业1   
1.修改第一个例子猜谜游戏里面的最终代码，使用fmt.Scanf来简化代码实现  

```
package main

import (
	// "bufio"
	"fmt"
	"math/rand"
	// "os"
	// "strconv"
	// "strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	// reader := bufio.NewReader(os.Stdin)
	for {
		// input, err := reader.ReadString('\n')
		// if err != nil {
		// 	fmt.Println("An error occured while reading input. Please try again", err)
		// 	continue
		// }
		// input = strings.Trim(input, "\r\n")
		var guess int
		fmt.Scanf("%d",&guess)


		// guess, err := strconv.Atoi(input)
		// if err != nil {
		// 	fmt.Println("Invalid input. Please enter an integer value")
		// 	continue
		// }
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}

```

2修改第二个例子命令行词典里面的最终代码，增加另一种翻译引擎的支持   
使用腾讯的 [](https://fanyi.qq.com/)
```
package main

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "net/http"
   "strings"
)

type AutoGenerated struct {
   SessionUUID string `json:"sessionUuid"`
   Translate struct {
      ErrCode int `json:"errCode"`
      ErrMsg string `json:"errMsg"`
      SessionUUID string `json:"sessionUuid"`
      Source string `json:"source"`
      Target string `json:"target"`
      Records []struct {
         SourceText string `json:"sourceText"`
         TargetText string `json:"targetText"`
         TraceID string `json:"traceId"`
      } `json:"records"`
      Full bool `json:"full"`
      Options struct {
      } `json:"options"`
   } `json:"translate"`
   Dict struct {
      Word string `json:"word"`
      DetailURL string `json:"detailUrl"`
      Abstract []struct {
         Ps string `json:"ps"`
         Explanation []string `json:"explanation"`
      } `json:"abstract"`
      Data []struct {
         EnHash string `json:"en_hash"`
      } `json:"data"`
      Type string `json:"type"`
      ErrCode int `json:"errCode"`
      ErrMsg string `json:"errMsg"`
   } `json:"dict"`
   Suggest struct {
      Data []struct {
         ID string `json:"id"`
         Word string `json:"word"`
         Abstract []struct {
            Exp []string `json:"exp"`
            Ps string `json:"ps"`
         } `json:"abstract"`
         SuggestTranslation string `json:"suggest_translation"`
         PhJSON string `json:"ph_json"`
         TransformPl string `json:"transform_pl"`
         TransformIs interface{} `json:"transform_is"`
         TransformWere interface{} `json:"transform_were"`
         TransformBeen interface{} `json:"transform_been"`
         TransformBeing interface{} `json:"transform_being"`
         TransformEr interface{} `json:"transform_er"`
         TransformEst interface{} `json:"transform_est"`
         Groups string `json:"groups"`
         AmEForm interface{} `json:"AmE_form"`
         BrEForm interface{} `json:"BrE_form"`
         ExamplesJSON string `json:"examples_json"`
         WordOfPhrase interface{} `json:"word_of_phrase"`
         Frq int `json:"frq"`
         Valid int `json:"valid"`
         TransformSing interface{} `json:"transform_sing"`
         PrxPhAmE string `json:"prx_ph_AmE"`
         PrxPhBrE string `json:"prx_ph_BrE"`
      } `json:"data"`
      ErrCode int `json:"errCode"`
      ErrMsg string `json:"errMsg"`
   } `json:"suggest"`
   ErrCode int `json:"errCode"`
   ErrMsg string `json:"errMsg"`
}

func translate(word string) {
   url := "https://fanyi.qq.com/api/translate"
   method := "POST"
   payload := strings.NewReader(`source=auto&target=zh&sourceText=` + word +`&qtv=5a3d5ab458f88908
&qtk=9PfIShCScN2DELc%2FuDCt5%2BeP86QHhPu6tSZB6ITACggmPKdhkSh3bwyi6LJqKiEfUhOTcngbToukd5UTNQQ66RbIRRmWQ%2FUaIE1ahXL0vjPQLok3HcfmdvTEZbS%2Fi3f512lBudDDZVAvswD7QZ2B3Q%3D%3D
&ticket=&randstr=&sessionUuid=translate_uuid1652253538206`)
   client := &http.Client {}
   req, err := http.NewRequest(method, url, payload)
   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
   req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
   req.Header.Add("Connection", "keep-alive")
   req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
   req.Header.Add("Cookie", "pgv_pvid=5948909680; pac_uid=0_935f2dad4e09f; tvfe_boss_uuid=8b4f775e020c647a; eas_sid=21M6a4h7p9P4c8m8J668T1x2z2; RK=oksQmnU1XX; ptcz=3f7da84564a30539d2db26c20ab25bb3ba481b2a2eeb87318fd1dc0755ef6aa3; ptui_loginuin=1260828355@qq.com; fy_guid=366184b4-5d6d-4da7-97d8-c32c19b64fa8; ADHOC_MEMBERSHIP_CLIENT_ID1.0=6ac9805b-e21f-b7e7-8b93-cbda95f4b43a; openCount=1; qtv=5a3d5ab458f88908; qtk=9PfIShCScN2DELc/uDCt5+eP86QHhPu6tSZB6ITACggmPKdhkSh3bwyi6LJqKiEfUhOTcngbToukd5UTNQQ66RbIRRmWQ/UaIE1ahXL0vjPQLok3HcfmdvTEZbS/i3f512lBudDDZVAvswD7QZ2B3Q==; gr_user_id=5b66730c-2ffc-489c-b366-3c6c404734ba; 8507d3409e6fad23_gr_session_id=e72f0b98-87b0-44b7-a42a-a2d69b60f6e7; 8c66aca9f0d1ff2e_gr_session_id=6132b52a-e285-41cc-84bb-6bc05d3deb73; 8c66aca9f0d1ff2e_gr_session_id_6132b52a-e285-41cc-84bb-6bc05d3deb73=true")
   req.Header.Add("Origin", "https://fanyi.qq.com")
   req.Header.Add("Referer", "https://fanyi.qq.com/")
   req.Header.Add("Sec-Fetch-Dest", "empty")
   req.Header.Add("Sec-Fetch-Mode", "cors")
   req.Header.Add("Sec-Fetch-Site", "same-origin")
   req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36")
   req.Header.Add("X-Requested-With", "XMLHttpRequest")
//    req.Header.Add("sec-ch-ua", "" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"")
   req.Header.Add("sec-ch-ua-mobile", "?0")
   req.Header.Add("sec-ch-ua-platform", "Windows")
   req.Header.Add("uc", "GLYdIz3WjseABEHSwmZpKBGZ9Afq6AxIJcPpg7ZK1j0=")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   //fmt.Println(string(body))
   var dictRes AutoGenerated
   err = json.Unmarshal(body,&dictRes)
   if err != nil {
      fmt.Println(err)
      return
   }
   for _, item := range dictRes.Translate.Records {
      fmt.Println(item.TargetText)
   }
   //fmt.Println(dictRes.Translate.Records)
}

func main(){
   fmt.Println("Welcome to use translate App!")
   var flag = "yes"
   for{
      fmt.Print("Please enter what you want to translate: ")
      var word string
      fmt.Scanln(&word)
      translate(word)

      fmt.Print("want to continue?(if wouldn't,enter q to exit): ")
      fmt.Scanln(&flag)
      if flag == "q"{
         fmt.Print("It's been an honor serving you.")
         break
      }
   }
}
```



3.在上一步骤的基础上，修改代码实现并行请求两个翻译引擎来提高响应速度  

```
package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io/ioutil"
   "net/http"
   "strings"
   "sync"
)

type DictReq struct {
   TransType string `json:"trans_type"`
   Source string `json:"source"`
}

type AutoGenerated struct {
   Rc int `json:"rc"`
   Wiki struct {
      KnownInLaguages int `json:"known_in_laguages"`
      Description struct {
         Source string `json:"source"`
         Target interface{} `json:"target"`
      } `json:"description"`
      ID string `json:"id"`
      Item struct {
         Source string `json:"source"`
         Target string `json:"target"`
      } `json:"item"`
      ImageURL string `json:"image_url"`
      IsSubject string `json:"is_subject"`
      Sitelink string `json:"sitelink"`
   } `json:"wiki"`
   Dictionary struct {
      Prons struct {
         EnUs string `json:"en-us"`
         En string `json:"en"`
      } `json:"prons"`
      Explanations []string `json:"explanations"`
      Synonym []string `json:"synonym"`
      Antonym []string `json:"antonym"`
      WqxExample [][]string `json:"wqx_example"`
      Entry string `json:"entry"`
      Type string `json:"type"`
      Related []interface{} `json:"related"`
      Source string `json:"source"`
   } `json:"dictionary"`
}

func translate(word string,wg *sync.WaitGroup) {
   defer wg.Done()
   url := "https://api.interpreter.caiyunai.com/v1/dict"
   method := "POST"
   request := DictReq{TransType: "en2zh",Source: word}
   buf, err := json.Marshal(request)
   if err != nil{
      fmt.Println(err)
   }
   data := bytes.NewReader(buf)
   client := &http.Client {}
   req, err := http.NewRequest(method, url, data)
   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Accept", "application/json, text/plain, */*")
   req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
   req.Header.Add("Connection", "keep-alive")
   req.Header.Add("Content-Type", "application/json;charset=UTF-8")
   req.Header.Add("Origin", "https://fanyi.caiyunapp.com")
   req.Header.Add("Referer", "https://fanyi.caiyunapp.com/")
   req.Header.Add("Sec-Fetch-Dest", "empty")
   req.Header.Add("Sec-Fetch-Mode", "cors")
   req.Header.Add("Sec-Fetch-Site", "cross-site")
   req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36")
   req.Header.Add("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
   req.Header.Add("app-name", "xy")
   req.Header.Add("device-id", "")
   req.Header.Add("os-type", "web")
   req.Header.Add("os-version", "")
   req.Header.Add("sec-ch-ua", "" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"")
   req.Header.Add("sec-ch-ua-mobile", "?0")
   req.Header.Add("sec-ch-ua-platform", ""Windows"")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   if res.StatusCode != 200{
      fmt.Println("bad StatusCode:", res.StatusCode, "body", string(body))
      return
   }
   var dictRes AutoGenerated
   err = json.Unmarshal(body,&dictRes)
   if err != nil{
      fmt.Println(err)
      return
   }
   fmt.Println("彩云小译")
   fmt.Println(word, "UK:", dictRes.Dictionary.Prons.En, "US:", dictRes.Dictionary.Prons.EnUs)
   for _, item := range dictRes.Dictionary.Explanations {
      fmt.Println(item)
   }
}


type AutoGenerated2 struct {
   SessionUUID string `json:"sessionUuid"`
   Translate struct {
      ErrCode int `json:"errCode"`
      ErrMsg string `json:"errMsg"`
      SessionUUID string `json:"sessionUuid"`
      Source string `json:"source"`
      Target string `json:"target"`
      Records []struct {
         SourceText string `json:"sourceText"`
         TargetText string `json:"targetText"`
         TraceID string `json:"traceId"`
      } `json:"records"`
      Full bool `json:"full"`
      Options struct {
      } `json:"options"`
   } `json:"translate"`
   Dict struct {
      Word string `json:"word"`
      DetailURL string `json:"detailUrl"`
      Abstract []struct {
         Ps string `json:"ps"`
         Explanation []string `json:"explanation"`
      } `json:"abstract"`
      Data []struct {
         EnHash string `json:"en_hash"`
      } `json:"data"`
      Type string `json:"type"`
      ErrCode int `json:"errCode"`
      ErrMsg string `json:"errMsg"`
   } `json:"dict"`
   Suggest struct {
      Data []struct {
         ID string `json:"id"`
         Word string `json:"word"`
         Abstract []struct {
            Exp []string `json:"exp"`
            Ps string `json:"ps"`
         } `json:"abstract"`
         SuggestTranslation string `json:"suggest_translation"`
         PhJSON string `json:"ph_json"`
         TransformPl string `json:"transform_pl"`
         TransformIs interface{} `json:"transform_is"`
         TransformWere interface{} `json:"transform_were"`
         TransformBeen interface{} `json:"transform_been"`
         TransformBeing interface{} `json:"transform_being"`
         TransformEr interface{} `json:"transform_er"`
         TransformEst interface{} `json:"transform_est"`
         Groups string `json:"groups"`
         AmEForm interface{} `json:"AmE_form"`
         BrEForm interface{} `json:"BrE_form"`
         ExamplesJSON string `json:"examples_json"`
         WordOfPhrase interface{} `json:"word_of_phrase"`
         Frq int `json:"frq"`
         Valid int `json:"valid"`
         TransformSing interface{} `json:"transform_sing"`
         PrxPhAmE string `json:"prx_ph_AmE"`
         PrxPhBrE string `json:"prx_ph_BrE"`
      } `json:"data"`
      ErrCode int `json:"errCode"`
      ErrMsg string `json:"errMsg"`
   } `json:"suggest"`
   ErrCode int `json:"errCode"`
   ErrMsg string `json:"errMsg"`
}

func translate2(word string,wg *sync.WaitGroup) {
   defer wg.Done()
   url := "https://fanyi.qq.com/api/translate"
   method := "POST"
   payload := strings.NewReader(`source=auto&target=zh&sourceText=` + word +`&qtv=5a3d5ab458f88908
&qtk=9PfIShCScN2DELc%2FuDCt5%2BeP86QHhPu6tSZB6ITACggmPKdhkSh3bwyi6LJqKiEfUhOTcngbToukd5UTNQQ66RbIRRmWQ%2FUaIE1ahXL0vjPQLok3HcfmdvTEZbS%2Fi3f512lBudDDZVAvswD7QZ2B3Q%3D%3D
&ticket=&randstr=&sessionUuid=translate_uuid1652253538206`)
   client := &http.Client {}
   req, err := http.NewRequest(method, url, payload)
   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
   req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
   req.Header.Add("Connection", "keep-alive")
   req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
   req.Header.Add("Cookie", "pgv_pvid=5948909680; pac_uid=0_935f2dad4e09f; tvfe_boss_uuid=8b4f775e020c647a; eas_sid=21M6a4h7p9P4c8m8J668T1x2z2; RK=oksQmnU1XX; ptcz=3f7da84564a30539d2db26c20ab25bb3ba481b2a2eeb87318fd1dc0755ef6aa3; ptui_loginuin=1260828355@qq.com; fy_guid=366184b4-5d6d-4da7-97d8-c32c19b64fa8; ADHOC_MEMBERSHIP_CLIENT_ID1.0=6ac9805b-e21f-b7e7-8b93-cbda95f4b43a; openCount=1; qtv=5a3d5ab458f88908; qtk=9PfIShCScN2DELc/uDCt5+eP86QHhPu6tSZB6ITACggmPKdhkSh3bwyi6LJqKiEfUhOTcngbToukd5UTNQQ66RbIRRmWQ/UaIE1ahXL0vjPQLok3HcfmdvTEZbS/i3f512lBudDDZVAvswD7QZ2B3Q==; gr_user_id=5b66730c-2ffc-489c-b366-3c6c404734ba; 8507d3409e6fad23_gr_session_id=e72f0b98-87b0-44b7-a42a-a2d69b60f6e7; 8c66aca9f0d1ff2e_gr_session_id=6132b52a-e285-41cc-84bb-6bc05d3deb73; 8c66aca9f0d1ff2e_gr_session_id_6132b52a-e285-41cc-84bb-6bc05d3deb73=true")
   req.Header.Add("Origin", "https://fanyi.qq.com")
   req.Header.Add("Referer", "https://fanyi.qq.com/")
   req.Header.Add("Sec-Fetch-Dest", "empty")
   req.Header.Add("Sec-Fetch-Mode", "cors")
   req.Header.Add("Sec-Fetch-Site", "same-origin")
   req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36")
   req.Header.Add("X-Requested-With", "XMLHttpRequest")
   req.Header.Add("sec-ch-ua", "" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"")
   req.Header.Add("sec-ch-ua-mobile", "?0")
   req.Header.Add("sec-ch-ua-platform", ""Windows"")
   req.Header.Add("uc", "GLYdIz3WjseABEHSwmZpKBGZ9Afq6AxIJcPpg7ZK1j0=")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   //fmt.Println(string(body))
   var dictRes AutoGenerated2
   err = json.Unmarshal(body,&dictRes)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println("腾讯翻译君")
   for _, item := range dictRes.Translate.Records {
      fmt.Println(item.TargetText)
   }
}

func main(){
   var word string
   fmt.Scanln(&word)
   wg := sync.WaitGroup{}
   wg.Add(2)
   translate(word,&wg)
   translate2(word,&wg)
}
```

#### book的作业1
练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。  

```
package main

import (
    "fmt"
	"os"
	"strings"
)

func main() {
	// 法一
	fmt.Println(strings.Join(os.Args, " "));
	// 法二
	fmt.Println(os.Args[0])
	
}
```

练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。

```
package main

import (
    "fmt"
	"os"
)

func main() {
	for index, value := range os.Args {
		fmt.Println(index,value)
	}
	
}
```

练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）

#### GoRoutine

#### 位运算符





##### Printf

`Printf`有一大堆这种转换，Go程序员称之为*动词（verb）*。下面的表格虽然远不是完整的规范，但展示了可用的很多特性：

```
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
```

#### dup1.go

**map**存储了键/值（key/value）的集合，对集合元素，提供常数时间的存、取或测试操作。键可以是任意类型，只要其值能用`==`运算符比较，最常见的例子是字符串；值则可以是任意类型。这个例子中的键是字符串，值是整数。内置函数`make`创建空`map`

`map`中不含某个键时不用担心，首次读到新行时，等号右边的表达式`counts[line]`的值将被计算为其类型的零值，对于`int`即0。

每次`dup`读取一行输入，该行被当做键存入`map`，其对应的值递增。`counts[input.Text()]++`语句等价下面两句：

```go
line := input.Text()
counts[line] = counts[line] + 1
```

```go
// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    //定义映射
    counts := make(map[string]int)
    //定义
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```

#### dup2.go

```go
// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
}
```

`map`是一个由`make`函数创建的数据结构的引用。`map`作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），被调用函数对`map`底层数据结构的任何修改，调用者函数都可以通过持有的`map`引用看到。在我们的例子中，`countLines`函数向`counts`插入的值，也会被`main`函数看到。（译注：类似于C++里的引用传递，实际上指针是另一个指针了，但内部存的值指向同一块内存）

**练习 1.4：** 修改`dup2`，出现重复的行时打印文件名称。

test1.4.go

```
// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type nameAndCount struct {
	filename []string
	count    int
}

func main() {
	//定义结构体 其中键值对为 出现的文本-结构体
	counts := make(map[string]*nameAndCount)
	files := os.Args[1:]
	if len(files) == 0 {
		print("files的长度为0")
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	// 输出
	for k, v := range counts {
		if v.count > 1 {
			fmt.Printf("\n 出现的次数：%d 序号：%s 文件名称：%v\n", v.count, k, v.filename)
		}
	}
}

func countLines(f *os.File, counts map[string]*nameAndCount) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		print("当前扫描的值为：" + input.Text() + "\n")
		//如果存在值，追加
		if _, ok := counts[key]; ok {
			counts[key].count++
			counts[key].filename = append(counts[key].filename, f.Name())
		} else {
			//不存在值
			//手动添加值
			counts[key] = &nameAndCount{
				filename: make([]string, 1), //手动创建数值
				count:    1,
			}
			counts[key].filename[0] = f.Name()
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}


```

#### url1.go

```go
// Fetch prints the content found at a URL.
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        b, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }
}
```



```go
// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		newUrl := url
		prefix := strings.HasPrefix(url, "http://")
		//
		if !prefix {
			println("没有前缀,这里自动添加")
			newUrl = "http://" + url
		}

		resp, err := http.Get(newUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

```

test1.9.go

```go
// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	// "io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		newUrl := url
		prefix := strings.HasPrefix(url, "http://")
		//
		if !prefix {
			println("没有前缀,这里自动添加")
			newUrl = "http://" + url
		}

		resp, err := http.Get(newUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)
		println(resp.StatusCode)

	}
}

```

#### slice.go

```go
package main

import "fmt"

func main() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))
}

```

请注意，切片的长度不变，但容量不同。 我们来了解 `quarter2` 切片。 声明此切片时，你指出希望切片从位置编号 3 开始，最后一个元素位于位置编号 6。 切片长度为 3 个元素，但容量为 9，原因是基础数组有更多元素或位置可供使用，但对切片而言不可见。 例如，如果你尝试打印类似 `fmt.Println(quarter2[3])` 的内容，会出现以下错误：`panic: runtime error: index out of range [3] with length 3`。

**切片容量仅指出切片可扩展的程度。**

如下：可以拓展

```go
package main

import "fmt"

func main() {
    months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
    quarter2 := months[3:6]
    quarter2Extended := quarter2[:4]
    fmt.Println(quarter2, len(quarter2), cap(quarter2))
    fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))
}
```

请注意在声明 `quarter2Extended` 变量时，无需指定初始位置 (`[:4]`)。 执行此操作时，Go 会假定你想要切片的第一个位置。 你可对最后一个位置 (`[1:]`) 执行相同的操作。 Go 将假定你要引用所有元素，直到切片的最新位置 (`len()-1`)。

Go 提供了内置函数 `append(slice, element)`，便于你向切片添加元素。 将要修改的切片和要追加的元素作为值发送给该函数。 然后，`append` 函数会返回一个新的切片，将其存储在变量中。 对于要更改的切片，变量可能相同。

#### append.go

```go
package main

import "fmt"

func main() {
    var numbers []int
    for i := 0; i < 10; i++ {
        numbers = append(numbers, i)
        fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
    }
}
```

#### structjson.go

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    ID        int
    FirstName string `json:"name"`
    LastName  string
    Address   string `json:"address,omitempty"`
}

type Employee struct {
    Person
    ManagerID int
}

type Contractor struct {
    Person
    CompanyID int
}

func main() {
    employees := []Employee{
        Employee{
            Person: Person{
                LastName: "Doe", FirstName: "John",
            },
        },
        Employee{
            Person: Person{
                LastName: "Campbell", FirstName: "David",
            },
        },
    }

    data, _ := json.Marshal(employees)
    fmt.Printf("%s\n", data)

    var decoded []Employee
    json.Unmarshal(data, &decoded)
    fmt.Printf("%v", decoded)
}
```

```
[{"ID":0,"name":"John","LastName":"Doe","ManagerID":0},{"ID":0,"name":"David","LastName":"Campbell","ManagerID":0}]
[{{0 John Doe } 0} {{0 David Campbell } 0}]
```

fibonacci.go

```go
/*你将编写一个程序来计算某个数字的斐波纳契数列。
你将编写一个函数，它返回一个包含按斐波纳契数列排列的所有数字的切片，而这些数字是通过根据用户输入的大于 2 的数字计算得到的。
让我们假设小于 2 的数字会导致错误，并返回一个 nil 切片。

请记住，斐波那契数列是一个数字列表，其中每个数字是前两个斐波那契数字之和。
例如，数字 6 的序列是 1,1,2,3,5,8，
数字 7 的序列是 1,1,2,3,5,8,13，
数字 8 的序列是 1,1,2,3,5,8,13,21，
以此类推。
*/

package main

import "fmt"

func fibonacci(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}
	nums := make([]int, n)
	//初始化
	nums[0] = 1
	nums[1] = 1
	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums
}

func main() {
	//定义输入的数据
	var num int
	fmt.Print("What's the Fibonacci sequence you want? ")
	fmt.Scanln(&num)
	fmt.Println("The Fibonacci sequence is:", fibonacci(num))
}

```

golang中海油一个***byte\***数据类型与rune相似，它们都是用来表示字符类型的变量类型。它们的不同在于：

- byte 等同于int8，常用来处理ascii字符
- **rune** 等同于int32,常用来处理unicode或utf-8字符



romanToArabic.go

```go
package main

import "fmt"

/*
编写一个程序来转换罗马数字（例如将 MCLX 转换成 1,160）。
使用映射加载要用于将字符串字符转换为数字的基本罗马数字。
例如，M 将是映射中的键，其值将为 1000。 使用以下字符串字符映射表列表：

M => 1000
D => 500
C => 100
L => 50
X => 10
V => 5
I => 1
如果用户输入的字母与上述列表中的不同，则打印一个错误。
*/
//初始化
var romanMap = map[rune]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func romanToArabic(s string) int {
	arabicVals := make([]int, len(s)+1)
	for index, digit := range s {
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			fmt.Printf("Error: The roman s %s has a bad digit: %c\n", s, digit)
			return 0
		}
	}

	total := 0

	for index := 0; index < len(s); index++ {
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}

	return total
}
func main() {
	//var x string
	//fmt.Scanf("%s", &x)
	//fmt.Printf(string(romanToArabic(x)))
	fmt.Println("MCLX is", romanToArabic("MCLX"))
	fmt.Println("MCMXCIX is ", romanToArabic("MCMXCIX"))
	fmt.Println("MCMZ is", romanToArabic("MCMZ"))
}

```

#### http.go

```go
// Server1 is a minimal "echo" server.
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler) // each request calls handler
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
```

我们只用了八九行代码就实现了一个Web服务程序，这都是多亏了标准库里的方法已经帮我们完成了大量工作。main函数将所有发送到/路径下的请求和handler函数关联起来，/开头的请求其实就是所有发送到当前站点上的请求，服务监听8000端口。发送到这个服务的“请求”是一个http.Request类型的对象，这个对象中包含了请求中的一系列相关字段，其中就包括我们需要的URL。当请求到达服务器时，这个请求会被传给handler函数来处理，这个函数会将/hello这个路径从请求的URL中解析出来，然后把其发送到响应中，这里我们用的是标准输出流的fmt.Fprintf。



```go
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, myCycles float64) { //接收 cycles参数
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	if myCycles == 0 {
		myCycles = cycles //如果为零，则使用常亮定义的值
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < myCycles*2*math.Pi; t += res { //使用myCycles变量
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handler_gif(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if r.Form["cycles"] != nil { // 获取cycles参数，并转为int型
		cycles, err := strconv.Atoi(r.Form["cycles"][0])
		if err != nil {
			lissajous(w, float64(cycles))
		}
	}
	lissajous(w, float64(0))
}
func main() {
	http.HandleFunc("/", handler)        // each request calls handler
	http.HandleFunc("/gif", handler_gif) // each request calls handler
	log.Println("localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

```

通过get传参数

如下：

![image-20230115200727029](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/12640/image-20230115200727029.png)

![image-20230115200814185](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/12640/image-20230115200814185.png)

#### new函数

另一个创建变量的方法是调用内建的new函数。表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为`*T`。

```Go
p := new(int)   // p, *int 类型, 指向匿名的 int 变量
fmt.Println(*p) // "0"
*p = 2          // 设置 int 匿名变量的值为 2
fmt.Println(*p) // "2"
```

下面的两个newInt函数有着相同的行为：

```Go
func newInt() *int {
    return new(int)
}

func newInt() *int {
    var dummy int
    return &dummy
}
```

每次调用new函数都是返回一个新的变量的地址，因此下面两个地址是不同的：

```Go
p := new(int)
q := new(int)
fmt.Println(p == q) // "false"
```

#### **package包管理**

结构如下

![image-20230115205545643](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/12640/image-20230115205545643.png)

##### AbsoluteZeroC.go

```go
package main

import (
	"Go/src/tempconv"
	"fmt"
)

func main() {
	fmt.Println("AbsoluteZero-K:", tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Println("Freezing-K:", tempconv.CToK(tempconv.FreezingC))
	fmt.Println("Boilinig-K:", tempconv.CToK(tempconv.BoilingC))
}

```

##### tempconv.go

```go
// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// 摄氏度
type Celsius float64

// 华氏度
type Fahrenheit float64

// Kelvin
type Kelvin float64

const (
	//绝对0度
	AbsoluteZeroC Celsius = -273.15
	//冰点
	FreezingC Celsius = 0
	//沸点
	BoilingC Celsius = 100
)

// 输出 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

//处理Kelvin绝对温度的转换 Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的

```

##### cony.go

```go
package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC) }

func KToC(k Kelvin) Celsius { return Celsius(k) - AbsoluteZeroC }

```

## web

### gin框架

**github仓库:**

[gin-gonic/gin: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. (github.com)](https://github.com/gin-gonic/gin#quick-start)

**使用go mod管理依赖**

**go mod init**

初始化gomod文件

**go get 参数**

参数介绍：

- -d 只下载不安装
- -f 只有在你包含了 -u 参数的时候才有效，不让 -u 去验证 import 中的每一个都已经获取了，这对于本地 fork 的包特别有用
- -fix 在获取源码之后先运行 fix，然后再去做其他的事情
- -t 同时也下载需要为运行测试所需要的包
- -u 强制使用网络去更新包和它的依赖包
- -v 显示执行的命令

示例

```shell
go get -u github.com/gin-gonic/gin
```

# 高质量编程简介及编码规范

高质量:

- 各种边界条件考虑完备
- 异常情况处理，稳定性
- 易读易维护

编程原则

- 简单性
- 可读性
- 生产力

## 编码规范

### 公共符号始终要注释

*例外：实现接口的方法不需要注释*

### 格式化

使用**gofmt**（官方工具）自动格式化

### 注释

- 代码作用（适合公共符号）
- 代码如何实现 （适合注释实现过程）
- 代码实现的原因（适合解释代码的外部因素和提供额外的上下文）
- 代码什么情况下出错（适合代码的限制条件）

公共符号始终要注释·包中声明的每个公共的符号:
变量、常量、函数以及结构都需要添加注释
.任何既不明显也不简短的公
共功能必须予以注释
无论长度或复杂程度如何，
对库中的任何函数都必须进行注释

### **命名规范**

#### **变量**

**缩略词全大写，但当其位于变量开头且不需要导出时，使用全小写**

- 例如使用ServeHTTP而不是ServeHttp
- 使用XMLHTTPRequest或者xmlHTTPRequest

- 变量距离其被使用的地方越远，则需要携带越多的上下文信息
- 全局变量在其名字中需要更多的上下文信息，使得在不同地方可以轻易辨认出其含义

#### **函数**

函数名不携带包名的上下文信息，因为包名和函数名总是成对出现的·函数名尽量简短
当名为foo的包某个函数返回类型Foo时，可以省略类型信息而不导致歧义
当名为foo的包某个函数返回类型T时(T并不是Foo)，可以在函数名中加入类型信息

#### package

- 只由小写字母组成。不包含大写字母和下划线等字符·
- 简短并包含一定的上下文信息。例如schema、task 等·
- 不要与标准库同名。例如不要使用sync或者strings

以下规则尽量满足，以标准库包名为例

- 不使用常用变量名作为包名。例如使用bufio而不是buf·使用单数而不是复数。例如使用encoding而不是encodings
- 谨慎地使用缩写。例如使用fmt 在不破坏上下文的情况下比 format 更加简短

### 控制流程

避免嵌套

尽量保存为最小缩进

### 错误处理

#### 简单错误

- 简单的错误指的是仅出现一次的错误，且在其他地方不需要捕获该错误
- 优先使用errors.New来创建匿名变量来直接表示简单错误
- 如果有格式化的需求，使用fmt.Errorf

#### 错误的Wrap和 Unwrap

·错误的Wrap 实际上是提供了一个error嵌套另一个error的能力，从而生成一个error的跟踪链
·在 fmt.Errorf中使用:%w关键字来将一个错误关联至
错误链中

#### 错误判定

- 判定一个错误是否为特定错误，使用errors.Is
- 不同于使用==，使用该方法可以判定错误链上的所有错误是否含有特定的错误

#### panic

- 不建议在业务代码中使用panic
- ·调用函数不包含recover会造成程序崩溃·若问题可以被屏蔽或解决，建议使用error代替panic
- 当程序启动阶段发生不可逆转的错误时，可以在init 或 main函数中使用panic

# 性能优化

### **benchmark工具**

![image-20230117105733060](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/12640/image-20230117105733060.png)

### slice

**提前指定大小**

**在大切片上创建小切片，使用copy代替**

### string

**使用strings.builder** 和java类似

### 空结构体

使用空结构体struct{}实列不占用空间

### map

map 预分配内存分析

- 不断向map中添加元素的操作会触发map的扩容·
- 提前分配好空间可以减少内存拷贝和Rehash 的消耗·
- 建议根据实际需求提前预估好需要的空间

### 使用atomic 包

锁的实现是通过操作系统来实现，属于系统调用.atomic 操作是通过硬件实现，

效率比锁高sync.Mutex应该用来保护一段逻辑，不仅仅用于保护一个变量。

对于非数值操作，可以使用atomic.Value，能承载一个interface}

# 实战

直接拉取仓库

[wolfogre/go-pprof-practice: go pprof practice. (github.com)](https://github.com/wolfogre/go-pprof-practice)

分析的博客:

[golang pprof 实战 | Wolfogre's Blog](https://blog.wolfogre.com/posts/go-ppof-practice/)

**性能分析工具 pprof**

项目目录

![image-20230117112122043](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/12640/image-20230117112122043.png)

没有外部依赖,直接运行即可

保持程序运行，打开浏览器访问 `http://localhost:6060/debug/pprof/`，可以看到如下页面：

![image-20230117114525219](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/12640/image-20230117114525219.png)

页面上展示了可用的程序运行采样数据，分别有：

| 类型         | 描述                       | 备注                                                         |
| :----------- | :------------------------- | :----------------------------------------------------------- |
| allocs       | 内存分配情况的采样信息     | 可以用浏览器打开，但可读性不高                               |
| blocks       | 阻塞操作情况的采样信息     | 可以用浏览器打开，但可读性不高                               |
| cmdline      | 显示程序启动命令及参数     | 可以用浏览器打开，这里会显示 `./go-pprof-practice`           |
| goroutine    | 当前所有协程的堆栈信息     | 可以用浏览器打开，但可读性不高                               |
| heap         | 堆上内存使用情况的采样信息 | 可以用浏览器打开，但可读性不高                               |
| mutex        | 锁争用情况的采样信息       | 可以用浏览器打开，但可读性不高                               |
| profile      | CPU 占用情况的采样信息     | 浏览器打开会下载文件                                         |
| threadcreate | 系统线程创建情况的采样信息 | 可以用浏览器打开，但可读性不高                               |
| trace        | 程序运行跟踪信息           | 浏览器打开会下载文件，本文不涉及，可另行参阅[《深入浅出 Go trace》](https://blog.wolfogre.com/redirect/v3/AwBGKjtUXC4lQ2UdNqHTCoMSAwM8Cv46xcUtPG6RCPoPbv8CcXH9xQrF_wJJOfr_AlNN-lP_AjMyHP8IQUxTTlFBTjhBFgn-UTESAwM8Cv46xcVaFgY7bkEGFtw7If3FPAZNCsU7Bsw8PAXMPIIcSojF) |



**命令行**

```shell
go tool pprof http://localhost:6060/debug/pprof/profile
```

**topN 查看占用最多的函数**

```shell
(pprof) top
Showing nodes accounting for 8.91s, 98.67% of 9.03s total
Dropped 35 nodes (cum <= 0.05s)
      flat  flat%   sum%        cum   cum%
     8.91s 98.67% 98.67%      8.93s 98.89%  github.com/wolfogre/go-pprof-practice/animal/felidae/tiger.(*Tiger).Eat
         0     0% 98.67%      8.93s 98.89%  github.com/wolfogre/go-pprof-practice/animal/felidae/tiger.(*Tiger).Live
         0     0% 98.67%      8.97s 99.34%  main.main
         0     0% 98.67%      8.97s 99.34%  runtime.main
         0     0% 98.67%      0.05s  0.55%  runtime.systemstack

```

flat=Cum 函数中没有调用其他函数

flat=0 函数中只有其他函数的调用

**输入 `list Eat`，查看问题具体在代码的哪一个位置：根据指定的正则表达式查找**

```shell
(pprof) list Eat 
Total: 9.03s
ROUTINE ======================== github.com/wolfogre/go-pprof-practice/animal/canidae/dog.(*Dog).Eat in H:\go-pprof-practice\animal\canidae\dog\dog.go
         0       10ms (flat, cum)  0.11% of Total
         .          .     26:   d.Pee()
         .          .     27:   d.Run()
         .          .     28:   d.Howl()
         .          .     29:}
         .          .     30:func (d *Dog) Eat() {
         .       10ms     31:   log.Println(d.Name(), "eat")
         .          .     32:}
         .          .     33:
         .          .     34:func (d *Dog) Drink() {
         .          .     35:   log.Println(d.Name(), "drink")
         .          .     36:}
ROUTINE ======================== github.com/wolfogre/go-pprof-practice/animal/felidae/tiger.(*Tiger).Eat in H:\go-pprof-practice\animal\felidae\tiger\tiger.go
     8.91s      8.93s (flat, cum) 98.89% of Total
         .          .     26:无效的循环
         .          .     27:*/
         .          .     28:func (t *Tiger) Eat() {
         .          .     29:   log.Println(t.Name(), "eat")
         .          .     30:   loop := 10000000000
     8.91s      8.93s     31:   for i := 0; i < loop; i++ {
         .          .     32:           // do nothing
         .          .     33:   }
         .          .     34:}
         .          .     35:
         .          .     36:func (t *Tiger) Drink() {

```

**web 调用关系可视化**

可以访问 [graphviz 官网](https://blog.wolfogre.com/redirect/v3/A421Yoc_xEV4GG_UO8tV1nMSAwM8Cv46xcU7gjwSbQjbbjsviVpukMUYBkEJFgboxTESAwM8Cv46xcVaFgY7bkEGFtw7If3FPAZNCsU7Bsw8PAXMPIIcSojF)寻找适合自己操作系统的安装方法

**调查内存**

```shell
go tool pprof http://localhost:6060/debug/pprof/heap
```

top

```shell
(pprof) top
Showing nodes accounting for 1.20GB, 100% of 1.20GB total
Dropped 4 nodes (cum <= 0.01GB)
      flat  flat%   sum%        cum   cum%
    1.20GB   100%   100%     1.20GB   100%  github.com/wolfogre/go-pprof-practice/animal/muridae/mouse.(*Mouse).Steal
         0     0%   100%     1.20GB   100%  github.com/wolfogre/go-pprof-practice/animal/muridae/mouse.(*Mouse).Live
         0     0%   100%     1.20GB   100%  main.main
         0     0%   100%     1.20GB   100%  runtime.main

```

查看到占用1G多的内存

**内存回收**

```shell
go tool pprof http://localhost:6060/debug/pprof/allocs
```

top

```
(pprof) top
Showing nodes accounting for 592MB, 99.75% of 593.50MB total
Dropped 16 nodes (cum <= 2.97MB)
      flat  flat%   sum%        cum   cum%
     592MB 99.75% 99.75%      592MB 99.75%  github.com/wolfogre/go-pprof-practice/animal/canidae/dog.(*Dog).Run (inline)
         0     0% 99.75%      592MB 99.75%  github.com/wolfogre/go-pprof-practice/animal/canidae/dog.(*Dog).Live
         0     0% 99.75%      592MB 99.75%  main.main
         0     0% 99.75%   592.50MB 99.83%  runtime.main

```

可以看到 `github.com/wolfogre/go-pprof-practice/animal/canidae/dog.(*Dog).Run` 会进行无意义的内存申请，而这个函数又会被频繁调用，这才导致程序不停地进行 GC:

```go
func (d *Dog) Run() {
	log.Println(d.Name(), "run")
	_ = make([]byte, 16 * constant.Mi)
}
```

**排查协程泄露**

```shell
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

```
(pprof) top
Showing nodes accounting for 103, 99.04% of 104 total
Showing top 10 nodes out of 33
      flat  flat%   sum%        cum   cum%
       102 98.08% 98.08%        102 98.08%  runtime.gopark
         1  0.96% 99.04%          1  0.96%  runtime.goroutineProfileWithLabels
         0     0% 99.04%        100 96.15%  github.com/wolfogre/go-pprof-practice/animal/canidae/wolf.(*Wolf).Drink.func1
         0     0% 99.04%          1  0.96%  github.com/wolfogre/go-pprof-practice/animal/felidae/cat.(*Cat).Live
         0     0% 99.04%          1  0.96%  github.com/wolfogre/go-pprof-practice/animal/felidae/cat.(*Cat).Pee
         0     0% 99.04%          1  0.96%  internal/poll.(*FD).Accept
         0     0% 99.04%          1  0.96%  internal/poll.(*FD).acceptOne
         0     0% 99.04%          1  0.96%  internal/poll.(*pollDesc).wait
         0     0% 99.04%          1  0.96%  internal/poll.execIO
```

**排查锁的争用**

```
go tool pprof http://localhost:6060/debug/pprof/mutex
```

```
(pprof) top
Showing nodes accounting for 126.40s, 100% of 126.40s total
      flat  flat%   sum%        cum   cum%
   126.40s   100%   100%    126.40s   100%  sync.(*Mutex).Unlock (inline)
         0     0%   100%    126.40s   100%  github.com/wolfogre/go-pprof-practice/animal/canidae/wolf.(*Wolf).Howl.func1

```

```go
func (w *Wolf) Howl() {
	log.Println(w.Name(), "howl")

	m := &sync.Mutex{}
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
}
```

可以看到，这个锁由主协程 Lock，并启动子协程去 Unlock，主协程会阻塞在第二次 Lock 这儿等待子协程完成任务，但由于子协程足足睡眠了一秒，导致主协程等待这个锁释放足足等了一秒钟。

