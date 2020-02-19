package platformclientv2
import (
	"time"
	"encoding/json"
)

// Script
type Script struct { 
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


	// PublishedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	PublishedDate *time.Time `json:"publishedDate,omitempty"`


	// VersionDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	VersionDate *time.Time `json:"versionDate,omitempty"`


	// StartPageId
	StartPageId *string `json:"startPageId,omitempty"`


	// StartPageName
	StartPageName *string `json:"startPageName,omitempty"`


	// Features
	Features *map[string]interface{} `json:"features,omitempty"`


	// Variables
	Variables *map[string]interface{} `json:"variables,omitempty"`


	// CustomActions
	CustomActions *map[string]interface{} `json:"customActions,omitempty"`


	// Pages
	Pages *[]Page `json:"pages,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Script) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
