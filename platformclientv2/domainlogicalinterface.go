package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Domainlogicalinterface) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
