package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowversion
type Flowversion struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// VirtualAgentEnabled
	VirtualAgentEnabled *bool `json:"virtualAgentEnabled,omitempty"`

	// AgenticVirtualAgentEnabled
	AgenticVirtualAgentEnabled *bool `json:"agenticVirtualAgentEnabled,omitempty"`

	// DatePublished - The date this version became the published version of the flow. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublished *time.Time `json:"datePublished,omitempty"`

	// DatePublishedEnd - The date this version was no longer the published version of the flow. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublishedEnd *time.Time `json:"datePublishedEnd,omitempty"`

	// NluInfo - Information about the natural language understanding configuration for the flow version
	NluInfo *Nluinfo `json:"nluInfo,omitempty"`

	// SupportedLanguages - List of supported languages for this version of the flow
	SupportedLanguages *[]Supportedlanguage `json:"supportedLanguages,omitempty"`

	// CompatibleFlowTypes - Compatible flow types designate which flow types are allowed to embed a flow’s configuration within their own flow configuration.  Currently the only flows that can be embedded are Common Module flows and the embedding flow can invoke them using the Call Common Module action.
	CompatibleFlowTypes *[]string `json:"compatibleFlowTypes,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowversion) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Flowversion) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DatePublished","DatePublishedEnd", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowversion
	
	DatePublished := new(string)
	if o.DatePublished != nil {
		
		*DatePublished = timeutil.Strftime(o.DatePublished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DatePublished = nil
	}
	
	DatePublishedEnd := new(string)
	if o.DatePublishedEnd != nil {
		
		*DatePublishedEnd = timeutil.Strftime(o.DatePublishedEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DatePublishedEnd = nil
	}
	
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
		
		VirtualAgentEnabled *bool `json:"virtualAgentEnabled,omitempty"`
		
		AgenticVirtualAgentEnabled *bool `json:"agenticVirtualAgentEnabled,omitempty"`
		
		DatePublished *string `json:"datePublished,omitempty"`
		
		DatePublishedEnd *string `json:"datePublishedEnd,omitempty"`
		
		NluInfo *Nluinfo `json:"nluInfo,omitempty"`
		
		SupportedLanguages *[]Supportedlanguage `json:"supportedLanguages,omitempty"`
		
		CompatibleFlowTypes *[]string `json:"compatibleFlowTypes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
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
		
		VirtualAgentEnabled: o.VirtualAgentEnabled,
		
		AgenticVirtualAgentEnabled: o.AgenticVirtualAgentEnabled,
		
		DatePublished: DatePublished,
		
		DatePublishedEnd: DatePublishedEnd,
		
		NluInfo: o.NluInfo,
		
		SupportedLanguages: o.SupportedLanguages,
		
		CompatibleFlowTypes: o.CompatibleFlowTypes,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
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
	
	if VirtualAgentEnabled, ok := FlowversionMap["virtualAgentEnabled"].(bool); ok {
		o.VirtualAgentEnabled = &VirtualAgentEnabled
	}
    
	if AgenticVirtualAgentEnabled, ok := FlowversionMap["agenticVirtualAgentEnabled"].(bool); ok {
		o.AgenticVirtualAgentEnabled = &AgenticVirtualAgentEnabled
	}
    
	if datePublishedString, ok := FlowversionMap["datePublished"].(string); ok {
		DatePublished, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePublishedString)
		o.DatePublished = &DatePublished
	}
	
	if datePublishedEndString, ok := FlowversionMap["datePublishedEnd"].(string); ok {
		DatePublishedEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePublishedEndString)
		o.DatePublishedEnd = &DatePublishedEnd
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
