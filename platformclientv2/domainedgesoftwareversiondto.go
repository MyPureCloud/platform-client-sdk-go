package platformclientv2
import (
	"time"
	"encoding/json"
)

// Domainedgesoftwareversiondto
type Domainedgesoftwareversiondto struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// EdgeVersion
	EdgeVersion *string `json:"edgeVersion,omitempty"`


	// PublishDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PublishDate *time.Time `json:"publishDate,omitempty"`


	// EdgeUri
	EdgeUri *string `json:"edgeUri,omitempty"`


	// LatestRelease
	LatestRelease *bool `json:"latestRelease,omitempty"`


	// Current
	Current *bool `json:"current,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Domainedgesoftwareversiondto) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
