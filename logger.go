/*
 * Copyright 2021 Jesko Schnepel
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package logger

import (
	"fmt"
	"time"
)

const FORMAT = "%s%s %s [%s] %s" + NC + "\n"
const TIME_FORMAT = "2006-01-02 15:04:05"

type Channel struct {
	name      string
	debugging bool
}

func NewChannel(name string) *Channel {
	return &Channel{
		name:      name,
		debugging: false,
	}
}

func (ch *Channel) EnableDebugging() {
	ch.debugging = true
}

func (ch *Channel) DisableDebugging() {
	ch.debugging = false
}

func (ch *Channel) print(level, color string, msg interface{}) {
	fmt.Printf(FORMAT, color, level, time.Now().Format(TIME_FORMAT), ch.name, msg)
}

func (ch *Channel) Info(msg interface{}) {
	ch.print("INFO", LIGHT_GREY, msg)
}

func (ch *Channel) InfoF(format string, a ...interface{}) {
	ch.Info(fmt.Sprintf(format, a...))
}

func (ch *Channel) Warn(msg interface{}) {
	ch.print("WARN", LIGHT_RED, msg)
}

func (ch *Channel) WarnF(format string, a ...interface{}) {
	ch.Warn(fmt.Sprintf(format, a...))
}

func (ch *Channel) Debug(msg interface{}) {
	if ch.debugging {
		ch.print("DEBG", CYAN, msg)
	}
}

func (ch *Channel) Fatal(msg interface{}) {
	ch.print("FATL", RED, msg)
}

func (ch *Channel) IfError(err error) bool {
	if err != nil {
		ch.print("FATL", RED, err)
		return true
	}
	return false
}