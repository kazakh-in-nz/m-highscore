package main

import (
	"flag"
	"time"

	pbhighscore "github.com/kazakh-in-nz/m_apis/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"golang.org/x/net/context"
)

func main() {
	var addressPtr = flag.String("address", "localhost:60051", "address to connect")

	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("address", *addressPtr).Msg("Failed to close connection")
		}
	}()

	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial e-highscore gRPC service")
	}

	c := pbhighscore.NewGameClient(conn)

	if c == nil {
		log.Info().Msg("Client nil")
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := c.GetHighScore(timeoutCtx, &pbhighscore.GetHighScoreRequest{})

	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response")
	}

	if r != nil {
		log.Info().Interface("highscore", r.GetHighScore()).Msg("Highscore from m-highscore microservice")
	} else {
		log.Error().Msg("Couldn't get highscore")
	}
}
