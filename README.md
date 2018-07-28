# KBO

Go언어로 만들어진 비공식 KBO 경기 결과 API

## 설치

1. 만약 서버까지 설치하고 싶으시면
```
$ go get github.com/seeeturtle/kbo/...
```

2. 만약 라이브러리만 설치하고 싶으시면
```
$ go get github.com/seeeturtle/kbo
```

## 실행

```
$ kbo-api -help
Usage of kbo-api:
  -log string
        file name to log (default "kbo-api.log")
  -port string
        port to listen (default "8080")
```

## 예제

```go
import (
    "fmt"
    "time"

    "github.com/seeeturtle/kbo"
)

func main() {
    parser := kbo.NewParser(
        kbo.URL,
        &http.Client{Timeout: 10*time.Second}, // client에 특정한 설정을 걸 수 있습니다.
    )

    games := parser.Parse(time.Date(2018, 5, 11, 0, 0, 0, 0, time.UTC))
    fmt.Println(games)
}
```
