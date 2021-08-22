package main

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	nexproto "github.com/PretendoNetwork/nex-protocols-go"
)

var nexServer *nex.Server

func main() {
	nexServer = nex.NewServer()
	nexServer.SetPrudpVersion(1)
	nexServer.SetNexVersion(4)
	nexServer.SetKerberosKeySize(32)
	nexServer.SetAccessKey("9f2b4678")

	nexServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==SMM1 - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	authenticationServer := nexproto.NewAuthenticationProtocol(nexServer)

	// Handle LoginEx RMC method
	authenticationServer.LoginEx(loginEx)

	// Handle RequestTicket RMC method
	authenticationServer.RequestTicket(requestTicket)

	nexServer.Listen(":60002")
}
