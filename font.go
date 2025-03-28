package main

import (
	_ "embed"
	"fmt"
	"golang.org/x/image/font/basicfont"
	"sync"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var (
	// Embed the font file directly into the binary
	//go:embed NotoSans-Regular.ttf
	notoSansFontBytes []byte
	fontCache         = make(map[float64]*truetype.Font)
	fontCacheLock     sync.RWMutex
)

func GetFont(fontSize float64) (*truetype.Font, error) {
	fontCacheLock.RLock()
	cachedFont, exists := fontCache[fontSize]
	fontCacheLock.RUnlock()

	if exists {
		return cachedFont, nil
	}

	// Parse the font
	ft, err := truetype.Parse(notoSansFontBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %w", err)
	}

	fontCacheLock.Lock()
	fontCache[fontSize] = ft
	fontCacheLock.Unlock()

	return ft, nil
}

func CreateFontFace(fontSize float64) (font.Face, error) {
	ft, err := GetFont(fontSize)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ft, &truetype.Options{
		Size: fontSize,
		DPI:  72,
	}), nil
}

func FallbackFontFace(size float64) font.Face {
	return basicfont.Face7x13
}
