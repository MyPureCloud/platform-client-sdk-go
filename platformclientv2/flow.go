package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flow
type Flow struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The flow identifier
	Id *string `json:"id,omitempty"`

	// Name - The flow name
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// LockedUser - User that has the flow locked.
	LockedUser *User `json:"lockedUser,omitempty"`

	// LockedClient - OAuth client that has the flow locked.
	LockedClient *Domainentityref `json:"lockedClient,omitempty"`

	// Active
	Active *bool `json:"active,omitempty"`

	// System
	System *bool `json:"system,omitempty"`

	// Deleted
	Deleted *bool `json:"deleted,omitempty"`

	// PublishedVersion
	PublishedVersion *Flowversion `json:"publishedVersion,omitempty"`

	// SavedVersion
	SavedVersion *Flowversion `json:"savedVersion,omitempty"`

	// InputSchema - json schema describing the inputs for the flow
	InputSchema *interface{} `json:"inputSchema,omitempty"`

	// OutputSchema - json schema describing the outputs for the flow
	OutputSchema *interface{} `json:"outputSchema,omitempty"`

	// CheckedInVersion
	CheckedInVersion *Flowversion `json:"checkedInVersion,omitempty"`

	// DebugVersion
	DebugVersion *Flowversion `json:"debugVersion,omitempty"`

	// PublishedBy
	PublishedBy *User `json:"publishedBy,omitempty"`

	// CurrentOperation
	CurrentOperation *Operation `json:"currentOperation,omitempty"`

	// NluInfo - Information about the natural language understanding configuration for the published version of the flow
	NluInfo *Nluinfo `json:"nluInfo,omitempty"`

	// SupportedLanguages - List of supported languages for the published version of the flow.
	SupportedLanguages *[]Supportedlanguage `json:"supportedLanguages,omitempty"`

	// CompatibleFlowTypes - Compatible flow types designate which flow types are allowed to embed a flow’s configuration within their own flow configuration.  Currently the only flows that can be embedded are Common Module flows and the embedding flow can invoke them using the Call Common Module action.
	CompatibleFlowTypes *[]string `json:"compatibleFlowTypes,omitempty"`

	// WorktypeId
	WorktypeId *string `json:"worktypeId,omitempty"`

	// VirtualAgentEnabled
	VirtualAgentEnabled *bool `json:"virtualAgentEnabled,omitempty"`

	// AgenticVirtualAgentEnabled
	AgenticVirtualAgentEnabled *bool `json:"agenticVirtualAgentEnabled,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flow) SetField(field string, fieldValue interface{}) {
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

func (o Flow) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Flow
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		LockedUser *User `json:"lockedUser,omitempty"`
		
		LockedClient *Domainentityref `json:"lockedClient,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		System *bool `json:"system,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		PublishedVersion *Flowversion `json:"publishedVersion,omitempty"`
		
		SavedVersion *Flowversion `json:"savedVersion,omitempty"`
		
		InputSchema *interface{} `json:"inputSchema,omitempty"`
		
		OutputSchema *interface{} `json:"outputSchema,omitempty"`
		
		CheckedInVersion *Flowversion `json:"checkedInVersion,omitempty"`
		
		DebugVersion *Flowversion `json:"debugVersion,omitempty"`
		
		PublishedBy *User `json:"publishedBy,omitempty"`
		
		CurrentOperation *Operation `json:"currentOperation,omitempty"`
		
		NluInfo *Nluinfo `json:"nluInfo,omitempty"`
		
		SupportedLanguages *[]Supportedlanguage `json:"supportedLanguages,omitempty"`
		
		CompatibleFlowTypes *[]string `json:"compatibleFlowTypes,omitempty"`
		
		WorktypeId *string `json:"worktypeId,omitempty"`
		
		VirtualAgentEnabled *bool `json:"virtualAgentEnabled,omitempty"`
		
		AgenticVirtualAgentEnabled *bool `json:"agenticVirtualAgentEnabled,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		VarType: o.VarType,
		
		LockedUser: o.LockedUser,
		
		LockedClient: o.LockedClient,
		
		Active: o.Active,
		
		System: o.System,
		
		Deleted: o.Deleted,
		
		PublishedVersion: o.PublishedVersion,
		
		SavedVersion: o.SavedVersion,
		
		InputSchema: o.InputSchema,
		
		OutputSchema: o.OutputSchema,
		
		CheckedInVersion: o.CheckedInVersion,
		
		DebugVersion: o.DebugVersion,
		
		PublishedBy: o.PublishedBy,
		
		CurrentOperation: o.CurrentOperation,
		
		NluInfo: o.NluInfo,
		
		SupportedLanguages: o.SupportedLanguages,
		
		CompatibleFlowTypes: o.CompatibleFlowTypes,
		
		WorktypeId: o.WorktypeId,
		
		VirtualAgentEnabled: o.VirtualAgentEnabled,
		
		AgenticVirtualAgentEnabled: o.AgenticVirtualAgentEnabled,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Flow) UnmarshalJSON(b []byte) error {
	var FlowMap map[string]interface{}
	err := json.Unmarshal(b, &FlowMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FlowMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := FlowMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := FlowMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if VarType, ok := FlowMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if LockedUser, ok := FlowMap["lockedUser"].(map[string]interface{}); ok {
		LockedUserString, _ := json.Marshal(LockedUser)
		json.Unmarshal(LockedUserString, &o.LockedUser)
	}
	
	if LockedClient, ok := FlowMap["lockedClient"].(map[string]interface{}); ok {
		LockedClientString, _ := json.Marshal(LockedClient)
		json.Unmarshal(LockedClientString, &o.LockedClient)
	}
	
	if Active, ok := FlowMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if System, ok := FlowMap["system"].(bool); ok {
		o.System = &System
	}
    
	if Deleted, ok := FlowMap["deleted"].(bool); ok {
		o.Deleted = &Deleted
	}
    
	if PublishedVersion, ok := FlowMap["publishedVersion"].(map[string]interface{}); ok {
		PublishedVersionString, _ := json.Marshal(PublishedVersion)
		json.Unmarshal(PublishedVersionString, &o.PublishedVersion)
	}
	
	if SavedVersion, ok := FlowMap["savedVersion"].(map[string]interface{}); ok {
		SavedVersionString, _ := json.Marshal(SavedVersion)
		json.Unmarshal(SavedVersionString, &o.SavedVersion)
	}
	
	if InputSchema, ok := FlowMap["inputSchema"].(map[string]interface{}); ok {
		InputSchemaString, _ := json.Marshal(InputSchema)
		json.Unmarshal(InputSchemaString, &o.InputSchema)
	}
	
	if OutputSchema, ok := FlowMap["outputSchema"].(map[string]interface{}); ok {
		OutputSchemaString, _ := json.Marshal(OutputSchema)
		json.Unmarshal(OutputSchemaString, &o.OutputSchema)
	}
	
	if CheckedInVersion, ok := FlowMap["checkedInVersion"].(map[string]interface{}); ok {
		CheckedInVersionString, _ := json.Marshal(CheckedInVersion)
		json.Unmarshal(CheckedInVersionString, &o.CheckedInVersion)
	}
	
	if DebugVersion, ok := FlowMap["debugVersion"].(map[string]interface{}); ok {
		DebugVersionString, _ := json.Marshal(DebugVersion)
		json.Unmarshal(DebugVersionString, &o.DebugVersion)
	}
	
	if PublishedBy, ok := FlowMap["publishedBy"].(map[string]interface{}); ok {
		PublishedByString, _ := json.Marshal(PublishedBy)
		json.Unmarshal(PublishedByString, &o.PublishedBy)
	}
	
	if CurrentOperation, ok := FlowMap["currentOperation"].(map[string]interface{}); ok {
		CurrentOperationString, _ := json.Marshal(CurrentOperation)
		json.Unmarshal(CurrentOperationString, &o.CurrentOperation)
	}
	
	if NluInfo, ok := FlowMap["nluInfo"].(map[string]interface{}); ok {
		NluInfoString, _ := json.Marshal(NluInfo)
		json.Unmarshal(NluInfoString, &o.NluInfo)
	}
	
	if SupportedLanguages, ok := FlowMap["supportedLanguages"].([]interface{}); ok {
		SupportedLanguagesString, _ := json.Marshal(SupportedLanguages)
		json.Unmarshal(SupportedLanguagesString, &o.SupportedLanguages)
	}
	
	if CompatibleFlowTypes, ok := FlowMap["compatibleFlowTypes"].([]interface{}); ok {
		CompatibleFlowTypesString, _ := json.Marshal(CompatibleFlowTypes)
		json.Unmarshal(CompatibleFlowTypesString, &o.CompatibleFlowTypes)
	}
	
	if WorktypeId, ok := FlowMap["worktypeId"].(string); ok {
		o.WorktypeId = &WorktypeId
	}
    
	if VirtualAgentEnabled, ok := FlowMap["virtualAgentEnabled"].(bool); ok {
		o.VirtualAgentEnabled = &VirtualAgentEnabled
	}
    
	if AgenticVirtualAgentEnabled, ok := FlowMap["agenticVirtualAgentEnabled"].(bool); ok {
		o.AgenticVirtualAgentEnabled = &AgenticVirtualAgentEnabled
	}
    
	if SelfUri, ok := FlowMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
