package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmoduleinformsteprequest - Learning module inform steps request
type Learningmoduleinformsteprequest struct { 
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

func (o *Learningmoduleinformsteprequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmoduleinformsteprequest
	
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

func (o *Learningmoduleinformsteprequest) UnmarshalJSON(b []byte) error {
	var LearningmoduleinformsteprequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmoduleinformsteprequestMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := LearningmoduleinformsteprequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Name, ok := LearningmoduleinformsteprequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Value, ok := LearningmoduleinformsteprequestMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if SharingUri, ok := LearningmoduleinformsteprequestMap["sharingUri"].(string); ok {
		o.SharingUri = &SharingUri
	}
	
	if ContentType, ok := LearningmoduleinformsteprequestMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
	
	if Order, ok := LearningmoduleinformsteprequestMap["order"].(float64); ok {
		OrderInt := int(Order)
		o.Order = &OrderInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmoduleinformsteprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
