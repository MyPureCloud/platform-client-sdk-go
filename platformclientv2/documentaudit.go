package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentaudit
type Documentaudit struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// User
	User *Domainentityref `json:"user,omitempty"`

	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`

	// TransactionId
	TransactionId *string `json:"transactionId,omitempty"`

	// TransactionInitiator
	TransactionInitiator *bool `json:"transactionInitiator,omitempty"`

	// Application
	Application *string `json:"application,omitempty"`

	// ServiceName
	ServiceName *string `json:"serviceName,omitempty"`

	// Level
	Level *string `json:"level,omitempty"`

	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// ActionContext
	ActionContext *string `json:"actionContext,omitempty"`

	// Action
	Action *string `json:"action,omitempty"`

	// Entity
	Entity *Auditentityreference `json:"entity,omitempty"`

	// Changes
	Changes *[]Auditchange `json:"changes,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentaudit) SetField(field string, fieldValue interface{}) {
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

func (o Documentaudit) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Timestamp", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Documentaudit
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User *Domainentityref `json:"user,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		TransactionId *string `json:"transactionId,omitempty"`
		
		TransactionInitiator *bool `json:"transactionInitiator,omitempty"`
		
		Application *string `json:"application,omitempty"`
		
		ServiceName *string `json:"serviceName,omitempty"`
		
		Level *string `json:"level,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ActionContext *string `json:"actionContext,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		Entity *Auditentityreference `json:"entity,omitempty"`
		
		Changes *[]Auditchange `json:"changes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		User: o.User,
		
		Workspace: o.Workspace,
		
		TransactionId: o.TransactionId,
		
		TransactionInitiator: o.TransactionInitiator,
		
		Application: o.Application,
		
		ServiceName: o.ServiceName,
		
		Level: o.Level,
		
		Timestamp: Timestamp,
		
		Status: o.Status,
		
		ActionContext: o.ActionContext,
		
		Action: o.Action,
		
		Entity: o.Entity,
		
		Changes: o.Changes,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Documentaudit) UnmarshalJSON(b []byte) error {
	var DocumentauditMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentauditMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DocumentauditMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DocumentauditMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if User, ok := DocumentauditMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Workspace, ok := DocumentauditMap["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if TransactionId, ok := DocumentauditMap["transactionId"].(string); ok {
		o.TransactionId = &TransactionId
	}
    
	if TransactionInitiator, ok := DocumentauditMap["transactionInitiator"].(bool); ok {
		o.TransactionInitiator = &TransactionInitiator
	}
    
	if Application, ok := DocumentauditMap["application"].(string); ok {
		o.Application = &Application
	}
    
	if ServiceName, ok := DocumentauditMap["serviceName"].(string); ok {
		o.ServiceName = &ServiceName
	}
    
	if Level, ok := DocumentauditMap["level"].(string); ok {
		o.Level = &Level
	}
    
	if timestampString, ok := DocumentauditMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if Status, ok := DocumentauditMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ActionContext, ok := DocumentauditMap["actionContext"].(string); ok {
		o.ActionContext = &ActionContext
	}
    
	if Action, ok := DocumentauditMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if Entity, ok := DocumentauditMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if Changes, ok := DocumentauditMap["changes"].([]interface{}); ok {
		ChangesString, _ := json.Marshal(Changes)
		json.Unmarshal(ChangesString, &o.Changes)
	}
	
	if SelfUri, ok := DocumentauditMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Documentaudit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
