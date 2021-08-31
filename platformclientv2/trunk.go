package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunk
type Trunk struct { 
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


	// TrunkType - The type of this trunk.
	TrunkType *string `json:"trunkType,omitempty"`


	// Edge - The Edge using this trunk.
	Edge *Domainentityref `json:"edge,omitempty"`


	// TrunkBase - The trunk base configuration used on this trunk.
	TrunkBase *Domainentityref `json:"trunkBase,omitempty"`


	// TrunkMetabase - The metabase used to create this trunk.
	TrunkMetabase *Domainentityref `json:"trunkMetabase,omitempty"`


	// EdgeGroup - The edge group associated with this trunk.
	EdgeGroup *Domainentityref `json:"edgeGroup,omitempty"`


	// InService - True if this trunk is in-service.  This comes from the trunk_enabled property of the referenced trunk base.
	InService *bool `json:"inService,omitempty"`


	// Enabled - True if the Edge used by this trunk is in-service
	Enabled *bool `json:"enabled,omitempty"`


	// LogicalInterface - The Logical Interface on the Edge to which the trunk is assigned.
	LogicalInterface *Domainentityref `json:"logicalInterface,omitempty"`


	// ConnectedStatus - The connected status of the trunk
	ConnectedStatus *Trunkconnectedstatus `json:"connectedStatus,omitempty"`


	// OptionsStatus - The trunk optionsStatus
	OptionsStatus *[]Trunkmetricsoptions `json:"optionsStatus,omitempty"`


	// RegistersStatus - The trunk registersStatus
	RegistersStatus *[]Trunkmetricsregisters `json:"registersStatus,omitempty"`


	// IpStatus - The trunk ipStatus
	IpStatus *Trunkmetricsnetworktypeip `json:"ipStatus,omitempty"`


	// OptionsEnabledStatus - Returns Enabled when the trunk base supports the availability interval and it has a value greater than 0.
	OptionsEnabledStatus *string `json:"optionsEnabledStatus,omitempty"`


	// RegistersEnabledStatus - Returns Enabled when the trunk base supports the registration interval and it has a value greater than 0.
	RegistersEnabledStatus *string `json:"registersEnabledStatus,omitempty"`


	// Family - The IP Network Family of the trunk
	Family *int `json:"family,omitempty"`


	// ProxyAddressList - The list of proxy addresses (ports if provided) for the trunk
	ProxyAddressList *[]string `json:"proxyAddressList,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Trunk) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunk
	
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
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ModifiedByApp *string `json:"modifiedByApp,omitempty"`
		
		CreatedByApp *string `json:"createdByApp,omitempty"`
		
		TrunkType *string `json:"trunkType,omitempty"`
		
		Edge *Domainentityref `json:"edge,omitempty"`
		
		TrunkBase *Domainentityref `json:"trunkBase,omitempty"`
		
		TrunkMetabase *Domainentityref `json:"trunkMetabase,omitempty"`
		
		EdgeGroup *Domainentityref `json:"edgeGroup,omitempty"`
		
		InService *bool `json:"inService,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		LogicalInterface *Domainentityref `json:"logicalInterface,omitempty"`
		
		ConnectedStatus *Trunkconnectedstatus `json:"connectedStatus,omitempty"`
		
		OptionsStatus *[]Trunkmetricsoptions `json:"optionsStatus,omitempty"`
		
		RegistersStatus *[]Trunkmetricsregisters `json:"registersStatus,omitempty"`
		
		IpStatus *Trunkmetricsnetworktypeip `json:"ipStatus,omitempty"`
		
		OptionsEnabledStatus *string `json:"optionsEnabledStatus,omitempty"`
		
		RegistersEnabledStatus *string `json:"registersEnabledStatus,omitempty"`
		
		Family *int `json:"family,omitempty"`
		
		ProxyAddressList *[]string `json:"proxyAddressList,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		State: o.State,
		
		ModifiedByApp: o.ModifiedByApp,
		
		CreatedByApp: o.CreatedByApp,
		
		TrunkType: o.TrunkType,
		
		Edge: o.Edge,
		
		TrunkBase: o.TrunkBase,
		
		TrunkMetabase: o.TrunkMetabase,
		
		EdgeGroup: o.EdgeGroup,
		
		InService: o.InService,
		
		Enabled: o.Enabled,
		
		LogicalInterface: o.LogicalInterface,
		
		ConnectedStatus: o.ConnectedStatus,
		
		OptionsStatus: o.OptionsStatus,
		
		RegistersStatus: o.RegistersStatus,
		
		IpStatus: o.IpStatus,
		
		OptionsEnabledStatus: o.OptionsEnabledStatus,
		
		RegistersEnabledStatus: o.RegistersEnabledStatus,
		
		Family: o.Family,
		
		ProxyAddressList: o.ProxyAddressList,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunk) UnmarshalJSON(b []byte) error {
	var TrunkMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TrunkMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := TrunkMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := TrunkMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := TrunkMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := TrunkMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := TrunkMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := TrunkMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := TrunkMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := TrunkMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := TrunkMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := TrunkMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if TrunkType, ok := TrunkMap["trunkType"].(string); ok {
		o.TrunkType = &TrunkType
	}
	
	if Edge, ok := TrunkMap["edge"].(map[string]interface{}); ok {
		EdgeString, _ := json.Marshal(Edge)
		json.Unmarshal(EdgeString, &o.Edge)
	}
	
	if TrunkBase, ok := TrunkMap["trunkBase"].(map[string]interface{}); ok {
		TrunkBaseString, _ := json.Marshal(TrunkBase)
		json.Unmarshal(TrunkBaseString, &o.TrunkBase)
	}
	
	if TrunkMetabase, ok := TrunkMap["trunkMetabase"].(map[string]interface{}); ok {
		TrunkMetabaseString, _ := json.Marshal(TrunkMetabase)
		json.Unmarshal(TrunkMetabaseString, &o.TrunkMetabase)
	}
	
	if EdgeGroup, ok := TrunkMap["edgeGroup"].(map[string]interface{}); ok {
		EdgeGroupString, _ := json.Marshal(EdgeGroup)
		json.Unmarshal(EdgeGroupString, &o.EdgeGroup)
	}
	
	if InService, ok := TrunkMap["inService"].(bool); ok {
		o.InService = &InService
	}
	
	if Enabled, ok := TrunkMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if LogicalInterface, ok := TrunkMap["logicalInterface"].(map[string]interface{}); ok {
		LogicalInterfaceString, _ := json.Marshal(LogicalInterface)
		json.Unmarshal(LogicalInterfaceString, &o.LogicalInterface)
	}
	
	if ConnectedStatus, ok := TrunkMap["connectedStatus"].(map[string]interface{}); ok {
		ConnectedStatusString, _ := json.Marshal(ConnectedStatus)
		json.Unmarshal(ConnectedStatusString, &o.ConnectedStatus)
	}
	
	if OptionsStatus, ok := TrunkMap["optionsStatus"].([]interface{}); ok {
		OptionsStatusString, _ := json.Marshal(OptionsStatus)
		json.Unmarshal(OptionsStatusString, &o.OptionsStatus)
	}
	
	if RegistersStatus, ok := TrunkMap["registersStatus"].([]interface{}); ok {
		RegistersStatusString, _ := json.Marshal(RegistersStatus)
		json.Unmarshal(RegistersStatusString, &o.RegistersStatus)
	}
	
	if IpStatus, ok := TrunkMap["ipStatus"].(map[string]interface{}); ok {
		IpStatusString, _ := json.Marshal(IpStatus)
		json.Unmarshal(IpStatusString, &o.IpStatus)
	}
	
	if OptionsEnabledStatus, ok := TrunkMap["optionsEnabledStatus"].(string); ok {
		o.OptionsEnabledStatus = &OptionsEnabledStatus
	}
	
	if RegistersEnabledStatus, ok := TrunkMap["registersEnabledStatus"].(string); ok {
		o.RegistersEnabledStatus = &RegistersEnabledStatus
	}
	
	if Family, ok := TrunkMap["family"].(float64); ok {
		FamilyInt := int(Family)
		o.Family = &FamilyInt
	}
	
	if ProxyAddressList, ok := TrunkMap["proxyAddressList"].([]interface{}); ok {
		ProxyAddressListString, _ := json.Marshal(ProxyAddressList)
		json.Unmarshal(ProxyAddressListString, &o.ProxyAddressList)
	}
	
	if SelfUri, ok := TrunkMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunk) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
