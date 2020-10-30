package snow

import (
    "testing"
    "fmt"
    "strconv"
    // "io"
    // "math/rand"
)

func TestWare(t *testing.T) {
    var m map[string]Seq = map[string]Seq{};
    co := Seq{name:"test1.txt"};
    co.Init();
    m["test1.txt"] = co;
    co2 := m["test1.txt"];
    
    fmt.Println(co2.Incr());
    i, _ := strconv.ParseInt("../GH45aA", 10, 64)
    j, _ := strconv.ParseInt("123", 10, 64)
    fmt.Println(fmt.Sprintf("%d", i + j));
}    

