package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimmetadata - Defines the SCIM metadata.
type Scimmetadata struct { 
	// ResourceType - The type of SCIM resource.
	ResourceType *string `json:"resourceType,omitempty"`


	// LastModified - The last time that the resource was modified. Date time is represented as an \"ISO-8601 string\", for example, yyyy-MM-ddTHH:mm:ss.SSSZ. Not included with \"Schema\" and \"ResourceType\" resources.
	LastModified *time.Time `json:"lastModified,omitempty"`


	// Location - The URI of the resource.
	Location *string `json:"location,omitempty"`


	// Version - The version of the resource. Matches the ETag HTTP response header. Not included with \"Schema\" and \"ResourceType\" resources.
	Version *string `json:"version,omitempty"`

}

func (o *Scimmetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimmetadata
	
	LastModified := new(string)
	if o.LastModified != nil {
		
		*LastModified = timeutil.Strftime(o.LastModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastModified = nil
	}
	
	return json.Marshal(&struct { 
		ResourceType *string `json:"resourceType,omitempty"`
		
		LastModified *string `json:"lastModified,omitempty"`
		
		Location *string `json:"location,omitempty"`
		
		Version *string `json:"version,omitempty"`
		*Alias
	}{ 
		ResourceType: o.ResourceType,
		
		LastModified: LastModified,
		
		Location: o.Location,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimmetadata) UnmarshalJSON(b []byte) error {
	var ScimmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ScimmetadataMap)
	if err != nil {
		return err
	}
	
	if ResourceType, ok := ScimmetadataMap["resourceType"].(string); ok {
		o.ResourceType = &ResourceType
	}
	
	if lastModifiedString, ok := ScimmetadataMap["lastModified"].(string); ok {
		LastModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", lastModifiedString)
		o.LastModified = &LastModified
	}
	
	if Location, ok := ScimmetadataMap["location"].(string); ok {
		o.Location = &Location
	}
	
	if Version, ok := ScimmetadataMap["version"].(string); ok {
		o.Version = &Version
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
