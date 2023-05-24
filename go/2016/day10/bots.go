package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/maths"
)

type Bot struct {
	id    int
	value int

	low  *Bot
	high *Bot
}

type StartingBot struct {
	*Bot
	bots  []Bot
	outs  []Bot
	start int
}

func (this *Bot) Receive(value int) {
	if this.value > 0 {
		low, high := maths.MinMax(value, this.value)
		this.low.Receive(low)
		this.high.Receive(high)
		if low == 17 && high == 61 {
			fmt.Println("Part 1 - Special Bot:", this.id)
		}
	} else {
		this.value = value
	}
}

func Parse(lines []string) (starter *StartingBot) {
	bots := make([]Bot, len(lines))
	outs := make([]Bot, len(lines))

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		} else if strings.HasPrefix(line, "value") {
			var value, botID int
			scanf(line, valueGoesToBot, &value, &botID)
			bot := &bots[botID]
			if bot.value > 0 {
				starter = &StartingBot{Bot: bot, start: value, bots: bots, outs: outs}
			} else {
				bot.value = value
				bot.id = botID
			}
		} else if strings.Contains(line, "low to bot") && strings.Contains(line, "high to bot") {
			var giverID, lowID, highID int
			scanf(line, botGivesToBotAndBot, &giverID, &lowID, &highID)
			bot := &bots[giverID]
			bot.id = giverID
			bot.low = &bots[lowID]
			bot.high = &bots[highID]
		} else if strings.Contains(line, "low to bot") && strings.Contains(line, "high to output") {
			var giverID, lowID, highID int
			scanf(line, botGivesToBotAndOut, &giverID, &lowID, &highID)
			bot := &bots[giverID]
			bot.id = giverID
			bot.low = &bots[lowID]
			bot.high = &outs[highID]
		} else if strings.Contains(line, "low to output") && strings.Contains(line, "high to bot") {
			var giverID, lowID, highID int
			scanf(line, botGivesToOutAndBot, &giverID, &lowID, &highID)
			bot := &bots[giverID]
			bot.id = giverID
			bot.low = &outs[lowID]
			bot.high = &bots[highID]
		} else if strings.Contains(line, "low to output") && strings.Contains(line, "high to output") {
			var giverID, lowID, highID int
			scanf(line, botGivesToOutAndOut, &giverID, &lowID, &highID)
			bot := &bots[giverID]
			bot.id = giverID
			bot.low = &outs[lowID]
			bot.high = &outs[highID]
		} else {
			panic("Invalid instruction: " + line)
		}
	}

	if starter == nil {
		panic("NIL")
	}
	return starter
}

func scanf(input, format string, args ...interface{}) {
	scanned, err := fmt.Sscanf(input, format, args...)
	if err != nil || scanned != len(args) {
		panic("Didn't work")
	}
}

const (
	valueGoesToBot      = "value %d goes to bot %d"
	botGivesToBotAndBot = "bot %d gives low to bot %d and high to bot %d"
	botGivesToBotAndOut = "bot %d gives low to bot %d and high to output %d"
	botGivesToOutAndBot = "bot %d gives low to output %d and high to bot %d"
	botGivesToOutAndOut = "bot %d gives low to output %d and high to output %d"
)
