package jsons2arrow

import (
	"bufio"
	"io"
	"iter"
	"os"

	ag "github.com/apache/arrow-go/v18/arrow"
	aa "github.com/apache/arrow-go/v18/arrow/array"
)

type JSONReader struct{ *aa.JSONReader }

func (j JSONReader) ToIter() iter.Seq2[ag.RecordBatch, error] {
	return aa.IterFromReader(j.JSONReader)
}

const ChunkSizeDefault int = 1024

type ReadOptions struct {
	*ag.Schema

	Options []aa.Option
}

func (o ReadOptions) ToReader(rdr io.Reader) JSONReader {
	return JSONReader{JSONReader: aa.NewJSONReader(
		rdr,
		o.Schema,
		o.Options...,
	)}
}

func (o ReadOptions) FromStdin() JSONReader {
	var br io.Reader = bufio.NewReader(os.Stdin)
	return o.ToReader(br)
}

type Schema struct{ *ag.Schema }

func (s Schema) ToOptionsDefault() ReadOptions {
	return ReadOptions{
		Schema: s.Schema,
		Options: []aa.Option{
			aa.WithChunk(ChunkSizeDefault),
		},
	}
}

func (s Schema) StdinToIterDefault() iter.Seq2[ag.RecordBatch, error] {
	return s.
		ToOptionsDefault().
		FromStdin().
		ToIter()
}
