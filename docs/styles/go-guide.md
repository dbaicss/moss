### 包名

package 命名的几条规则：
- 全小写。不包含大写字母或者下划线。
- 简洁。
- 不要使用复数。比如，使用 net/url，而不是 net/urls。
- 避免："common", "util", "shared", "lib"，不解释。

更多参考[Package Names] 和 [Style guideline for Go packages].

  [Package Names]: https://blog.golang.org/package-names
  [Style guideline for Go packages]: https://rakyll.org/style-packages/
 

### 声明语句分组

go支持分组声明

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
import "a"
import "b"
```

</td><td>

```go
import (
  "a"
  "b"
)
```

</td></tr>
</tbody></table>

同样也适用于常量、变量、type的申明
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go

const a = 1
const b = 2



var a = 1
var b = 2



type Area float64
type Volume float64
```

</td><td>

```go
const (
  a = 1
  b = 2
)

var (
  a = 1
  b = 2
)

type (
  Area float64
  Volume float64
)
```

</td></tr>
</tbody></table>

### 使用原生字符串，避免转义

Go 支持使用反引号，也就是 "`" 来表示原生字符串，在需要转义的场景下，我们应该尽量使用这种方案来替换。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
wantError := "unknown name:\"test\""
```

</td><td>

```go
wantError := `unknown error:"test"`
```

</td></tr>
</tbody></table>
