package main

import (
	"fmt"
	"iter"
	"log"
	"os"

	ag "github.com/apache/arrow-go/v18/arrow"
	av "github.com/apache/arrow-go/v18/arrow/avro"
	ha "github.com/hamba/avro/v2"
	ja "github.com/takanoriyanagitani/go-jsons2arrow"
)

func must[T any](t T, e error) T {
	if nil != e {
		panic(e)
	}

	return t
}

var avroSchemaName string = os.Getenv("ENV_AVRO_SCHEMA_NAME")
var avroSchemaBytes []byte = must(os.ReadFile(avroSchemaName)) //nolint:gosec

var avroSchema ha.Schema = ha.MustParse(string(avroSchemaBytes))

var arrowSchema *ag.Schema = must(av.ArrowSchemaFromAvro(avroSchema))

var jschema ja.Schema = ja.Schema{Schema: arrowSchema}

var ibatch iter.Seq2[ag.RecordBatch, error] = jschema.StdinToIterDefault()

func printAll(ib iter.Seq2[ag.RecordBatch, error]) error {
	for rbat, e := range ib {
		if nil != e {
			return e
		}

		fmt.Printf("%v\n", rbat) //nolint:forbidigo
	}

	return nil
}

func main() {
	e := printAll(ibatch)
	if nil != e {
		log.Fatal(e)
	}
}
