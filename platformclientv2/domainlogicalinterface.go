package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainlogicalinterface
type Domainlogicalinterface struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


	// Description - The resource's description.
	Description *string `json:"description,omitempty"`


	// Version - The current version of the resource.
	Version *int `json:"version,omitempty"`


	// DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The ID of the user that last modified the resource.
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// CreatedBy - The ID of the user that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`


	// State - Indicates if the resource is active, inactive, or deleted.
	State *string `json:"state,omitempty"`


	// ModifiedByApp - The application that last modified the resource.
	ModifiedByApp *string `json:"modifiedByApp,omitempty"`


	// CreatedByApp - The application that created the resource.
	CreatedByApp *string `json:"createdByApp,omitempty"`


	// EdgeUri
	EdgeUri *string `json:"edgeUri,omitempty"`


	// EdgeAssignedId
	EdgeAssignedId *string `json:"edgeAssignedId,omitempty"`


	// FriendlyName - Friendly Name
	FriendlyName *string `json:"friendlyName,omitempty"`


	// VlanTagId
	VlanTagId *int `json:"vlanTagId,omitempty"`


	// HardwareAddress - Hardware Address
	HardwareAddress *string `json:"hardwareAddress,omitempty"`


	// PhysicalAdapterId - Physical Adapter Id
	PhysicalAdapterId *string `json:"physicalAdapterId,omitempty"`


	// IfStatus
	IfStatus *string `json:"ifStatus,omitempty"`


	// InterfaceType - The type of this network interface.
	InterfaceType *string `json:"interfaceType,omitempty"`


	// PublicNatAddressIpV4 - IPv4 NENT IP Address
	PublicNatAddressIpV4 *string `json:"publicNatAddressIpV4,omitempty"`


	// PublicNatAddressIpV6 - IPv6 NENT IP Address
	PublicNatAddressIpV6 *string `json:"publicNatAddressIpV6,omitempty"`


	// Routes - The list of routes assigned to this interface.
	Routes *[]Domainnetworkroute `json:"routes,omitempty"`


	// Addresses - The list of IP addresses on this interface.  Priority of dns addresses are based on order in the list.
	Addresses *[]Domainnetworkaddress `json:"addresses,omitempty"`


	// Ipv4Capabilities - IPv4 interface settings.
	Ipv4Capabilities *Domaincapabilities `json:"ipv4Capabilities,omitempty"`


	// Ipv6Capabilities - IPv6 interface settings.
	Ipv6Capabilities *Domaincapabilities `json:"ipv6Capabilities,omitempty"`


	// CurrentState
	CurrentState *string `json:"currentState,omitempty"`


	// LastModifiedUserId
	LastModifiedUserId *string `json:"lastModifiedUserId,omitempty"`


	// LastModifiedCorrelationId
	LastModifiedCorrelationId *string `json:"lastModifiedCorrelationId,omitempty"`


	// CommandResponses
	CommandResponses *[]Domainnetworkcommandresponse `json:"commandResponses,omitempty"`


	// InheritPhoneTrunkBasesIPv4 - The IPv4 phone trunk base assignment will be inherited from the Edge Group.
	InheritPhoneTrunkBasesIPv4 *bool `json:"inheritPhoneTrunkBasesIPv4,omitempty"`


	// InheritPhoneTrunkBasesIPv6 - The IPv6 phone trunk base assignment will be inherited from the Edge Group.
	InheritPhoneTrunkBasesIPv6 *bool `json:"inheritPhoneTrunkBasesIPv6,omitempty"`


	// UseForInternalEdgeCommunication - This interface will be used for all internal edge-to-edge communication using settings from the edgeTrunkBaseAssignment on the Edge Group.
	UseForInternalEdgeCommunication *bool `json:"useForInternalEdgeCommunication,omitempty"`


	// UseForIndirectEdgeCommunication - Site Interconnects using the \"Indirect\" method will communicate using the Public IP Address specified on the interface. Use this option when a NAT enabled firewall is between the Edge and the far end.
	UseForIndirectEdgeCommunication *bool `json:"useForIndirectEdgeCommunication,omitempty"`


	// UseForCloudProxyEdgeCommunication - Site Interconnects using the \"Cloud Proxy\" method will broker the connection between them with a Cloud Proxy. This method is required for connections between one or more Sites using Cloud Media, but can optionally be used between two premises Sites if Direct or Indirect are not an option.
	UseForCloudProxyEdgeCommunication *bool `json:"useForCloudProxyEdgeCommunication,omitempty"`


	// UseForWanInterface - This interface will be used for all communication with the internet.
	UseForWanInterface *bool `json:"useForWanInterface,omitempty"`


	// ExternalTrunkBaseAssignments - External trunk base settings to use for external communication from this interface.
	ExternalTrunkBaseAssignments *[]Trunkbaseassignment `json:"externalTrunkBaseAssignments,omitempty"`


	// PhoneTrunkBaseAssignments - Phone trunk base settings to use for phone communication from this interface.  These settings will be ignored when \"inheritPhoneTrunkBases\" is true.
	PhoneTrunkBaseAssignments *[]Trunkbaseassignment `json:"phoneTrunkBaseAssignments,omitempty"`


	// TraceEnabled
	TraceEnabled *bool `json:"traceEnabled,omitempty"`


	// StartDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Domainlogicalinterface) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainlogicalinterface
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ModifiedByApp *string `json:"modifiedByApp,omitempty"`
		
		CreatedByApp *string `json:"createdByApp,omitempty"`
		
		EdgeUri *string `json:"edgeUri,omitempty"`
		
		EdgeAssignedId *string `json:"edgeAssignedId,omitempty"`
		
		FriendlyName *string `json:"friendlyName,omitempty"`
		
		VlanTagId *int `json:"vlanTagId,omitempty"`
		
		HardwareAddress *string `json:"hardwareAddress,omitempty"`
		
		PhysicalAdapterId *string `json:"physicalAdapterId,omitempty"`
		
		IfStatus *string `json:"ifStatus,omitempty"`
		
		InterfaceType *string `json:"interfaceType,omitempty"`
		
		PublicNatAddressIpV4 *string `json:"publicNatAddressIpV4,omitempty"`
		
		PublicNatAddressIpV6 *string `json:"publicNatAddressIpV6,omitempty"`
		
		Routes *[]Domainnetworkroute `json:"routes,omitempty"`
		
		Addresses *[]Domainnetworkaddress `json:"addresses,omitempty"`
		
		Ipv4Capabilities *Domaincapabilities `json:"ipv4Capabilities,omitempty"`
		
		Ipv6Capabilities *Domaincapabilities `json:"ipv6Capabilities,omitempty"`
		
		CurrentState *string `json:"currentState,omitempty"`
		
		LastModifiedUserId *string `json:"lastModifiedUserId,omitempty"`
		
		LastModifiedCorrelationId *string `json:"lastModifiedCorrelationId,omitempty"`
		
		CommandResponses *[]Domainnetworkcommandresponse `json:"commandResponses,omitempty"`
		
		InheritPhoneTrunkBasesIPv4 *bool `json:"inheritPhoneTrunkBasesIPv4,omitempty"`
		
		InheritPhoneTrunkBasesIPv6 *bool `json:"inheritPhoneTrunkBasesIPv6,omitempty"`
		
		UseForInternalEdgeCommunication *bool `json:"useForInternalEdgeCommunication,omitempty"`
		
		UseForIndirectEdgeCommunication *bool `json:"useForIndirectEdgeCommunication,omitempty"`
		
		UseForCloudProxyEdgeCommunication *bool `json:"useForCloudProxyEdgeCommunication,omitempty"`
		
		UseForWanInterface *bool `json:"useForWanInterface,omitempty"`
		
		ExternalTrunkBaseAssignments *[]Trunkbaseassignment `json:"externalTrunkBaseAssignments,omitempty"`
		
		PhoneTrunkBaseAssignments *[]Trunkbaseassignment `json:"phoneTrunkBaseAssignments,omitempty"`
		
		TraceEnabled *bool `json:"traceEnabled,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		State: o.State,
		
		ModifiedByApp: o.ModifiedByApp,
		
		CreatedByApp: o.CreatedByApp,
		
		EdgeUri: o.EdgeUri,
		
		EdgeAssignedId: o.EdgeAssignedId,
		
		FriendlyName: o.FriendlyName,
		
		VlanTagId: o.VlanTagId,
		
		HardwareAddress: o.HardwareAddress,
		
		PhysicalAdapterId: o.PhysicalAdapterId,
		
		IfStatus: o.IfStatus,
		
		InterfaceType: o.InterfaceType,
		
		PublicNatAddressIpV4: o.PublicNatAddressIpV4,
		
		PublicNatAddressIpV6: o.PublicNatAddressIpV6,
		
		Routes: o.Routes,
		
		Addresses: o.Addresses,
		
		Ipv4Capabilities: o.Ipv4Capabilities,
		
		Ipv6Capabilities: o.Ipv6Capabilities,
		
		CurrentState: o.CurrentState,
		
		LastModifiedUserId: o.LastModifiedUserId,
		
		LastModifiedCorrelationId: o.LastModifiedCorrelationId,
		
		CommandResponses: o.CommandResponses,
		
		InheritPhoneTrunkBasesIPv4: o.InheritPhoneTrunkBasesIPv4,
		
		InheritPhoneTrunkBasesIPv6: o.InheritPhoneTrunkBasesIPv6,
		
		UseForInternalEdgeCommunication: o.UseForInternalEdgeCommunication,
		
		UseForIndirectEdgeCommunication: o.UseForIndirectEdgeCommunication,
		
		UseForCloudProxyEdgeCommunication: o.UseForCloudProxyEdgeCommunication,
		
		UseForWanInterface: o.UseForWanInterface,
		
		ExternalTrunkBaseAssignments: o.ExternalTrunkBaseAssignments,
		
		PhoneTrunkBaseAssignments: o.PhoneTrunkBaseAssignments,
		
		TraceEnabled: o.TraceEnabled,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainlogicalinterface) UnmarshalJSON(b []byte) error {
	var DomainlogicalinterfaceMap map[string]interface{}
	err := json.Unmarshal(b, &DomainlogicalinterfaceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DomainlogicalinterfaceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DomainlogicalinterfaceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := DomainlogicalinterfaceMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := DomainlogicalinterfaceMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := DomainlogicalinterfaceMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := DomainlogicalinterfaceMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DomainlogicalinterfaceMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := DomainlogicalinterfaceMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := DomainlogicalinterfaceMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := DomainlogicalinterfaceMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := DomainlogicalinterfaceMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := DomainlogicalinterfaceMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if EdgeUri, ok := DomainlogicalinterfaceMap["edgeUri"].(string); ok {
		o.EdgeUri = &EdgeUri
	}
	
	if EdgeAssignedId, ok := DomainlogicalinterfaceMap["edgeAssignedId"].(string); ok {
		o.EdgeAssignedId = &EdgeAssignedId
	}
	
	if FriendlyName, ok := DomainlogicalinterfaceMap["friendlyName"].(string); ok {
		o.FriendlyName = &FriendlyName
	}
	
	if VlanTagId, ok := DomainlogicalinterfaceMap["vlanTagId"].(float64); ok {
		VlanTagIdInt := int(VlanTagId)
		o.VlanTagId = &VlanTagIdInt
	}
	
	if HardwareAddress, ok := DomainlogicalinterfaceMap["hardwareAddress"].(string); ok {
		o.HardwareAddress = &HardwareAddress
	}
	
	if PhysicalAdapterId, ok := DomainlogicalinterfaceMap["physicalAdapterId"].(string); ok {
		o.PhysicalAdapterId = &PhysicalAdapterId
	}
	
	if IfStatus, ok := DomainlogicalinterfaceMap["ifStatus"].(string); ok {
		o.IfStatus = &IfStatus
	}
	
	if InterfaceType, ok := DomainlogicalinterfaceMap["interfaceType"].(string); ok {
		o.InterfaceType = &InterfaceType
	}
	
	if PublicNatAddressIpV4, ok := DomainlogicalinterfaceMap["publicNatAddressIpV4"].(string); ok {
		o.PublicNatAddressIpV4 = &PublicNatAddressIpV4
	}
	
	if PublicNatAddressIpV6, ok := DomainlogicalinterfaceMap["publicNatAddressIpV6"].(string); ok {
		o.PublicNatAddressIpV6 = &PublicNatAddressIpV6
	}
	
	if Routes, ok := DomainlogicalinterfaceMap["routes"].([]interface{}); ok {
		RoutesString, _ := json.Marshal(Routes)
		json.Unmarshal(RoutesString, &o.Routes)
	}
	
	if Addresses, ok := DomainlogicalinterfaceMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if Ipv4Capabilities, ok := DomainlogicalinterfaceMap["ipv4Capabilities"].(map[string]interface{}); ok {
		Ipv4CapabilitiesString, _ := json.Marshal(Ipv4Capabilities)
		json.Unmarshal(Ipv4CapabilitiesString, &o.Ipv4Capabilities)
	}
	
	if Ipv6Capabilities, ok := DomainlogicalinterfaceMap["ipv6Capabilities"].(map[string]interface{}); ok {
		Ipv6CapabilitiesString, _ := json.Marshal(Ipv6Capabilities)
		json.Unmarshal(Ipv6CapabilitiesString, &o.Ipv6Capabilities)
	}
	
	if CurrentState, ok := DomainlogicalinterfaceMap["currentState"].(string); ok {
		o.CurrentState = &CurrentState
	}
	
	if LastModifiedUserId, ok := DomainlogicalinterfaceMap["lastModifiedUserId"].(string); ok {
		o.LastModifiedUserId = &LastModifiedUserId
	}
	
	if LastModifiedCorrelationId, ok := DomainlogicalinterfaceMap["lastModifiedCorrelationId"].(string); ok {
		o.LastModifiedCorrelationId = &LastModifiedCorrelationId
	}
	
	if CommandResponses, ok := DomainlogicalinterfaceMap["commandResponses"].([]interface{}); ok {
		CommandResponsesString, _ := json.Marshal(CommandResponses)
		json.Unmarshal(CommandResponsesString, &o.CommandResponses)
	}
	
	if InheritPhoneTrunkBasesIPv4, ok := DomainlogicalinterfaceMap["inheritPhoneTrunkBasesIPv4"].(bool); ok {
		o.InheritPhoneTrunkBasesIPv4 = &InheritPhoneTrunkBasesIPv4
	}
	
	if InheritPhoneTrunkBasesIPv6, ok := DomainlogicalinterfaceMap["inheritPhoneTrunkBasesIPv6"].(bool); ok {
		o.InheritPhoneTrunkBasesIPv6 = &InheritPhoneTrunkBasesIPv6
	}
	
	if UseForInternalEdgeCommunication, ok := DomainlogicalinterfaceMap["useForInternalEdgeCommunication"].(bool); ok {
		o.UseForInternalEdgeCommunication = &UseForInternalEdgeCommunication
	}
	
	if UseForIndirectEdgeCommunication, ok := DomainlogicalinterfaceMap["useForIndirectEdgeCommunication"].(bool); ok {
		o.UseForIndirectEdgeCommunication = &UseForIndirectEdgeCommunication
	}
	
	if UseForCloudProxyEdgeCommunication, ok := DomainlogicalinterfaceMap["useForCloudProxyEdgeCommunication"].(bool); ok {
		o.UseForCloudProxyEdgeCommunication = &UseForCloudProxyEdgeCommunication
	}
	
	if UseForWanInterface, ok := DomainlogicalinterfaceMap["useForWanInterface"].(bool); ok {
		o.UseForWanInterface = &UseForWanInterface
	}
	
	if ExternalTrunkBaseAssignments, ok := DomainlogicalinterfaceMap["externalTrunkBaseAssignments"].([]interface{}); ok {
		ExternalTrunkBaseAssignmentsString, _ := json.Marshal(ExternalTrunkBaseAssignments)
		json.Unmarshal(ExternalTrunkBaseAssignmentsString, &o.ExternalTrunkBaseAssignments)
	}
	
	if PhoneTrunkBaseAssignments, ok := DomainlogicalinterfaceMap["phoneTrunkBaseAssignments"].([]interface{}); ok {
		PhoneTrunkBaseAssignmentsString, _ := json.Marshal(PhoneTrunkBaseAssignments)
		json.Unmarshal(PhoneTrunkBaseAssignmentsString, &o.PhoneTrunkBaseAssignments)
	}
	
	if TraceEnabled, ok := DomainlogicalinterfaceMap["traceEnabled"].(bool); ok {
		o.TraceEnabled = &TraceEnabled
	}
	
	if startDateString, ok := DomainlogicalinterfaceMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := DomainlogicalinterfaceMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if SelfUri, ok := DomainlogicalinterfaceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainlogicalinterface) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
