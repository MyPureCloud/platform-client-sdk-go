package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxonheartbeatalertstopicheartbeatalert
type Klaxonheartbeatalertstopicheartbeatalert struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// SenderId
	SenderId *string `json:"senderId,omitempty"`

	// HeartBeatTimeoutInMinutes
	HeartBeatTimeoutInMinutes *float32 `json:"heartBeatTimeoutInMinutes,omitempty"`

	// RuleId
	RuleId *string `json:"ruleId,omitempty"`

	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`

	// NotificationUsers
	NotificationUsers *[]Klaxonheartbeatalertstopicnotificationuser `json:"notificationUsers,omitempty"`

	// AlertTypes
	AlertTypes *[]string `json:"alertTypes,omitempty"`

	// RuleType
	RuleType *string `json:"ruleType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Klaxonheartbeatalertstopicheartbeatalert) SetField(field string, fieldValue interface{}) {
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

func (o Klaxonheartbeatalertstopicheartbeatalert) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","EndDate", }
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
	type Alias Klaxonheartbeatalertstopicheartbeatalert
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SenderId *string `json:"senderId,omitempty"`
		
		HeartBeatTimeoutInMinutes *float32 `json:"heartBeatTimeoutInMinutes,omitempty"`
		
		RuleId *string `json:"ruleId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		NotificationUsers *[]Klaxonheartbeatalertstopicnotificationuser `json:"notificationUsers,omitempty"`
		
		AlertTypes *[]string `json:"alertTypes,omitempty"`
		
		RuleType *string `json:"ruleType,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SenderId: o.SenderId,
		
		HeartBeatTimeoutInMinutes: o.HeartBeatTimeoutInMinutes,
		
		RuleId: o.RuleId,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		NotificationUsers: o.NotificationUsers,
		
		AlertTypes: o.AlertTypes,
		
		RuleType: o.RuleType,
		Alias:    (Alias)(o),
	})
}

func (o *Klaxonheartbeatalertstopicheartbeatalert) UnmarshalJSON(b []byte) error {
	var KlaxonheartbeatalertstopicheartbeatalertMap map[string]interface{}
	err := json.Unmarshal(b, &KlaxonheartbeatalertstopicheartbeatalertMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KlaxonheartbeatalertstopicheartbeatalertMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := KlaxonheartbeatalertstopicheartbeatalertMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SenderId, ok := KlaxonheartbeatalertstopicheartbeatalertMap["senderId"].(string); ok {
		o.SenderId = &SenderId
	}
    
	if HeartBeatTimeoutInMinutes, ok := KlaxonheartbeatalertstopicheartbeatalertMap["heartBeatTimeoutInMinutes"].(float64); ok {
		HeartBeatTimeoutInMinutesFloat32 := float32(HeartBeatTimeoutInMinutes)
		o.HeartBeatTimeoutInMinutes = &HeartBeatTimeoutInMinutesFloat32
	}
    
	if RuleId, ok := KlaxonheartbeatalertstopicheartbeatalertMap["ruleId"].(string); ok {
		o.RuleId = &RuleId
	}
    
	if startDateString, ok := KlaxonheartbeatalertstopicheartbeatalertMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := KlaxonheartbeatalertstopicheartbeatalertMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if NotificationUsers, ok := KlaxonheartbeatalertstopicheartbeatalertMap["notificationUsers"].([]interface{}); ok {
		NotificationUsersString, _ := json.Marshal(NotificationUsers)
		json.Unmarshal(NotificationUsersString, &o.NotificationUsers)
	}
	
	if AlertTypes, ok := KlaxonheartbeatalertstopicheartbeatalertMap["alertTypes"].([]interface{}); ok {
		AlertTypesString, _ := json.Marshal(AlertTypes)
		json.Unmarshal(AlertTypesString, &o.AlertTypes)
	}
	
	if RuleType, ok := KlaxonheartbeatalertstopicheartbeatalertMap["ruleType"].(string); ok {
		o.RuleType = &RuleType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Klaxonheartbeatalertstopicheartbeatalert) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
