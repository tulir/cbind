# cbind
[![GoDoc](https://gitlab.com/tslocum/godoc-static/-/raw/master/badge.svg)](https://docs.rocketnine.space/gitlab.com/tslocum/cbind)
[![CI status](https://gitlab.com/tslocum/cbind/badges/master/pipeline.svg)](https://gitlab.com/tslocum/cbind/commits/master)
[![Donate](https://img.shields.io/liberapay/receives/rocketnine.space.svg?logo=liberapay)](https://liberapay.com/rocketnine.space)

Key event handling library for tcell

## Usage

```go
// Create a new input configuration to store the keybinds.
c := cbind.NewConfiguration()

// Set keybind Alt+s.
c.SetRune(tcell.ModAlt, 's', func(ev *tcell.EventKey) *tcell.EventKey {
    // Save
    return nil
})

// Set keybind Alt+o.
c.SetRune(tcell.ModAlt, 'o', func(ev *tcell.EventKey) *tcell.EventKey {
    // Open
    return nil
})

// Set keybind Escape.
c.SetKey(tcell.ModNone, tcell.KeyEscape, func(ev *tcell.EventKey) *tcell.EventKey {
    // Exit
    return nil
})

// Capture input. This will differ based on the framework in use (if any).
// When using tview or cview, call Application.SetInputCapture before calling
// Application.Run.
app.SetInputCapture(c.Capture)
```

## Documentation

Documentation is available via [gdooc](https://docs.rocketnine.space/gitlab.com/tslocum/cbind).

You may use `whichkeybind` to determine and validate key combinations.

```bash
go get gitlab.com/tslocum/cbind/whichkeybind
```

## Support

Please share issues and suggestions [here](https://gitlab.com/tslocum/cbind/issues).
