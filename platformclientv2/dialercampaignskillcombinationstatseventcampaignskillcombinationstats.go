package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignskillcombinationstatseventcampaignskillcombinationstats
type Dialercampaignskillcombinationstatseventcampaignskillcombinationstats struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Campaign
	Campaign *Dialercampaignskillcombinationstatseventurireference `json:"campaign,omitempty"`

	// ContactList - A UriReference for a resource
	ContactList *Dialercampaignskillcombinationstatseventurireference `json:"contactList,omitempty"`

	// TotalSkillCombinations - The total number of unique skill combinations
	TotalSkillCombinations *int `json:"totalSkillCombinations,omitempty"`

	// TotalRemainingContacts - The total number of remaining contacts
	TotalRemainingContacts *int `json:"totalRemainingContacts,omitempty"`

	// TotalProcessedContacts - The total number of processed contacts
	TotalProcessedContacts *int `json:"totalProcessedContacts,omitempty"`

	// SkillCombinationDetails - Details for each skill combination
	SkillCombinationDetails *[]Dialercampaignskillcombinationstatseventskillcombinationdetailnotification `json:"skillCombinationDetails,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialercampaignskillcombinationstatseventcampaignskillcombinationstats) SetField(field string, fieldValue interface{}) {
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

func (o Dialercampaignskillcombinationstatseventcampaignskillcombinationstats) MarshalJSON() ([]byte, error) {
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
	type Alias Dialercampaignskillcombinationstatseventcampaignskillcombinationstats
	
	return json.Marshal(&struct { 
		Campaign *Dialercampaignskillcombinationstatseventurireference `json:"campaign,omitempty"`
		
		ContactList *Dialercampaignskillcombinationstatseventurireference `json:"contactList,omitempty"`
		
		TotalSkillCombinations *int `json:"totalSkillCombinations,omitempty"`
		
		TotalRemainingContacts *int `json:"totalRemainingContacts,omitempty"`
		
		TotalProcessedContacts *int `json:"totalProcessedContacts,omitempty"`
		
		SkillCombinationDetails *[]Dialercampaignskillcombinationstatseventskillcombinationdetailnotification `json:"skillCombinationDetails,omitempty"`
		Alias
	}{ 
		Campaign: o.Campaign,
		
		ContactList: o.ContactList,
		
		TotalSkillCombinations: o.TotalSkillCombinations,
		
		TotalRemainingContacts: o.TotalRemainingContacts,
		
		TotalProcessedContacts: o.TotalProcessedContacts,
		
		SkillCombinationDetails: o.SkillCombinationDetails,
		Alias:    (Alias)(o),
	})
}

func (o *Dialercampaignskillcombinationstatseventcampaignskillcombinationstats) UnmarshalJSON(b []byte) error {
	var DialercampaignskillcombinationstatseventcampaignskillcombinationstatsMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignskillcombinationstatseventcampaignskillcombinationstatsMap)
	if err != nil {
		return err
	}
	
	if Campaign, ok := DialercampaignskillcombinationstatseventcampaignskillcombinationstatsMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if ContactList, ok := DialercampaignskillcombinationstatseventcampaignskillcombinationstatsMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if TotalSkillCombinations, ok := DialercampaignskillcombinationstatseventcampaignskillcombinationstatsMap["totalSkillCombinations"].(float64); ok {
		TotalSkillCombinationsInt := int(TotalSkillCombinations)
		o.TotalSkillCombinations = &TotalSkillCombinationsInt
	}
	
	if TotalRemainingContacts, ok := DialercampaignskillcombinationstatseventcampaignskillcombinationstatsMap["totalRemainingContacts"].(float64); ok {
		TotalRemainingContactsInt := int(TotalRemainingContacts)
		o.TotalRemainingContacts = &TotalRemainingContactsInt
	}
	
	if TotalProcessedContacts, ok := DialercampaignskillcombinationstatseventcampaignskillcombinationstatsMap["totalProcessedContacts"].(float64); ok {
		TotalProcessedContactsInt := int(TotalProcessedContacts)
		o.TotalProcessedContacts = &TotalProcessedContactsInt
	}
	
	if SkillCombinationDetails, ok := DialercampaignskillcombinationstatseventcampaignskillcombinationstatsMap["skillCombinationDetails"].([]interface{}); ok {
		SkillCombinationDetailsString, _ := json.Marshal(SkillCombinationDetails)
		json.Unmarshal(SkillCombinationDetailsString, &o.SkillCombinationDetails)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignskillcombinationstatseventcampaignskillcombinationstats) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
