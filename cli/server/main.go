package main

import (
	"flag"

	"github.com/rs/zerolog/log"

	grpcSetup "github.com/kazakh-in-nz/m-highscore/internal/server/grpc"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "address where you can connect with m-highscore service")
	flag.Parse()

	s := grpcSetup.NewServer(*addressPtr)
	err := s.ListenAndServe()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start grpc server of m-highscore")
	}
}
