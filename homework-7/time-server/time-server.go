package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type message []byte

func main() {
	var kbdin string
	//cancel := make(chan string)
	//tick := make(<-chan time.Time)
	//tick = time.Tick(1 * time.Second)
	//При испльзовании select для выхода из программы time-сервер перестаёт быть параллельным.
	//Второе и последующие подключения к time-серверу не работают.
	go func() {
		for {
			fmt.Println("Type 'exit' for exit")
			fmt.Scanln(&kbdin)
			if kbdin == "exit" {
				//	cancel <- kbdin
				os.Exit(0)
			}
		}
	}()

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)

		//	select {
		//	case <-cancel:
		//		return
		//	case <-tick:
		//		continue
		//	}

	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}
