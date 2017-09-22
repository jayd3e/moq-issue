# Moq Issue

I created a repo around the problem I'm seeing with Moq.  Check out the following output:

```
vagrant /opt/code/Rain/src/github.com/project/moq_test $ go build -o moq_test .
vagrant /opt/code/Rain/src/github.com/project/moq_test $ go generate
example.go:5:2: could not import github.com/project/moq_test/subpackage (can't find import: github.com/project/moq_test/subpackage)
moq [flags] destination interface [interface2 [interface3 [...]]]
  -out string
    	output file (default stdout)
  -pkg string
    	package name (default will infer)
example.go:8: running "moq": exit status 1
```

Here is my environment:

```
vagrant /opt/code/Rain/src/github.com/project/moq_test $ go env
GOARCH="386"
GOBIN=""
GOEXE=""
GOHOSTARCH="386"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/opt/code/Rain"
GORACE=""
GOROOT="/usr/local/go"
GOTOOLDIR="/usr/local/go/pkg/tool/linux_386"
GCCGO="gccgo"
GO386=""
CC="gcc"
GOGCCFLAGS="-fPIC -m32 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build086048455=/tmp/go-build"
CXX="g++"
CGO_ENABLED="1"
PKG_CONFIG="pkg-config"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
```

The problem occurs when I try adding a subpackage to your example found [here](https://github.com/matryer/moq/blob/master/example/example.go).
As you can see, this repo matches your example almost exactly, however introduces a single subpackage using the standard import scheme.
The code compiles via `go build`, but `go generate` does not work, as `moq` errors out.
