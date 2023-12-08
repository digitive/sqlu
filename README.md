# sqlu
`sqlu` is a minimalist SQL update query builder package.

## Why
The awesome package [SQLC](https://github.com/sqlc-dev/sqlc) can handle most of the SQL job with one exception which is updating tables dynamically. `sqlu` can be used as a complementary package which handles the dynamic update job.

## Example

```go
package main

import (
    "github.com/digitive/sqlu"
)

func main() {
    if _, err := sqlu.Update("people").Set("age", 18).Where("name=?", "John Smith"); err != nil {
        panic(err)
    }

    if _, err := sqlu.Update("people").Setm(sqlu.Fields{
      "age": 18,
      "mobile": "1234567890",
    }).Where("name=?", "John Smith"); err != nil {
        panic(err)
    }
}
```
