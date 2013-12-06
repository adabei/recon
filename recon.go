package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/adabei/goldenbot/rcon"
	_ "github.com/adabei/goldenbot/rcon/q3"
	"github.com/howeyc/gopass"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var hosts map[string]Host

var (
	// flags
	protocol = flag.String("pr", "q3", "protocol name (depends on game, see documentation)")
)

type Host struct {
	Type     string
	Addr     string
	Password string
}

func main() {
	flag.Parse()

	addr := "127.0.0.1:28960"
	password := ""

	// ip:port pair
	if strings.Contains(flag.Arg(0), ":") {
		addr = flag.Arg(0)
		fmt.Print("Password: ")
		password = string(gopass.GetPasswd())
	} else if matches, _ := regexp.MatchString("[:alnum:]", flag.Arg(0)); matches {
		load(homePath() + "/" + ".recon.cfg")
		host := hosts[flag.Arg(0)]
		addr = host.Addr
		password = host.Password
		*protocol = host.Type
	}

	queries := make(chan rcon.RCONQuery)
	ez := rcon.EasyQuery(queries)
  go rcon.Relay(*protocol, addr, password, queries)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for scanner.Scan() {
		res := ez(scanner.Text())
		fmt.Println(string(res))
		fmt.Print(">")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from standard input: ", err)
		os.Exit(1)
	}
}

func load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &hosts)
	if err != nil {
		log.Fatal(err)
	}
}

// Naive retrival of path for home
func homePath() string {
  p := os.Getenv("HOME") // unix
  if p == "" {
    p = strings.Replace(os.Getenv("HOMEPATH"), "\\", "/", -1) // windows
  }
  return p
}
