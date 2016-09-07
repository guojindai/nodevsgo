package main 

import (
  "net/http"
  "strconv"
  "io/ioutil"
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
  data, err := ioutil.ReadFile("../data/resgo")
  if err != nil {
    w.Write([]byte(err.Error()))
  } else {
    w.Write(data)
  }
}