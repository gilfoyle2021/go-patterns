package bridge

import (
	"errors"
	"fmt"
	"io"
)

//PrinterAPI 打印机具体的打印方式
type PrinterAPI interface {
	//具体的打印方式实现
	PrintMsg(string) error
}

//PrinterAPIImpl1 打印机具体的打印方式1
type PrinterAPIImpl1 struct {
}

func (pi *PrinterAPIImpl1) PrintMsg(msg string) error {
	fmt.Printf("Printer impl1 msg: %s \n", msg)
	return nil
}

//PrinterAPIImpl2 打印机具体的打印方式2
type PrinterAPIImpl2 struct {
	Writer io.Writer
}

func (pi *PrinterAPIImpl2) PrintMsg(msg string) error {
	if pi.Writer == nil {
		return errors.New("has not been impl yet")
	}
	fmt.Fprintf(pi.Writer, "Printer impl2 msg:%s", msg)
	return nil
}

//PrinterAbstraction 打印机的抽象类，代表打印机
type PrinterAbstraction interface {
	//Print 打印操作
	Print() error
}

//NormalPrinter 普通的打印机，可以有不同的打印方式
type NormalPrinter struct {
	//Msg print msg
	Msg     string
	Printer PrinterAPI
}

func (np *NormalPrinter) Print() error {

	return np.Printer.PrintMsg(np.Msg)
}

//PacktPrinter 数据包的打印机，可以有不同的打印方式
type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (np *PacktPrinter) Print() error {
	np.Printer.PrintMsg(fmt.Sprintf("Message from packet:%s", np.Msg))
	return nil
}
