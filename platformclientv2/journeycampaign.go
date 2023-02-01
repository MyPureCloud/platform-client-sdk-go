package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeycampaign
type Journeycampaign struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Content - Differentiate ads or links that point to the same URL (e.g. textlink).
	Content *string `json:"content,omitempty"`

	// Medium - Identify a medium such as email or cost-per-click (e.g. CPC).
	Medium *string `json:"medium,omitempty"`

	// Name - Identify a specific product promotion or strategic campaign (e.g. 320banner).
	Name *string `json:"name,omitempty"`

	// Source - Identify a search engine, newsletter name, or other source (e.g. Google).
	Source *string `json:"source,omitempty"`

	// Term - Note the keywords for this ad (e.g. running+shoes).
	Term *string `json:"term,omitempty"`

	// ClickId - The click ID (unique number that is generated when a potential customer clicks on an affiliate link).
	ClickId *string `json:"clickId,omitempty"`

	// Network - The ad network to which the click ID belongs.
	Network *string `json:"network,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeycampaign) SetField(field string, fieldValue interface{}) {
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

func (o Journeycampaign) MarshalJSON() ([]byte, error) {
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
	type Alias Journeycampaign
	
	return json.Marshal(&struct { 
		Content *string `json:"content,omitempty"`
		
		Medium *string `json:"medium,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		Term *string `json:"term,omitempty"`
		
		ClickId *string `json:"clickId,omitempty"`
		
		Network *string `json:"network,omitempty"`
		Alias
	}{ 
		Content: o.Content,
		
		Medium: o.Medium,
		
		Name: o.Name,
		
		Source: o.Source,
		
		Term: o.Term,
		
		ClickId: o.ClickId,
		
		Network: o.Network,
		Alias:    (Alias)(o),
	})
}

func (o *Journeycampaign) UnmarshalJSON(b []byte) error {
	var JourneycampaignMap map[string]interface{}
	err := json.Unmarshal(b, &JourneycampaignMap)
	if err != nil {
		return err
	}
	
	if Content, ok := JourneycampaignMap["content"].(string); ok {
		o.Content = &Content
	}
    
	if Medium, ok := JourneycampaignMap["medium"].(string); ok {
		o.Medium = &Medium
	}
    
	if Name, ok := JourneycampaignMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Source, ok := JourneycampaignMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if Term, ok := JourneycampaignMap["term"].(string); ok {
		o.Term = &Term
	}
    
	if ClickId, ok := JourneycampaignMap["clickId"].(string); ok {
		o.ClickId = &ClickId
	}
    
	if Network, ok := JourneycampaignMap["network"].(string); ok {
		o.Network = &Network
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeycampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
