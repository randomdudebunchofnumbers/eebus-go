package ship

import (
	"sync"
	"time"

	"github.com/enbility/eebus-go/spine"
	spineModel "github.com/enbility/eebus-go/spine/model"
)

type dataHandlerTest struct {
	sentMessage []byte

	mux sync.Mutex

	handleConnectionClosedInvoked bool
}

func (s *dataHandlerTest) lastMessage() []byte {
	s.mux.Lock()
	defer s.mux.Unlock()

	return s.sentMessage
}

var _ ShipDataConnection = (*dataHandlerTest)(nil)

func (s *dataHandlerTest) InitDataProcessing(dataProcessing ShipDataProcessing) {}

func (s *dataHandlerTest) WriteMessageToDataConnection(message []byte) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.sentMessage = message

	return nil
}

func (s *dataHandlerTest) CloseDataConnection(int, string)       {}
func (w *dataHandlerTest) IsDataConnectionClosed() (bool, error) { return false, nil }

var _ ShipServiceDataProvider = (*dataHandlerTest)(nil)

func (s *dataHandlerTest) IsRemoteServiceForSKIPaired(string) bool { return true }
func (s *dataHandlerTest) HandleConnectionClosed(*ShipConnection, bool) {
	s.handleConnectionClosedInvoked = true
}
func (s *dataHandlerTest) ReportServiceShipID(string, string)               {}
func (s *dataHandlerTest) AllowWaitingForTrust(string) bool                 { return false }
func (s *dataHandlerTest) HandleShipHandshakeStateUpdate(string, ShipState) {}

func initTest(role shipRole) (*ShipConnection, *dataHandlerTest) {
	localDevice := spine.NewDeviceLocalImpl("TestBrandName", "TestDeviceModel", "TestSerialNumber", "TestDeviceCode",
		"TestDeviceAddress", spineModel.DeviceTypeTypeEnergyManagementSystem, spineModel.NetworkManagementFeatureSetTypeSmart, time.Second*4)

	dataHandler := &dataHandlerTest{}
	conhandler := NewConnectionHandler(dataHandler, dataHandler, localDevice, role, "LocalShipID", "RemoveDevice", "RemoteShipID")

	return conhandler, dataHandler
}

func shutdownTest(conhandler *ShipConnection) {
	conhandler.stopHandshakeTimer()
}
