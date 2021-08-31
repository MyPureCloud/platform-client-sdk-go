package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainphysicalinterface
type Domainphysicalinterface struct { 
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


	// FriendlyName
	FriendlyName *string `json:"friendlyName,omitempty"`


	// HardwareAddress
	HardwareAddress *string `json:"hardwareAddress,omitempty"`


	// PortLabel
	PortLabel *string `json:"portLabel,omitempty"`


	// PhysicalCapabilities
	PhysicalCapabilities *Domainphysicalcapabilities `json:"physicalCapabilities,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Domainphysicalinterface) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainphysicalinterface
	
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
		
		EdgeUri *string `json:"edgeUri,omitempty"`
		
		FriendlyName *string `json:"friendlyName,omitempty"`
		
		HardwareAddress *string `json:"hardwareAddress,omitempty"`
		
		PortLabel *string `json:"portLabel,omitempty"`
		
		PhysicalCapabilities *Domainphysicalcapabilities `json:"physicalCapabilities,omitempty"`
		
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
		
		EdgeUri: o.EdgeUri,
		
		FriendlyName: o.FriendlyName,
		
		HardwareAddress: o.HardwareAddress,
		
		PortLabel: o.PortLabel,
		
		PhysicalCapabilities: o.PhysicalCapabilities,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainphysicalinterface) UnmarshalJSON(b []byte) error {
	var DomainphysicalinterfaceMap map[string]interface{}
	err := json.Unmarshal(b, &DomainphysicalinterfaceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DomainphysicalinterfaceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DomainphysicalinterfaceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := DomainphysicalinterfaceMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := DomainphysicalinterfaceMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := DomainphysicalinterfaceMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DomainphysicalinterfaceMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := DomainphysicalinterfaceMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := DomainphysicalinterfaceMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := DomainphysicalinterfaceMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := DomainphysicalinterfaceMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := DomainphysicalinterfaceMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if EdgeUri, ok := DomainphysicalinterfaceMap["edgeUri"].(string); ok {
		o.EdgeUri = &EdgeUri
	}
	
	if FriendlyName, ok := DomainphysicalinterfaceMap["friendlyName"].(string); ok {
		o.FriendlyName = &FriendlyName
	}
	
	if HardwareAddress, ok := DomainphysicalinterfaceMap["hardwareAddress"].(string); ok {
		o.HardwareAddress = &HardwareAddress
	}
	
	if PortLabel, ok := DomainphysicalinterfaceMap["portLabel"].(string); ok {
		o.PortLabel = &PortLabel
	}
	
	if PhysicalCapabilities, ok := DomainphysicalinterfaceMap["physicalCapabilities"].(map[string]interface{}); ok {
		PhysicalCapabilitiesString, _ := json.Marshal(PhysicalCapabilities)
		json.Unmarshal(PhysicalCapabilitiesString, &o.PhysicalCapabilities)
	}
	
	if SelfUri, ok := DomainphysicalinterfaceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainphysicalinterface) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
