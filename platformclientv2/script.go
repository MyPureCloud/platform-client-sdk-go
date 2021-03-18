package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Script
type Script struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VersionId
	VersionId *string `json:"versionId,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// PublishedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PublishedDate *time.Time `json:"publishedDate,omitempty"`


	// VersionDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
