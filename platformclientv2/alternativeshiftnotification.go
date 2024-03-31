package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Alternativeshiftnotification
type Alternativeshiftnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// WeekDate - The start date of the schedule with which this trade is associated. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`

	// Granularity - The granularity of alternative shifts to be traded
	Granularity *string `json:"granularity,omitempty"`

	// NewState - The new state of the alternative shift trade, null if there was no change
	NewState *string `json:"newState,omitempty"`

	// InitiatingUser - The user who initiated the alternative shift trade
	InitiatingUser *Userreference `json:"initiatingUser,omitempty"`

	// InitiatingShiftDate - The start date and time of the initiating shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	InitiatingShiftDate *time.Time `json:"initiatingShiftDate,omitempty"`

	// ReceivingUser - The user on the receiving this alternative shift trade
	ReceivingUser *Userreference `json:"receivingUser,omitempty"`

	// ReceivingShiftDate - The start date and time of the receiving alternative shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReceivingShiftDate *time.Time `json:"receivingShiftDate,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Alternativeshiftnotification) SetField(field string, fieldValue interface{}) {
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

func (o Alternativeshiftnotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "InitiatingShiftDate","ReceivingShiftDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "WeekDate", }

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
	type Alias Alternativeshiftnotification
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	
	InitiatingShiftDate := new(string)
	if o.InitiatingShiftDate != nil {
		
		*InitiatingShiftDate = timeutil.Strftime(o.InitiatingShiftDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		InitiatingShiftDate = nil
	}
	
	ReceivingShiftDate := new(string)
	if o.ReceivingShiftDate != nil {
		
		*ReceivingShiftDate = timeutil.Strftime(o.ReceivingShiftDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReceivingShiftDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		NewState *string `json:"newState,omitempty"`
		
		InitiatingUser *Userreference `json:"initiatingUser,omitempty"`
		
		InitiatingShiftDate *string `json:"initiatingShiftDate,omitempty"`
		
		ReceivingUser *Userreference `json:"receivingUser,omitempty"`
		
		ReceivingShiftDate *string `json:"receivingShiftDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		WeekDate: WeekDate,
		
		Granularity: o.Granularity,
		
		NewState: o.NewState,
		
		InitiatingUser: o.InitiatingUser,
		
		InitiatingShiftDate: InitiatingShiftDate,
		
		ReceivingUser: o.ReceivingUser,
		
		ReceivingShiftDate: ReceivingShiftDate,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Alternativeshiftnotification) UnmarshalJSON(b []byte) error {
	var AlternativeshiftnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &AlternativeshiftnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AlternativeshiftnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if weekDateString, ok := AlternativeshiftnotificationMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02", weekDateString)
		o.WeekDate = &WeekDate
	}
	
	if Granularity, ok := AlternativeshiftnotificationMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if NewState, ok := AlternativeshiftnotificationMap["newState"].(string); ok {
		o.NewState = &NewState
	}
    
	if InitiatingUser, ok := AlternativeshiftnotificationMap["initiatingUser"].(map[string]interface{}); ok {
		InitiatingUserString, _ := json.Marshal(InitiatingUser)
		json.Unmarshal(InitiatingUserString, &o.InitiatingUser)
	}
	
	if initiatingShiftDateString, ok := AlternativeshiftnotificationMap["initiatingShiftDate"].(string); ok {
		InitiatingShiftDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", initiatingShiftDateString)
		o.InitiatingShiftDate = &InitiatingShiftDate
	}
	
	if ReceivingUser, ok := AlternativeshiftnotificationMap["receivingUser"].(map[string]interface{}); ok {
		ReceivingUserString, _ := json.Marshal(ReceivingUser)
		json.Unmarshal(ReceivingUserString, &o.ReceivingUser)
	}
	
	if receivingShiftDateString, ok := AlternativeshiftnotificationMap["receivingShiftDate"].(string); ok {
		ReceivingShiftDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", receivingShiftDateString)
		o.ReceivingShiftDate = &ReceivingShiftDate
	}
	
	if SelfUri, ok := AlternativeshiftnotificationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Alternativeshiftnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
