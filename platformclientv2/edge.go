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

func (o *Edge) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edge
	
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
		
		Interfaces: o.Interfaces,
		
		Make: o.Make,
		
		Model: o.Model,
		
		ApiVersion: o.ApiVersion,
		
		SoftwareVersion: o.SoftwareVersion,
		
		SoftwareVersionTimestamp: o.SoftwareVersionTimestamp,
		
		SoftwareVersionPlatform: o.SoftwareVersionPlatform,
		
		SoftwareVersionConfiguration: o.SoftwareVersionConfiguration,
		
		FullSoftwareVersion: o.FullSoftwareVersion,
		
		PairingId: o.PairingId,
		
		Fingerprint: o.Fingerprint,
		
		FingerprintHint: o.FingerprintHint,
		
		CurrentVersion: o.CurrentVersion,
		
		StagedVersion: o.StagedVersion,
		
		Patch: o.Patch,
		
		StatusCode: o.StatusCode,
		
		EdgeGroup: o.EdgeGroup,
		
		Site: o.Site,
		
		SoftwareStatus: o.SoftwareStatus,
		
		OnlineStatus: o.OnlineStatus,
		
		SerialNumber: o.SerialNumber,
		
		PhysicalEdge: o.PhysicalEdge,
		
		Managed: o.Managed,
		
		EdgeDeploymentType: o.EdgeDeploymentType,
		
		CallDrainingState: o.CallDrainingState,
		
		ConversationCount: o.ConversationCount,
		
		Proxy: o.Proxy,
		
		OfflineConfigCalled: o.OfflineConfigCalled,
		
		OsName: o.OsName,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Edge) UnmarshalJSON(b []byte) error {
	var EdgeMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EdgeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := EdgeMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := EdgeMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Version, ok := EdgeMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := EdgeMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := EdgeMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := EdgeMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    
	if CreatedBy, ok := EdgeMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
    
	if State, ok := EdgeMap["state"].(string); ok {
		o.State = &State
	}
    
	if ModifiedByApp, ok := EdgeMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
    
	if CreatedByApp, ok := EdgeMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
    
	if Interfaces, ok := EdgeMap["interfaces"].([]interface{}); ok {
		InterfacesString, _ := json.Marshal(Interfaces)
		json.Unmarshal(InterfacesString, &o.Interfaces)
	}
	
	if Make, ok := EdgeMap["make"].(string); ok {
		o.Make = &Make
	}
    
	if Model, ok := EdgeMap["model"].(string); ok {
		o.Model = &Model
	}
    
	if ApiVersion, ok := EdgeMap["apiVersion"].(string); ok {
		o.ApiVersion = &ApiVersion
	}
    
	if SoftwareVersion, ok := EdgeMap["softwareVersion"].(string); ok {
		o.SoftwareVersion = &SoftwareVersion
	}
    
	if SoftwareVersionTimestamp, ok := EdgeMap["softwareVersionTimestamp"].(string); ok {
		o.SoftwareVersionTimestamp = &SoftwareVersionTimestamp
	}
    
	if SoftwareVersionPlatform, ok := EdgeMap["softwareVersionPlatform"].(string); ok {
		o.SoftwareVersionPlatform = &SoftwareVersionPlatform
	}
    
	if SoftwareVersionConfiguration, ok := EdgeMap["softwareVersionConfiguration"].(string); ok {
		o.SoftwareVersionConfiguration = &SoftwareVersionConfiguration
	}
    
	if FullSoftwareVersion, ok := EdgeMap["fullSoftwareVersion"].(string); ok {
		o.FullSoftwareVersion = &FullSoftwareVersion
	}
    
	if PairingId, ok := EdgeMap["pairingId"].(string); ok {
		o.PairingId = &PairingId
	}
    
	if Fingerprint, ok := EdgeMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
    
	if FingerprintHint, ok := EdgeMap["fingerprintHint"].(string); ok {
		o.FingerprintHint = &FingerprintHint
	}
    
	if CurrentVersion, ok := EdgeMap["currentVersion"].(string); ok {
		o.CurrentVersion = &CurrentVersion
	}
    
	if StagedVersion, ok := EdgeMap["stagedVersion"].(string); ok {
		o.StagedVersion = &StagedVersion
	}
    
	if Patch, ok := EdgeMap["patch"].(string); ok {
		o.Patch = &Patch
	}
    
	if StatusCode, ok := EdgeMap["statusCode"].(string); ok {
		o.StatusCode = &StatusCode
	}
    
	if EdgeGroup, ok := EdgeMap["edgeGroup"].(map[string]interface{}); ok {
		EdgeGroupString, _ := json.Marshal(EdgeGroup)
		json.Unmarshal(EdgeGroupString, &o.EdgeGroup)
	}
	
	if Site, ok := EdgeMap["site"].(map[string]interface{}); ok {
		SiteString, _ := json.Marshal(Site)
		json.Unmarshal(SiteString, &o.Site)
	}
	
	if SoftwareStatus, ok := EdgeMap["softwareStatus"].(map[string]interface{}); ok {
		SoftwareStatusString, _ := json.Marshal(SoftwareStatus)
		json.Unmarshal(SoftwareStatusString, &o.SoftwareStatus)
	}
	
	if OnlineStatus, ok := EdgeMap["onlineStatus"].(string); ok {
		o.OnlineStatus = &OnlineStatus
	}
    
	if SerialNumber, ok := EdgeMap["serialNumber"].(string); ok {
		o.SerialNumber = &SerialNumber
	}
    
	if PhysicalEdge, ok := EdgeMap["physicalEdge"].(bool); ok {
		o.PhysicalEdge = &PhysicalEdge
	}
    
	if Managed, ok := EdgeMap["managed"].(bool); ok {
		o.Managed = &Managed
	}
    
	if EdgeDeploymentType, ok := EdgeMap["edgeDeploymentType"].(string); ok {
		o.EdgeDeploymentType = &EdgeDeploymentType
	}
    
	if CallDrainingState, ok := EdgeMap["callDrainingState"].(string); ok {
		o.CallDrainingState = &CallDrainingState
	}
    
	if ConversationCount, ok := EdgeMap["conversationCount"].(float64); ok {
		ConversationCountInt := int(ConversationCount)
		o.ConversationCount = &ConversationCountInt
	}
	
	if Proxy, ok := EdgeMap["proxy"].(string); ok {
		o.Proxy = &Proxy
	}
    
	if OfflineConfigCalled, ok := EdgeMap["offlineConfigCalled"].(bool); ok {
		o.OfflineConfigCalled = &OfflineConfigCalled
	}
    
	if OsName, ok := EdgeMap["osName"].(string); ok {
		o.OsName = &OsName
	}
    
	if SelfUri, ok := EdgeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
