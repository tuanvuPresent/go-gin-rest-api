## Go

---
### Start
- Hello.go
```go
// hello.go
package main

import "fmt"

func main() {
    fmt.Printf("Hello World\n")
}
```
---
### Array
```go
var a [3]int
var b = [...]int{1, 2, 3}
var c = [...]int{2: 3, 1: 2}
var d = [...]int{1, 2, 4: 5, 6}
```
- Khi biến array được gán hoặc truyền, thì toàn bộ array sẽ được sao chép. Nếu kích thước của array lớn, thì phép gán array sẽ chịu tổn phí lớn. Để tránh việc overhead (tổn phí) trong việc sao chép array, bạn có thể truyền con trỏ của array.
```go
var a = [...]int{1, 2, 3}
// b là một con trỏ tới array a
var b = &a
```
- Duyệt array
```go
for i := range a {
fmt.Printf("a[%d]: %d\n", i, a[i])
}
for i, v := range a {
fmt.Printf("a[%d]: %d\n", i, v)
}
for i := 0; i < len(a); i++ {
fmt.Printf("c[%d]: %d\n", i, a[i])
}
```
- Slice: Khi chúng ta sử dụng cú pháp tạo slice từ một slice cho trước như sau: d = c[:2] thì chúng ta nên lưu ý ở điểm sau. Slice sẽ không sao chép dữ liệu của slice gốc c qua d mà nó tạo ra một giá trị slice mới trỏ đến mảng ban đầu. Do đó, sửa đổi các phần tử của slice vừa được tạo sẽ sửa đổi các phần tử của slice gốc.
- Trong trường hợp slice ban đầu không đủ sức chứa khi thêm vào phần tử, hàm append sẽ hiện thực cấp phát lại vùng nhớ có kích thước:
  + Nếu kích thước cũ (cap) < 1024: cấp phát gấp đôi (x2) vùng nhớ cũ.
  + Nếu kích thước cũ >= 1024: cấp phát 1.25x vùng nhớ cũ.
---
### String
- string cũng là một array của các byte dữ liệu, nhưng khác với array những phần tử của string là immutable.
```go
var s = "hello world"
hello := s[:5]
world := s[6:]
fmt.Println(hello)
fmt.Println(world)
```
---
### Function
- Một hàm trong ngôn ngữ Go có thể có nhiều tham số và nhiều giá trị trả về. Cả tham số và giá trị trả về trao đổi dữ liệu với hàm theo cách truyền vào giá trị (pass by value)
```go
func Add(a, b int) int {
    return a + b
}
var Add = func(a, b int) int {
    return a + b
}
func Sum(a int, more ...int) int {
  for _, v := range more {
    a += v
  }
  return a
}
```
- Defer trong Function: Lệnh defer trì hoãn việc thực thi hàm cho tới khi hàm bao ngoài nó return. Các đối số trong lời gọi defer được đánh giá ngay lặp tức nhưng lời gọi không được thực thi cho tới khi hàm bao ngoài nó return.
- Truyền vào con trỏ ngầm định (slice) khiến nội dung của biến x bị thay đổi
```go
func change(x []int) {
    for i := range x {
        x[i] *= 2
    }
}
data := [3]int{8,9,0}
change(data[0:])
```
---
### Type
- Đối với một kiểu nhất định, tên của mỗi phương thức phải là duy nhất và các phương thức cũng như hàm đều không hỗ trợ overload.
- trong Go khái niệm đóng gói khá đơn giản: chữ cái đầu tên viết hoa thì mở, còn viết thường thì đóng. Quy tắt này áp dụng cho hằng, biến, hàm, trường, phương thức, v.v... Có điều trong Go, khái niệm mở hay đóng chỉ áp dụng bên ngoài package. Trong package, mọi cái đều mở dù tên viết hoa hay viết thường.


```go
package main

import (
	"fmt"
	"math"
)

type Complex struct {
	a, b float64
}

func (c Complex) Module() float64 {
	return math.Sqrt(c.a*c.a + c.b*c.b)
}

func main() {
	c := Complex{3, 4}
	fmt.Println(c.Module())
}
```

```go
type Human struct {
    FirstName   string
    LastName    string
    CanSwim     bool
}
// Amy được embedded với kiểu Human
// và do đó Amy có thể gọi các phương thức của Human
type Amy struct {
    Human
}
```
- Interface
```go
type Human interface {
	Run() string
}
type Amy struct {
	Human
}

func (a Amy) Run() string {
	fmt.Println("Run")
}
```
---
### Go và web (Gin)
```go
package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func ping(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
}

func main() {
  router := gin.Default()
  router.GET("/ping", ping)
  router.Run("0.0.0.0:8006")
}
```
