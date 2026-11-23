package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchdraftrequest
type Patchdraftrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Subject - The subject of the email draft. Replaces the existing subject.
	Subject *string `json:"subject,omitempty"`

	// To - The To recipients of the email draft. Replaces the entire To list.
	To *[]Emailaddress `json:"to,omitempty"`

	// Cc - The CC recipients of the email draft. Replaces the entire CC list.
	Cc *[]Emailaddress `json:"cc,omitempty"`

	// Bcc - The BCC recipients of the email draft. Replaces the entire BCC list.
	Bcc *[]Emailaddress `json:"bcc,omitempty"`

	// OverrideHeaders - Custom headers to set on the outbound email. Replaces the entire override headers list. Header keys are case-sensitive: keys differing only in case (e.g. MSIP and msip) are kept as distinct headers, while exact duplicate keys result in a 400 Bad Request.
	OverrideHeaders *[]Headerentry `json:"overrideHeaders,omitempty"`

	// HistoryIncluded - Whether to include message history in the draft. Once set to true, cannot be reverted to false.
	HistoryIncluded *bool `json:"historyIncluded,omitempty"`

	// Translation - A set of definitions to translate email attributes and correctly display date and time, for a given language. Used when historyIncluded is set to true.
	Translation *Historyheaderstranslation `json:"translation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Patchdraftrequest) SetField(field string, fieldValue interface{}) {
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

func (o Patchdraftrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Patchdraftrequest
	
	return json.Marshal(&struct { 
		Subject *string `json:"subject,omitempty"`
		
		To *[]Emailaddress `json:"to,omitempty"`
		
		Cc *[]Emailaddress `json:"cc,omitempty"`
		
		Bcc *[]Emailaddress `json:"bcc,omitempty"`
		
		OverrideHeaders *[]Headerentry `json:"overrideHeaders,omitempty"`
		
		HistoryIncluded *bool `json:"historyIncluded,omitempty"`
		
		Translation *Historyheaderstranslation `json:"translation,omitempty"`
		Alias
	}{ 
		Subject: o.Subject,
		
		To: o.To,
		
		Cc: o.Cc,
		
		Bcc: o.Bcc,
		
		OverrideHeaders: o.OverrideHeaders,
		
		HistoryIncluded: o.HistoryIncluded,
		
		Translation: o.Translation,
		Alias:    (Alias)(o),
	})
}

func (o *Patchdraftrequest) UnmarshalJSON(b []byte) error {
	var PatchdraftrequestMap map[string]interface{}
	err := json.Unmarshal(b, &PatchdraftrequestMap)
	if err != nil {
		return err
	}
	
	if Subject, ok := PatchdraftrequestMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if To, ok := PatchdraftrequestMap["to"].([]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if Cc, ok := PatchdraftrequestMap["cc"].([]interface{}); ok {
		CcString, _ := json.Marshal(Cc)
		json.Unmarshal(CcString, &o.Cc)
	}
	
	if Bcc, ok := PatchdraftrequestMap["bcc"].([]interface{}); ok {
		BccString, _ := json.Marshal(Bcc)
		json.Unmarshal(BccString, &o.Bcc)
	}
	
	if OverrideHeaders, ok := PatchdraftrequestMap["overrideHeaders"].([]interface{}); ok {
		OverrideHeadersString, _ := json.Marshal(OverrideHeaders)
		json.Unmarshal(OverrideHeadersString, &o.OverrideHeaders)
	}
	
	if HistoryIncluded, ok := PatchdraftrequestMap["historyIncluded"].(bool); ok {
		o.HistoryIncluded = &HistoryIncluded
	}
    
	if Translation, ok := PatchdraftrequestMap["translation"].(map[string]interface{}); ok {
		TranslationString, _ := json.Marshal(Translation)
		json.Unmarshal(TranslationString, &o.Translation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchdraftrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
