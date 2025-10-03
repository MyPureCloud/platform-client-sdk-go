package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundvoicecampaignpostcontactvoiceattributes
type Outboundvoicecampaignpostcontactvoiceattributes struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContactPhoneNumber
	ContactPhoneNumber *string `json:"contactPhoneNumber,omitempty"`

	// ContactPhoneType
	ContactPhoneType *string `json:"contactPhoneType,omitempty"`

	// CallerIdPhoneNumber
	CallerIdPhoneNumber *string `json:"callerIdPhoneNumber,omitempty"`

	// CallerIdName
	CallerIdName *string `json:"callerIdName,omitempty"`

	// AgentOwnedColumnName
	AgentOwnedColumnName *string `json:"agentOwnedColumnName,omitempty"`

	// PreviewModeColumnName
	PreviewModeColumnName *string `json:"previewModeColumnName,omitempty"`

	// IsDeliveredAsPreview
	IsDeliveredAsPreview *bool `json:"isDeliveredAsPreview,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outboundvoicecampaignpostcontactvoiceattributes) SetField(field string, fieldValue interface{}) {
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

func (o Outboundvoicecampaignpostcontactvoiceattributes) MarshalJSON() ([]byte, error) {
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
	type Alias Outboundvoicecampaignpostcontactvoiceattributes
	
	return json.Marshal(&struct { 
		ContactPhoneNumber *string `json:"contactPhoneNumber,omitempty"`
		
		ContactPhoneType *string `json:"contactPhoneType,omitempty"`
		
		CallerIdPhoneNumber *string `json:"callerIdPhoneNumber,omitempty"`
		
		CallerIdName *string `json:"callerIdName,omitempty"`
		
		AgentOwnedColumnName *string `json:"agentOwnedColumnName,omitempty"`
		
		PreviewModeColumnName *string `json:"previewModeColumnName,omitempty"`
		
		IsDeliveredAsPreview *bool `json:"isDeliveredAsPreview,omitempty"`
		Alias
	}{ 
		ContactPhoneNumber: o.ContactPhoneNumber,
		
		ContactPhoneType: o.ContactPhoneType,
		
		CallerIdPhoneNumber: o.CallerIdPhoneNumber,
		
		CallerIdName: o.CallerIdName,
		
		AgentOwnedColumnName: o.AgentOwnedColumnName,
		
		PreviewModeColumnName: o.PreviewModeColumnName,
		
		IsDeliveredAsPreview: o.IsDeliveredAsPreview,
		Alias:    (Alias)(o),
	})
}

func (o *Outboundvoicecampaignpostcontactvoiceattributes) UnmarshalJSON(b []byte) error {
	var OutboundvoicecampaignpostcontactvoiceattributesMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundvoicecampaignpostcontactvoiceattributesMap)
	if err != nil {
		return err
	}
	
	if ContactPhoneNumber, ok := OutboundvoicecampaignpostcontactvoiceattributesMap["contactPhoneNumber"].(string); ok {
		o.ContactPhoneNumber = &ContactPhoneNumber
	}
    
	if ContactPhoneType, ok := OutboundvoicecampaignpostcontactvoiceattributesMap["contactPhoneType"].(string); ok {
		o.ContactPhoneType = &ContactPhoneType
	}
    
	if CallerIdPhoneNumber, ok := OutboundvoicecampaignpostcontactvoiceattributesMap["callerIdPhoneNumber"].(string); ok {
		o.CallerIdPhoneNumber = &CallerIdPhoneNumber
	}
    
	if CallerIdName, ok := OutboundvoicecampaignpostcontactvoiceattributesMap["callerIdName"].(string); ok {
		o.CallerIdName = &CallerIdName
	}
    
	if AgentOwnedColumnName, ok := OutboundvoicecampaignpostcontactvoiceattributesMap["agentOwnedColumnName"].(string); ok {
		o.AgentOwnedColumnName = &AgentOwnedColumnName
	}
    
	if PreviewModeColumnName, ok := OutboundvoicecampaignpostcontactvoiceattributesMap["previewModeColumnName"].(string); ok {
		o.PreviewModeColumnName = &PreviewModeColumnName
	}
    
	if IsDeliveredAsPreview, ok := OutboundvoicecampaignpostcontactvoiceattributesMap["isDeliveredAsPreview"].(bool); ok {
		o.IsDeliveredAsPreview = &IsDeliveredAsPreview
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundvoicecampaignpostcontactvoiceattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
