package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationcaseassociation
type Journeysessioneventsnotificationcaseassociation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CaseId
	CaseId *string `json:"caseId,omitempty"`

	// CaseReference
	CaseReference *string `json:"caseReference,omitempty"`

	// DateAssociated
	DateAssociated *time.Time `json:"dateAssociated,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeysessioneventsnotificationcaseassociation) SetField(field string, fieldValue interface{}) {
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

func (o Journeysessioneventsnotificationcaseassociation) MarshalJSON() ([]byte, error) {
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
	type Alias Journeysessioneventsnotificationcaseassociation
	
	DateAssociated := new(string)
	if o.DateAssociated != nil {
		
		*DateAssociated = timeutil.Strftime(o.DateAssociated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAssociated = nil
	}
	
	return json.Marshal(&struct { 
		CaseId *string `json:"caseId,omitempty"`
		
		CaseReference *string `json:"caseReference,omitempty"`
		
		DateAssociated *string `json:"dateAssociated,omitempty"`
		Alias
	}{ 
		CaseId: o.CaseId,
		
		CaseReference: o.CaseReference,
		
		DateAssociated: DateAssociated,
		Alias:    (Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationcaseassociation) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationcaseassociationMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationcaseassociationMap)
	if err != nil {
		return err
	}
	
	if CaseId, ok := JourneysessioneventsnotificationcaseassociationMap["caseId"].(string); ok {
		o.CaseId = &CaseId
	}
    
	if CaseReference, ok := JourneysessioneventsnotificationcaseassociationMap["caseReference"].(string); ok {
		o.CaseReference = &CaseReference
	}
    
	if dateAssociatedString, ok := JourneysessioneventsnotificationcaseassociationMap["dateAssociated"].(string); ok {
		DateAssociated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAssociatedString)
		o.DateAssociated = &DateAssociated
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationcaseassociation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
