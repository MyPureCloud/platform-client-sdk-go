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


	// DateCheckedIn
	DateCheckedIn *int `json:"dateCheckedIn,omitempty"`


	// DateSaved
	DateSaved *int `json:"dateSaved,omitempty"`


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

func (o *Flowversion) MarshalJSON() ([]byte, error) {
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
		
		DateCheckedIn *int `json:"dateCheckedIn,omitempty"`
		
		DateSaved *int `json:"dateSaved,omitempty"`
		
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
		Id: o.Id,
		
		Name: o.Name,
		
		CommitVersion: o.CommitVersion,
		
		ConfigurationVersion: o.ConfigurationVersion,
		
		VarType: o.VarType,
		
		Secure: o.Secure,
		
		Debug: o.Debug,
		
		CreatedBy: o.CreatedBy,
		
		CreatedByClient: o.CreatedByClient,
		
		ConfigurationUri: o.ConfigurationUri,
		
		DateCreated: o.DateCreated,
		
		DateCheckedIn: o.DateCheckedIn,
		
		DateSaved: o.DateSaved,
		
		GenerationId: o.GenerationId,
		
		PublishResultUri: o.PublishResultUri,
		
		InputSchema: o.InputSchema,
		
		OutputSchema: o.OutputSchema,
		
		NluInfo: o.NluInfo,
		
		SupportedLanguages: o.SupportedLanguages,
		
		CompatibleFlowTypes: o.CompatibleFlowTypes,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowversion) UnmarshalJSON(b []byte) error {
	var FlowversionMap map[string]interface{}
	err := json.Unmarshal(b, &FlowversionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowversionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FlowversionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if CommitVersion, ok := FlowversionMap["commitVersion"].(string); ok {
		o.CommitVersion = &CommitVersion
	}
    
	if ConfigurationVersion, ok := FlowversionMap["configurationVersion"].(string); ok {
		o.ConfigurationVersion = &ConfigurationVersion
	}
    
	if VarType, ok := FlowversionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Secure, ok := FlowversionMap["secure"].(bool); ok {
		o.Secure = &Secure
	}
    
	if Debug, ok := FlowversionMap["debug"].(bool); ok {
		o.Debug = &Debug
	}
    
	if CreatedBy, ok := FlowversionMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if CreatedByClient, ok := FlowversionMap["createdByClient"].(map[string]interface{}); ok {
		CreatedByClientString, _ := json.Marshal(CreatedByClient)
		json.Unmarshal(CreatedByClientString, &o.CreatedByClient)
	}
	
	if ConfigurationUri, ok := FlowversionMap["configurationUri"].(string); ok {
		o.ConfigurationUri = &ConfigurationUri
	}
    
	if DateCreated, ok := FlowversionMap["dateCreated"].(float64); ok {
		DateCreatedInt := int(DateCreated)
		o.DateCreated = &DateCreatedInt
	}
	
	if DateCheckedIn, ok := FlowversionMap["dateCheckedIn"].(float64); ok {
		DateCheckedInInt := int(DateCheckedIn)
		o.DateCheckedIn = &DateCheckedInInt
	}
	
	if DateSaved, ok := FlowversionMap["dateSaved"].(float64); ok {
		DateSavedInt := int(DateSaved)
		o.DateSaved = &DateSavedInt
	}
	
	if GenerationId, ok := FlowversionMap["generationId"].(string); ok {
		o.GenerationId = &GenerationId
	}
    
	if PublishResultUri, ok := FlowversionMap["publishResultUri"].(string); ok {
		o.PublishResultUri = &PublishResultUri
	}
    
	if InputSchema, ok := FlowversionMap["inputSchema"].(map[string]interface{}); ok {
		InputSchemaString, _ := json.Marshal(InputSchema)
		json.Unmarshal(InputSchemaString, &o.InputSchema)
	}
	
	if OutputSchema, ok := FlowversionMap["outputSchema"].(map[string]interface{}); ok {
		OutputSchemaString, _ := json.Marshal(OutputSchema)
		json.Unmarshal(OutputSchemaString, &o.OutputSchema)
	}
	
	if NluInfo, ok := FlowversionMap["nluInfo"].(map[string]interface{}); ok {
		NluInfoString, _ := json.Marshal(NluInfo)
		json.Unmarshal(NluInfoString, &o.NluInfo)
	}
	
	if SupportedLanguages, ok := FlowversionMap["supportedLanguages"].([]interface{}); ok {
		SupportedLanguagesString, _ := json.Marshal(SupportedLanguages)
		json.Unmarshal(SupportedLanguagesString, &o.SupportedLanguages)
	}
	
	if CompatibleFlowTypes, ok := FlowversionMap["compatibleFlowTypes"].([]interface{}); ok {
		CompatibleFlowTypesString, _ := json.Marshal(CompatibleFlowTypes)
		json.Unmarshal(CompatibleFlowTypesString, &o.CompatibleFlowTypes)
	}
	
	if SelfUri, ok := FlowversionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flowversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
