package main

import (
    "os"
    "fmt"

    "xmlvisitor/xmlvisitor"

    flags "github.com/jessevdk/go-flags"
)

type xmlVisitor struct {
}

func (xv *xmlVisitor) HandleStart(tagName *string, attrp *map[string]string, xp *xmlvisitor.XmlParser) error {
    fmt.Printf("Start: [%s]\n", *tagName)

    return nil
}

func (xv *xmlVisitor) HandleEnd(tagName *string, xp *xmlvisitor.XmlParser) error {
    fmt.Printf("Stop: [%s]\n", *tagName)

    return nil
}

func (xv *xmlVisitor) HandleCharData(data *string, xp *xmlvisitor.XmlParser) error {
    if xp.LastState() != xmlvisitor.XmlPartStartTag {
        return nil
    }

    fmt.Printf("CharData: [%s]\n", *data)

    return nil
}

func (xv *xmlVisitor) HandleComment(comment *string, xp *xmlvisitor.XmlParser) error {
    fmt.Printf("Comment: [%s]\n", *comment)

    return nil
}

func (xv *xmlVisitor) HandleProcessingInstruction(target *string, instruction *string, xp *xmlvisitor.XmlParser) error {
    fmt.Printf("Processing Instruction: [%s] [%s]\n", *target, *instruction)

    return nil
}

func (xv *xmlVisitor) HandleDirective(directive *string, xp *xmlvisitor.XmlParser) error {
    fmt.Printf("Directive: [%s]\n", *directive)

    return nil
}

func newXmlVisitor() (*xmlVisitor) {
    return &xmlVisitor {}
}

type options struct {
    XmlFilepath string  `short:"f" long:"xml-filepath" description:"XML file-path" required:"true"`
}

func readOptions () *options {
    o := options {}

    _, err := flags.Parse(&o)
    if err != nil {
        os.Exit(1)
    }

    return &o
}

func main() {
    var xmlFilepath string

    o := readOptions()

    xmlFilepath = o.XmlFilepath

    v := newXmlVisitor()
    p := xmlvisitor.NewXmlParser(&xmlFilepath, v)

    defer p.Close()

    err := p.Parse()
    if err != nil {
        print("Error: %s\n", err.Error())
        os.Exit(1)
    }
}
