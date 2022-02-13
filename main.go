package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/tach200/go-discord-chat/proto"

	"google.golang.org/grpc"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	ChannelID string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&ChannelID, "c", "942196758115659817", "Channel ID")

	flag.Parse()
}

func main() {
	// Create the Discord Session
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Printf("could not create new discord session: %v", err)
	}

	//Start gRPC server
	// TODO set flag here for port
	lis, err := net.Listen("tcp", "localhost:1122")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDiscordMessageServer()
	grpcServer.Serve(lis)

	// Shut the server down properly
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

// messageChannel sends a message to with the given payload to a specified channel
func messageChannel(s *discordgo.Session, channelID, msgPayload string) error {
	_, err := s.ChannelMessageSend(ChannelID, msgPayload)
	if err != nil {
		return err
	}

	return nil
}
