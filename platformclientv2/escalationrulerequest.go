package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Escalationrulerequest
type Escalationrulerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the escalation rule.
	Name *string `json:"name,omitempty"`

	// MatchCriteria - The criteria that defines when a social media message should be escalated.
	MatchCriteria *string `json:"matchCriteria,omitempty"`

	// Priority - The priority of the escalation rule. The lower the number the higer the priority. Once a rule is matched others are skipped.
	Priority *int `json:"priority,omitempty"`

	// DivisionId - The ID of the division the social escalation rule belongs to.
	DivisionId *string `json:"divisionId,omitempty"`

	// Description - A description of the social escalation rule.
	Description *string `json:"description,omitempty"`

	// Status - The status of the escalation rule.
	Status *string `json:"status,omitempty"`

	// OpenEscalation - The target integration configuration used for an open message escalation.
	OpenEscalation *Escalationtarget `json:"openEscalation,omitempty"`

	// FacebookEscalation - The target integration configuration used for a Facebook message escalation.
	FacebookEscalation *Escalationtarget `json:"facebookEscalation,omitempty"`

	// InstagramEscalation - The target integration configuration used for an Instagram message escalation.
	InstagramEscalation *Escalationtarget `json:"instagramEscalation,omitempty"`

	// TwitterEscalation - The target integration configuration used for a X (formerly Twitter) message escalation.
	TwitterEscalation *Escalationtarget `json:"twitterEscalation,omitempty"`

	// GoogleBusinessProfileEscalation - The target integration configuration used for a Google Business Profile message escalation.
	GoogleBusinessProfileEscalation *Escalationtarget `json:"googleBusinessProfileEscalation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Escalationrulerequest) SetField(field string, fieldValue interface{}) {
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

func (o Escalationrulerequest) MarshalJSON() ([]byte, error) {
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
	type Alias Escalationrulerequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		MatchCriteria *string `json:"matchCriteria,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		OpenEscalation *Escalationtarget `json:"openEscalation,omitempty"`
		
		FacebookEscalation *Escalationtarget `json:"facebookEscalation,omitempty"`
		
		InstagramEscalation *Escalationtarget `json:"instagramEscalation,omitempty"`
		
		TwitterEscalation *Escalationtarget `json:"twitterEscalation,omitempty"`
		
		GoogleBusinessProfileEscalation *Escalationtarget `json:"googleBusinessProfileEscalation,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		MatchCriteria: o.MatchCriteria,
		
		Priority: o.Priority,
		
		DivisionId: o.DivisionId,
		
		Description: o.Description,
		
		Status: o.Status,
		
		OpenEscalation: o.OpenEscalation,
		
		FacebookEscalation: o.FacebookEscalation,
		
		InstagramEscalation: o.InstagramEscalation,
		
		TwitterEscalation: o.TwitterEscalation,
		
		GoogleBusinessProfileEscalation: o.GoogleBusinessProfileEscalation,
		Alias:    (Alias)(o),
	})
}

func (o *Escalationrulerequest) UnmarshalJSON(b []byte) error {
	var EscalationrulerequestMap map[string]interface{}
	err := json.Unmarshal(b, &EscalationrulerequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := EscalationrulerequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if MatchCriteria, ok := EscalationrulerequestMap["matchCriteria"].(string); ok {
		o.MatchCriteria = &MatchCriteria
	}
    
	if Priority, ok := EscalationrulerequestMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if DivisionId, ok := EscalationrulerequestMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if Description, ok := EscalationrulerequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Status, ok := EscalationrulerequestMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if OpenEscalation, ok := EscalationrulerequestMap["openEscalation"].(map[string]interface{}); ok {
		OpenEscalationString, _ := json.Marshal(OpenEscalation)
		json.Unmarshal(OpenEscalationString, &o.OpenEscalation)
	}
	
	if FacebookEscalation, ok := EscalationrulerequestMap["facebookEscalation"].(map[string]interface{}); ok {
		FacebookEscalationString, _ := json.Marshal(FacebookEscalation)
		json.Unmarshal(FacebookEscalationString, &o.FacebookEscalation)
	}
	
	if InstagramEscalation, ok := EscalationrulerequestMap["instagramEscalation"].(map[string]interface{}); ok {
		InstagramEscalationString, _ := json.Marshal(InstagramEscalation)
		json.Unmarshal(InstagramEscalationString, &o.InstagramEscalation)
	}
	
	if TwitterEscalation, ok := EscalationrulerequestMap["twitterEscalation"].(map[string]interface{}); ok {
		TwitterEscalationString, _ := json.Marshal(TwitterEscalation)
		json.Unmarshal(TwitterEscalationString, &o.TwitterEscalation)
	}
	
	if GoogleBusinessProfileEscalation, ok := EscalationrulerequestMap["googleBusinessProfileEscalation"].(map[string]interface{}); ok {
		GoogleBusinessProfileEscalationString, _ := json.Marshal(GoogleBusinessProfileEscalation)
		json.Unmarshal(GoogleBusinessProfileEscalationString, &o.GoogleBusinessProfileEscalation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Escalationrulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
