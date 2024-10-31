# Go Get Start

1. 函数中定义变量的值

```go
func main() {
    a := 10
    fmt.Println(a)
}
```

2. 基本数值类型

- 整数 : `int` , `uint`
- 浮点数 : `float32` , `float64`
- 布尔值 : `bool`
- 字符串 : `string`

3.  运算符
- 加法 : `+`
- 减法 : `-`
- 乘法 : `*`
- 除法 : `/`
- 取余 : `%`
- 自增 : `++`
- 自减 : `--`
- 位与 : `&`
- 位或 : `|`
- 非 : `!`

4. 条件语句

- if-else 语句  
```go 
func main() {
    a := 10
    if a > 0 {
        fmt.Println("a is positive")
    } else {
        fmt.Println("a is negative")
    }
}
```

在`if条件语句`中可以初始化变量，在`else`中也可以使用变量。


- switch...case 语句  

```go
a := 10
switch a {
    case 10:
        fmt.Println("a is 10")
    case 20:
        fmt.Println("a is 20")
    default:
        fmt.Println("a is not 10 or 20")
}
```

5. 循环语句

- for 循环
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

6. 函数

```go
func sum(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println(sum(10, 20))
}
```

go 语言中的`函数`可以返回多个值, 只需要把返回值放在`()`中即可。

```go 
func swap(a,b int) (int,int){
    return b,a
}

func main() {
    x,y := swap(10,20)
    fmt.Println(x,y)
}
```

函数也可以作为值传递给其他函数。

```go
func sum (a int, b int, 
    transform func(int) int) int {
    return transform(a+b)
}
```

7. 数组

数组在go中长度固定，不能动态扩容。

```go
arr := [5]int{1,2,3,4,5}
```

可变长度的数组需要使用`Slice`。

```go
arr := []int{1,2,3,4,5}
```

8. 键值对

Map :在go 中使用map可以方便的存储键值对。

```go
m := map[string]int{"a":1,"b":2}
fmt.Println(m["a"]) // 1
```

如果是没有初始值的空数组，可以使用`make`函数。

```go
arr := make([]int, 5)
```

9. 结构

使用type关键字可以定义结构体。

```go 
type Person struct {
    name string
    age int
}

func main() {
    p := Person{"Alice", 20}
    fmt.Println(p.name) // Alice
}
```

10. 指针

```go
q := &p
p.age = 3
fmt.Println(*q) // 3
fmt.Println(p) // 3
```

11. 方法

结构可以添加方法，方法可以访问结构体的字段。

```go
func (p *Person) sayHello() {
    fmt.Println("Hello, my name is", p.name)
}

p.sayHello()
```

