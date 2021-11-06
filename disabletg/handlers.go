// disabletg library Project
// Copyright (C) 2021 ALiwoto <aminnimaj@gmail.com>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package disabletg

import (
	"strings"

	"github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (d *Disabler) disablerFilter(msg *gotgbot.Message) bool {
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

	if msg.Text != "" {
		cmd = strings.Fields(msg.Text)[0]
	} else if msg.Caption != "" {
		cmd = strings.Fields(msg.Caption)[0]
	}

	cmd = strongStringGo.Split(cmd, " ", "@", "/")[0]

	if d.IsDisabled(msg.Chat.Id, cmd) {
		return ext.EndGroups
	}

	return ext.ContinueGroups
}
