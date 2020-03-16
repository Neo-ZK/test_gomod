package test_gomod

// #cgo darwin CFLAGS: -I${SRCDIR}/include -Wno-deprecated-declarations
// #cgo darwin LDFLAGS: -L${SRCDIR}/lib -lssl -lcrypto
//#include <openssl/ssl.h>
import "C"

import(
    "unsafe"
)

func PrintSSL() {
    var a *C.SSL
    print(unsafe.Sizeof(a))
}
