package snow

import (
    "testing"
    "fmt"
    // "io"
    // "math/rand"
)

func TestMin(t *testing.T) {
    
    seq := Seq{name:"snow.txt", log_cnt: 8};
    snow := Snow{seq:seq, shard_id:5}
    snow.Init();
    for i:=0;i<=10;i++ {
        
        fmt.Println(snow.Gen());
    }

}
