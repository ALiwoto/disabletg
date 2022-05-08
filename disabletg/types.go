// disabletg library Project
// Copyright (C) 2021-2022 ALiwoto <woto@kaizoku.cyou>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package disabletg

import (
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

type DisablerConfig struct {
	Edits             bool
	Channels          bool
	UseInternal       bool
	Caption           bool
	Triggers          []rune
	HandlerGroup      int
	GlobalIgnoreChats []int64
	Core              DisableCore
}

type Disabler struct {
	config           *DisablerConfig
	internalCommands map[int64]map[string]bool
	internalMutex    *sync.Mutex
	filter           filters.Message

	handler handlers.Response

	// msgHandler is the original message handler of this limiter.
	// it should remain private.
	msgHandler *handlers.Message
}

type DisableCore interface {
	IsDisabled(chatID int64, command string) bool
	IsGlobalDisabled(command string) bool
}
