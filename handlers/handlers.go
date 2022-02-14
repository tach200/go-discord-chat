package handlers

import (
	"context"
	"fmt"

	pb "github.com/achmang/go-discord-chat/proto"
	"github.com/bwmarrin/discordgo"
)

type DiscordBotServer struct {
	pb.UnimplementedDiscordMessageServer
	*discordgo.Session
	ChannelID string
}

func (s *DiscordBotServer) SendChanMessage(ctx context.Context, payload *pb.MessageChannel) (*pb.ServerResponse, error) {

	msg := fmt.Sprintf("Subject : %v \n %v", payload.Subject, payload.Content)

	err := messageChannel(s.Session, s.ChannelID, msg)
	if err != nil {
		return &pb.ServerResponse{
			Message: "Send FAILED",
		}, err
	}

	return &pb.ServerResponse{
		Message: "Send SUCCESS",
	}, nil
}

// messageChannel sends a message to with the given payload to a specified channel
func messageChannel(s *discordgo.Session, channelID, msgPayload string) error {
	_, err := s.ChannelMessageSend(channelID, msgPayload)
	if err != nil {
		return err
	}

	return nil
}
