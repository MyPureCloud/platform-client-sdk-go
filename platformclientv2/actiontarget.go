package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actiontarget
type Actiontarget struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// UserData - Additional user data associated with the target in key/value format.
	UserData *[]Keyvalue `json:"userData,omitempty"`


	// SupportedMediaTypes - Supported media types of the target.
	SupportedMediaTypes *[]string `json:"supportedMediaTypes,omitempty"`


	// State - Indicates the state of the target.
	State *string `json:"state,omitempty"`


	// Description - Description of the target.
	Description *string `json:"description,omitempty"`


	// ServiceLevel - Service Level of the action target. Chat offers for the target will be throttled with the aim of achieving this service level.
	ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`


	// ShortAbandonThreshold - Indicates the non-default short abandon threshold
	ShortAbandonThreshold *int `json:"shortAbandonThreshold,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate - The date the target was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - The date the target was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (o *Actiontarget) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actiontarget
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UserData *[]Keyvalue `json:"userData,omitempty"`
		
		SupportedMediaTypes *[]string `json:"supportedMediaTypes,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`
		
		ShortAbandonThreshold *int `json:"shortAbandonThreshold,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		UserData: o.UserData,
		
		SupportedMediaTypes: o.SupportedMediaTypes,
		
		State: o.State,
		
		Description: o.Description,
		
		ServiceLevel: o.ServiceLevel,
		
		ShortAbandonThreshold: o.ShortAbandonThreshold,
		
		SelfUri: o.SelfUri,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Actiontarget) UnmarshalJSON(b []byte) error {
	var ActiontargetMap map[string]interface{}
	err := json.Unmarshal(b, &ActiontargetMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActiontargetMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ActiontargetMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if UserData, ok := ActiontargetMap["userData"].([]interface{}); ok {
		UserDataString, _ := json.Marshal(UserData)
		json.Unmarshal(UserDataString, &o.UserData)
	}
	
	if SupportedMediaTypes, ok := ActiontargetMap["supportedMediaTypes"].([]interface{}); ok {
		SupportedMediaTypesString, _ := json.Marshal(SupportedMediaTypes)
		json.Unmarshal(SupportedMediaTypesString, &o.SupportedMediaTypes)
	}
	
	if State, ok := ActiontargetMap["state"].(string); ok {
		o.State = &State
	}
    
	if Description, ok := ActiontargetMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ServiceLevel, ok := ActiontargetMap["serviceLevel"].(map[string]interface{}); ok {
		ServiceLevelString, _ := json.Marshal(ServiceLevel)
		json.Unmarshal(ServiceLevelString, &o.ServiceLevel)
	}
	
	if ShortAbandonThreshold, ok := ActiontargetMap["shortAbandonThreshold"].(float64); ok {
		ShortAbandonThresholdInt := int(ShortAbandonThreshold)
		o.ShortAbandonThreshold = &ShortAbandonThresholdInt
	}
	
	if SelfUri, ok := ActiontargetMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if createdDateString, ok := ActiontargetMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := ActiontargetMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actiontarget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
