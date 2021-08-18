package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Dataschema) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dataschema

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		AppliesTo *[]string `json:"appliesTo,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		JsonSchema *Jsonschemadocument `json:"jsonSchema,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Version: u.Version,
		
		AppliesTo: u.AppliesTo,
		
		Enabled: u.Enabled,
		
		CreatedBy: u.CreatedBy,
		
		DateCreated: DateCreated,
		
		JsonSchema: u.JsonSchema,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dataschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
