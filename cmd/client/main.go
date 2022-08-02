package main

import (
	"crypto/sha256"
	"flag"
	"net"
	"tcp-pow/internal/calculators"
	"tcp-pow/internal/handlers"
	"tcp-pow/internal/mappers"
)

func main() {
	connectionMapper := mappers.NewConnectionMapper()
	solutionCalculator := calculators.NewSolutionCalculator(sha256.New())
	challengeSentHandler := handlers.NewChallengeSentHandler(solutionCalculator)
	quoteSentHandler := handlers.NewQuoteSentHandler()
	errorMapper := mappers.NewErrorMapper()
	clientHandler := handlers.NewClientHandler(connectionMapper, challengeSentHandler, quoteSentHandler, challengeSentHandler, errorMapper)

	var url string
	flag.StringVar(&url, "url", "localhost:8080", "url on which server is located. With port!")
	flag.Parse()

	conn, err := net.Dial("tcp", url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = clientHandler.SendInitialRequest(conn)
	if err != nil {
		panic(err)
	}

	for {
		request, done, err := clientHandler.Handle(conn)
		if err != nil {
			panic(err)
		}

		if done {
			return
		}

		request = append(request, '\n')
		_, err = conn.Write(request)
		if err != nil {
			panic(err)
		}
	}

}
