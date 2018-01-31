package xmlvisitor

import (
    "os"
    "path"
    "testing"

    "encoding/json"

    "github.com/dsoprea/go-logging"
)

var (
    testingAssetsPath = ""
)

type xmlVisitor struct {
    visits [][3]string
}

func (xv *xmlVisitor) HandleStart(tagName *string, attrp *map[string]string, xp *XmlParser) error {
    xv.visits = append(xv.visits, [3]string { "start", *tagName, "" })

    return nil
}

func (xv *xmlVisitor) HandleEnd(tagName *string, xp *XmlParser) error {
    xv.visits = append(xv.visits, [3]string { "stop", *tagName, "" })

    return nil
}

func (xv *xmlVisitor) HandleValue(tagName *string, value *string, xp *XmlParser) error {
    xv.visits = append(xv.visits, [3]string { "value", *tagName, *value })

    return nil
}

/*
func (xv *xmlVisitor) HandleComment(comment *string, xp *XmlParser) error {
    fmt.Printf("Comment: [%s]\n", *comment)

    return nil
}

func (xv *xmlVisitor) HandleProcessingInstruction(target *string, instruction *string, xp *XmlParser) error {
    fmt.Printf("Processing Instruction: [%s] [%s]\n", *target, *instruction)

    return nil
}

func (xv *xmlVisitor) HandleDirective(directive *string, xp *XmlParser) error {
    fmt.Printf("Directive: [%s]\n", *directive)

    return nil
}
*/

func TestParse(t *testing.T) {
    xmlFilepath := path.Join(testingAssetsPath, "20130729.gpx")

    f, err := os.Open(xmlFilepath)
    log.PanicIf(err)

    defer f.Close()

    xv := &xmlVisitor{
        visits: make([][3]string, 0),
    }

    p := NewXmlParser(f, xv)

    err = p.Parse()
    log.PanicIf(err)

    expectedRaw := `[
  [
    "start",
    "gpx",
    ""
  ],
  [
    "start",
    "time",
    ""
  ],
  [
    "stop",
    "time",
    ""
  ],
  [
    "value",
    "time",
    "2013-07-30T02:38:29Z"
  ],
  [
    "start",
    "bounds",
    ""
  ],
  [
    "stop",
    "bounds",
    ""
  ],
  [
    "start",
    "trk",
    ""
  ],
  [
    "start",
    "trkseg",
    ""
  ],
  [
    "start",
    "trkpt",
    ""
  ],
  [
    "start",
    "ele",
    ""
  ],
  [
    "stop",
    "ele",
    ""
  ],
  [
    "value",
    "ele",
    "-18.5"
  ],
  [
    "start",
    "course",
    ""
  ],
  [
    "stop",
    "course",
    ""
  ],
  [
    "value",
    "course",
    "0.0"
  ],
  [
    "start",
    "speed",
    ""
  ],
  [
    "stop",
    "speed",
    ""
  ],
  [
    "value",
    "speed",
    "0.75"
  ],
  [
    "start",
    "hdop",
    ""
  ],
  [
    "stop",
    "hdop",
    ""
  ],
  [
    "value",
    "hdop",
    "5.8"
  ],
  [
    "start",
    "src",
    ""
  ],
  [
    "stop",
    "src",
    ""
  ],
  [
    "value",
    "src",
    "gps"
  ],
  [
    "start",
    "sat",
    ""
  ],
  [
    "stop",
    "sat",
    ""
  ],
  [
    "value",
    "sat",
    "7"
  ],
  [
    "start",
    "time",
    ""
  ],
  [
    "stop",
    "time",
    ""
  ],
  [
    "value",
    "time",
    "2013-07-30T02:38:29Z"
  ],
  [
    "stop",
    "trkpt",
    ""
  ],
  [
    "start",
    "trkpt",
    ""
  ],
  [
    "start",
    "ele",
    ""
  ],
  [
    "stop",
    "ele",
    ""
  ],
  [
    "value",
    "ele",
    "-43.900001525878906"
  ],
  [
    "start",
    "hdop",
    ""
  ],
  [
    "stop",
    "hdop",
    ""
  ],
  [
    "value",
    "hdop",
    "47.6"
  ],
  [
    "start",
    "src",
    ""
  ],
  [
    "stop",
    "src",
    ""
  ],
  [
    "value",
    "src",
    "gps"
  ],
  [
    "start",
    "sat",
    ""
  ],
  [
    "stop",
    "sat",
    ""
  ],
  [
    "value",
    "sat",
    "4"
  ],
  [
    "start",
    "time",
    ""
  ],
  [
    "stop",
    "time",
    ""
  ],
  [
    "value",
    "time",
    "2013-07-30T02:39:15Z"
  ],
  [
    "stop",
    "trkpt",
    ""
  ],
  [
    "start",
    "trkpt",
    ""
  ],
  [
    "start",
    "ele",
    ""
  ],
  [
    "stop",
    "ele",
    ""
  ],
  [
    "value",
    "ele",
    "-8.100000381469727"
  ],
  [
    "start",
    "hdop",
    ""
  ],
  [
    "stop",
    "hdop",
    ""
  ],
  [
    "value",
    "hdop",
    "23.6"
  ],
  [
    "start",
    "src",
    ""
  ],
  [
    "stop",
    "src",
    ""
  ],
  [
    "value",
    "src",
    "gps"
  ],
  [
    "start",
    "sat",
    ""
  ],
  [
    "stop",
    "sat",
    ""
  ],
  [
    "value",
    "sat",
    "5"
  ],
  [
    "start",
    "time",
    ""
  ],
  [
    "stop",
    "time",
    ""
  ],
  [
    "value",
    "time",
    "2013-07-30T02:40:17Z"
  ],
  [
    "stop",
    "trkpt",
    ""
  ],
  [
    "start",
    "trkpt",
    ""
  ],
  [
    "start",
    "ele",
    ""
  ],
  [
    "stop",
    "ele",
    ""
  ],
  [
    "value",
    "ele",
    "-24.600000381469727"
  ],
  [
    "start",
    "hdop",
    ""
  ],
  [
    "stop",
    "hdop",
    ""
  ],
  [
    "value",
    "hdop",
    "21.6"
  ],
  [
    "start",
    "src",
    ""
  ],
  [
    "stop",
    "src",
    ""
  ],
  [
    "value",
    "src",
    "gps"
  ],
  [
    "start",
    "sat",
    ""
  ],
  [
    "stop",
    "sat",
    ""
  ],
  [
    "value",
    "sat",
    "5"
  ],
  [
    "start",
    "time",
    ""
  ],
  [
    "stop",
    "time",
    ""
  ],
  [
    "value",
    "time",
    "2013-07-30T02:41:19Z"
  ],
  [
    "stop",
    "trkpt",
    ""
  ],
  [
    "stop",
    "trkseg",
    ""
  ],
  [
    "stop",
    "trk",
    ""
  ],
  [
    "stop",
    "gpx",
    ""
  ]
]`

    expected := [][3]string{}
    err = json.Unmarshal([]byte(expectedRaw), &expected)
    log.PanicIf(err)

    if len(expected) != len(xv.visits) {
        t.Fatalf("actual had (%d) entries and expected had (%d) entries", len(xv.visits), len(expected))
    }

    for i, x := range xv.visits {
        if expected[i] != x {
            t.Fatalf("entry (%d) did not match: [%v] != [%v]", i, x, expected[i])
        }
    }
}

func init() {
    goPath := os.Getenv("GOPATH")
    testingAssetsPath = path.Join(goPath, "src", "github.com", "dsoprea", "go-xmlvisitor", "assets")
}
