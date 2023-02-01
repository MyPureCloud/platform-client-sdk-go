package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignprogresseventcampaignprogress
type Outboundmessagingmessagingcampaignprogresseventcampaignprogress struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Campaign
	Campaign *Outboundmessagingmessagingcampaignprogresseventurireference `json:"campaign,omitempty"`

	// NumberOfContactsCalled - The number of contacts that have been called so far
	NumberOfContactsCalled *float32 `json:"numberOfContactsCalled,omitempty"`

	// NumberOfContactsMessaged - The number of contacts that have been messaged so far
	NumberOfContactsMessaged *float32 `json:"numberOfContactsMessaged,omitempty"`

	// TotalNumberOfContacts - The total number of contacts in the contact list
	TotalNumberOfContacts *float32 `json:"totalNumberOfContacts,omitempty"`

	// Percentage - numberOfContactsContacted/totalNumberOfContacts*100
	Percentage *int `json:"percentage,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outboundmessagingmessagingcampaignprogresseventcampaignprogress) SetField(field string, fieldValue interface{}) {
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

func (o Outboundmessagingmessagingcampaignprogresseventcampaignprogress) MarshalJSON() ([]byte, error) {
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
	type Alias Outboundmessagingmessagingcampaignprogresseventcampaignprogress
	
	return json.Marshal(&struct { 
		Campaign *Outboundmessagingmessagingcampaignprogresseventurireference `json:"campaign,omitempty"`
		
		NumberOfContactsCalled *float32 `json:"numberOfContactsCalled,omitempty"`
		
		NumberOfContactsMessaged *float32 `json:"numberOfContactsMessaged,omitempty"`
		
		TotalNumberOfContacts *float32 `json:"totalNumberOfContacts,omitempty"`
		
		Percentage *int `json:"percentage,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		Alias
	}{ 
		Campaign: o.Campaign,
		
		NumberOfContactsCalled: o.NumberOfContactsCalled,
		
		NumberOfContactsMessaged: o.NumberOfContactsMessaged,
		
		TotalNumberOfContacts: o.TotalNumberOfContacts,
		
		Percentage: o.Percentage,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignprogresseventcampaignprogress) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap)
	if err != nil {
		return err
	}
	
	if Campaign, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if NumberOfContactsCalled, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["numberOfContactsCalled"].(float64); ok {
		NumberOfContactsCalledFloat32 := float32(NumberOfContactsCalled)
		o.NumberOfContactsCalled = &NumberOfContactsCalledFloat32
	}
    
	if NumberOfContactsMessaged, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["numberOfContactsMessaged"].(float64); ok {
		NumberOfContactsMessagedFloat32 := float32(NumberOfContactsMessaged)
		o.NumberOfContactsMessaged = &NumberOfContactsMessagedFloat32
	}
    
	if TotalNumberOfContacts, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["totalNumberOfContacts"].(float64); ok {
		TotalNumberOfContactsFloat32 := float32(TotalNumberOfContacts)
		o.TotalNumberOfContacts = &TotalNumberOfContactsFloat32
	}
    
	if Percentage, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["percentage"].(float64); ok {
		PercentageInt := int(Percentage)
		o.Percentage = &PercentageInt
	}
	
	if AdditionalProperties, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignprogresseventcampaignprogress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
