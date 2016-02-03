# gotrees

I wanted to remind myself of how red/black trees work so I decided to create an implementation in GO. The implementation is still a work in progress and currently only supports insertions and extraction into a sorted slice. I will add deletion once I am able to remember how to correctly handle all double black cases.

The tests are meant to be run by the Ginkgo BDD testing framework. Before you can run them, you need to retrieve the Ginkgo and Gomega libraries with go get

```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega
```

Once Ginkgo and Gomega are available on your system, just navigate to the root folder and run ginkgo

```
$ ginkgo
```
