# ziptest

A Golang demonstration of reading specific entries within a large, remote, `.zip` file without downloading the entire file.

Uses:

* Go
* [snabb/httpreaderat](https://github.com/snabb/httpreaderat), a library providing a HTTP `io.ReaderAt` interface.
* [avvmoto/buf-readerat](https://github.com/avvmoto/buf-readerat), a library that buffers `io.ReaderAt.ReadAt` reads. This reduces the number of HTTP `Partial` calls made.
* a `httptrace` tracer context to know if HTTP persistent connections are being used.

## Usage

    git clone https://github.com/ogri-la/ziptest
    cd ziptest
    go run .

## Related

Similar libraries in this space:

* [krolaw/zipstream](https://github.com/krolaw/zipstream)
* [zhyee/zipstream](https://github.com/zhyee/zipstream)
* [ozkatz/cloudzip](https://github.com/ozkatz/cloudzip)

## Licence

Public domain, no licence required, no attribution necessary, no copyright asserted. Go nuts.
