package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Routepathrequest
type Routepathrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QueueId - The ID of the queue to associate with the route path
	QueueId *string `json:"queueId,omitempty"`

	// MediaType - The media type of the given queue to associate with the route path
	MediaType *string `json:"mediaType,omitempty"`

	// LanguageId - The ID of the language to associate with the route path
	LanguageId *string `json:"languageId,omitempty"`

	// SkillIds - The set of skill IDs to associate with the route path
	SkillIds *[]string `json:"skillIds,omitempty"`

	// SourcePlanningGroup - The planning group from which to take route paths. This property is only needed if a route path already exists in another planning group.Note that taking a route path from another planning group will modify the other planning group
	SourcePlanningGroup *Sourceplanninggrouprequest `json:"sourcePlanningGroup,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Routepathrequest) SetField(field string, fieldValue interface{}) {
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

func (o Routepathrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Routepathrequest
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		SourcePlanningGroup *Sourceplanninggrouprequest `json:"sourcePlanningGroup,omitempty"`
		Alias
	}{ 
		QueueId: o.QueueId,
		
		MediaType: o.MediaType,
		
		LanguageId: o.LanguageId,
		
		SkillIds: o.SkillIds,
		
		SourcePlanningGroup: o.SourcePlanningGroup,
		Alias:    (Alias)(o),
	})
}

func (o *Routepathrequest) UnmarshalJSON(b []byte) error {
	var RoutepathrequestMap map[string]interface{}
	err := json.Unmarshal(b, &RoutepathrequestMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := RoutepathrequestMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if MediaType, ok := RoutepathrequestMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if LanguageId, ok := RoutepathrequestMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if SkillIds, ok := RoutepathrequestMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if SourcePlanningGroup, ok := RoutepathrequestMap["sourcePlanningGroup"].(map[string]interface{}); ok {
		SourcePlanningGroupString, _ := json.Marshal(SourcePlanningGroup)
		json.Unmarshal(SourcePlanningGroupString, &o.SourcePlanningGroup)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routepathrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
