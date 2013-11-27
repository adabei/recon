package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
  "strings"
  "github.com/howeyc/gopass"
  "github.com/adabei/goldenbot/rcon"
  "github.com/adabei/goldenbot/rcon/q3"
  "regexp"
  "log"
)

func main(){
  port := flag.Int("p", 28960, "the port to connect to")
  configPath := flag.String("f", "recon.cfg", "the config file to read connections from")
  flag.Parse()
  address := flag.Args()[0]
  password := ""

  matches, _ := regexp.MatchString("[:alnum:]", os.Args[1])
  if len(os.Args) == 2 && matches {
    fi, err := os.Open(*configPath)

    if err != nil {
      log.Fatal("Couldn't open config file: ", err)
    }
    sc := bufio.NewScanner(fi)

    for sc.Scan() {
      line := sc.Text()
      values := strings.Split(line, ";")
      if values[0] == os.Args[1] {
        address = values[1]
        password = values[2]
        fi.Close()
        break
      }
    }
  }

  // naive check
  if !strings.Contains(address, ":") {
    address = address + string(*port)
  }
  
  if password == "" {
    fmt.Print("password: ")
    password = string(gopass.GetPasswd())
  }


  requests := make(chan RCONRequest)
  response := make(chan string)
	r := q3.NewRCON(address, password, requests)
  go r.Relay()
  query := rcon.EasyQuery(requests)
 
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print(">")
  for scanner.Scan() {
    res := query(scanner.Text())
    fmt.Println(res)
    fmt.Print(">")
  }

  if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
    os.Exit(1)
	}
  
  os.Exit(0)
}
