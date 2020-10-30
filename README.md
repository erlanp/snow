# snow
生成id 这是一个练手项目, 如用于生产,参见postgre的协议。 这是一个go项目， java项目 https://github.com/erlanp/snow-id
例子
```go
package main

import (
	"fmt"
	"log"

	"github.com/erlanp/snow"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)



func Index(ctx *fasthttp.RequestCtx) {
	seq := snow.NewSeq(filter2(ctx.FormValue("a")));
	
	fmt.Fprint(ctx, fmt.Sprintf("%d", (*seq).Incr()));
}

func Index2(ctx *fasthttp.RequestCtx) {
	snow := snow.NewSnow(filter2(ctx.FormValue("a")));
	fmt.Fprint(ctx, fmt.Sprintf("%d", (*snow).Gen()));
}


func main() {
	router := fasthttprouter.New()

	router.GET("/", Index);
	router.GET("/index2", Index2)

	log.Fatal(fasthttp.ListenAndServe("127.0.0.1:8486", router.Handler))
}

func filter2(b []byte) string {
	j := len(b);
	for i:=0; i<j; i++ {
		if (b[i] < 48 || b[i] > 57) {
			return "";
		}
	}
	return string(b);
}

```