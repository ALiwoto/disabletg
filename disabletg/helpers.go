// disabletg library Project
// Copyright (C) 2021 ALiwoto <aminnimaj@gmail.com>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package disabletg

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func NewDisabler(dispatcher *ext.Dispatcher, config *DisablerConfig) {
	d := new(Disabler)
	if config == nil {
		config = GetDefaultConfig()
	}
	d.config = config

	d.filter = d.disablerFilter
	d.handler = d.disablerHandler

	h := handlers.NewMessage(d.filter, d.handler)

	d.msgHandler = &h
	d.msgHandler.AllowChannel = config.Channels
	d.msgHandler.AllowEdited = config.Edits

	if config.UseInternal {
		d.setupInternal()
	}

	dispatcher.AddHandlerToGroup(d.msgHandler, d.config.HandlerGroup)
}

func GetDefaultConfig() *DisablerConfig {
	return &DisablerConfig{
		Triggers: []rune{'/', '!'},
	} //TODO: add a default config.
}
