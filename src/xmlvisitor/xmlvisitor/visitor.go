// XML parser/visitor logic

package xmlvisitor

import (
    "os"
    "strings"
    "errors"

    "encoding/xml"
)

const (
    XmlPart_Initial = iota
    XmlPartStartTag = iota
    XmlPartEndTag   = iota
    XmlPartCharData = iota
)

type XmlVisitor interface {

// TODO(dustin): !! Add support for xml.Comment, xml.ProcInst, and xml.Directive .

    HandleStart(tagName *string, attrp *map[string]string, xp *XmlParser) error
    HandleEnd(tagName *string, xp *XmlParser) error
    HandleCharData(data *string, xp *XmlParser) error
}

type XmlParser struct {
    f *os.File
    decoder *xml.Decoder
    ns *Stack
    v XmlVisitor
    lastState int
}

func (xp *XmlParser) NodeStack() *Stack {
    return xp.ns
}

// Create parser.
func NewXmlParser(filepath *string, visitor XmlVisitor) *XmlParser {
    f, err := os.Open(*filepath)
    if err != nil {
        panic(err)
    }

    decoder := xml.NewDecoder(f)
    ns := newStack()

    return &XmlParser {
            f: f,
            decoder: decoder,
            ns: ns,
            v: visitor,
            lastState: XmlPart_Initial,
    }
}

func (xp *XmlParser) LastState() int {
    return xp.lastState
}

func (xp *XmlParser) LastStateName() string {
    if xp.lastState == XmlPart_Initial {
        return ""
    } else if xp.lastState == XmlPartStartTag {
        return "StartTag"
    } else if xp.lastState == XmlPartEndTag {
        return "EndTag"
    } else if xp.lastState == XmlPartCharData {
        return "CharData"
    } else {
        panic(errors.New("Invalid XML state."))
    }
}

// Close resources.
func (xp *XmlParser) Close() {
    xp.f.Close()
}

// Run the parse with a minimal memory footprint.
func (xp *XmlParser) Parse() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = r.(error)
        }
    }()

    for {
        token, err := xp.decoder.Token()
        if err != nil {
            break
        }
  
        switch t := token.(type) {
        case xml.StartElement:
            elmt := xml.StartElement(t)
            name := elmt.Name.Local

            xp.ns.push(name)

            var attributes map[string]string = make(map[string]string)
            for _, a := range t.Attr {
                attributes[a.Name.Local] = a.Value
            }

            err := xp.v.HandleStart(&name, &attributes, xp)
            if err != nil {
                panic(err)
            }

            xp.lastState = XmlPartStartTag

        case xml.EndElement:
            xp.ns.pop()

            elmt := xml.EndElement(t)
            name := elmt.Name.Local

            err := xp.v.HandleEnd(&name, xp)
            if err != nil {
                panic(err)
            }

            xp.lastState = XmlPartEndTag

        case xml.CharData:
            bytes := xml.CharData(t)
            s := strings.TrimSpace(string([]byte(bytes)))

            err := xp.v.HandleCharData(&s, xp)
            if err != nil {
                panic(err)
            }

            xp.lastState = XmlPartCharData

        case xml.Comment:
        case xml.ProcInst:
        case xml.Directive:
        }
    }

    return nil
}
