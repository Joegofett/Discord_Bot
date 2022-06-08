package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	v "github.com/joegofett/discord_bot/emoji_voting"
	t "github.com/joegofett/discord_bot/tradingView"
)

//These

func main() {
	dg, err := discordgo.New("Bot " + "ODQxMzcxMDU1MTQyMDc2NDQ2.YJlx2w.LDjnOLwsjZMfdcIeXEGYwBGKids")
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
		return
	}
	// Register the messageCreate func asa callback for messageCreate events.
	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("bot is now running. press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

var day = 5
var hour = 9
var min = 5

// func fridayWins(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	if day == int(time.Now().Weekday()) {
// 		if hour == time.Now().Hour() {
// 			if min == time.Now().Minute() {
// 				room := "847835494343770113"
// 				s.ChannelMessageSend(room, "testing")
// 			}
// 		}
// 	}
// }

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		v.Emoji(s, m)
		return
	}

	if strings.Contains(m.Content, "%") {

		t.Crypto(s, m)
	}

	if strings.Contains(m.Content, "$") {
		valid := strings.SplitAfter(m.Content, "$")
		if len(valid[1]) > 6 {
			s.ChannelMessageSend(m.ChannelID, "Meow? hmmm this doesn't seem right")
			return
		}
		t.Message(s, m)
	}
	if m.Content == "&Among Us" {
		s.ChannelMessageSend(m.ChannelID, "Meow Meow! What time is everyone available for Among Us? All times EST. Whomever Kills Joe I'mma fite you @Killers (Among us) ")
	}

	if m.Content == "&Time" {
		s.ChannelMessageSend(m.ChannelID, "Meow Meow! What time is everyone available? All times EST")
	}
	if m.Content == "&Day" {

		s.ChannelMessageSend(m.ChannelID, "Meow: What day is everyone available?")
	}
	// if day == int(time.Now().Weekday()) {
	// 	if hour == time.Now().Hour() {
	// 		if min == time.Now().Minute() {
	if m.Content == "&Friday" {

		s.ChannelMessageSend(m.ChannelID, "Meow!!! Meow!!!!!! Now that I have everyone's attention it's Friday! It's wins of the week! so let's hear everyone's wins! My win was jumping on the high beam in the house :smiley_cat: :heart_eyes_cat: ")
		fmt.Println(m.ChannelID)
	}
	// }
	// }
}
