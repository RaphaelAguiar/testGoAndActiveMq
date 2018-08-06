package main

import (
	"fmt"

	"github.com/go-stomp/stomp"
)

func main() {
	var fila = "queue/test"

	conn, errCon := stomp.Dial("tcp", "localhost:61616")

	if errCon != nil {
		fmt.Println(errCon)
		return
	} else {
		fmt.Println("Conexão realizada com sucesso!")
	}

	var quit = make(chan bool)

	sub, errSub := conn.Subscribe(fila, stomp.AckClient)
	if errSub != nil {
		fmt.Println(errCon)
		return
	}

	go func() {
		msg := <-sub.C
		fmt.Println("Mensagem recebida: " + string(msg.Body))
		quit <- true
		return
	}()

	errSend := conn.Send(
		fila,
		"text/plain",
		[]byte("### Esta aqui é a mensagem! ###"))

	if errSend != nil {
		fmt.Println(errSend)
	} else {
		fmt.Println("Mensagem enviada com sucesso!")
	}

	for {
		select {
		case <-quit:
			sub.Unsubscribe()
			return
		}
	}
}
