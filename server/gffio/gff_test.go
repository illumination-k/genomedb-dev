package gffio_test

import (
	"genomedb/gffio"
	"reflect"
	"testing"
)

func TestBasicGffParserConsumeLine(t *testing.T) {
	commentLine := "##gff-version 3"
	gffLine := "chr1\tfeature\trRNA\t1\t1967\t.\t-\t.\tID=Mp1g00005a.1;geneID=Mp1g00005a"

	gffRecordParser := gffio.NewGffParser()

	commentParseErr := gffRecordParser.ConsumeLine(commentLine)
	t.Run("Check a Comment Line Parsing", func(t *testing.T) {
		if commentParseErr != nil {
			t.Errorf("Error: %v", commentParseErr)
		}
	})

	err := gffRecordParser.ConsumeLine(gffLine)

	t.Run("Check a Record Line Parsing", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	multimap := gffio.NewMultiMap()
	multimap.Put("ID", "Mp1g00005a.1")
	multimap.Put("geneID", "Mp1g00005a")
	expected := []gffio.GffRecord{
		*gffio.NewGffRecord("chr1", "feature", "rRNA", 1, 1967, ".", "-", ".", multimap),
	}

	t.Run("Basic consume line test", func(t *testing.T) {
		if !reflect.DeepEqual(expected, gffRecordParser.Records) {
			t.Errorf("\nError in parsing line: %v\nExpected: %v\nActual: %v", gffLine, expected, gffRecordParser.Records)
		}
	})
}
