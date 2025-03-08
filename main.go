package main

import (
    "fmt"
    "log"
    "net/http"
    "net"
)

func main() {

    //returning this server ip
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, GetLocalIP())
    })

    //response with "ok"
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "OKs")
    })

    log.Fatal(http.ListenAndServe(":80", nil))

}

func GetLocalIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddress := conn.LocalAddr().(*net.UDPAddr)

    return localAddress.IP
}
