// disabletg library Project
// Copyright (C) 2021-2022 ALiwoto <aminnimaj@gmail.com>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package disabletg

import "sync"

func (d *Disabler) setupInternal() {
	if d.internalCommands == nil {
		d.internalCommands = make(map[int64]map[string]bool)
		d.internalMutex = new(sync.Mutex)
	}
}

func (d *Disabler) ConsiderCaption() bool {
	if d.config != nil {
		return d.config.Caption
	}
	return false
}

func (d *Disabler) ConsiderChannels() bool {
	if d.config != nil {
		return d.config.Channels
	}
	return false
}

func (d *Disabler) ConsiderChannelsAndEdits() bool {
	return d.ConsiderChannels() && d.ConsiderEdits()
}

func (d *Disabler) ConsiderEdits() bool {
	if d.config != nil {
		return d.config.Edits
	}
	return false
}

func (d *Disabler) IsUsingInternals() bool {
	if d.config != nil {
		return d.config.UseInternal
	}
	return false
}

func (d *Disabler) IsInternalDisabled(chatId int64, command string) bool {
	if d.config == nil || !d.config.UseInternal || d.internalMutex == nil {
		return false
	}

	d.internalMutex.Lock()
	if len(d.internalCommands) == 0 {
		d.internalMutex.Unlock()
		return false
	}

	chat := d.internalCommands[chatId]
	if len(chat) == 0 {
		d.internalMutex.Unlock()
		return false
	}

	disabled := chat[command]
	d.internalMutex.Unlock()
	return disabled
}

func (d *Disabler) IsDisabled(chatId int64, command string) bool {
	if d.config.Core != nil {
		if d.config.Core.IsDisabled(chatId, command) {
			return true
		}
		if d.IsGlobalDisabled(command) && !d.IsGlobalIgnored(chatId) {
			return true
		}
	}

	// now check for internal disabled commands
	return d.IsInternalDisabled(chatId, command)
}

func (d *Disabler) GetTriggers() []rune {
	if d.config != nil {
		return d.config.Triggers
	}
	return []rune{'/', '!'}
}

func (d *Disabler) IsGlobalDisabled(command string) bool {
	if d.config.Core != nil {
		return d.config.Core.IsGlobalDisabled(command)
	}
	return false
}

func (d *Disabler) IsGlobalIgnored(chatId int64) bool {
	if d.config.Core != nil && len(d.config.GlobalIgnoreChats) > 0 {
		for _, id := range d.config.GlobalIgnoreChats {
			if id == chatId {
				return true
			}
		}
	}
	return false
}

func (d *Disabler) GetGlobalIgnoredChats() []int64 {
	if d.config.Core != nil {
		return d.config.GlobalIgnoreChats
	}
	return nil
}
