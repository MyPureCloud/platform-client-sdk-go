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

func (u *Trunk) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunk

	
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
		
		TrunkType: u.TrunkType,
		
		Edge: u.Edge,
		
		TrunkBase: u.TrunkBase,
		
		TrunkMetabase: u.TrunkMetabase,
		
		EdgeGroup: u.EdgeGroup,
		
		InService: u.InService,
		
		Enabled: u.Enabled,
		
		LogicalInterface: u.LogicalInterface,
		
		ConnectedStatus: u.ConnectedStatus,
		
		OptionsStatus: u.OptionsStatus,
		
		RegistersStatus: u.RegistersStatus,
		
		IpStatus: u.IpStatus,
		
		OptionsEnabledStatus: u.OptionsEnabledStatus,
		
		RegistersEnabledStatus: u.RegistersEnabledStatus,
		
		Family: u.Family,
		
		ProxyAddressList: u.ProxyAddressList,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunk) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
