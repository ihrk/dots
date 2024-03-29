package dots

import "io"

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

// Render writes all image data into buffer, then writes buffered data into wr.
func (r *Renderer) Render(w io.Writer, p *Image) error {
	n := p.ByteLen()

	r.checkBuf(n)

	p.read(r.buf)

	_, err := w.Write(r.buf[:n])

	return err
}
