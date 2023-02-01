package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignprogress
type Campaignprogress struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Campaign - Identifier of the campaign
	Campaign *Domainentityref `json:"campaign,omitempty"`

	// ContactList - Identifier of the contact list
	ContactList *Domainentityref `json:"contactList,omitempty"`

	// NumberOfContactsCalled - Number of contacts called during the campaign
	NumberOfContactsCalled *int `json:"numberOfContactsCalled,omitempty"`

	// NumberOfContactsMessaged - Number of contacts messaged during the campaign
	NumberOfContactsMessaged *int `json:"numberOfContactsMessaged,omitempty"`

	// TotalNumberOfContacts - Total number of contacts in the campaign
	TotalNumberOfContacts *int `json:"totalNumberOfContacts,omitempty"`

	// Percentage - Percentage of contacts processed during the campaign
	Percentage *int `json:"percentage,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaignprogress) SetField(field string, fieldValue interface{}) {
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

func (o Campaignprogress) MarshalJSON() ([]byte, error) {
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
	type Alias Campaignprogress
	
	return json.Marshal(&struct { 
		Campaign *Domainentityref `json:"campaign,omitempty"`
		
		ContactList *Domainentityref `json:"contactList,omitempty"`
		
		NumberOfContactsCalled *int `json:"numberOfContactsCalled,omitempty"`
		
		NumberOfContactsMessaged *int `json:"numberOfContactsMessaged,omitempty"`
		
		TotalNumberOfContacts *int `json:"totalNumberOfContacts,omitempty"`
		
		Percentage *int `json:"percentage,omitempty"`
		Alias
	}{ 
		Campaign: o.Campaign,
		
		ContactList: o.ContactList,
		
		NumberOfContactsCalled: o.NumberOfContactsCalled,
		
		NumberOfContactsMessaged: o.NumberOfContactsMessaged,
		
		TotalNumberOfContacts: o.TotalNumberOfContacts,
		
		Percentage: o.Percentage,
		Alias:    (Alias)(o),
	})
}

func (o *Campaignprogress) UnmarshalJSON(b []byte) error {
	var CampaignprogressMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignprogressMap)
	if err != nil {
		return err
	}
	
	if Campaign, ok := CampaignprogressMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if ContactList, ok := CampaignprogressMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if NumberOfContactsCalled, ok := CampaignprogressMap["numberOfContactsCalled"].(float64); ok {
		NumberOfContactsCalledInt := int(NumberOfContactsCalled)
		o.NumberOfContactsCalled = &NumberOfContactsCalledInt
	}
	
	if NumberOfContactsMessaged, ok := CampaignprogressMap["numberOfContactsMessaged"].(float64); ok {
		NumberOfContactsMessagedInt := int(NumberOfContactsMessaged)
		o.NumberOfContactsMessaged = &NumberOfContactsMessagedInt
	}
	
	if TotalNumberOfContacts, ok := CampaignprogressMap["totalNumberOfContacts"].(float64); ok {
		TotalNumberOfContactsInt := int(TotalNumberOfContacts)
		o.TotalNumberOfContacts = &TotalNumberOfContactsInt
	}
	
	if Percentage, ok := CampaignprogressMap["percentage"].(float64); ok {
		PercentageInt := int(Percentage)
		o.Percentage = &PercentageInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignprogress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
