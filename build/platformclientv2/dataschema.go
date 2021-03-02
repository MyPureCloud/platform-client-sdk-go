package platformclientv2
import (
	"time"
	"encoding/json"
)

// Dataschema
type Dataschema struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Version - The schema's version, a positive integer. Required for updates.
	Version *int `json:"version,omitempty"`


	// AppliesTo - One of \"CONTACT\" or \"EXTERNAL_ORGANIZATION\".  Indicates the built-in entity type to which this schema applies.
	AppliesTo *[]string `json:"appliesTo,omitempty"`


	// Enabled - The schema's enabled/disabled status. A disabled schema cannot be assigned to any other entities, but the data on those entities from the schema still exists.
	Enabled *bool `json:"enabled,omitempty"`


	// CreatedBy - The URI of the user that created this schema.
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// DateCreated - The date and time this schema was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// JsonSchema - A JSON schema defining the extension to the built-in entity type.
	JsonSchema *Jsonschemadocument `json:"jsonSchema,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dataschema) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
