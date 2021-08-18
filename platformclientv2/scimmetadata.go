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

func (u *Scimmetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimmetadata

	
	LastModified := new(string)
	if u.LastModified != nil {
		
		*LastModified = timeutil.Strftime(u.LastModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		ResourceType: u.ResourceType,
		
		LastModified: LastModified,
		
		Location: u.Location,
		
		Version: u.Version,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
