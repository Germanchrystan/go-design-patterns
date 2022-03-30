package firstexample

/*
Let suppose that you have some sort of Document type,
and you want to make an interface that allows people to build the different machines,
different constructs for operating on a document. Doing things like printing the document, sending it as a fax, etc.
*/

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(d Document) {

}

func (m *MultiFunctionPrinter) Fax(d Document) {

}

func (m *MultiFunctionPrinter) Scan(d Document) {

}

/*
The old printer doesn't have any Faxing or Scanning capabilities.
But because you want to implement this interface for whatever reason
(maybe some other API rely on the machine interface),
you have to implement this anyway.
*/

type OldFashionedPrinter struct {
}

func (o *OldFashionedPrinter) Print(d Document) {

}

// The old printer can't implement this methods
func (o *OldFashionedPrinter) Fax(d Document) {
	panic("Operation not supported")
}

//Deprecated: ...
//(the above comment is a lie)
func (o *OldFashionedPrinter) Scan(d Document) {
	panic("Operation not supported")
}

func main() {
	ofp := OldFashionedPrinter{}
	ofp.Scan(Document{})
}
