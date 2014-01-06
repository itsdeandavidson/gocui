package gocui

import (
	"github.com/nsf/termbox-go"
)

type (
	Key          termbox.Key
	Modifier     termbox.Modifier
	KeybindingCB func(*Gui, *View) error
)

const (
	KeyF1         Key = Key(termbox.KeyF1)
	KeyF2             = Key(termbox.KeyF2)
	KeyF3             = Key(termbox.KeyF3)
	KeyF4             = Key(termbox.KeyF4)
	KeyF5             = Key(termbox.KeyF5)
	KeyF6             = Key(termbox.KeyF6)
	KeyF7             = Key(termbox.KeyF7)
	KeyF8             = Key(termbox.KeyF8)
	KeyF9             = Key(termbox.KeyF9)
	KeyF10            = Key(termbox.KeyF10)
	KeyF11            = Key(termbox.KeyF11)
	KeyF12            = Key(termbox.KeyF12)
	KeyInsert         = Key(termbox.KeyInsert)
	KeyDelete         = Key(termbox.KeyDelete)
	KeyHome           = Key(termbox.KeyHome)
	KeyEnd            = Key(termbox.KeyEnd)
	KeyPgup           = Key(termbox.KeyPgup)
	KeyPgdn           = Key(termbox.KeyPgdn)
	KeyArrowUp        = Key(termbox.KeyArrowUp)
	KeyArrowDown      = Key(termbox.KeyArrowDown)
	KeyArrowLeft      = Key(termbox.KeyArrowLeft)
	KeyArrowRight     = Key(termbox.KeyArrowRight)
)

const (
	KeyCtrlTilde      Key = Key(termbox.KeyCtrlTilde)
	KeyCtrl2              = Key(termbox.KeyCtrl2)
	KeyCtrlSpace          = Key(termbox.KeyCtrlSpace)
	KeyCtrlA              = Key(termbox.KeyCtrlA)
	KeyCtrlB              = Key(termbox.KeyCtrlB)
	KeyCtrlC              = Key(termbox.KeyCtrlC)
	KeyCtrlD              = Key(termbox.KeyCtrlD)
	KeyCtrlE              = Key(termbox.KeyCtrlE)
	KeyCtrlF              = Key(termbox.KeyCtrlF)
	KeyCtrlG              = Key(termbox.KeyCtrlG)
	KeyBackspace          = Key(termbox.KeyBackspace)
	KeyCtrlH              = Key(termbox.KeyCtrlH)
	KeyTab                = Key(termbox.KeyTab)
	KeyCtrlI              = Key(termbox.KeyCtrlI)
	KeyCtrlJ              = Key(termbox.KeyCtrlJ)
	KeyCtrlK              = Key(termbox.KeyCtrlK)
	KeyCtrlL              = Key(termbox.KeyCtrlL)
	KeyEnter              = Key(termbox.KeyEnter)
	KeyCtrlM              = Key(termbox.KeyCtrlM)
	KeyCtrlN              = Key(termbox.KeyCtrlN)
	KeyCtrlO              = Key(termbox.KeyCtrlO)
	KeyCtrlP              = Key(termbox.KeyCtrlP)
	KeyCtrlQ              = Key(termbox.KeyCtrlQ)
	KeyCtrlR              = Key(termbox.KeyCtrlR)
	KeyCtrlS              = Key(termbox.KeyCtrlS)
	KeyCtrlT              = Key(termbox.KeyCtrlT)
	KeyCtrlU              = Key(termbox.KeyCtrlU)
	KeyCtrlV              = Key(termbox.KeyCtrlV)
	KeyCtrlW              = Key(termbox.KeyCtrlW)
	KeyCtrlX              = Key(termbox.KeyCtrlX)
	KeyCtrlY              = Key(termbox.KeyCtrlY)
	KeyCtrlZ              = Key(termbox.KeyCtrlZ)
	KeyEsc                = Key(termbox.KeyEsc)
	KeyCtrlLsqBracket     = Key(termbox.KeyCtrlLsqBracket)
	KeyCtrl3              = Key(termbox.KeyCtrl3)
	KeyCtrl4              = Key(termbox.KeyCtrl4)
	KeyCtrlBackslash      = Key(termbox.KeyCtrlBackslash)
	KeyCtrl5              = Key(termbox.KeyCtrl5)
	KeyCtrlRsqBracket     = Key(termbox.KeyCtrlRsqBracket)
	KeyCtrl6              = Key(termbox.KeyCtrl6)
	KeyCtrl7              = Key(termbox.KeyCtrl7)
	KeyCtrlSlash          = Key(termbox.KeyCtrlSlash)
	KeyCtrlUnderscore     = Key(termbox.KeyCtrlUnderscore)
	KeySpace              = Key(termbox.KeySpace)
	KeyBackspace2         = Key(termbox.KeyBackspace2)
	KeyCtrl8              = Key(termbox.KeyCtrl8)
)

const (
	ModAlt Modifier = Modifier(termbox.ModAlt)
)

type Keybinding struct {
	ViewName string
	Key      Key
	Ch       rune
	Mod      Modifier
	CB       KeybindingCB
}

func NewKeybinding(viewname string, key Key, ch rune, mod Modifier, cb KeybindingCB) (kb *Keybinding) {
	kb = &Keybinding{
		ViewName: viewname,
		Key:      key,
		Ch:       ch,
		Mod:      mod,
		CB:       cb,
	}
	return kb
}