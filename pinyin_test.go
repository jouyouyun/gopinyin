/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

import (
	"fmt"
	"testing"
)

func TestRangeArray(t *testing.T) {
	a1 := []string{"12", "34"}
	a2 := []string{"ab", "cd"}
	rets := RangeArray(a1, a2)
	if len(rets) == 0 {
		t.Error("Range Array Failed!")
		return
	}
	fmt.Println(rets)
}

func TestPinyin(t *testing.T) {
	m := &Manager{}
	rets := m.PinyinFromKey("重")
	if len(rets) == 0 {
		t.Error("Get Pinyin Failed!")
		return
	}
	fmt.Println(rets)

	rets = m.PinyinFromKey("你好")
	if len(rets) == 0 {
		t.Error("Get Pinyin Failed!")
		return
	}
	fmt.Println(rets)

	rets = m.PinyinFromKey("藏重")
	if len(rets) == 0 {
		t.Error("Get Pinyin Failed!")
		return
	}
	fmt.Println(rets)
}
