# CD-Man

A desktop maze game for **Windows, Linux, and macOS**, written in Go.
Collect dots, avoid enemies, and explore five worlds. The game includes
single-player and two-player modes, a built-in demo, high scores, and in-game help.
Starting a game goes directly to the level without an identification prompt.

## Features

- Five worlds with distinct graphics and obstacles.
- Pixel-art rendering with nearest-neighbor scaling and a 4:3 display area.
- Double-buffered Windows rendering to prevent black flashes between frames.
- Fullscreen and windowed modes.
- Keyboard input, configurable player controls, and joystick support.
- Sound effects and persistent high scores.
- Embedded game assets: no separate resource installation is required.

## Requirements

| Platform | Build requirements | Runtime requirements |
| --- | --- | --- |
| Windows x64 / ARM64 | Go 1.22 or newer | System Windows libraries only |
| Linux | Go 1.22+, a C compiler | SDL2 |
| macOS | Go 1.22+, Xcode Command Line Tools | SDL2 |

Build each platform on that platform. Scripts target the architecture selected
by the installed Go toolchain. Windows uses native Win32/GDI/WinMM APIs and does
not require a C compiler. Linux and macOS use a small cgo adapter that loads SDL2
at runtime; SDL development headers are not needed. SDL3 is not a substitute.

On macOS, install the Xcode Command Line Tools with `xcode-select --install` if
needed. Install Go and SDL2 using your preferred package manager or installer.
The adapter supports standard Homebrew library paths and SDL2.framework.
On Linux, install Go, a C compiler, and the SDL2 runtime using your distribution's
package manager.

## Build and run

### Windows

Open Command Prompt or PowerShell in the project directory:

```powershell
.\build-windows.bat
.\build\CDMan.exe
```

This is a batch file, so no PowerShell script execution-policy change is required.
Alternatively, build directly:

```powershell
go test ./...
go build -buildvcs=false -trimpath -ldflags "-H=windowsgui" -o build/CDMan.exe ./cmd/cdman
```

### Linux

```sh
sh build-linux.sh
./build/cdman
```

### macOS

```sh
sh build-macos.sh
./build/cdman
```

All three scripts run the tests first and stop on failure. Compiled binaries
are written to `build/`; binaries are not included in the source distribution.
On Linux and macOS, keep SDL2 installed when running the resulting executable.

## Controls

| Action | Control |
| --- | --- |
| Open the menu from the title sequence | Space |
| Navigate menus | Arrow keys |
| Select a menu item | Enter |
| Move the first player with default keyboard controls | Arrow keys |
| Return to the menu during a game | Esc |
| Resume a paused game | Select **Continue** |
| Toggle fullscreen | Alt+Enter |

Use the player settings to configure controls for each player. Open the in-game
**Help** page for gameplay rules and additional control information.

## Saved data

High scores and recorded demo data are stored in the `CDMan-Go` folder inside
the operating system's user configuration directory:

| Platform | Default location |
| --- | --- |
| Windows | `%APPDATA%\CDMan-Go` |
| Linux | `$XDG_CONFIG_HOME/CDMan-Go`, or `~/.config/CDMan-Go` |
| macOS | `~/Library/Application Support/CDMan-Go` |

Existing `HIGHSC.CDM` and `DEMO.CDM` files in this folder remain compatible.
Replacing the application or rebuilding the project does not remove saved data.
The game saves these files through its high-score and demo workflows; closing
the window does not create a full gameplay checkpoint.

## Project layout

- `cmd/cdman/` — application entry point and persistent storage.
- `internal/game/` — game state, gameplay, rendering, sound, and Go tests.
- `internal/assets/` — embedded images, fonts, maps, menu data, and demo data.
- `internal/platform/` — native window, input, and audio adapters.
- `build-windows.bat`, `build-linux.sh`, `build-macos.sh` — build scripts.

Run `go test ./...` to check arithmetic, all 17 graphical resources, and the
menu → direct game start → movement → pause → continue workflow.

## Credits and licensing

Game artwork and content retain the rights of their respective copyright holders.
The bundled 8×14 bitmap font (`internal/game/ega14.bin`) is from the DOSBox-X
`int10_font_14` table, copyright © 2002–2020 The DOSBox Team, licensed under
GPL-2.0-or-later. Its license is provided in [LICENSE-font.txt](LICENSE-font.txt).
The font license does not grant a separate license to the other game content.
