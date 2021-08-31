package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgetrunkbase
type Edgetrunkbase struct { 
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


	// TrunkMetabase - The meta-base this trunk is based on.
	TrunkMetabase *Domainentityref `json:"trunkMetabase,omitempty"`


	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// TrunkType - The type of this trunk base.
	TrunkType *string `json:"trunkType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Edgetrunkbase) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgetrunkbase
	
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
		
		TrunkMetabase *Domainentityref `json:"trunkMetabase,omitempty"`
		
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		TrunkType *string `json:"trunkType,omitempty"`
		
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
		
		TrunkMetabase: o.TrunkMetabase,
		
		Properties: o.Properties,
		
		TrunkType: o.TrunkType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgetrunkbase) UnmarshalJSON(b []byte) error {
	var EdgetrunkbaseMap map[string]interface{}
	err := json.Unmarshal(b, &EdgetrunkbaseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgetrunkbaseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := EdgetrunkbaseMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := EdgetrunkbaseMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := EdgetrunkbaseMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := EdgetrunkbaseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := EdgetrunkbaseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := EdgetrunkbaseMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := EdgetrunkbaseMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := EdgetrunkbaseMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := EdgetrunkbaseMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := EdgetrunkbaseMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if TrunkMetabase, ok := EdgetrunkbaseMap["trunkMetabase"].(map[string]interface{}); ok {
		TrunkMetabaseString, _ := json.Marshal(TrunkMetabase)
		json.Unmarshal(TrunkMetabaseString, &o.TrunkMetabase)
	}
	
	if Properties, ok := EdgetrunkbaseMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if TrunkType, ok := EdgetrunkbaseMap["trunkType"].(string); ok {
		o.TrunkType = &TrunkType
	}
	
	if SelfUri, ok := EdgetrunkbaseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgetrunkbase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
