
/*
Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)

*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	isUDP := flag.Bool("u", false, "Use UDP instead of the default option of TCP.")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Hostname and port required")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	connType := "tcp"

	if *isUDP {
		connType = "upd"
	}

	conn, err := net.Dial(connType, net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal(err)
	}

	osSignals := make(chan os.Signal, 1)
	listenErr := make(chan error, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	go write(conn, listenErr, osSignals)
	go read(conn, listenErr, osSignals)

	select {
	case <-osSignals:
		conn.Close()
	case err = <-listenErr:
		if err != nil {
			log.Fatal(err)
		}
	}

}

func write(conn net.Conn, listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				osSignals <- syscall.SIGTERM
			}
			listenErr <- err
		}

		fmt.Fprintf(conn, text+"\n")
	}
}

func read(conn net.Conn, listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(conn)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				osSignals <- syscall.SIGTERM
			}
			listenErr <- err
		}

		fmt.Print(text)
	}
}
