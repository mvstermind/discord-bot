package handler

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mvstermind/halset/generator"
)

//var token string = "your token"

func New(token string) {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating session,", err)
		return
	}

	dg.AddHandler(MessageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("BOT IS WORKINNNNG!!!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	help := `
	COMMANDS:
		-generate -Generate MIDI with Bpm
		-g - Same as -generate
		-bpm - Generate BPM only
		-b - Same as -bpm
		-scale - Generte scale only
		-s - Same as -scle	
		-sbpm - Generate scale and BPM
		-chords - Generate text chord progression
		-c - Same as -chords
	`

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "-h" || m.Content == "--help" {
		s.ChannelMessageSend(m.ChannelID, help)
	}
	if m.Content == "-bpm" || m.Content == "-b" {
		gned := generator.Gen{}
		bpmstr := strconv.Itoa(gned.GetBpm()) // had to do this cuz bot accepts only string
		wholeMessage := fmt.Sprintf("BPM: %v", bpmstr)
		s.ChannelMessageSend(m.ChannelID, wholeMessage)

	}
	if m.Content == "-sbpm" {
		gned := generator.Gen{}
		bpmstr := strconv.Itoa(gned.GetBpm()) // had to do this cuz bot accepts only string
		scale := gned.GetKey()
		wholeMessage := fmt.Sprintf("BPM: %v\nScale: %v", bpmstr, scale)
		s.ChannelMessageSend(m.ChannelID, wholeMessage)

	}

	if m.Content == "-c" || m.Content == "-chords" {
		gned := generator.Gen{}
		chords := gned.GetChords()
		wholeMessage := fmt.Sprintf("Chords: %v", chords)
		s.ChannelMessageSend(m.ChannelID, wholeMessage)
	}

	if m.Content == "-scale" || m.Content == "-s" {
		gned := generator.Gen{}
		scale := gned.GetKey()
		wholeMessage := fmt.Sprintf("Scale %v", scale)
		s.ChannelMessageSend(m.ChannelID, wholeMessage)
	}

	if m.Content == "-generate" || m.Content == "-g" {
		gned := generator.Gen{}
		bpmstr := strconv.Itoa(gned.GetBpm())
		wholeMessage := fmt.Sprintf("BPM: %v", bpmstr)
		filename, filethingo := Midi()
		s.ChannelFileSendWithMessage(m.ChannelID, wholeMessage, filename, filethingo)
	}
}
