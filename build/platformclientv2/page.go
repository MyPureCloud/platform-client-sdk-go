package platformclientv2
import (
	"time"
	"encoding/json"
)

// Page
type Page struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VersionId
	VersionId *string `json:"versionId,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// RootContainer
	RootContainer *map[string]interface{} `json:"rootContainer,omitempty"`


	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Page) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
