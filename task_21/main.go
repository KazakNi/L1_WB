package main

import "fmt"

/* Реализовать паттерн «адаптер» на любом примере. */
type WiFiPrinter struct{}

func (w *WiFiPrinter) PrintPaperbyWiFi() {
	fmt.Println("Printing by WiFi.")
}

type CableTypePrinter struct {
}

func (c *CableTypePrinter) PrintPaper() {
	fmt.Println("Printing paper via cable connected printer")
}

type Printer interface {
	PrintPaper()
}

type Office struct {
}

func (p *Office) PrintPaper(prnt Printer) {
	fmt.Println("Client asks to print paper.")
	prnt.PrintPaper()
}

type PrinterAdapter struct {
	printerMachine *WiFiPrinter
}

func (p *PrinterAdapter) PrintPaper() {
	fmt.Println("Cable signal is converted to WiFi.")
	p.printerMachine.PrintPaperbyWiFi()
}

func main() {

	office := &Office{}
	cable_printer := &CableTypePrinter{}

	office.PrintPaper(cable_printer)

	wifi_printer := &WiFiPrinter{}
	PrinterAdapter := &PrinterAdapter{
		printerMachine: wifi_printer,
	}

	office.PrintPaper(PrinterAdapter)

}
