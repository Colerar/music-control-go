package main

import (
	"log"
	"os"
	"os/exec"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

// Take from: /Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/System/Library/Frameworks/Carbon.framework/Versions/A/Frameworks/HIToolbox.framework/Versions/A/Headers/Events.h
const (
	LeftArrow  hotkey.Key = 0x7B
	RightArrow hotkey.Key = 0x7C
	DownArrow  hotkey.Key = 0x7D
	UpArrow    hotkey.Key = 0x7E
)

const (
	playpause = "Application('Music').playpause();"
	volumeUp  = `var app = Application('Music');
var volume = app.soundVolume();
app.soundVolume = volume + 10 < 100 ? volume + 10 : 100;
`
	volumeDown = `var app = Application('Music');
var volume = app.soundVolume();
app.soundVolume = volume - 10 > 0 ? volume - 10 : 0;`
	nextTrack     = "Application('Music').nextTrack();"
	previousTrack = "Application('Music').previousTrack();"
)

func main() { mainthread.Init(fn) }

func fn() {
	mods := []hotkey.Modifier{hotkey.ModCtrl, hotkey.ModCmd, hotkey.ModOption}
	go func() {
		listenHotkey(hotkey.KeyP, mods, func() {
			println("playpause()")
			executeJax(playpause)
		})
	}()
	go func() {
		listenHotkey(UpArrow, mods, func() {
			println("volumeUp()")
			executeJax(volumeUp)
		})
	}()
	go func() {
		listenHotkey(DownArrow, mods, func() {
			println("volumeDown()")
			executeJax(volumeDown)
		})
	}()
	go func() {
		listenHotkey(RightArrow, mods, func() {
			println("nextTrack()")
			executeJax(nextTrack)
		})
	}()
	go func() {
		listenHotkey(LeftArrow, mods, func() {
			println("previousTrack()")
			executeJax(previousTrack)
		})
	}()

	select {} // block forever
}

func listenHotkey(key hotkey.Key, modifiers []hotkey.Modifier, callback func()) {
	hk := hotkey.New(modifiers, key)
	err := hk.Register()
	if err != nil {
		log.Println(err)
	}
	for {
		<-hk.Keydown()
		callback()
	}
}

func executeJax(code string) {
	cmd := exec.Command("osascript", "-l", "JavaScript", "-e", code)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n\n", err)
	}
}
