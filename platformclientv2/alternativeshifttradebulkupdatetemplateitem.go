package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Alternativeshifttradebulkupdatetemplateitem
type Alternativeshifttradebulkupdatetemplateitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TradeId - The ID of this alternative shift trade
	TradeId *string `json:"tradeId,omitempty"`

	// State - The current state of this alternative shift trade request
	State *string `json:"state,omitempty"`

	// FailureReason - The reason the update failed, if applicable
	FailureReason *string `json:"failureReason,omitempty"`

	// AdminDateReviewed - The timestamp of when the trade request was manually reviewed by an admin in ISO-8601 format
	AdminDateReviewed *time.Time `json:"adminDateReviewed,omitempty"`

	// AdminReviewedBy - The admin who manually reviewed this alternative shift trade after system denial
	AdminReviewedBy *Userreference `json:"adminReviewedBy,omitempty"`

	// Metadata - Version metadata for this alternative shift trade
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Alternativeshifttradebulkupdatetemplateitem) SetField(field string, fieldValue interface{}) {
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

func (o Alternativeshifttradebulkupdatetemplateitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "AdminDateReviewed", }
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
	type Alias Alternativeshifttradebulkupdatetemplateitem
	
	AdminDateReviewed := new(string)
	if o.AdminDateReviewed != nil {
		
		*AdminDateReviewed = timeutil.Strftime(o.AdminDateReviewed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AdminDateReviewed = nil
	}
	
	return json.Marshal(&struct { 
		TradeId *string `json:"tradeId,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		
		AdminDateReviewed *string `json:"adminDateReviewed,omitempty"`
		
		AdminReviewedBy *Userreference `json:"adminReviewedBy,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		TradeId: o.TradeId,
		
		State: o.State,
		
		FailureReason: o.FailureReason,
		
		AdminDateReviewed: AdminDateReviewed,
		
		AdminReviewedBy: o.AdminReviewedBy,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Alternativeshifttradebulkupdatetemplateitem) UnmarshalJSON(b []byte) error {
	var AlternativeshifttradebulkupdatetemplateitemMap map[string]interface{}
	err := json.Unmarshal(b, &AlternativeshifttradebulkupdatetemplateitemMap)
	if err != nil {
		return err
	}
	
	if TradeId, ok := AlternativeshifttradebulkupdatetemplateitemMap["tradeId"].(string); ok {
		o.TradeId = &TradeId
	}
    
	if State, ok := AlternativeshifttradebulkupdatetemplateitemMap["state"].(string); ok {
		o.State = &State
	}
    
	if FailureReason, ok := AlternativeshifttradebulkupdatetemplateitemMap["failureReason"].(string); ok {
		o.FailureReason = &FailureReason
	}
    
	if adminDateReviewedString, ok := AlternativeshifttradebulkupdatetemplateitemMap["adminDateReviewed"].(string); ok {
		AdminDateReviewed, _ := time.Parse("2006-01-02T15:04:05.999999Z", adminDateReviewedString)
		o.AdminDateReviewed = &AdminDateReviewed
	}
	
	if AdminReviewedBy, ok := AlternativeshifttradebulkupdatetemplateitemMap["adminReviewedBy"].(map[string]interface{}); ok {
		AdminReviewedByString, _ := json.Marshal(AdminReviewedBy)
		json.Unmarshal(AdminReviewedByString, &o.AdminReviewedBy)
	}
	
	if Metadata, ok := AlternativeshifttradebulkupdatetemplateitemMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Alternativeshifttradebulkupdatetemplateitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
