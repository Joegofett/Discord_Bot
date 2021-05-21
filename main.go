package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	t "github.com/joegofett/discord_bot/tradingView"
	v "github.com/joegofett/discord_bot/emoji_voting"
	"github.com/bwmarrin/discordgo"
)

//These
var day = 5
var hour = 7
var min = 15

// Emoji's need to be unicode so making them all var's here at the top
var number4 = "4️⃣"
var number5 = "5️⃣"
var number6 = "6️⃣"
var number7 = "7️⃣"
var number8 = "8️⃣"
var number9 = "9️⃣"
var number10 = "🔟"

func main() {
	dg, err := discordgo.New("Bot " + "")
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
func EmojiMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		// This is for polls for what day's they want to play
		v.emoji_voting.voting.Emoji(s, m)
			return
		}
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		EmojiMessage(s, m)
		return
	}

	if strings.Contains(m.Content, "$") {

		stonk := strings.SplitAfter(m.Content, "$")
		final := t.tradingView(stonk[1])

		s.ChannelMessageSend(m.ChannelID, final)
	}

	if m.Content == "&Time" {
		s.ChannelMessageSend(m.ChannelID, "Meow Meow! What time is everyone available tommorrow for a happy hour? All times EST")
	}
	if m.Content == "&Day" {

		s.ChannelMessageSend(m.ChannelID, "Meow: What day is everyone available?")
	}

	if day == int(time.Now().Weekday()) {
		if hour == time.Now().Hour() {
			if min == time.Now().Minute() {

				s.ChannelMessageSend(m.ChannelID, "Meow!!! Meow!!!!!! Now that I have everyone's attention it's Friday! It's wins of the week! so let's hear everyone's wins! My win was finding the cat nip bag and getting into when Joe wasn't looking :eyes:")
			}
		}
	}
}
