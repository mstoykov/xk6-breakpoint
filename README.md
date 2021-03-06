# xk6-breakpoint

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
| ---------------------------------------------------------------------------------------------------------------------------- |

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [gvm](https://github.com/moovweb/gvm)
- [Git](https://git-scm.com/)

Then, install [xk6](https://github.com/k6io/xk6) and build your custom k6 binary with the Kafka extension:

1. Install `xk6`:
  ```shell
  $ go install github.com/k6io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  $ xk6 build --with github.com/mstoykov/xk6-breakpoint@latest
  ```

# example

```javascript
import breakpoint from "k6/x/breakpoint"
export default function() {
    var s = "something";
    breakpoint.break()
}
```

Some asciinema showing how to (ab)use it:

[![basic](https://asciinema.org/a/GvhZaBg8WiL1yh9gdsLveJaPB.svg)](https://asciinema.org/a/GvhZaBg8WiL1yh9gdsLveJaPB)
[![advanced](https://asciinema.org/a/dPKP4kcotM6RU5kgWUkfGFbR6.svg)](https://asciinema.org/a/dPKP4kcotM6RU5kgWUkfGFbR6)
[![multiple VUs (not recommended)](https://asciinema.org/a/hzEZdFHVkDhQyGXiOyOwhjmfc.svg)](https://asciinema.org/a/hzEZdFHVkDhQyGXiOyOwhjmfc)
