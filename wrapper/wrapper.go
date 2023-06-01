package wrapper

import (
	"fmt"
	"github.com/hadrianl/ibapi"
)

type Wrapper struct {
	ibapi.Wrapper
	DebugMode bool
}

func (w *Wrapper) TickSize(reqID int64, tickType int64, size int64) {
	if w.DebugMode {
		w.Wrapper.TickSize(reqID, tickType, size)
	}
}

func (w *Wrapper) MarketDataType(reqID int64, marketDataType int64) {
	if w.DebugMode {
		w.Wrapper.MarketDataType(reqID, marketDataType)
	}
}

func (w *Wrapper) TickReqParams(tickerID int64, minTick float64, bboExchange string, snapshotPermissions int64) {
	if w.DebugMode {
		w.Wrapper.TickReqParams(tickerID, minTick, bboExchange, snapshotPermissions)
	}
}

func (w *Wrapper) ContractDetails(reqID int64, conDetails *ibapi.ContractDetails) {
	if w.DebugMode {
		w.Wrapper.ContractDetails(reqID, conDetails)
	}
}

func (w *Wrapper) ContractDetailsEnd(reqID int64) {
	if w.DebugMode {
		w.Wrapper.ContractDetailsEnd(reqID)
	}
}

func (w *Wrapper) Error(reqID int64, errCode int64, errString string) {
	if w.DebugMode {
		w.Wrapper.Error(reqID, errCode, errString)
	}
}

func (w *Wrapper) TickPrice(reqID int64, tickType int64, price float64, attrib ibapi.TickAttrib) {
	if tickType == 37 {
		fmt.Printf("Calculated price: %.2f EUR\n", price)
	} else if tickType == 67 {
		fmt.Printf("Last ask price: %.2f EUR\n", price)
	} else if tickType == 66 {
		fmt.Printf("Last bid price: %.2f EUR\n", price)
	} else {
		if w.DebugMode {
			w.Wrapper.TickPrice(reqID, tickType, price, attrib)
		}
	}
}
