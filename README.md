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
c.Request("GET", "/user", q, nil, &user)
```
