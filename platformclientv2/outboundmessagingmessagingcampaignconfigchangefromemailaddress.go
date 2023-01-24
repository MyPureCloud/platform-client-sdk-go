package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangefromemailaddress - An outbound-messaging messaging campaign email address
type Outboundmessagingmessagingcampaignconfigchangefromemailaddress struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FriendlyName - The friendly name of the email address.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// LocalPart - The local part of the email address.
	LocalPart *string `json:"localPart,omitempty"`

	// Domain - A UriReference for a resource
	Domain *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"domain,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outboundmessagingmessagingcampaignconfigchangefromemailaddress) SetField(field string, fieldValue interface{}) {
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

func (o Outboundmessagingmessagingcampaignconfigchangefromemailaddress) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Outboundmessagingmessagingcampaignconfigchangefromemailaddress
	
	return json.Marshal(&struct { 
		FriendlyName *string `json:"friendlyName,omitempty"`
		
		LocalPart *string `json:"localPart,omitempty"`
		
		Domain *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"domain,omitempty"`
		Alias
	}{ 
		FriendlyName: o.FriendlyName,
		
		LocalPart: o.LocalPart,
		
		Domain: o.Domain,
		Alias:    (Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignconfigchangefromemailaddress) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignconfigchangefromemailaddressMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignconfigchangefromemailaddressMap)
	if err != nil {
		return err
	}
	
	if FriendlyName, ok := OutboundmessagingmessagingcampaignconfigchangefromemailaddressMap["friendlyName"].(string); ok {
		o.FriendlyName = &FriendlyName
	}
    
	if LocalPart, ok := OutboundmessagingmessagingcampaignconfigchangefromemailaddressMap["localPart"].(string); ok {
		o.LocalPart = &LocalPart
	}
    
	if Domain, ok := OutboundmessagingmessagingcampaignconfigchangefromemailaddressMap["domain"].(map[string]interface{}); ok {
		DomainString, _ := json.Marshal(Domain)
		json.Unmarshal(DomainString, &o.Domain)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangefromemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
