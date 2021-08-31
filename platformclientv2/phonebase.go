package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonebase
type Phonebase struct { 
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


	// PhoneMetaBase - A phone metabase is essentially a database for storing phone configuration settings, which simplifies the configuration process.
	PhoneMetaBase *Domainentityref `json:"phoneMetaBase,omitempty"`


	// Lines - The list of linebases associated with the phone base.
	Lines *[]Linebase `json:"lines,omitempty"`


	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// Capabilities
	Capabilities *Phonecapabilities `json:"capabilities,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Phonebase) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonebase
	
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
		
		PhoneMetaBase *Domainentityref `json:"phoneMetaBase,omitempty"`
		
		Lines *[]Linebase `json:"lines,omitempty"`
		
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		Capabilities *Phonecapabilities `json:"capabilities,omitempty"`
		
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
		
		PhoneMetaBase: o.PhoneMetaBase,
		
		Lines: o.Lines,
		
		Properties: o.Properties,
		
		Capabilities: o.Capabilities,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Phonebase) UnmarshalJSON(b []byte) error {
	var PhonebaseMap map[string]interface{}
	err := json.Unmarshal(b, &PhonebaseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PhonebaseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PhonebaseMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := PhonebaseMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := PhonebaseMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := PhonebaseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := PhonebaseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := PhonebaseMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := PhonebaseMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := PhonebaseMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := PhonebaseMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := PhonebaseMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if PhoneMetaBase, ok := PhonebaseMap["phoneMetaBase"].(map[string]interface{}); ok {
		PhoneMetaBaseString, _ := json.Marshal(PhoneMetaBase)
		json.Unmarshal(PhoneMetaBaseString, &o.PhoneMetaBase)
	}
	
	if Lines, ok := PhonebaseMap["lines"].([]interface{}); ok {
		LinesString, _ := json.Marshal(Lines)
		json.Unmarshal(LinesString, &o.Lines)
	}
	
	if Properties, ok := PhonebaseMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if Capabilities, ok := PhonebaseMap["capabilities"].(map[string]interface{}); ok {
		CapabilitiesString, _ := json.Marshal(Capabilities)
		json.Unmarshal(CapabilitiesString, &o.Capabilities)
	}
	
	if SelfUri, ok := PhonebaseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Phonebase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
