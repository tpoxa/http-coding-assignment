package dataanalyser

import (
	"bytes"
	"context"
	"github.com/qredo-external/go-maksym-trofimenko/internal/data-analyser/calcs"
	"io"
)

type Analyser struct {
}

func NewAnalyser() *Analyser {
	return &Analyser{}
}

func (a Analyser) CalcSum(_ context.Context, r io.Reader) (float64, error) {
	b := bytes.Buffer{}
	_, err := b.ReadFrom(r)
	if err != nil {
		return 0, err
	}
	return calcs.SumJsonNumbers(b.Bytes())
}
