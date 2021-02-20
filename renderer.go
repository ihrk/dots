package dots

import (
	"io"
)

type Renderer struct {
	buf []byte
}

func NewRenderer(buf []byte) *Renderer {
	return &Renderer{buf: buf}
}

func (r *Renderer) checkBuf(n int) {
	if n <= len(r.buf) {
		return
	}
	r.buf = make([]byte, len(r.buf)+n)
}

//Render writes all image data into
//buffer, then writes buffered data into wr.
func (r *Renderer) Render(wr io.Writer, img *Image) (err error) {
	n := img.ByteLen()

	r.checkBuf(n)

	img.read(r.buf)

	_, err = wr.Write(r.buf[:n])

	return
}
