package main

import (
	"crypto/sha256"
	"flag"
	"net"
	"tcp-pow/internal/calculators"
	"tcp-pow/internal/handlers"
	"tcp-pow/internal/infrastructure/cache"
	"tcp-pow/internal/mappers"
	"tcp-pow/internal/servers"
	"tcp-pow/internal/validators"
)

const zenquotesUrl = "https://zenquotes.io/api/random"

func main() {
	var url string
	var difficulty int

	flag.StringVar(&url, "url", ":8080", "url on which to serve server")
	flag.IntVar(&difficulty, "difficulty", 5, "how many zeros do we want hash to start with")
	flag.Parse()

	connectionMapper := mappers.NewConnectionMapper()

	errorMapper := mappers.NewErrorMapper()

	cache := cache.NewCache()
	nonceValidator := validators.NewNonceValidator(cache)

	nonceLength := 10
	nonceCalculator := calculators.NewNonceCalculator(nonceLength)
	challengeRequestedHandler := handlers.NewChallengeRequestedHandler(nonceCalculator, difficulty)

	quoteHandler := handlers.NewQuoteHandler(zenquotesUrl)

	solutionSentHandler := handlers.NewSolutionSentHandler(sha256.New(), difficulty, quoteHandler)

	connectionHandler := handlers.NewConnectionHandler(connectionMapper, errorMapper, nonceValidator, challengeRequestedHandler, solutionSentHandler)

	server, err := servers.NewServer(url)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := server.GetConnection()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			for {
				response, done, err := connectionHandler.Handle(conn)
				if err != nil {
					panic(err)
				}

				if done {
					return
				}

				response = append(response, '\n')
				_, err = conn.Write(response)
				if err != nil {
					panic(err)
				}
			}
		}(conn)
	}
}
