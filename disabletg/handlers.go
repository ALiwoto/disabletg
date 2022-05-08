// disabletg library Project
// Copyright (C) 2021-2022 ALiwoto <woto@kaizoku.cyou>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package disabletg

import (
	"strings"

	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (d *Disabler) disablerFilter(msg *gotgbot.Message) bool {
	var cmd string
	if msg.Text != "" {
		cmd = strings.Fields(msg.Text)[0]
	} else if msg.Caption != "" && d.ConsiderCaption() {
		cmd = strings.Fields(msg.Caption)[0]
	}
	if len(cmd) == 0 {
		return false
	}

	pre := ([]rune(cmd))[0]
	for _, current := range d.GetTriggers() {
		if pre == current {
			return true
		}
	}
	return true
}

func (d *Disabler) disablerHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	var msg *gotgbot.Message
	var cmd string

	switch {
	case ctx.EffectiveMessage != nil:
		msg = ctx.Message

	case ctx.EditedMessage != nil && d.ConsiderEdits():
		msg = ctx.EditedMessage

	case ctx.ChannelPost != nil && d.ConsiderChannels():
		msg = ctx.ChannelPost

	case ctx.EditedChannelPost != nil && d.ConsiderChannelsAndEdits():
		msg = ctx.EditedChannelPost
	}

	if msg == nil {
		return ext.ContinueGroups
	}

	var tmpArray []string
	if msg.Text != "" {
		tmpArray = strings.Fields(msg.Text)
	} else if msg.Caption != "" && d.ConsiderCaption() {
		tmpArray = strings.Fields(msg.Caption)
	}

	if len(tmpArray) < 1 {
		return ext.ContinueGroups
	}

	cmd = tmpArray[0]

	if len(cmd) < 1 {
		return ext.ContinueGroups
	}

	tmpArray = ssg.Split(cmd, " ", "@", "/", "-")
	if len(tmpArray) < 1 {
		return ext.ContinueGroups
	}

	cmd = tmpArray[0]

	if d.IsDisabled(msg.Chat.Id, cmd) {
		return ext.EndGroups
	}

	return ext.ContinueGroups
}
