package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerwrapupcodemappingconfigchangewrapupcodemapping
type Dialerwrapupcodemappingconfigchangewrapupcodemapping struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// DefaultSet
	DefaultSet *[]string `json:"defaultSet,omitempty"`


	// Mapping
	Mapping *map[string][]string `json:"mapping,omitempty"`

}

func (o *Dialerwrapupcodemappingconfigchangewrapupcodemapping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerwrapupcodemappingconfigchangewrapupcodemapping
	
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
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DefaultSet *[]string `json:"defaultSet,omitempty"`
		
		Mapping *map[string][]string `json:"mapping,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		DefaultSet: o.DefaultSet,
		
		Mapping: o.Mapping,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerwrapupcodemappingconfigchangewrapupcodemapping) UnmarshalJSON(b []byte) error {
	var DialerwrapupcodemappingconfigchangewrapupcodemappingMap map[string]interface{}
	err := json.Unmarshal(b, &DialerwrapupcodemappingconfigchangewrapupcodemappingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialerwrapupcodemappingconfigchangewrapupcodemappingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialerwrapupcodemappingconfigchangewrapupcodemappingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DialerwrapupcodemappingconfigchangewrapupcodemappingMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialerwrapupcodemappingconfigchangewrapupcodemappingMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialerwrapupcodemappingconfigchangewrapupcodemappingMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if DefaultSet, ok := DialerwrapupcodemappingconfigchangewrapupcodemappingMap["defaultSet"].([]interface{}); ok {
		DefaultSetString, _ := json.Marshal(DefaultSet)
		json.Unmarshal(DefaultSetString, &o.DefaultSet)
	}
	
	if Mapping, ok := DialerwrapupcodemappingconfigchangewrapupcodemappingMap["mapping"].(map[string]interface{}); ok {
		MappingString, _ := json.Marshal(Mapping)
		json.Unmarshal(MappingString, &o.Mapping)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerwrapupcodemappingconfigchangewrapupcodemapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
