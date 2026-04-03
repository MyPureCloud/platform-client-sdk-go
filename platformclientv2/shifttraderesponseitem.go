package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttraderesponseitem
type Shifttraderesponseitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of this shift trade
	Id *string `json:"id,omitempty"`

	// State - The state of this shift trade
	State *string `json:"state,omitempty"`

	// ExpirationDate - When this shift trade will expire. Date time is represented as an ISO-8601 string
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// Initiating - Details about the initiating user involved in this shift trade
	Initiating *Shifttradeinitiatingsideresponseitem `json:"initiating,omitempty"`

	// Receiving - Details about the receiving user involved in this shift trade
	Receiving *Shifttradereceivingsideresponseitem `json:"receiving,omitempty"`

	// AcceptableIntervals - Time frames when the initiating user is willing to accept trades. Empty means giving up the shift
	AcceptableIntervals *[]Requireddaterange `json:"acceptableIntervals,omitempty"`

	// OneSided - Whether this is a one-sided shift trade (e.g. the initiating user is not asking for a shift in return)
	OneSided *bool `json:"oneSided,omitempty"`

	// Target - The user to whom the shift trade request was sent in a direct trade, or the user with whom a shift trade was Matched
	Target *Shifttradetargetresponseitem `json:"target,omitempty"`

	// ReviewedBy - The admin who approved or denied this shift trade
	ReviewedBy *Userreference `json:"reviewedBy,omitempty"`

	// ReviewedDate - The timestamp of when the trade request was reviewed by an admin in ISO-8601 format
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`

	// Metadata - Version metadata for this shift trade
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Shifttraderesponseitem) SetField(field string, fieldValue interface{}) {
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

func (o Shifttraderesponseitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ExpirationDate","ReviewedDate", }
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
	type Alias Shifttraderesponseitem
	
	ExpirationDate := new(string)
	if o.ExpirationDate != nil {
		
		*ExpirationDate = timeutil.Strftime(o.ExpirationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExpirationDate = nil
	}
	
	ReviewedDate := new(string)
	if o.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(o.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ExpirationDate *string `json:"expirationDate,omitempty"`
		
		Initiating *Shifttradeinitiatingsideresponseitem `json:"initiating,omitempty"`
		
		Receiving *Shifttradereceivingsideresponseitem `json:"receiving,omitempty"`
		
		AcceptableIntervals *[]Requireddaterange `json:"acceptableIntervals,omitempty"`
		
		OneSided *bool `json:"oneSided,omitempty"`
		
		Target *Shifttradetargetresponseitem `json:"target,omitempty"`
		
		ReviewedBy *Userreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		ExpirationDate: ExpirationDate,
		
		Initiating: o.Initiating,
		
		Receiving: o.Receiving,
		
		AcceptableIntervals: o.AcceptableIntervals,
		
		OneSided: o.OneSided,
		
		Target: o.Target,
		
		ReviewedBy: o.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Shifttraderesponseitem) UnmarshalJSON(b []byte) error {
	var ShifttraderesponseitemMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttraderesponseitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ShifttraderesponseitemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := ShifttraderesponseitemMap["state"].(string); ok {
		o.State = &State
	}
    
	if expirationDateString, ok := ShifttraderesponseitemMap["expirationDate"].(string); ok {
		ExpirationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationDateString)
		o.ExpirationDate = &ExpirationDate
	}
	
	if Initiating, ok := ShifttraderesponseitemMap["initiating"].(map[string]interface{}); ok {
		InitiatingString, _ := json.Marshal(Initiating)
		json.Unmarshal(InitiatingString, &o.Initiating)
	}
	
	if Receiving, ok := ShifttraderesponseitemMap["receiving"].(map[string]interface{}); ok {
		ReceivingString, _ := json.Marshal(Receiving)
		json.Unmarshal(ReceivingString, &o.Receiving)
	}
	
	if AcceptableIntervals, ok := ShifttraderesponseitemMap["acceptableIntervals"].([]interface{}); ok {
		AcceptableIntervalsString, _ := json.Marshal(AcceptableIntervals)
		json.Unmarshal(AcceptableIntervalsString, &o.AcceptableIntervals)
	}
	
	if OneSided, ok := ShifttraderesponseitemMap["oneSided"].(bool); ok {
		o.OneSided = &OneSided
	}
    
	if Target, ok := ShifttraderesponseitemMap["target"].(map[string]interface{}); ok {
		TargetString, _ := json.Marshal(Target)
		json.Unmarshal(TargetString, &o.Target)
	}
	
	if ReviewedBy, ok := ShifttraderesponseitemMap["reviewedBy"].(map[string]interface{}); ok {
		ReviewedByString, _ := json.Marshal(ReviewedBy)
		json.Unmarshal(ReviewedByString, &o.ReviewedBy)
	}
	
	if reviewedDateString, ok := ShifttraderesponseitemMap["reviewedDate"].(string); ok {
		ReviewedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", reviewedDateString)
		o.ReviewedDate = &ReviewedDate
	}
	
	if Metadata, ok := ShifttraderesponseitemMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttraderesponseitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
