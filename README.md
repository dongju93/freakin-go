# Go 언어 정복기

저는 Intermediate 레벨의 Python 백엔드 개발자입니다. 하지만 스크립트 언어를 넘어 low-level 프로그래밍의 세계에 더 깊이 들어가고 싶다는 열망이 있었습니다. 그래서 Go를 선택하게 되었습니다.

프로그래밍의 기본은 이미 알고 있지만, 새로운 언어인 만큼 변수, 함수, 반복문과 같은 가장 기초적인 부분부터 학습하며 진행할 것입니다.

새로운 언어를 배우는 과정이기 때문에, VS Code 에디터의 GitHub Copilot: Auto-complete 기능으 의도적으로 비활성화하고 학습을 진행하고 있습니다.
하지만 모르는 개념을 만나거나 이해가 어려운 부분에 부딪혔을 때는 GitHub Copilot: Chat-Ask 기능을 적극적으로 활용하여 질문하고 해답을 찾고 있습니다. 이처럼 적절한 AI 도구 활용은 학습 효율을 크게 높여줄 수 있다고 생각합니다

### Go를 선택한 5가지 이유

1.  **간단하고 강력한 동시성 처리**: Go의 고루틴(goroutine)과 채널(channel)은 복잡한 동시성 프로그래밍을 매우 쉽게 만들어줍니다. Python의 GIL(Global Interpreter Lock)의 한계를 넘어 진정한 병렬 처리를 경험하고 싶었습니다.
2.  **압도적인 성능**: Go는 컴파일 언어로서 C/C++에 버금가는 빠른 실행 속도를 자랑합니다. 스크립트 언어인 Python과 비교했을 때, 특히 대규모 트래픽을 처리하는 백엔드 시스템에서 성능상의 이점이 분명합니다.
3.  **단순하고 명료한 문법**: Go는 의도적으로 언어의 복잡성을 줄여, 배우기 쉽고 코드를 읽고 이해하기 편합니다. 이는 "하나의 일을 하는 데에는 하나의 명확한 방법이 있어야 한다"는 철학을 가진 Python 개발자에게 매우 매력적인 부분입니다.
4.  **정적 타이핑의 안정성**: 컴파일 시점에 타입 에러를 잡을 수 있는 정적 타이핑은 코드의 안정성과 신뢰성을 크게 높여줍니다. Python의 동적 타이핑이 주는 유연함도 좋지만, 대규모 프로젝트에서는 정적 타이핑이 더 큰 이점을 가진다고 생각합니다.
5.  **강력한 표준 라이브러리와 툴체인**: Go는 웹 서버, 암호화 등 필수적인 기능을 포함한 풍부한 표준 라이브러리를 제공합니다. 또한 `go fmt` (코드 포맷팅), `go test` (테스팅) 등 내장된 툴체인은 개발 생산성을 극대화합니다.

하나의 언어나 프레임워크에 깊이 파고들려면 새로운 버전에 대한 정보를 잘 알고 있어야 한다고 믿습니다. 그래서 Go의 버전별 주요 변경 사항을 꾸준히 기록하고 정리할 예정입니다.

---

# Go 버전별 변경 사항 (Changelog)

## Go 1.22

### 새로운 구문 (New Syntax)

- **`for` 루프 변수 의미 체계 변경**: `for` 루프의 각 반복에서 새로운 변수를 생성하여 의도치 않은 변수 공유 버그를 해결했습니다.

  ```go
  // Before Go 1.22 (shared variable bug):
  var funcs []func()
  for i := 0; i < 3; i++ {
      funcs = append(funcs, func() { println(i) }) // Always prints 3
  }

  // Go 1.22+ (each iteration creates new variable):
  var funcs []func()
  for i := 0; i < 3; i++ {
      funcs = append(funcs, func() { println(i) }) // Prints 0, 1, 2
  }
  ```

- **정수 범위 반복**: `for` 루프에서 정수 범위를 직접 반복할 수 있게 되어 고정된 횟수의 반복을 더 간결하게 작성할 수 있습니다.

  ```go
  // Before:
  for i := 0; i < 10; i++ {
      doSomething(i)
  }

  // Go 1.22+:
  for i := range 10 {
      doSomething(i) // i goes from 0 to 9
  }
  ```

### 핵심 라이브러리 (Core Libraries)

- **`math/rand/v2` 패키지 추가**: 더 깨끗하고 일관된 API를 제공하며, 고품질의 빠른 의사 난수 생성 알고리즘을 사용하는 새로운 `math/rand/v2` 패키지가 도입되었습니다.

  ```go
  // Old math/rand:
  import "math/rand"
  rand.Seed(time.Now().UnixNano())
  n := rand.Intn(100)

  // New math/rand/v2:
  import "math/rand/v2"
  n := rand.IntN(100) // No seeding required, better quality
  ```

- **`net/http.ServeMux` 향상**: HTTP 라우팅이 더 표현력이 풍부해져, `net/http.ServeMux` 패턴이 메서드와 와일드카드(예: `GET /task/{id}/`)를 허용합니다.

  ```go
  mux := http.NewServeMux()

  // Method-specific routing:
  mux.HandleFunc("GET /users/{id}", getUserHandler)
  mux.HandleFunc("POST /users", createUserHandler)
  mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

  // Wildcard patterns:
  mux.HandleFunc("GET /files/{path...}", serveFileHandler)
  ```

- **`database/sql.Null[T]`**: `database/sql`에 새로운 `Null[T]` 타입을 추가하여 nullable 컬럼을 스캔할 수 있는 방법을 제공합니다.

  ```go
  // Before:
  var name sql.NullString
  var age sql.NullInt64
  err := row.Scan(&name, &age)

  // Go 1.22+:
  var name sql.Null[string]
  var age sql.Null[int64]
  err := row.Scan(&name, &age)

  if name.Valid {
      fmt.Println(name.V) // .V instead of .String
  }
  ```

- **`slices.Concat`**: `slices` 패키지에 여러 슬라이스를 연결하는 `Concat` 함수가 추가되었습니다.

  ```go
  import "slices"

  s1 := []int{1, 2, 3}
  s2 := []int{4, 5}
  s3 := []int{6, 7, 8}

  result := slices.Concat(s1, s2, s3) // [1 2 3 4 5 6 7 8]
  ```

### 성능 (Performance)

- **런타임 메모리 최적화**: 런타임이 유형 기반 가비지 컬렉션 메타데이터를 힙 객체에 더 가깝게 유지하여 CPU 성능을 1-3% 향상시키고 대부분의 Go 프로그램에서 메모리 오버헤드를 약 1% 줄였습니다.
- **프로필 기반 최적화 (PGO)**: PGO 빌드는 더 많은 호출을 가상화 해제할 수 있게 되어, PGO가 활성화된 대부분의 프로그램에서 2-14%의 런타임 향상을 가져옵니다.

## Go 1.23

### 새로운 구문 (New Syntax)

- **`for-range` 루프의 이터레이터 함수**: `for-range` 루프의 `range` 절이 이제 이터레이터 함수(예: `func(func(K) bool)` 또는 `func(func(K, V) bool)`)를 허용하여 사용자 정의 시퀀스에 대한 이터레이터를 만들 수 있습니다.

  ```go
  // Single-value iterator:
  func fibonacci(max int) func(func(int) bool) {
      return func(yield func(int) bool) {
          a, b := 0, 1
          for a < max {
              if !yield(a) {
                  return
              }
              a, b = b, a+b
          }
      }
  }

  // Usage:
  for n := range fibonacci(100) {
      fmt.Println(n) // 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89
  }

  // Key-value iterator:
  func enumerate[T any](slice []T) func(func(int, T) bool) {
      return func(yield func(int, T) bool) {
          for i, v := range slice {
              if !yield(i, v) {
                  return
              }
          }
      }
  }

  // Usage:
  for i, v := range enumerate([]string{"a", "b", "c"}) {
      fmt.Printf("%d: %s\n", i, v)
  }
  ```

### 핵심 라이브러리 (Core Libraries)

- **`time.Timer` 및 `time.Ticker` 변경**:

  - 프로그램에서 더 이상 참조되지 않는 타이머와 티커는 `Stop` 메서드가 호출되지 않았더라도 즉시 가비지 컬렉션 대상이 되어 메모리 누수를 방지합니다.
  - 타이머 채널이 버퍼링되지 않도록(용량 0) 변경되어 `Reset` 또는 `Stop` 호출 후 오래된 값이 전송되거나 수신되지 않도록 보장합니다.

  ```go
  // Before Go 1.23 (potential memory leak):
  func createManyTimers() {
      for i := 0; i < 1000; i++ {
          timer := time.NewTimer(time.Hour)
          // Forgot to call timer.Stop() - memory leak!
      }
  } // Timers remain in memory until they fire

  // Go 1.23+ (automatic cleanup):
  func createManyTimers() {
      for i := 0; i < 1000; i++ {
          timer := time.NewTimer(time.Hour)
          // timer automatically GC'd when out of scope
      }
  } // Timers are cleaned up immediately

  // Timer channel behavior change:
  timer := time.NewTimer(time.Second)
  timer.Stop()
  // Before: timer.C might still deliver a value
  // Go 1.23+: timer.C will never deliver after Stop()
  ```

- **새로운 패키지**:

  - `iter`: 이터레이터 작업을 위한 함수를 제공합니다.

    ```go
    import "iter"

    // Pull converts push-style iterator to pull-style
    next, stop := iter.Pull(fibonacci(100))
    defer stop()

    for {
        v, ok := next()
        if !ok {
            break
        }
        fmt.Println(v)
    }
    ```

  - `structs`: 구조체 속성을 수정하기 위한 마커 유형을 정의합니다.

    ```go
    import "structs"

    type Config struct {
        Name    string
        Timeout time.Duration `structs:"timeout"`
        _       structs.NoUnkeyedLiteral // Prevents struct{name, timeout}
        _       structs.NoComparable    // Makes struct not comparable
    }
    ```

  - `unique`: 비교 가능한 값을 정식화("인터닝")하여 메모리를 절약하고 비교 효율성을 높이는 기능을 제공합니다.

    ```go
    import "unique"

    // Intern strings to save memory
    handle1 := unique.Make("hello")
    handle2 := unique.Make("hello")

    // handle1 == handle2 (same underlying pointer)
    // Fast comparison, reduced memory usage
    fmt.Println(handle1.Value()) // "hello"
    ```

- **`maps` 패키지 추가**: `All`, `Keys`, `Values`, `Insert`, `Collect`와 같이 이터레이터와 함께 작동하는 여러 새로운 함수가 `maps` 패키지에 추가되었습니다.

  ```go
  import "maps"

  m := map[string]int{"a": 1, "b": 2, "c": 3}

  // Iterate over key-value pairs
  for k, v := range maps.All(m) {
      fmt.Printf("%s: %d\n", k, v)
  }

  // Iterate over keys only
  for k := range maps.Keys(m) {
      fmt.Println("key:", k)
  }

  // Iterate over values only
  for v := range maps.Values(m) {
      fmt.Println("value:", v)
  }

  // Insert multiple key-value pairs
  newEntries := map[string]int{"d": 4, "e": 5}
  maps.Insert(m, maps.All(newEntries))

  // Collect iterator into map
  doubled := maps.Collect(func(yield func(string, int) bool) {
      for k, v := range maps.All(m) {
          if !yield(k, v*2) {
              return
          }
      }
  })
  ```

### 도구 (Tooling)

- **Go 원격 측정 (Telemetry)**: Go 1.23부터 Go 툴체인은 사용 및 중단 통계를 수집하여 Go 팀이 툴체인 사용 및 성능을 이해하는 데 도움을 줍니다. (선택적 참여)

## Go 1.24

### 새로운 구문 (New Syntax)

- **제네릭 타입 별칭 전체 지원**: 타입 별칭이 정의된 타입처럼 매개변수화될 수 있도록 제네릭 타입 별칭을 완벽하게 지원합니다.

  ```go
  // Generic type alias:
  type Map[K comparable, V any] = map[K]V
  type StringMap[V any] = Map[string, V]

  // Usage:
  var userAges StringMap[int] = Map[string, int]{
      "alice": 30,
      "bob":   25,
  }

  // Function type alias:
  type Handler[T any] = func(T) error
  type StringHandler = Handler[string]

  func process(h StringHandler, data string) error {
      return h(data)
  }
  ```

### 핵심 라이브러리 (Core Libraries)

- **FIPS 140-3 준수**: Go 암호화 모듈 및 `GOFIPS140` 환경 변수를 포함하여 FIPS 140-3 준수를 위한 메커니즘을 도입했습니다.

  ```bash
  # Enable FIPS mode:
  export GOFIPS140=1
  go run myapp.go  # Uses FIPS-validated crypto implementations
  ```

- **새로운 암호화 패키지**: `crypto/mlkem`, `crypto/hkdf`, `crypto/pbkdf2`, `crypto/sha3`가 추가되었습니다.

  ```go
  import (
      "crypto/hkdf"
      "crypto/pbkdf2"
      "crypto/sha3"
      "crypto/mlkem"
  )

  // HKDF key derivation:
  key := hkdf.Extract(sha256.New, []byte("secret"), []byte("salt"))
  derivedKey := hkdf.Expand(sha256.New, key, []byte("info"), 32)

  // PBKDF2 password-based key derivation:
  key := pbkdf2.Key([]byte("password"), []byte("salt"), 10000, 32, sha256.New)

  // SHA-3 hashing:
  hash := sha3.Sum256([]byte("data"))

  // ML-KEM (post-quantum cryptography):
  publicKey, privateKey, err := mlkem.GenerateKey768()
  ciphertext, sharedSecret, err := mlkem.Encapsulate(publicKey)
  recoveredSecret, err := mlkem.Decapsulate(privateKey, ciphertext)
  ```

- **`os.Root`**: 특정 디렉토리 내에서만 파일 시스템 접근을 허용하는 `os.Root` 타입을 추가하여 경로가 외부 위치를 참조하는 것을 방지합니다.

  ```go
  // Create a sandboxed filesystem root:
  root, err := os.NewRoot("/safe/directory")
  if err != nil {
      log.Fatal(err)
  }

  // All operations are constrained to /safe/directory:
  file, err := root.Open("file.txt")          // Opens /safe/directory/file.txt
  err = root.MkdirAll("subdir/nested", 0755)  // Creates /safe/directory/subdir/nested

  // These will fail - path escapes:
  _, err = root.Open("../../../etc/passwd")   // Error: path escapes root
  _, err = root.Open("/etc/passwd")           // Error: absolute path not allowed
  ```

- **`bytes` 및 `strings` 패키지 개선**: 새로운 이터레이터 함수가 추가되었습니다.

  ```go
  import (
      "bytes"
      "strings"
  )

  // Iterate over lines in a string:
  for line := range strings.Lines("line1\nline2\nline3") {
      fmt.Println(line)
  }

  // Iterate over bytes:
  for b := range bytes.Values([]byte("hello")) {
      fmt.Printf("%c ", b) // h e l l o
  }

  // Split with iterator:
  for part := range strings.Split("a,b,c", ",") {
      fmt.Println(part) // a, b, c
  }
  ```

- **`runtime.AddCleanup`**: `SetFinalizer`보다 더 유연하고 효율적이며 오류 발생 가능성이 적은 정리 메커니즘을 제공합니다.

  ```go
  import "runtime"

  type Resource struct {
      fd int
  }

  func NewResource() *Resource {
      r := &Resource{fd: syscall.Open("file")}

      // Old way with SetFinalizer:
      runtime.SetFinalizer(r, (*Resource).Close)

      // New way with AddCleanup (Go 1.24+):
      runtime.AddCleanup(r, func() {
          if r.fd != -1 {
              syscall.Close(r.fd)
              r.fd = -1
          }
      })

      return r
  }

  func (r *Resource) Close() {
      if r.fd != -1 {
          syscall.Close(r.fd)
          r.fd = -1
          // Cleanup automatically removed when called
      }
  }
  ```

### 성능 (Performance)

- **CPU 오버헤드 감소**: 새로운 Swiss Tables 기반 맵 구현, 더 효율적인 작은 객체 메모리 할당, 새로운 런타임 내부 뮤텍스 구현을 통해 벤치마크 전반에 걸쳐 평균 2-3%의 CPU 오버헤드를 줄였습니다.

### 도구 (Tooling)

- **실행 파일 종속성 추적**: `go.mod`의 `tool` 지시문을 사용하여 실행 파일 종속성을 추적할 수 있어 `tools.go` 빈 임포트가 필요 없어졌습니다.

  ```go
  // Before Go 1.24 (tools.go):
  //go:build tools
  package tools

  import (
      _ "github.com/golangci/golangci-lint/cmd/golangci-lint"
      _ "github.com/swaggo/swag/cmd/swag"
  )

  // Go 1.24+ (go.mod):
  module myproject

  go 1.24

  tool (
      github.com/golangci/golangci-lint/cmd/golangci-lint v1.54.0
      github.com/swaggo/swag/cmd/swag v1.16.0
  )

  // Usage:
  // go run github.com/golangci/golangci-lint/cmd/golangci-lint run
  ```

## Go 1.25

### 핵심 라이브러리 (Core Libraries)

- **실험적인 `encoding/json/v2` 패키지**: 성능 향상을 목표로 하는 새로운 JSON 구현인 `encoding/json/v2` 패키지를 실험적으로 도입했습니다. (`GOEXPERIMENT=jsonv2`로 활성화)

  ```bash
  # Enable experimental JSON v2:
  export GOEXPERIMENT=jsonv2
  go build -tags=jsonv2 myapp.go
  ```

  ```go
  import "encoding/json/v2"

  type User struct {
      Name  string `json:"name"`
      Email string `json:"email,omitempty"`
  }

  // Improved performance and better API:
  user := User{Name: "Alice", Email: "alice@example.com"}

  // Marshal with options:
  data, err := json.Marshal(user, json.WithIndent("", "  "))

  // Unmarshal with validation:
  var decoded User
  err = json.Unmarshal(data, &decoded, json.RejectUnknownMembers())
  ```

- **`testing/synctest` 안정화**: 동시성 코드 테스트를 지원하는 `testing/synctest` 패키지가 안정화되었습니다.

  ```go
  import "testing/synctest"

  func TestConcurrentCode(t *testing.T) {
      synctest.Run(func() {
          // Test concurrent code with deterministic scheduling
          ch := make(chan int, 1)

          go func() {
              ch <- 42
          }()

          go func() {
              val := <-ch
              assert.Equal(t, 42, val)
          }()

          // synctest ensures deterministic execution order
      })
  }
  ```

### 성능 (Performance)

- **DWARF5 지원**: 컴파일러와 링커가 DWARF 버전 5를 사용하여 디버그 정보를 생성하여 바이너리의 디버깅 정보 공간을 줄이고 링크 속도를 높입니다.

  ```bash
  # Compile with DWARF5 debug info (default in Go 1.25):
  go build -gcflags="-dwarfversion=5" myapp.go

  # Results in smaller binaries and faster linking
  ```

- **컨테이너 인식 `GOMAXPROCS`**: Linux에서 런타임이 cgroup의 CPU 대역폭 제한을 고려하여 `GOMAXPROCS`의 기본값을 설정합니다.

  ```go
  // Before Go 1.25 (container with 0.5 CPU limit):
  runtime.GOMAXPROCS(0) // Returns number of host CPUs (e.g., 8)

  // Go 1.25+ (container aware):
  runtime.GOMAXPROCS(0) // Returns 1 (respects cgroup CPU limit of 0.5)

  // Manual override still works:
  runtime.GOMAXPROCS(2) // Force 2 goroutines
  ```

- **실험적인 가비지 컬렉터**: 작은 객체의 마킹 및 스캐닝 성능을 향상시키는 새로운 실험적인 가비지 컬렉터를 사용할 수 있습니다.

  ```bash
  # Enable experimental GC:
  export GOEXPERIMENT=gcexperiment
  go run myapp.go

  # Or at build time:
  go build -tags=gcexperiment myapp.go
  ```

### 도구 (Tooling)

- **`go.mod ignore` 지시문**: `go` 명령이 무시할 디렉토리를 지정할 수 있는 `go.mod ignore` 지시문을 추가했습니다.

  ```go
  // go.mod:
  module myproject

  go 1.25

  ignore (
      testdata/     // Ignore test data directories
      vendor/       // Ignore vendor directory
      .git/         // Ignore git directories
      node_modules/ // Ignore npm modules
  )
  ```

- **`go doc -http`**: 문서 서버를 시작하고 브라우저에서 여는 `go doc -http` 옵션을 추가했습니다.

  ```bash
  # Start documentation server and open browser:
  go doc -http=:8080

  # Automatically opens http://localhost:8080 in default browser
  # Shows documentation for current module and dependencies
  ```
