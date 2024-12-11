package Encoding

import (
	"fmt"
	hashes "go-api/Hashes"
	"image"
	"image/color"
	"image/png"
	"encoding/base64"
	"bytes"
)

type Identicon struct {
	Hex       []byte
	Color     color.RGBA
	Grid      []Cell
	PixelMap  []Pixel
	ImageSize int
	CellSize  int
}

type Cell struct {
	Value int
	Index int
}

type Pixel struct {
	TopLeft     image.Point
	BottomRight image.Point
}

func hash(data, hashName string) []byte {
	hashList := hashes.GetHashFunc()
	hashFunc := hashList[hashName]
	hasher, _ := hashFunc()
	hasher.Write([]byte(data))

	return hasher.Sum(nil)
}

func HashInput(input, hashName string) Identicon {
	hash := hash(input, hashName)
	return Identicon{
		Hex: hash[:],
		ImageSize: 250,
		CellSize: 50,
	}
}

func PickColor(identicon Identicon) Identicon {
	identicon.Color = color.RGBA{
		R: identicon.Hex[0],
		G: identicon.Hex[1],
		B: identicon.Hex[2],
		A: 255,
	}
	return identicon
}

func BuildGrid(identicon Identicon) Identicon {
	var grid []Cell
	for i := 0; i < 15; i += 3 {
		row := []int{int(identicon.Hex[i]), int(identicon.Hex[i+1]), int(identicon.Hex[i+2])}
		row = append(row, row[1], row[0])
		for _, value := range row {
			grid = append(grid, Cell{Value: value, Index: len(grid)})
		}
	}
	identicon.Grid = grid
	return identicon
}

func FilterOddSquares(identicon Identicon) Identicon {
	var filteredGrid []Cell
	for _, cell := range identicon.Grid {
		if cell.Value%2 == 0 { 
			filteredGrid = append(filteredGrid, cell)
		}
	}
	identicon.Grid = filteredGrid
	return identicon
}

func BuildPixelMap(identicon Identicon) Identicon {
	var pixelMap []Pixel
	for _, cell := range identicon.Grid {
		x := (cell.Index % 5) * identicon.CellSize
		y := (cell.Index / 5) * identicon.CellSize
		topLeft := image.Point{x, y}
		bottomRight := image.Point{x + identicon.CellSize, y + identicon.CellSize}
		pixelMap = append(pixelMap, Pixel{TopLeft: topLeft, BottomRight: bottomRight})
	}
	identicon.PixelMap = pixelMap
	return identicon
}

func DrawImage(identicon Identicon) (string, error) {
	img := image.NewRGBA(image.Rect(0, 0, identicon.ImageSize, identicon.ImageSize))

	for _, pixel := range identicon.PixelMap {
		for y := pixel.TopLeft.Y; y < pixel.BottomRight.Y; y++ {
			for x := pixel.TopLeft.X; x < pixel.BottomRight.X; x++ {
				img.Set(x, y, identicon.Color)
			}
		}
	}

	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return "", fmt.Errorf("error encoding image: %v", err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Str, nil
}

func IdenticonMain(input, hashName string) (string, error) {
	identicon := HashInput(input, hashName)
	identicon = PickColor(identicon)
	identicon = BuildGrid(identicon)
	identicon = FilterOddSquares(identicon)
	identicon = BuildPixelMap(identicon)
	base64Image, err := DrawImage(identicon)
	if err != nil {
		return "", fmt.Errorf("cannot create the image: %v", err)
	}

	return base64Image, nil
}
