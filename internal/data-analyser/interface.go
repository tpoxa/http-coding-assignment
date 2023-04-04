package dataanalyser

import (
	"context"
	"io"
)

//go:generate mockery --name=IDataAnalyser --filename=analyser.go
type IDataAnalyser interface {
	CalcSum(ctx context.Context, r io.Reader) (float64, error)
}
