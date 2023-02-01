package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingsticker
type Messagingsticker struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ProviderStickerId - The sticker Id of the sticker, assigned by the sticker provider.
	ProviderStickerId *int `json:"providerStickerId,omitempty"`

	// ProviderPackageId - The package Id of the sticker, assigned by the sticker provider.
	ProviderPackageId *int `json:"providerPackageId,omitempty"`

	// PackageName - The package name of the sticker, assigned by the sticker provider.
	PackageName *string `json:"packageName,omitempty"`

	// MessengerType - The type of the messenger provider.
	MessengerType *string `json:"messengerType,omitempty"`

	// StickerType - The type of the sticker.
	StickerType *string `json:"stickerType,omitempty"`

	// ProviderVersion - The version of the sticker, assigned by the provider.
	ProviderVersion *int `json:"providerVersion,omitempty"`

	// UriLocation
	UriLocation *string `json:"uriLocation,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messagingsticker) SetField(field string, fieldValue interface{}) {
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

func (o Messagingsticker) MarshalJSON() ([]byte, error) {
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
	type Alias Messagingsticker
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ProviderStickerId *int `json:"providerStickerId,omitempty"`
		
		ProviderPackageId *int `json:"providerPackageId,omitempty"`
		
		PackageName *string `json:"packageName,omitempty"`
		
		MessengerType *string `json:"messengerType,omitempty"`
		
		StickerType *string `json:"stickerType,omitempty"`
		
		ProviderVersion *int `json:"providerVersion,omitempty"`
		
		UriLocation *string `json:"uriLocation,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ProviderStickerId: o.ProviderStickerId,
		
		ProviderPackageId: o.ProviderPackageId,
		
		PackageName: o.PackageName,
		
		MessengerType: o.MessengerType,
		
		StickerType: o.StickerType,
		
		ProviderVersion: o.ProviderVersion,
		
		UriLocation: o.UriLocation,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Messagingsticker) UnmarshalJSON(b []byte) error {
	var MessagingstickerMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingstickerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MessagingstickerMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MessagingstickerMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ProviderStickerId, ok := MessagingstickerMap["providerStickerId"].(float64); ok {
		ProviderStickerIdInt := int(ProviderStickerId)
		o.ProviderStickerId = &ProviderStickerIdInt
	}
	
	if ProviderPackageId, ok := MessagingstickerMap["providerPackageId"].(float64); ok {
		ProviderPackageIdInt := int(ProviderPackageId)
		o.ProviderPackageId = &ProviderPackageIdInt
	}
	
	if PackageName, ok := MessagingstickerMap["packageName"].(string); ok {
		o.PackageName = &PackageName
	}
    
	if MessengerType, ok := MessagingstickerMap["messengerType"].(string); ok {
		o.MessengerType = &MessengerType
	}
    
	if StickerType, ok := MessagingstickerMap["stickerType"].(string); ok {
		o.StickerType = &StickerType
	}
    
	if ProviderVersion, ok := MessagingstickerMap["providerVersion"].(float64); ok {
		ProviderVersionInt := int(ProviderVersion)
		o.ProviderVersion = &ProviderVersionInt
	}
	
	if UriLocation, ok := MessagingstickerMap["uriLocation"].(string); ok {
		o.UriLocation = &UriLocation
	}
    
	if SelfUri, ok := MessagingstickerMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingsticker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
