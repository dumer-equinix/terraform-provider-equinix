package equinix

import (
	"context"
	"log"
	"strings"
	"time"

	v4 "github.com/equinix-labs/fabric-go/fabric/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFabricPort() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(1 * time.Minute),
		},
		ReadContext: resourceFabricPortRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: readFabricPortResourceSchema(),

		Description: "Resource allows creation and management of Equinix Fabric	layer 2 connections",
	}
}

func resourceFabricPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Config).fabricClient
	ctx = context.WithValue(ctx, v4.ContextAccessToken, meta.(*Config).FabricAuthToken)
	port, _, err := client.PortsApi.GetPortByUuid(ctx, d.Id())
	log.Printf("Port read response %v", port)
	if err != nil {
		log.Printf("[WARN] Port %s not found , error %s", d.Id(), err)
		//TODO needs to see if I can trigger actual error and use exact error message
		if !strings.Contains(err.Error(), "500") {
			d.SetId("")
		}
		return diag.FromErr(err)
	}
	d.SetId(port.Uuid)
	return setFabricPortMap(d, port)
}

func setFabricPortMap(d *schema.ResourceData, port v4.Port) diag.Diagnostics {
	diags := diag.Diagnostics{}
	err := setMap(d, map[string]interface{}{
		"name":                   port.Name,
		"bandwidth":              port.Bandwidth,
		"available_bandwidth":    port.AvailableBandwidth,
		"used_bandwidth":         port.UsedBandwidth,
		"href":                   port.Href,
		"description":            port.Description,
		"type":                   port.Type_,
		"state":                  port.State,
		"service_type":           port.ServiceType,
		"cvp_id":                 port.CvpId,
		"asn":                    port.Asn,
		"physical_port_quantity": port.PhysicalPortQuantity,
		"operation":              portOperationToTerra(port.Operation),
		"order":                  orderMappingToTerra(port.Order),
		"redundancy":             portRedundancyToTerra(port.Redundancy),
		"account":                accountToTerra(port.Account),
		"change_log":             changeLogToTerra(port.Changelog),
		"location":               locationToTerra(port.Location),
		//"device":                 deviceToTerra(port.Device),
		//"interface":              portInterfaceToTerra(port.Interface_),
		//"tether":                 tetherToTerra(port.Tether),
		//"encapsulation":          encapsulationrToTerra(port.Encapsulation),
		//"lag":                    lagToTerra(port.Lag),
		//"settings":               settingsToTerra(port.Settings),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
