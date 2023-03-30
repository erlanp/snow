# snow
生成分布式id 这是一个练手项目, 如用于生产, 参见postgre的协议。
```go
package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/erlanp/snow"
	"github.com/valyala/fasthttp"
)

func Sequence(ctx *fasthttp.RequestCtx) {
	seq := snow.NewSeq(filterString(ctx.FormValue("name")))

	fmt.Fprint(ctx, fmt.Sprintf("%d", (*seq).Incr()))
}

func Snow(ctx *fasthttp.RequestCtx) {
	snow := snow.NewSnow(filterString(ctx.FormValue("name")), 5)
	fmt.Fprint(ctx, fmt.Sprintf("%d", (*snow).Gen()))
}

func main() {
	router := fasthttprouter.New()

	router.GET("/sequence", Sequence)
	router.GET("/snow", Snow)

	log.Fatal(fasthttp.ListenAndServe("127.0.0.1:8486", router.Handler))
}

func filterString(b []byte) string {
	j := len(b)
	for i := 0; i < j; i++ {
		if (b[i] < 48 || b[i] > 57) && (b[i] < 65 || b[i] > 90) && (b[i] < 97 || b[i] > 122) {
			return ""
		}
	}
	return string(b)
}
```