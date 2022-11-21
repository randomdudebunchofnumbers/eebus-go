package main

import (
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/DerAndereAndi/eebus-go/service"
	"github.com/DerAndereAndi/eebus-go/spine/model"
)

type evse struct {
	myService *service.EEBUSService
}

func (h *evse) run() {
	var err error
	var certificate tls.Certificate
	var remoteSki string

	if len(os.Args) == 5 {
		remoteSki = os.Args[2]

		certificate, err = tls.LoadX509KeyPair(os.Args[3], os.Args[4])
		if err != nil {
			usage()
			log.Fatal(err)
		}
	} else {
		certificate, err = service.CreateCertificate("Demo", "Demo", "DE", "Demo-Unit-02")
		if err != nil {
			log.Fatal(err)
		}

		pemdata := pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certificate.Certificate[0],
		})
		fmt.Println(string(pemdata))

		b, err := x509.MarshalECPrivateKey(certificate.PrivateKey.(*ecdsa.PrivateKey))
		if err != nil {
			log.Fatal(err)
		}
		pemdata = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
		fmt.Println(string(pemdata))
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
		log.Fatal(err)
	}

	serviceDescription, err := service.NewServiceDescription(
		"Demo", "Demo", "EVSE", "234567890",
		model.DeviceTypeTypeChargingStation, port, certificate)
	if err != nil {
		log.Fatal(err)
	}
	serviceDescription.SetAlternateIdentifier("Demo-EVSE-234567890")

	h.myService = service.NewEEBUSService(serviceDescription, h)
	h.myService.SetLogging(h)

	if err = h.myService.Setup(); err != nil {
		fmt.Println(err)
		return
	}

	if len(remoteSki) == 0 {
		os.Exit(0)
	}

	h.myService.Start()
	// defer h.myService.Shutdown()

	remoteService := service.ServiceDetails{
		SKI: remoteSki,
	}
	h.myService.RegisterRemoteService(remoteService)
}

// handle a request to trust a remote service
func (h *evse) RemoteServiceTrustRequested(service *service.EEBUSService, ski string) {
	// we directly trust it in this example
	h.myService.UpdateRemoteServiceTrust(ski, true)
}

// report the Ship ID of a newly trusted connection
func (h *evse) RemoteServiceShipIDReported(service *service.EEBUSService, ski string, shipID string) {
	// we should associated the Ship ID with the SKI and store it
	// so the next connection can start trusted
	fmt.Println("SKI", ski, "has Ship ID:", shipID)
}

func (h *evse) RemoteSKIConnected(service *service.EEBUSService, ski string) {}

func (h *evse) RemoteSKIDisconnected(service *service.EEBUSService, ski string) {}

// main app
func usage() {
	fmt.Println("First Run:")
	fmt.Println("  go run /cmd/evse/main.go <serverport>")
	fmt.Println()
	fmt.Println("General Usage:")
	fmt.Println("  go run /cmd/evse/main.go <serverport> <hems-ski> <crtfile> <keyfile>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	h := evse{}
	h.run()

	// Clean exit to make sure mdns shutdown is invoked
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	// User exit
}

// Logging interface

func (h *evse) Trace(args ...interface{}) {
	h.print("TRACE", args...)
}

func (h *evse) Tracef(format string, args ...interface{}) {
	h.printFormat("TRACE", format, args...)
}

func (h *evse) Debug(args ...interface{}) {
	h.print("DEBUG", args...)
}

func (h *evse) Debugf(format string, args ...interface{}) {
	h.printFormat("DEBUG", format, args...)
}

func (h *evse) Info(args ...interface{}) {
	h.print("INFO ", args...)
}

func (h *evse) Infof(format string, args ...interface{}) {
	h.printFormat("INFO ", format, args...)
}

func (h *evse) Error(args ...interface{}) {
	h.print("ERROR", args...)
}

func (h *evse) Errorf(format string, args ...interface{}) {
	h.printFormat("ERROR", format, args...)
}

func (h *evse) currentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (h *evse) print(msgType string, args ...interface{}) {
	fmt.Printf("%s %s ", h.currentTimestamp(), msgType)
	fmt.Println(args...)
}

func (h *evse) printFormat(msgType, format string, args ...interface{}) {
	value := fmt.Sprintf(format, args...)
	fmt.Println(h.currentTimestamp(), msgType, value)
}
