package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// 1、接收客户端 request，并将 request 中带的 header 写入 response header
	head := r.Header
	for k, vals := range head {
		for _, val := range vals {
			w.Header().Add(k, val)
		}
	}

	// 2、读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	w.Header().Add("VERSION", version)
	//fmt.Printf("%#+v\n", w.Header())

	// 3、Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	clientIP := r.RemoteAddr
	httpStatusOk := http.StatusOK
	fmt.Printf("client ip: %s, http code: %d", clientIP, httpStatusOk)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// 4、当访问 localhost/healthz 时，应返回 200
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok")
}
