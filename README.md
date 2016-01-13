# rest-client

Simple REST client helper written in Go.

## Usage

```
import (
    "net/url"
    "github.com/pulcy/rest-client"
)

c := restclient.NewRestClient(baseURL)
var user UserType
q := url.Values{}
q.Set("id", "some-user-id")
if err := c.Request("GET", "/user", q, nil, &user); err != nil {
    panic("GET /user failed")
}
// Modify user ...
if err := c.Request("POST", "/user", nil, user, nil); err != nil {
    panic("POST /user failed")
}
```
