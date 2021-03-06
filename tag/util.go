package tag

import (
	"bytes"
	window "github.com/as/ms/win"
	"image"
	"io"
	"io/ioutil"
	"fmt"
	"log"
	"os"

	"github.com/as/clip"
	"github.com/as/cursor"
	"github.com/as/frame"
	"golang.org/x/image/font/gofont/gomono"
)

func readfile(s string) (p []byte) {
	var err error
	if isdir(s){
		fi, err := ioutil.ReadDir(s)
		if err != nil{
			log.Println(err)
			return nil
		}
		b := new(bytes.Buffer)
		for _, v := range fi{
			fmt.Fprintf(b, "%s\t", v.Name())
		}
		return b.Bytes()
	}
	p, err = ioutil.ReadFile(s)
	if err != nil {
		log.Println(err)
	}
	return p
}
func writefile(s string, p []byte) {
	fd, err := os.Create(s)
	if err != nil {
		log.Println(err)
	}
	n, err := io.Copy(fd, bytes.NewReader(p))
	if err != nil {
		log.Fatalln(err)
	}
	println("wrote", n, "bytes")
}
func mkfont(size int) frame.Font {
	return frame.NewTTF(gomono.TTF, size)
}

func init() {
	var err error
	Clip, err = clip.New()
	if err != nil {
		panic(err)
	}
}
func moveMouse(pt image.Point) {
	cursor.MoveTo(window.ClientAbs().Min.Add(pt))
}
