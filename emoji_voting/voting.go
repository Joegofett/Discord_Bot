package voting

import (
	"github.com/bwmarrin/discordgo"

)

// Emoji's need to be unicode so making them all var's here at the top
var letterM = "🇲"
var letterT = "🇹"
var letterW = "🇼"
var letterH = "🇭"
var letterF = "🇫"
var letterS = "🇸"
var letterU = "⛪"
var number6 = "6️⃣"
var number7 = "7️⃣"
var number8 = "8️⃣"
var number9 = "9️⃣"
var number10 = "🔟"



func Emoji(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "Meow: What day is everyone available?" {
		s.MessageReactionAdd(m.ChannelID, m.ID, letterM)
		s.MessageReactionAdd(m.ChannelID, m.ID, letterT)
		s.MessageReactionAdd(m.ChannelID, m.ID, letterW)
		s.MessageReactionAdd(m.ChannelID, m.ID, letterH)
		s.MessageReactionAdd(m.ChannelID, m.ID, letterF)
		s.MessageReactionAdd(m.ChannelID, m.ID, letterS)
		s.MessageReactionAdd(m.ChannelID, m.ID, letterU)
	}

	//This is for polls for what time everyone is available all time EST
	if m.Content == "Meow Meow! What time is everyone available today? All times EST" {
		s.MessageReactionAdd(m.ChannelID, m.ID, number6)
		s.MessageReactionAdd(m.ChannelID, m.ID, number7)
		s.MessageReactionAdd(m.ChannelID, m.ID, number8)
		s.MessageReactionAdd(m.ChannelID, m.ID, number9)
		s.MessageReactionAdd(m.ChannelID, m.ID, number10)
	}

}
}