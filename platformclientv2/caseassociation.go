package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Caseassociation - Represents an association between a case and an interaction
type Caseassociation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the association.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// AssociationType - Association type.
	AssociationType *string `json:"associationType,omitempty"`

	// DateAssociated - Interaction association date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAssociated *time.Time `json:"dateAssociated,omitempty"`

	// Workitem - Associated workitem ID.
	Workitem *Workitemreference `json:"workitem,omitempty"`

	// Conversation - Associated conversation ID.
	Conversation *Conversationreference `json:"conversation,omitempty"`

	// Stage - The stage related to this association.
	Stage *Stagereference `json:"stage,omitempty"`

	// Step - The step related to this association.
	Step *Stepreference `json:"step,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// VarCase - Case ID
	VarCase *Casereference `json:"case,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Caseassociation) SetField(field string, fieldValue interface{}) {
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

func (o Caseassociation) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateAssociated", }
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
	type Alias Caseassociation
	
	DateAssociated := new(string)
	if o.DateAssociated != nil {
		
		*DateAssociated = timeutil.Strftime(o.DateAssociated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAssociated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		AssociationType *string `json:"associationType,omitempty"`
		
		DateAssociated *string `json:"dateAssociated,omitempty"`
		
		Workitem *Workitemreference `json:"workitem,omitempty"`
		
		Conversation *Conversationreference `json:"conversation,omitempty"`
		
		Stage *Stagereference `json:"stage,omitempty"`
		
		Step *Stepreference `json:"step,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		VarCase *Casereference `json:"case,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		AssociationType: o.AssociationType,
		
		DateAssociated: DateAssociated,
		
		Workitem: o.Workitem,
		
		Conversation: o.Conversation,
		
		Stage: o.Stage,
		
		Step: o.Step,
		
		SelfUri: o.SelfUri,
		
		VarCase: o.VarCase,
		Alias:    (Alias)(o),
	})
}

func (o *Caseassociation) UnmarshalJSON(b []byte) error {
	var CaseassociationMap map[string]interface{}
	err := json.Unmarshal(b, &CaseassociationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CaseassociationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CaseassociationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if AssociationType, ok := CaseassociationMap["associationType"].(string); ok {
		o.AssociationType = &AssociationType
	}
    
	if dateAssociatedString, ok := CaseassociationMap["dateAssociated"].(string); ok {
		DateAssociated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAssociatedString)
		o.DateAssociated = &DateAssociated
	}
	
	if Workitem, ok := CaseassociationMap["workitem"].(map[string]interface{}); ok {
		WorkitemString, _ := json.Marshal(Workitem)
		json.Unmarshal(WorkitemString, &o.Workitem)
	}
	
	if Conversation, ok := CaseassociationMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Stage, ok := CaseassociationMap["stage"].(map[string]interface{}); ok {
		StageString, _ := json.Marshal(Stage)
		json.Unmarshal(StageString, &o.Stage)
	}
	
	if Step, ok := CaseassociationMap["step"].(map[string]interface{}); ok {
		StepString, _ := json.Marshal(Step)
		json.Unmarshal(StepString, &o.Step)
	}
	
	if SelfUri, ok := CaseassociationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if VarCase, ok := CaseassociationMap["case"].(map[string]interface{}); ok {
		VarCaseString, _ := json.Marshal(VarCase)
		json.Unmarshal(VarCaseString, &o.VarCase)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Caseassociation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
