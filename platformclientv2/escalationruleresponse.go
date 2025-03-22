package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Escalationruleresponse
type Escalationruleresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - ID of the escalation rule.
	Id *string `json:"id,omitempty"`

	// Name - The name of the escalation rule.
	Name *string `json:"name,omitempty"`

	// MatchCriteria - The criteria that defines when a social media message should be escalated.
	MatchCriteria *string `json:"matchCriteria,omitempty"`

	// Priority - The priority of the escalation rule.
	Priority *int `json:"priority,omitempty"`

	// DivisionId - The ID of the division the social escalation rule belongs to.
	DivisionId *string `json:"divisionId,omitempty"`

	// Description - A description of the social escalation rule.
	Description *string `json:"description,omitempty"`

	// DateCreated - Timestamp indicating when the escalation rule was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Timestamp indicating when the escalation rule was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Status - The status of the escalation rule.
	Status *string `json:"status,omitempty"`

	// OpenEscalation - The target integration configuration used for an open social media message if the match criteria returns true.
	OpenEscalation *Escalationtarget `json:"openEscalation,omitempty"`

	// FacebookEscalation - The target integration configuration used for a Facebook social media message if the match criteria returns true.
	FacebookEscalation *Escalationtarget `json:"facebookEscalation,omitempty"`

	// InstagramEscalation - The target integration configuration used for an Instagram social media message if the match criteria returns true.
	InstagramEscalation *Escalationtarget `json:"instagramEscalation,omitempty"`

	// TwitterEscalation - The target integration configuration used for a X (formerly Twitter) social media message if the match criteria returns true.
	TwitterEscalation *Escalationtarget `json:"twitterEscalation,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Escalationruleresponse) SetField(field string, fieldValue interface{}) {
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

func (o Escalationruleresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Escalationruleresponse
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		MatchCriteria *string `json:"matchCriteria,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		OpenEscalation *Escalationtarget `json:"openEscalation,omitempty"`
		
		FacebookEscalation *Escalationtarget `json:"facebookEscalation,omitempty"`
		
		InstagramEscalation *Escalationtarget `json:"instagramEscalation,omitempty"`
		
		TwitterEscalation *Escalationtarget `json:"twitterEscalation,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		MatchCriteria: o.MatchCriteria,
		
		Priority: o.Priority,
		
		DivisionId: o.DivisionId,
		
		Description: o.Description,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Status: o.Status,
		
		OpenEscalation: o.OpenEscalation,
		
		FacebookEscalation: o.FacebookEscalation,
		
		InstagramEscalation: o.InstagramEscalation,
		
		TwitterEscalation: o.TwitterEscalation,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Escalationruleresponse) UnmarshalJSON(b []byte) error {
	var EscalationruleresponseMap map[string]interface{}
	err := json.Unmarshal(b, &EscalationruleresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EscalationruleresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EscalationruleresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if MatchCriteria, ok := EscalationruleresponseMap["matchCriteria"].(string); ok {
		o.MatchCriteria = &MatchCriteria
	}
    
	if Priority, ok := EscalationruleresponseMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if DivisionId, ok := EscalationruleresponseMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if Description, ok := EscalationruleresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateCreatedString, ok := EscalationruleresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := EscalationruleresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Status, ok := EscalationruleresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if OpenEscalation, ok := EscalationruleresponseMap["openEscalation"].(map[string]interface{}); ok {
		OpenEscalationString, _ := json.Marshal(OpenEscalation)
		json.Unmarshal(OpenEscalationString, &o.OpenEscalation)
	}
	
	if FacebookEscalation, ok := EscalationruleresponseMap["facebookEscalation"].(map[string]interface{}); ok {
		FacebookEscalationString, _ := json.Marshal(FacebookEscalation)
		json.Unmarshal(FacebookEscalationString, &o.FacebookEscalation)
	}
	
	if InstagramEscalation, ok := EscalationruleresponseMap["instagramEscalation"].(map[string]interface{}); ok {
		InstagramEscalationString, _ := json.Marshal(InstagramEscalation)
		json.Unmarshal(InstagramEscalationString, &o.InstagramEscalation)
	}
	
	if TwitterEscalation, ok := EscalationruleresponseMap["twitterEscalation"].(map[string]interface{}); ok {
		TwitterEscalationString, _ := json.Marshal(TwitterEscalation)
		json.Unmarshal(TwitterEscalationString, &o.TwitterEscalation)
	}
	
	if SelfUri, ok := EscalationruleresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Escalationruleresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
