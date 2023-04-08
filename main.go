package main

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-common-go/authentication"
)

var nexServer *nex.Server

func main() {
	nexServer = nex.NewServer()
	nexServer.SetPRUDPVersion(1)
	nexServer.SetPRUDPProtocolMinorVersion(2)
	nexServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 3,
		Minor: 8,
		Patch: 3,
	})
	nexServer.SetKerberosPassword(os.Getenv("KERBEROS_PASSWORD"))
	nexServer.SetAccessKey("9f2b4678")

	nexServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==SMM1 - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	authenticationProtocol := authentication.NewCommonAuthenticationProtocol(nexServer)

	secureStationURL := nex.NewStationURL("")
	secureStationURL.SetScheme("prudps")
	secureStationURL.SetAddress(os.Getenv("SECURE_SERVER_LOCATION"))
	secureStationURL.SetPort(os.Getenv("SECURE_SERVER_PORT"))
	secureStationURL.SetCID("1")
	secureStationURL.SetPID("2")
	secureStationURL.SetSID("1")
	secureStationURL.SetStream("10")
	secureStationURL.SetType("2")

	authenticationProtocol.SetSecureStationURL(secureStationURL)
	authenticationProtocol.SetBuildName("Pretendo SMM")
	authenticationProtocol.SetPasswordFromPIDFunction(passwordFromPID)

	nexServer.Listen(":60002")
}
