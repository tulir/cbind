package cbind

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
)

// Modifier labels
const (
	LabelCtrl  = "ctrl"
	LabelAlt   = "alt"
	LabelMeta  = "meta"
	LabelShift = "shift"
)

var fullKeyNames = map[string]string{
	"backspace2": "Backspace",
	"pgup":       "PageUp",
	"pgdn":       "PageDown",
	"esc":        "Escape",
}

// Decode decodes a string as a key or combination of keys.
func Decode(s string) (mod tcell.ModMask, key tcell.Key, ch rune, err error) {
	if len(s) == 0 {
		return 0, 0, 0, errors.New("empty string")
	}

	// Special case for plus rune decoding
	if s[len(s)-1:] == "+" {
		key = tcell.KeyRune
		ch = '+'

		if len(s) == 1 {
			return mod, key, ch, nil
		} else if len(s) == 2 {
			return 0, 0, 0, fmt.Errorf("invalid key %s", s)
		} else {
			s = s[:len(s)-2]
		}
	}

	split := strings.Split(s, "+")
DECODEPIECE:
	for _, piece := range split {
		// Decode modifiers
		pieceLower := strings.ToLower(piece)
		switch pieceLower {
		case LabelCtrl:
			mod |= tcell.ModCtrl
			continue
		case LabelAlt:
			mod |= tcell.ModAlt
			continue
		case LabelMeta:
			mod |= tcell.ModMeta
			continue
		case LabelShift:
			mod |= tcell.ModShift
			continue
		}

		// Decode key
		for shortKey, fullKey := range fullKeyNames {
			if pieceLower == strings.ToLower(fullKey) {
				pieceLower = shortKey
				break
			}
		}
		switch pieceLower {
		case "backspace":
			key = tcell.KeyBackspace2
			continue
		case "space", "spacebar":
			key = tcell.KeyRune
			ch = ' '
			continue
		}
		for k, keyName := range tcell.KeyNames {
			if pieceLower == strings.ToLower(keyName) {
				key = k
				continue DECODEPIECE
			}
		}

		// Decode rune
		if len(piece) > 1 {
			return 0, 0, 0, fmt.Errorf("unknown key name or invalid rune: %s", piece)
		}

		key = tcell.KeyRune
		ch = rune(piece[0])
	}

	return mod, key, ch, nil
}

// Encode encodes a key or combination of keys a string.
func Encode(mod tcell.ModMask, key tcell.Key, ch rune) (string, error) {
	var b strings.Builder
	var wrote bool

	// Encode modifiers
	if mod&tcell.ModCtrl != 0 {
		b.WriteString(upperFirst(LabelCtrl))
		wrote = true
	}
	if mod&tcell.ModAlt != 0 {
		if wrote {
			b.WriteRune('+')
		}
		b.WriteString(upperFirst(LabelAlt))
		wrote = true
	}
	if mod&tcell.ModMeta != 0 {
		if wrote {
			b.WriteRune('+')
		}
		b.WriteString(upperFirst(LabelMeta))
		wrote = true
	}
	if mod&tcell.ModShift != 0 {
		if wrote {
			b.WriteRune('+')
		}
		b.WriteString(upperFirst(LabelShift))
		wrote = true
	}

	if ch == ' ' {
		if wrote {
			b.WriteRune('+')
		}
		b.WriteString("Space")
	} else if key != tcell.KeyRune {
		// Encode key
		keyName := tcell.KeyNames[key]
		if keyName == "" {
			return "", fmt.Errorf("invalid or unknown key: %d", key)
		}
		fullKeyName := fullKeyNames[strings.ToLower(keyName)]
		if fullKeyName != "" {
			keyName = fullKeyName
		}

		if wrote {
			b.WriteRune('+')
		}
		b.WriteString(keyName)
	} else {
		// Encode rune
		if wrote {
			b.WriteRune('+')
		}
		b.WriteRune(ch)
	}

	return b.String(), nil
}

func upperFirst(s string) string {
	if len(s) <= 1 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
