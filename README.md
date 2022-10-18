# music-control-go

Global hotkey daemon for macOS Music.app

## Build

```bash
go build -ldflags="-s -w"
```

## Usage

### Run

```bash
music-control-go
```

For the first operation, a prompt will ask you for permission, just press `OK`.

### Shortcuts

- ⌃ ⌥ ⌘ P: play / pause
- ⌃ ⌥ ⌘ ↑: volume up
- ⌃ ⌥ ⌘ ↓: volume down
- ⌃ ⌥ ⌘ →: next track
- ⌃ ⌥ ⌘ ←: previous track

> **Note**
> Don't like such shortcuts?
>
> Fork and edit main.go, or even better,
> making a pull request about customizing shortcuts config. :)

### Run as a daemon

Create a `me.colerar.music-control.plist` and
put it into `~/Library/LaunchAgents`:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>me.colerar.music-control</string>
    <key>ProgramArguments</key>
    <array>
        <string>/path/to/music-control-go</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>
```

You can reboot your machine or run:

```bash
launchctl load -w ~/Library/LaunchAgents/me.colerar.music-control.plist
```

to make it work immediately.
