package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	DateCreated *int `json:"dateCreated,omitempty"`


	// GenerationId
	GenerationId *string `json:"generationId,omitempty"`


	// PublishResultUri
	PublishResultUri *string `json:"publishResultUri,omitempty"`


	// InputSchema
	InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`


	// OutputSchema
	OutputSchema *Jsonschemadocument `json:"outputSchema,omitempty"`


	// NluInfo - Information about the natural language understanding configuration for the flow version
	NluInfo *Nluinfo `json:"nluInfo,omitempty"`


	// SupportedLanguages - List of supported languages for this version of the flow
	SupportedLanguages *[]Supportedlanguage `json:"supportedLanguages,omitempty"`


	// CompatibleFlowTypes - Compatible flow types designate which flow types are allowed to embed a flow’s configuration within their own flow configuration.  Currently the only flows that can be embedded are Common Module flows and the embedding flow can invoke them using the Call Common Module action.
	CompatibleFlowTypes *[]string `json:"compatibleFlowTypes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Flowversion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowversion

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		CommitVersion *string `json:"commitVersion,omitempty"`
		
		ConfigurationVersion *string `json:"configurationVersion,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Secure *bool `json:"secure,omitempty"`
		
		Debug *bool `json:"debug,omitempty"`
		
		CreatedBy *User `json:"createdBy,omitempty"`
		
		CreatedByClient *Domainentityref `json:"createdByClient,omitempty"`
		
		ConfigurationUri *string `json:"configurationUri,omitempty"`
		
		DateCreated *int `json:"dateCreated,omitempty"`
		
		GenerationId *string `json:"generationId,omitempty"`
		
		PublishResultUri *string `json:"publishResultUri,omitempty"`
		
		InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`
		
		OutputSchema *Jsonschemadocument `json:"outputSchema,omitempty"`
		
		NluInfo *Nluinfo `json:"nluInfo,omitempty"`
		
		SupportedLanguages *[]Supportedlanguage `json:"supportedLanguages,omitempty"`
		
		CompatibleFlowTypes *[]string `json:"compatibleFlowTypes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		CommitVersion: u.CommitVersion,
		
		ConfigurationVersion: u.ConfigurationVersion,
		
		VarType: u.VarType,
		
		Secure: u.Secure,
		
		Debug: u.Debug,
		
		CreatedBy: u.CreatedBy,
		
		CreatedByClient: u.CreatedByClient,
		
		ConfigurationUri: u.ConfigurationUri,
		
		DateCreated: u.DateCreated,
		
		GenerationId: u.GenerationId,
		
		PublishResultUri: u.PublishResultUri,
		
		InputSchema: u.InputSchema,
		
		OutputSchema: u.OutputSchema,
		
		NluInfo: u.NluInfo,
		
		SupportedLanguages: u.SupportedLanguages,
		
		CompatibleFlowTypes: u.CompatibleFlowTypes,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Flowversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
