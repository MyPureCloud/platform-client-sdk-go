package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contestprizes
type Contestprizes struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Tier - The Contest Prize tier
	Tier *int `json:"tier,omitempty"`

	// Name - The Contest Prize name
	Name *string `json:"name,omitempty"`

	// Description - The Contest Prize description
	Description *string `json:"description,omitempty"`

	// ImageId - The Contest Prize image id
	ImageId *string `json:"imageId,omitempty"`

	// ImageUrl - The Contest Prize image url
	ImageUrl *string `json:"imageUrl,omitempty"`

	// WinnersCount - The Contest Prize winner Count
	WinnersCount *int `json:"winnersCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contestprizes) SetField(field string, fieldValue interface{}) {
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

func (o Contestprizes) MarshalJSON() ([]byte, error) {
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
	type Alias Contestprizes
	
	return json.Marshal(&struct { 
		Tier *int `json:"tier,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ImageId *string `json:"imageId,omitempty"`
		
		ImageUrl *string `json:"imageUrl,omitempty"`
		
		WinnersCount *int `json:"winnersCount,omitempty"`
		Alias
	}{ 
		Tier: o.Tier,
		
		Name: o.Name,
		
		Description: o.Description,
		
		ImageId: o.ImageId,
		
		ImageUrl: o.ImageUrl,
		
		WinnersCount: o.WinnersCount,
		Alias:    (Alias)(o),
	})
}

func (o *Contestprizes) UnmarshalJSON(b []byte) error {
	var ContestprizesMap map[string]interface{}
	err := json.Unmarshal(b, &ContestprizesMap)
	if err != nil {
		return err
	}
	
	if Tier, ok := ContestprizesMap["tier"].(float64); ok {
		TierInt := int(Tier)
		o.Tier = &TierInt
	}
	
	if Name, ok := ContestprizesMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ContestprizesMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ImageId, ok := ContestprizesMap["imageId"].(string); ok {
		o.ImageId = &ImageId
	}
    
	if ImageUrl, ok := ContestprizesMap["imageUrl"].(string); ok {
		o.ImageUrl = &ImageUrl
	}
    
	if WinnersCount, ok := ContestprizesMap["winnersCount"].(float64); ok {
		WinnersCountInt := int(WinnersCount)
		o.WinnersCount = &WinnersCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contestprizes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
