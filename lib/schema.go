package lib

type DetailDelivery struct {
	Title string
	Value string
}

type ResultDelivery []DetailDelivery

type DataHistory struct {
	Title string
	Date  string
}

type ReceiverInfo struct {
	Name         string
	Relationship string
}

type DetailHistory struct {
	Data     []DataHistory
	Receiver ReceiverInfo
}

type DetailShipment struct {
	Title string
	Value string
}

type ResultShipment []DetailShipment

type DetailStatus struct {
	Code    string
	Message string
}

type DetailTracking struct {
	Data struct {
		Delivery []DetailDelivery
		History  DetailHistory
		Shipment []DetailShipment
	}
	Status DetailStatus
}
