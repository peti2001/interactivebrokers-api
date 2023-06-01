package main

import (
	"fmt"
	"time"

	"github.com/hadrianl/ibapi"
	"lighthoseCapital/wrapper"
)

func main() {
	w := &wrapper.Wrapper{
		DebugMode: false,
	}

	// internal api log is zap log, you could use GetLogger to get the logger
	// besides, you could use SetAPILogger to set you own log option
	// or you can just use the other logger
	log := ibapi.GetLogger().Sugar()
	defer log.Sync()
	// implement your own IbWrapper to handle the msg delivered via tws or gateway
	// Wrapper{} below is a default implement which just log the msg
	ic := ibapi.NewIbClient(w)

	// tcp connect with tws or gateway
	// fail if tws or gateway had not yet set the trust IP
	if err := ic.Connect("127.0.0.1", 4002, 0); err != nil {
		log.Panic("Connect failed:", err)
	}

	// handshake with tws or gateway, send handshake protocol to tell tws or gateway the version of client
	// and receive the server version and connection time from tws or gateway.
	// fail if someone else had already connected to tws or gateway with same clientID
	if err := ic.HandShake(); err != nil {
		log.Panic("HandShake failed:", err)
	}

	//ic.ReqContractDetails(ic.GetReqID(), &ibapi.Contract{
	//	Symbol:       "GBL",
	//	SecurityType: "FUT",
	//	Currency:     "EUR",
	//	Exchange:     "EUREX",
	//})

	fmt.Println("Marked Data for \"Euro Bund (10 Year Bond)\":")
	ic.ReqMarketDataType(3)
	ic.ReqMktData(
		ic.GetReqID(),
		&ibapi.Contract{
			Symbol:       "GBL",
			SecurityType: "FUT",
			Currency:     "EUR",
			Exchange:     "EUREX",
			Expiry:       "20231207",
		},
		"221",
		false,
		false,
		nil,
	)

	// start to send req and receive msg from tws or gateway after this
	ic.Run()
	<-time.After(time.Second * 60)
	ic.Disconnect()
}
