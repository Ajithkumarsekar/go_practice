package imgcat

import (
	"encoding/base64"
	"errors"
	"io"
	"strings"
)

//https://iterm2.com/documentation-images.html

// Copy copies the given image reader and encodes it as an
// iTerm2 image into the writer.
func Copy(w io.Writer, r io.Reader) error {
	header := strings.NewReader("\033]1337;File=inline=1:")
	footer := strings.NewReader("\a\n")

	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()

		// write the contents to pw
		wc := base64.NewEncoder(base64.StdEncoding, pw)
		_, err := io.Copy(wc, r)
		if err != nil {
			pw.CloseWithError(errors.New("could not encode image"))
			return
		}

		if err := wc.Close(); err != nil {
			pw.CloseWithError(errors.New("could not close base64 encoder"))
			return
		}
	}()
	_, err := io.Copy(w, io.MultiReader(header, pr, footer))
	return err
}

// NewWriter returns a new imgcat writer.
func NewWriter(w io.Writer) io.WriteCloser {
	pr, pw := io.Pipe()

	wc := &writer{pw, make(chan struct{})}
	go func() {
		defer close(wc.done)
		err := Copy(w, pr)
		pr.CloseWithError(err)
	}()
	return wc
}

type writer struct {
	pw   *io.PipeWriter
	done chan struct{}
}

func (w *writer) Write(data []byte) (int, error) {
	return w.pw.Write(data)
}

func (w *writer) Close() error {
	if err := w.pw.Close(); err != nil {
		return err
	}
	<-w.done
	return nil
}
