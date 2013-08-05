package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
)

func main(){
  target := flag.String("target", "127.0.0.1:28960", "ip.address:port of the server")
  password := flag.String("password", "", "the RCON password")
  flag.Parse()

  requests := make(chan RCONRequest)
  response := make(chan string)
	rcon := NewRCON(*target, *password, requests)
  go rcon.Relay()

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    requests <- *NewRCONRequest(scanner.Text(), response)
    res := <- response
    fmt.Println(res)
  }

  if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
    os.Exit(1)
	}
  
  os.Exit(0)
}
