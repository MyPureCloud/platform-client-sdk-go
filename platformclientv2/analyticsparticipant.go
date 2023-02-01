package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsparticipant
type Analyticsparticipant struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ExternalContactId - External contact identifier
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// ExternalOrganizationId - External organization identifier
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`

	// FlaggedReason - Reason for which participant flagged conversation
	FlaggedReason *string `json:"flaggedReason,omitempty"`

	// ParticipantId - Unique identifier for the participant
	ParticipantId *string `json:"participantId,omitempty"`

	// ParticipantName - A human readable name identifying the participant
	ParticipantName *string `json:"participantName,omitempty"`

	// Purpose - The participant's purpose
	Purpose *string `json:"purpose,omitempty"`

	// TeamId - The team ID the user is a member of
	TeamId *string `json:"teamId,omitempty"`

	// UserId - Unique identifier for the user
	UserId *string `json:"userId,omitempty"`

	// Sessions - List of sessions associated to this participant
	Sessions *[]Analyticssession `json:"sessions,omitempty"`

	// Attributes - List of attributes associated to this participant
	Attributes *map[string]string `json:"attributes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsparticipant) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsparticipant) MarshalJSON() ([]byte, error) {
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
	type Alias Analyticsparticipant
	
	return json.Marshal(&struct { 
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		ParticipantName *string `json:"participantName,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Sessions *[]Analyticssession `json:"sessions,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		Alias
	}{ 
		ExternalContactId: o.ExternalContactId,
		
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		FlaggedReason: o.FlaggedReason,
		
		ParticipantId: o.ParticipantId,
		
		ParticipantName: o.ParticipantName,
		
		Purpose: o.Purpose,
		
		TeamId: o.TeamId,
		
		UserId: o.UserId,
		
		Sessions: o.Sessions,
		
		Attributes: o.Attributes,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsparticipant) UnmarshalJSON(b []byte) error {
	var AnalyticsparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsparticipantMap)
	if err != nil {
		return err
	}
	
	if ExternalContactId, ok := AnalyticsparticipantMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ExternalOrganizationId, ok := AnalyticsparticipantMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
    
	if FlaggedReason, ok := AnalyticsparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if ParticipantId, ok := AnalyticsparticipantMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if ParticipantName, ok := AnalyticsparticipantMap["participantName"].(string); ok {
		o.ParticipantName = &ParticipantName
	}
    
	if Purpose, ok := AnalyticsparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if TeamId, ok := AnalyticsparticipantMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
    
	if UserId, ok := AnalyticsparticipantMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Sessions, ok := AnalyticsparticipantMap["sessions"].([]interface{}); ok {
		SessionsString, _ := json.Marshal(Sessions)
		json.Unmarshal(SessionsString, &o.Sessions)
	}
	
	if Attributes, ok := AnalyticsparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
