package main 

import (
  "net/http"
  "strconv"
  "log"
)

const (
  SERVER_PORT = 35033
)

func main () {
  http.HandleFunc("/", handle)
  log.Println("Start server with listenting on port", SERVER_PORT)
  http.ListenAndServe(":" + strconv.Itoa(SERVER_PORT), nil)  
}

func handle (w http.ResponseWriter, r *http.Request) {
  sum := 0
  for i := 0; i < 10000; i ++ {
    sum += i
  }
  w.Write([]byte(strconv.Itoa(sum)))
}