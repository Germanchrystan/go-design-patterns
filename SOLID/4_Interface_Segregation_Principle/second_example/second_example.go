package secondexample

// ISP
/*
Try to break up interfaces into separate parts that will be needed.
In this example, there is no guarantee that if somebody need printing, they also need faxing.
It might make sense to separate the printing and the scanning into separate interfaces.
*/
type Document struct {
}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m *MyPrinter) Print(d Document) {

}

// Interfaces can be composed into different interfaces
type MultiFunctionPrinter interface {
	Printer
	Scanner
}

type Photocopier struct{}

func (p *Photocopier) Scan(d Document)  {}
func (p *Photocopier) Print(d Document) {}

// Decorator design pattern
// Granular definitions
type MultiFuncionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFuncionMachine) Print(d Document) {
	m.printer.Print(d)
}
