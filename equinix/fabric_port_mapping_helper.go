package equinix

import (
	v4 "github.com/equinix-labs/fabric-go/fabric/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func portRedundancy2ToTerra(redundancy *v4.PortRedundancy) *schema.Set {
	if redundancy == nil {
		return nil
	}
	redundancies := []*v4.PortRedundancy{redundancy}
	mappedRedundancies := make([]interface{}, 0)
	for _, redundancy := range redundancies {
		mappedRedundancy := make(map[string]interface{})
		mappedRedundancy["enabled"] = redundancy.Enabled
		mappedRedundancy["group"] = redundancy.Group
		mappedRedundancy["priority"] = string(*redundancy.Priority)
		mappedRedundancies = append(mappedRedundancies, mappedRedundancy)
	}
	redundancySet := schema.NewSet(
		schema.HashResource(readPortsRedundancyRes),
		mappedRedundancies,
	)
	return redundancySet
}

func portOperationToTerra(operation *v4.PortOperation) *schema.Set {
	if operation == nil {
		return nil
	}
	operations := []*v4.PortOperation{operation}
	mappedOperations := make([]interface{}, 0)
	for _, operation := range operations {
		mappedOperation := make(map[string]interface{})
		mappedOperation["operational_status"] = operation.OperationalStatus
		mappedOperation["connection_count"] = operation.ConnectionCount
		mappedOperation["op_status_changed_at"] = operation.OpStatusChangedAt.String()
		mappedOperations = append(mappedOperations, mappedOperation)
	}
	operationSet := schema.NewSet(
		schema.HashResource(createOperationRes),
		mappedOperations,
	)
	return operationSet
}

func portDeviceRedundancyToTerra(redundancy *v4.PortDeviceRedundancy) *schema.Set {
	if redundancy == nil {
		return nil
	}
	redundancies := []*v4.PortDeviceRedundancy{redundancy}
	mappedRedundancies := make([]interface{}, 0)
	for _, redundancy := range redundancies {
		mappedRedundancy := make(map[string]interface{})
		mappedRedundancy["group"] = redundancy.Group
		mappedRedundancy["priority"] = redundancy.Priority
		mappedRedundancies = append(mappedRedundancies, mappedRedundancy)
	}
	redundancySet := schema.NewSet(
		schema.HashResource(readPortDeviceRedundancyRes),
		mappedRedundancies,
	)
	return redundancySet
}

func deviceToTerra(device *v4.PortDevice) *schema.Set {
	if device == nil {
		return nil
	}
	devices := []*v4.PortDevice{device}
	mappedDevices := make([]interface{}, 0)
	for _, device := range devices {
		mappedDevice := make(map[string]interface{})
		mappedDevice["name"] = device.Name
		if device.Redundancy != nil {
			mappedDevice["redundancy"] = portDeviceRedundancyToTerra(device.Redundancy)
		}
		mappedDevices = append(mappedDevices, mappedDevice)
	}
	deviceSet := schema.NewSet(
		schema.HashResource(readPortDeviceRes),
		mappedDevices,
	)
	return deviceSet
}

func portInterfaceToTerra(portInterface *v4.PortInterface) *schema.Set {
	if portInterface == nil {
		return nil
	}
	portInterfaces := []*v4.PortInterface{portInterface}
	mappedPortInterfaces := make([]interface{}, 0)
	for _, portInterface := range portInterfaces {
		mappedPortInterface := make(map[string]interface{})
		mappedPortInterface["type"] = portInterface.Type_
		mappedPortInterface["ifIndex"] = portInterface.IfIndex
		mappedPortInterface["name"] = portInterface.Name
		mappedPortInterfaces = append(mappedPortInterfaces, mappedPortInterface)
	}
	portInterfaceSet := schema.NewSet(
		schema.HashResource(readPortInterfaceRes),
		mappedPortInterfaces,
	)
	return portInterfaceSet
}

func tetherToTerra(portTether *v4.PortTether) *schema.Set {
	if portTether == nil {
		return nil
	}
	portTethers := []*v4.PortTether{portTether}
	mappedPortTethers := make([]interface{}, 0)
	for _, portTether := range portTethers {
		mappedPortTether := make(map[string]interface{})
		mappedPortTether["cross_connectId"] = portTether.CrossConnectId
		mappedPortTether["cabinet_number"] = portTether.CabinetNumber
		mappedPortTether["system_name"] = portTether.SystemName
		mappedPortTether["patch_panel"] = portTether.PatchPanel
		mappedPortTether["patch_panel_port_a"] = portTether.PatchPanelPortA
		mappedPortTether["patch_panel_port_b"] = portTether.PatchPanelPortB
		mappedPortTethers = append(mappedPortTethers, mappedPortTether)
	}
	portTetherSet := schema.NewSet(
		schema.HashResource(readPortTetherRes),
		mappedPortTethers,
	)
	return portTetherSet
}

func encapsulationToTerra(portEncapsulation *v4.PortEncapsulation) *schema.Set {
	if portEncapsulation == nil {
		return nil
	}
	portEncapsulations := []*v4.PortEncapsulation{portEncapsulation}
	mappedPortEncapsulations := make([]interface{}, 0)
	for _, portEncapsulation := range portEncapsulations {
		mappedPortEncapsulation := make(map[string]interface{})
		mappedPortEncapsulation["type"] = portEncapsulation.Type_
		mappedPortEncapsulation["tag_protocol_id"] = portEncapsulation.TagProtocolId
		mappedPortEncapsulations = append(mappedPortEncapsulations, mappedPortEncapsulation)
	}
	portEncapsulationSet := schema.NewSet(
		schema.HashResource(readPortEncapsulationRes),
		mappedPortEncapsulations,
	)
	return portEncapsulationSet
}

func lagToTerra(portLag *v4.PortLag) *schema.Set {
	if portLag == nil {
		return nil
	}
	portLags := []*v4.PortLag{portLag}
	mappedPortLags := make([]interface{}, 0)
	for _, portLag := range portLags {
		mappedPortLag := make(map[string]interface{})
		mappedPortLag["enabled"] = portLag.Enabled
		mappedPortLag["id"] = portLag.Id
		mappedPortLag["name"] = portLag.Name
		mappedPortLag["member_status"] = portLag.MemberStatus
		mappedPortLags = append(mappedPortLags, mappedPortLag)
	}
	portLagSet := schema.NewSet(
		schema.HashResource(readPortLagRes),
		mappedPortLags,
	)
	return portLagSet
}

func settingsToTerra(portSetting *v4.PortSettings) *schema.Set {
	if portSetting == nil {
		return nil
	}
	portSettings := []*v4.PortSettings{portSetting}
	mappedPortSettings := make([]interface{}, 0)
	for _, portSetting := range portSettings {
		mappedPortSetting := make(map[string]interface{})
		mappedPortSetting["port_type"] = portSetting.PortType
		mappedPortSettings = append(mappedPortSettings, mappedPortSetting)
	}
	portSettingsSet := schema.NewSet(
		schema.HashResource(readPortSettingsRes),
		mappedPortSettings,
	)
	return portSettingsSet
}
