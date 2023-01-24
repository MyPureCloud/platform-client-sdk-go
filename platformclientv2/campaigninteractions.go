package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaigninteractions
type Campaigninteractions struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Campaign
	Campaign *Domainentityref `json:"campaign,omitempty"`

	// PendingInteractions
	PendingInteractions *[]Campaigninteraction `json:"pendingInteractions,omitempty"`

	// ProceedingInteractions
	ProceedingInteractions *[]Campaigninteraction `json:"proceedingInteractions,omitempty"`

	// PreviewingInteractions
	PreviewingInteractions *[]Campaigninteraction `json:"previewingInteractions,omitempty"`

	// InteractingInteractions
	InteractingInteractions *[]Campaigninteraction `json:"interactingInteractions,omitempty"`

	// ScheduledInteractions
	ScheduledInteractions *[]Campaigninteraction `json:"scheduledInteractions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaigninteractions) SetField(field string, fieldValue interface{}) {
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

func (o Campaigninteractions) MarshalJSON() ([]byte, error) {
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
	type Alias Campaigninteractions
	
	return json.Marshal(&struct { 
		Campaign *Domainentityref `json:"campaign,omitempty"`
		
		PendingInteractions *[]Campaigninteraction `json:"pendingInteractions,omitempty"`
		
		ProceedingInteractions *[]Campaigninteraction `json:"proceedingInteractions,omitempty"`
		
		PreviewingInteractions *[]Campaigninteraction `json:"previewingInteractions,omitempty"`
		
		InteractingInteractions *[]Campaigninteraction `json:"interactingInteractions,omitempty"`
		
		ScheduledInteractions *[]Campaigninteraction `json:"scheduledInteractions,omitempty"`
		Alias
	}{ 
		Campaign: o.Campaign,
		
		PendingInteractions: o.PendingInteractions,
		
		ProceedingInteractions: o.ProceedingInteractions,
		
		PreviewingInteractions: o.PreviewingInteractions,
		
		InteractingInteractions: o.InteractingInteractions,
		
		ScheduledInteractions: o.ScheduledInteractions,
		Alias:    (Alias)(o),
	})
}

func (o *Campaigninteractions) UnmarshalJSON(b []byte) error {
	var CampaigninteractionsMap map[string]interface{}
	err := json.Unmarshal(b, &CampaigninteractionsMap)
	if err != nil {
		return err
	}
	
	if Campaign, ok := CampaigninteractionsMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if PendingInteractions, ok := CampaigninteractionsMap["pendingInteractions"].([]interface{}); ok {
		PendingInteractionsString, _ := json.Marshal(PendingInteractions)
		json.Unmarshal(PendingInteractionsString, &o.PendingInteractions)
	}
	
	if ProceedingInteractions, ok := CampaigninteractionsMap["proceedingInteractions"].([]interface{}); ok {
		ProceedingInteractionsString, _ := json.Marshal(ProceedingInteractions)
		json.Unmarshal(ProceedingInteractionsString, &o.ProceedingInteractions)
	}
	
	if PreviewingInteractions, ok := CampaigninteractionsMap["previewingInteractions"].([]interface{}); ok {
		PreviewingInteractionsString, _ := json.Marshal(PreviewingInteractions)
		json.Unmarshal(PreviewingInteractionsString, &o.PreviewingInteractions)
	}
	
	if InteractingInteractions, ok := CampaigninteractionsMap["interactingInteractions"].([]interface{}); ok {
		InteractingInteractionsString, _ := json.Marshal(InteractingInteractions)
		json.Unmarshal(InteractingInteractionsString, &o.InteractingInteractions)
	}
	
	if ScheduledInteractions, ok := CampaigninteractionsMap["scheduledInteractions"].([]interface{}); ok {
		ScheduledInteractionsString, _ := json.Marshal(ScheduledInteractions)
		json.Unmarshal(ScheduledInteractionsString, &o.ScheduledInteractions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaigninteractions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
