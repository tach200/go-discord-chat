package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/achmang/go-discord-chat/handlers"
	pb "github.com/achmang/go-discord-chat/proto"
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
		log.Fatalf("could not create new discord session: %v", err)
	}
	log.Println("Discord session created")

	//Start gRPC server
	// TODO set flag here for port
	lis, err := net.Listen("tcp", "localhost:1122")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDiscordMessageServer(grpcServer, &handlers.DiscordBotServer{
		UnimplementedDiscordMessageServer: pb.UnimplementedDiscordMessageServer{},
		Session:                           discord,
		ChannelID:                         ChannelID,
	})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpc sever couldnt sevre %v", err)
	}

	// Shut the server down properly
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
