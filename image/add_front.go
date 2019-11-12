package main

import (
	//"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const (
	FontFile               = "./Alibaba-PuHuiTi-Medium.ttf"
	BackgroundFile         = "./icon.png"
	DPI                    = 72
	FontSize       float64 = 18.0
	Spacing        float64 = 1.5
	Text           string  = "咁大瘾"
)

func main() {
	fontBytes, err := ioutil.ReadFile(FontFile)
	if err != nil {
		log.Println(err)
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
	}
	bgFile, err := os.Open(BackgroundFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	bg, err := png.Decode(bgFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fg := image.Black
	rgba := image.NewRGBA(image.Rect(0, 0, bg.Bounds().Size().X, bg.Bounds().Size().Y))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(DPI)
	c.SetFont(f)
	c.SetFontSize(FontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	var x1, x2, y1, y2 int = 0, bg.Bounds().Size().X, 0, bg.Bounds().Size().Y

	//计算文字插入点
	all_width := 0
	min_y_offset := int(^uint(0) >> 1)
	max_y_offset := ^int(^uint(0) >> 1)
	min_x_offset := int(^uint(0) >> 1)
	max_x_offset := ^int(^uint(0) >> 1)
	all_hight := int(^uint(0) >> 1)
	opts := truetype.Options{}
	opts.Size = FontSize
	face := truetype.NewFace(f, &opts)

	need_judge := false
	need_judge_size := 0
	for _, x := range Text {
		gb, _, _ := face.GlyphBounds(rune(x))
		i_y_offset := int(math.Floor(float64(gb.Min.Y)/64 + 0.5))
		i_y_offset_max := int(math.Floor(float64(gb.Max.Y)/64 + 0.5))
		if i_y_offset_max-i_y_offset < int(FontSize/2) {
			need_judge_size++
		}

	}
	if need_judge_size != len([]rune(Text)) {
		need_judge = true
	}

	for _, x := range Text {

		gb, awidth, _ := face.GlyphBounds(rune(x))
		iwidthf := int(math.Floor(float64(awidth)/64 + 0.5))
		i_y_offset := int(math.Floor(float64(gb.Min.Y)/64 + 0.5))
		i_y_offset_max := int(math.Floor(float64(gb.Max.Y)/64 + 0.5))
		if need_judge && i_y_offset_max-i_y_offset < int(FontSize/2) {

			i_x_offset := int(math.Floor(float64(gb.Min.X)/64 + 0.5))
			if min_x_offset > i_x_offset {
				min_x_offset = i_x_offset
			}
			if max_x_offset < i_x_offset {
				max_x_offset = i_x_offset
			}

		} else {
			if min_y_offset > i_y_offset {
				min_y_offset = i_y_offset
			}
			if max_y_offset < i_y_offset {
				max_y_offset = i_y_offset
			}

			if all_hight > i_y_offset_max-i_y_offset {
				all_hight = i_y_offset_max - i_y_offset
			}

			i_x_offset := int(math.Floor(float64(gb.Min.X)/64 + 0.5))
			if min_x_offset > i_x_offset {
				min_x_offset = i_x_offset
			}
			if max_x_offset < i_x_offset {
				max_x_offset = i_x_offset
			}

		}
		all_width += iwidthf

	}

	pt := freetype.Pt(((x1+x2)/2)-all_width/2+(min_x_offset+max_x_offset)/2, ((y1+y2)/2)-all_hight/2-(min_y_offset+max_y_offset)/2)

	//var sX, sY int
	//sX = bg.Bounds().Size().X / 2
	////sY = bg.Bounds().Size().Y/2 - int(c.PointToFixed(FontSize))>>6*len(Text)
	//sY = bg.Bounds().Size().Y / 2
	//pt := freetype.Pt(sX, sY)
	_, err = c.DrawString(Text, pt)
	if err != nil {
		log.Println(err)
	}
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := &bytes.Buffer{}
	//b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//err = b.Flush()
	err, fileUrl := uploadFile("", bytes.NewReader(b.Bytes()))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("create png suc: url=%s", fileUrl))
}

type FileUploadResponse struct {
	Url string `json:"url"`
}

func uploadFile(url string, body io.Reader) (error, string) {
	resp, err := http.Post(url, "application/x-png", body)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println("upload resp: %s", string(respBody))
	if err != nil {
		log.Println(err)
		return err, ""
	}
	ret := FileUploadResponse{}
	err = json.Unmarshal(respBody, &ret)
	if err != nil {
		return err, ""
	}
	return nil, ret.Url
}
