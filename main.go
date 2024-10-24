package main

import "fmt"

func main() {
    fmt.Println("Hello, world.")
}

shadowsocks2-linux -c 'ss://AEAD_CHACHA20_POLY1305:eC2xO=9I9oHf!H(HDB2@>A3@(VH5@83.229.86.18:8488' -verbose -socks :1080 -u