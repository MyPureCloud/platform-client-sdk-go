package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// CallDrainingState - The current state of the Edge's call draining process before it can be safely rebooted or updated.
	CallDrainingState *string `json:"callDrainingState,omitempty"`


	// ConversationCount - The remaining number of conversations the Edge has to drain before it can be safely rebooted or updated. When an Edge is not draining conversations, this will be NULL or 0.
	ConversationCount *int `json:"conversationCount,omitempty"`


	// Proxy - Edge HTTP proxy configuration for the WAN port. The field can be a hostname, FQDN, IPv4 or IPv6 address. If port is not included, port 80 is assumed.
	Proxy *string `json:"proxy,omitempty"`


	// OfflineConfigCalled - True if the offline edge configuration endpoint has been called for this edge.
	OfflineConfigCalled *bool `json:"offlineConfigCalled,omitempty"`


	// OsName - The name provided by the operating system of the Edge.
	OsName *string `json:"osName,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Edge) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edge

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ModifiedByApp *string `json:"modifiedByApp,omitempty"`
		
		CreatedByApp *string `json:"createdByApp,omitempty"`
		
		Interfaces *[]Edgeinterface `json:"interfaces,omitempty"`
		
		Make *string `json:"make,omitempty"`
		
		Model *string `json:"model,omitempty"`
		
		ApiVersion *string `json:"apiVersion,omitempty"`
		
		SoftwareVersion *string `json:"softwareVersion,omitempty"`
		
		SoftwareVersionTimestamp *string `json:"softwareVersionTimestamp,omitempty"`
		
		SoftwareVersionPlatform *string `json:"softwareVersionPlatform,omitempty"`
		
		SoftwareVersionConfiguration *string `json:"softwareVersionConfiguration,omitempty"`
		
		FullSoftwareVersion *string `json:"fullSoftwareVersion,omitempty"`
		
		PairingId *string `json:"pairingId,omitempty"`
		
		Fingerprint *string `json:"fingerprint,omitempty"`
		
		FingerprintHint *string `json:"fingerprintHint,omitempty"`
		
		CurrentVersion *string `json:"currentVersion,omitempty"`
		
		StagedVersion *string `json:"stagedVersion,omitempty"`
		
		Patch *string `json:"patch,omitempty"`
		
		StatusCode *string `json:"statusCode,omitempty"`
		
		EdgeGroup *Edgegroup `json:"edgeGroup,omitempty"`
		
		Site *Site `json:"site,omitempty"`
		
		SoftwareStatus *Domainedgesoftwareupdatedto `json:"softwareStatus,omitempty"`
		
		OnlineStatus *string `json:"onlineStatus,omitempty"`
		
		SerialNumber *string `json:"serialNumber,omitempty"`
		
		PhysicalEdge *bool `json:"physicalEdge,omitempty"`
		
		Managed *bool `json:"managed,omitempty"`
		
		EdgeDeploymentType *string `json:"edgeDeploymentType,omitempty"`
		
		CallDrainingState *string `json:"callDrainingState,omitempty"`
		
		ConversationCount *int `json:"conversationCount,omitempty"`
		
		Proxy *string `json:"proxy,omitempty"`
		
		OfflineConfigCalled *bool `json:"offlineConfigCalled,omitempty"`
		
		OsName *string `json:"osName,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Version: u.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: u.ModifiedBy,
		
		CreatedBy: u.CreatedBy,
		
		State: u.State,
		
		ModifiedByApp: u.ModifiedByApp,
		
		CreatedByApp: u.CreatedByApp,
		
		Interfaces: u.Interfaces,
		
		Make: u.Make,
		
		Model: u.Model,
		
		ApiVersion: u.ApiVersion,
		
		SoftwareVersion: u.SoftwareVersion,
		
		SoftwareVersionTimestamp: u.SoftwareVersionTimestamp,
		
		SoftwareVersionPlatform: u.SoftwareVersionPlatform,
		
		SoftwareVersionConfiguration: u.SoftwareVersionConfiguration,
		
		FullSoftwareVersion: u.FullSoftwareVersion,
		
		PairingId: u.PairingId,
		
		Fingerprint: u.Fingerprint,
		
		FingerprintHint: u.FingerprintHint,
		
		CurrentVersion: u.CurrentVersion,
		
		StagedVersion: u.StagedVersion,
		
		Patch: u.Patch,
		
		StatusCode: u.StatusCode,
		
		EdgeGroup: u.EdgeGroup,
		
		Site: u.Site,
		
		SoftwareStatus: u.SoftwareStatus,
		
		OnlineStatus: u.OnlineStatus,
		
		SerialNumber: u.SerialNumber,
		
		PhysicalEdge: u.PhysicalEdge,
		
		Managed: u.Managed,
		
		EdgeDeploymentType: u.EdgeDeploymentType,
		
		CallDrainingState: u.CallDrainingState,
		
		ConversationCount: u.ConversationCount,
		
		Proxy: u.Proxy,
		
		OfflineConfigCalled: u.OfflineConfigCalled,
		
		OsName: u.OsName,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
