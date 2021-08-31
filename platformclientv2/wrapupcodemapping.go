package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wrapupcodemapping
type Wrapupcodemapping struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// DefaultSet - The default set of wrap-up flags. These will be used if there is no entry for a given wrap-up code in the mapping.
	DefaultSet *[]string `json:"defaultSet,omitempty"`


	// Mapping - A map from wrap-up code identifiers to a set of wrap-up flags.
	Mapping *map[string][]string `json:"mapping,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Wrapupcodemapping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wrapupcodemapping
	
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
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		DefaultSet: o.DefaultSet,
		
		Mapping: o.Mapping,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Wrapupcodemapping) UnmarshalJSON(b []byte) error {
	var WrapupcodemappingMap map[string]interface{}
	err := json.Unmarshal(b, &WrapupcodemappingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WrapupcodemappingMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := WrapupcodemappingMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := WrapupcodemappingMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := WrapupcodemappingMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := WrapupcodemappingMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if DefaultSet, ok := WrapupcodemappingMap["defaultSet"].([]interface{}); ok {
		DefaultSetString, _ := json.Marshal(DefaultSet)
		json.Unmarshal(DefaultSetString, &o.DefaultSet)
	}
	
	if Mapping, ok := WrapupcodemappingMap["mapping"].(map[string]interface{}); ok {
		MappingString, _ := json.Marshal(Mapping)
		json.Unmarshal(MappingString, &o.Mapping)
	}
	
	if SelfUri, ok := WrapupcodemappingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wrapupcodemapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
