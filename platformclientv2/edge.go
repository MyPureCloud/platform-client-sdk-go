package platformclientv2
import (
	"time"
	"encoding/json"
)

// Edge
type Edge struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// Description - The resource's description.
	Description *string `json:"description,omitempty"`


	// Version - The current version of the resource.
	Version *int32 `json:"version,omitempty"`


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


	// Interfaces - The list of interfaces for the edge. (Deprecated) Replaced by configuring trunks/ip info on the logical interface instead
	Interfaces *[]Edgeinterface `json:"interfaces,omitempty"`


	// Make
	Make *string `json:"make,omitempty"`


	// Model
	Model *string `json:"model,omitempty"`


	// ApiVersion
	ApiVersion *string `json:"apiVersion,omitempty"`


	// SoftwareVersion
	SoftwareVersion *string `json:"softwareVersion,omitempty"`


	// SoftwareVersionTimestamp
	SoftwareVersionTimestamp *string `json:"softwareVersionTimestamp,omitempty"`


	// SoftwareVersionPlatform
	SoftwareVersionPlatform *string `json:"softwareVersionPlatform,omitempty"`


	// SoftwareVersionConfiguration
	SoftwareVersionConfiguration *string `json:"softwareVersionConfiguration,omitempty"`


	// FullSoftwareVersion
	FullSoftwareVersion *string `json:"fullSoftwareVersion,omitempty"`


	// PairingId - The pairing Id for a hardware Edge in the format: 00000-00000-00000-00000-00000. This field is only required when creating an Edge with a deployment type of HARDWARE.
	PairingId *string `json:"pairingId,omitempty"`


	// Fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`


	// FingerprintHint
	FingerprintHint *string `json:"fingerprintHint,omitempty"`


	// CurrentVersion
	CurrentVersion *string `json:"currentVersion,omitempty"`


	// StagedVersion
	StagedVersion *string `json:"stagedVersion,omitempty"`


	// Patch
	Patch *string `json:"patch,omitempty"`


	// StatusCode - The current status of the Edge.
	StatusCode *string `json:"statusCode,omitempty"`


	// EdgeGroup
	EdgeGroup *Edgegroup `json:"edgeGroup,omitempty"`


	// Site - The Site to which the Edge is assigned.
	Site *Site `json:"site,omitempty"`


	// SoftwareStatus - Details about an in-progress or recently in-progress Edge software upgrade. This node appears only if a software upgrade was recently initiated for this Edge.
	SoftwareStatus *Domainedgesoftwareupdatedto `json:"softwareStatus,omitempty"`


	// OnlineStatus
	OnlineStatus *string `json:"onlineStatus,omitempty"`


	// SerialNumber
	SerialNumber *string `json:"serialNumber,omitempty"`


	// PhysicalEdge
	PhysicalEdge *bool `json:"physicalEdge,omitempty"`


	// Managed
	Managed *bool `json:"managed,omitempty"`


	// EdgeDeploymentType
	EdgeDeploymentType *string `json:"edgeDeploymentType,omitempty"`


	// CallDrainingState
	CallDrainingState *string `json:"callDrainingState,omitempty"`


	// ConversationCount
	ConversationCount *int32 `json:"conversationCount,omitempty"`


	// Proxy - Edge HTTP proxy configuration for the WAN port. The field can be a hostname, FQDN, IPv4 or IPv6 address. If port is not included, port 80 is assumed.
	Proxy *string `json:"proxy,omitempty"`


	// OfflineConfigCalled - True if the offline edge configuration endpoint has been called for this edge.
	OfflineConfigCalled *bool `json:"offlineConfigCalled,omitempty"`


	// OsName - The name provided by the operating system of the Edge.
	OsName *string `json:"osName,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edge) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
