package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmshifttradingupdatetradejobeventshifttradedetail
type Wfmshifttradingupdatetradejobeventshifttradedetail struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// ExpirationDate
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// AcceptableIntervals
	AcceptableIntervals *[]Wfmshifttradingupdatetradejobeventshifttradeinterval `json:"acceptableIntervals,omitempty"`

	// OneSided
	OneSided *bool `json:"oneSided,omitempty"`

	// Initiating
	Initiating *Wfmshifttradingupdatetradejobeventshifttradeinitiatingside `json:"initiating,omitempty"`

	// Receiving
	Receiving *Wfmshifttradingupdatetradejobeventshifttradereceivingside `json:"receiving,omitempty"`

	// Target
	Target *Wfmshifttradingupdatetradejobeventshifttradetarget `json:"target,omitempty"`

	// ReviewedBy
	ReviewedBy *Wfmshifttradingupdatetradejobeventuserreference `json:"reviewedBy,omitempty"`

	// ReviewedDate
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`

	// Metadata
	Metadata *Wfmshifttradingupdatetradejobeventwfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmshifttradingupdatetradejobeventshifttradedetail) SetField(field string, fieldValue interface{}) {
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

func (o Wfmshifttradingupdatetradejobeventshifttradedetail) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmshifttradingupdatetradejobeventshifttradedetail
	
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
		
		AcceptableIntervals *[]Wfmshifttradingupdatetradejobeventshifttradeinterval `json:"acceptableIntervals,omitempty"`
		
		OneSided *bool `json:"oneSided,omitempty"`
		
		Initiating *Wfmshifttradingupdatetradejobeventshifttradeinitiatingside `json:"initiating,omitempty"`
		
		Receiving *Wfmshifttradingupdatetradejobeventshifttradereceivingside `json:"receiving,omitempty"`
		
		Target *Wfmshifttradingupdatetradejobeventshifttradetarget `json:"target,omitempty"`
		
		ReviewedBy *Wfmshifttradingupdatetradejobeventuserreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		Metadata *Wfmshifttradingupdatetradejobeventwfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		ExpirationDate: ExpirationDate,
		
		AcceptableIntervals: o.AcceptableIntervals,
		
		OneSided: o.OneSided,
		
		Initiating: o.Initiating,
		
		Receiving: o.Receiving,
		
		Target: o.Target,
		
		ReviewedBy: o.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmshifttradingupdatetradejobeventshifttradedetail) UnmarshalJSON(b []byte) error {
	var WfmshifttradingupdatetradejobeventshifttradedetailMap map[string]interface{}
	err := json.Unmarshal(b, &WfmshifttradingupdatetradejobeventshifttradedetailMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["state"].(string); ok {
		o.State = &State
	}
    
	if expirationDateString, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["expirationDate"].(string); ok {
		ExpirationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationDateString)
		o.ExpirationDate = &ExpirationDate
	}
	
	if AcceptableIntervals, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["acceptableIntervals"].([]interface{}); ok {
		AcceptableIntervalsString, _ := json.Marshal(AcceptableIntervals)
		json.Unmarshal(AcceptableIntervalsString, &o.AcceptableIntervals)
	}
	
	if OneSided, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["oneSided"].(bool); ok {
		o.OneSided = &OneSided
	}
    
	if Initiating, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["initiating"].(map[string]interface{}); ok {
		InitiatingString, _ := json.Marshal(Initiating)
		json.Unmarshal(InitiatingString, &o.Initiating)
	}
	
	if Receiving, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["receiving"].(map[string]interface{}); ok {
		ReceivingString, _ := json.Marshal(Receiving)
		json.Unmarshal(ReceivingString, &o.Receiving)
	}
	
	if Target, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["target"].(map[string]interface{}); ok {
		TargetString, _ := json.Marshal(Target)
		json.Unmarshal(TargetString, &o.Target)
	}
	
	if ReviewedBy, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["reviewedBy"].(map[string]interface{}); ok {
		ReviewedByString, _ := json.Marshal(ReviewedBy)
		json.Unmarshal(ReviewedByString, &o.ReviewedBy)
	}
	
	if reviewedDateString, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["reviewedDate"].(string); ok {
		ReviewedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", reviewedDateString)
		o.ReviewedDate = &ReviewedDate
	}
	
	if Metadata, ok := WfmshifttradingupdatetradejobeventshifttradedetailMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmshifttradingupdatetradejobeventshifttradedetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
