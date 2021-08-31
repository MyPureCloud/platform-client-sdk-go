package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmoduleinformstep
type Learningmoduleinformstep struct { 
	// VarType - The learning module inform step type
	VarType *string `json:"type,omitempty"`


	// Name - The name of the inform step or content
	Name *string `json:"name,omitempty"`


	// Value - The value for inform step
	Value *string `json:"value,omitempty"`


	// SharingUri - The sharing uri for Content type inform step
	SharingUri *string `json:"sharingUri,omitempty"`


	// ContentType - The document type for Content type Inform step
	ContentType *string `json:"contentType,omitempty"`


	// Order - The order of inform step in a learning module
	Order *int `json:"order,omitempty"`

}

func (o *Learningmoduleinformstep) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmoduleinformstep
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		SharingUri *string `json:"sharingUri,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		Order *int `json:"order,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Name: o.Name,
		
		Value: o.Value,
		
		SharingUri: o.SharingUri,
		
		ContentType: o.ContentType,
		
		Order: o.Order,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningmoduleinformstep) UnmarshalJSON(b []byte) error {
	var LearningmoduleinformstepMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmoduleinformstepMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := LearningmoduleinformstepMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Name, ok := LearningmoduleinformstepMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Value, ok := LearningmoduleinformstepMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if SharingUri, ok := LearningmoduleinformstepMap["sharingUri"].(string); ok {
		o.SharingUri = &SharingUri
	}
	
	if ContentType, ok := LearningmoduleinformstepMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
	
	if Order, ok := LearningmoduleinformstepMap["order"].(float64); ok {
		OrderInt := int(Order)
		o.Order = &OrderInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmoduleinformstep) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
