package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/achmang/go-discord-chat/handlers"
	pb "github.com/achmang/go-discord-chat/proto"
	"github.com/bwmarrin/discordgo"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
)

var (
	Token     string
	ChannelID string
	GrpcPort  string
	RestPort  string
)

// restServer starts the http rest server to listen for incoming messages,
// pass in a discrod session so that incoming message can be sent to discord channel.
func restServer(discord *discordgo.Session, port string) {
	mux := runtime.NewServeMux()

	pb.RegisterDiscordMessageHandlerServer(context.Background(), mux, &handlers.DiscordBotServer{
		UnimplementedDiscordMessageServer: pb.UnimplementedDiscordMessageServer{},
		Session:                           discord,
		ChannelID:                         ChannelID,
	})

	addr := fmt.Sprintf("localhost:%s", port)
	http.ListenAndServe(addr, mux)
}

// grpcServer starts the grpc server to listen for incoming messages,
// pass in a discrod session so that incoming message can be sent to discord channel.
func grpcServer(discord *discordgo.Session, port string) {
	addr := fmt.Sprintf("localhost:%s", port)

	lis, err := net.Listen("tcp", addr)
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
}

func init() {
	viper.SetConfigName("env")
	viper.AddConfigPath("/configs/")
	viper.AddConfigPath("/configs/secrets")
	viper.AddConfigPath("$HOME/.configs")
	viper.AddConfigPath(".")
	viper.SetConfigFile("yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error: couldn't read config file: %s", err)
	}

	flag.StringVar(&Token, "t", viper.GetString("token"), "Bot Token")
	flag.StringVar(&ChannelID, "c", viper.GetString("channelID"), "Channel ID")
	flag.StringVar(&GrpcPort, "gp", viper.GetString("grpcPort"), "gRPC Port")
	flag.StringVar(&RestPort, "rp", viper.GetString("restPort"), "REST Port")

	flag.Parse()
}

func main() {
	// Create the Discord Session
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("could not create new discord session: %v", err)
	}
	log.Println("Discord session created")

	go restServer(discord, RestPort)
	go grpcServer(discord, GrpcPort)

	// Shut the server down properly
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
