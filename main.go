package main

import (
	"fmt"
	"os"

	"github.com/zalando/go-keyring"
)

func fatal(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(-1)
}

func fatalf(format string, v ...interface{}) {
	fatal(fmt.Sprintf(format, v...))
}

func main() {
	if len(os.Args) < 2 {
		fatalf("USAGE: %s (get|set)", os.Args[0])
	}
	switch os.Args[1] {
	case "get":
		if len(os.Args) < 4 {
			fatalf("USAGE: %s get <service> <username>", os.Args[0])
		}
		v, err := keyring.Get(os.Args[2], os.Args[3])
		if err != nil {
			fatal(err)
		}
		fmt.Println(v)
	case "set":
		if len(os.Args) < 5 {
			fatalf("USAGE: %s set <service> <username> <password>", os.Args[0])
		}
		if err := keyring.Set(os.Args[2], os.Args[3], os.Args[4]); err != nil {
			fatal(err)
		}
	}
}
