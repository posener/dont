# package dont

Here you can find the towel you need to travel the galaxy.

### dont.Panic

The `dont.Panic` function can help you handle errors that you expect not to have.

Usage example:

```go
package hitchhike

import (
    "os/user"
    "github.com/posener/dont"
)

var arthur = dont.Panic(user.Current())
```

> It is recommended only to use the `dont.Panic` function to initialize package scope variables,
> where in case of a panic, the program will crush on loading time.