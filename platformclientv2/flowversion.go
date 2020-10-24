package platformclientv2
import (
	"encoding/json"
)

// Flowversion
type Flowversion struct { 
	// Id - The flow version identifier
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// CommitVersion
	CommitVersion *string `json:"commitVersion,omitempty"`


	// ConfigurationVersion
	ConfigurationVersion *string `json:"configurationVersion,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Secure
	Secure *bool `json:"secure,omitempty"`


	// Debug
	Debug *bool `json:"debug,omitempty"`


	// CreatedBy
	CreatedBy *User `json:"createdBy,omitempty"`


	// CreatedByClient
	CreatedByClient *Domainentityref `json:"createdByClient,omitempty"`


	// ConfigurationUri
	ConfigurationUri *string `json:"configurationUri,omitempty"`


	// DateCreated
	DateCreated *int64 `json:"dateCreated,omitempty"`


	// GenerationId
	GenerationId *string `json:"generationId,omitempty"`


	// PublishResultUri
	PublishResultUri *string `json:"publishResultUri,omitempty"`


	// InputSchema
	InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`


	// OutputSchema
	OutputSchema *Jsonschemadocument `json:"outputSchema,omitempty"`


	// NluInfo - Information about the NLU domain version for the flow version
	NluInfo *Nluinfo `json:"nluInfo,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowversion) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
