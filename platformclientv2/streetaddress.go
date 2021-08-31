package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Streetaddress
type Streetaddress struct { 
	// Country - 2 Letter Country code, like US or GB
	Country *string `json:"country,omitempty"`


	// A1 - State or Province
	A1 *string `json:"A1,omitempty"`


	// A3 - City or township
	A3 *string `json:"A3,omitempty"`


	// RD - Number and street
	RD *string `json:"RD,omitempty"`


	// HNO - House Number
	HNO *string `json:"HNO,omitempty"`


	// LOC - extra location info like suite 300
	LOC *string `json:"LOC,omitempty"`


	// NAM - Name of the customer
	NAM *string `json:"NAM,omitempty"`


	// PC - Postal code
	PC *string `json:"PC,omitempty"`

}

func (o *Streetaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Streetaddress
	
	return json.Marshal(&struct { 
		Country *string `json:"country,omitempty"`
		
		A1 *string `json:"A1,omitempty"`
		
		A3 *string `json:"A3,omitempty"`
		
		RD *string `json:"RD,omitempty"`
		
		HNO *string `json:"HNO,omitempty"`
		
		LOC *string `json:"LOC,omitempty"`
		
		NAM *string `json:"NAM,omitempty"`
		
		PC *string `json:"PC,omitempty"`
		*Alias
	}{ 
		Country: o.Country,
		
		A1: o.A1,
		
		A3: o.A3,
		
		RD: o.RD,
		
		HNO: o.HNO,
		
		LOC: o.LOC,
		
		NAM: o.NAM,
		
		PC: o.PC,
		Alias:    (*Alias)(o),
	})
}

func (o *Streetaddress) UnmarshalJSON(b []byte) error {
	var StreetaddressMap map[string]interface{}
	err := json.Unmarshal(b, &StreetaddressMap)
	if err != nil {
		return err
	}
	
	if Country, ok := StreetaddressMap["country"].(string); ok {
		o.Country = &Country
	}
	
	if A1, ok := StreetaddressMap["A1"].(string); ok {
		o.A1 = &A1
	}
	
	if A3, ok := StreetaddressMap["A3"].(string); ok {
		o.A3 = &A3
	}
	
	if RD, ok := StreetaddressMap["RD"].(string); ok {
		o.RD = &RD
	}
	
	if HNO, ok := StreetaddressMap["HNO"].(string); ok {
		o.HNO = &HNO
	}
	
	if LOC, ok := StreetaddressMap["LOC"].(string); ok {
		o.LOC = &LOC
	}
	
	if NAM, ok := StreetaddressMap["NAM"].(string); ok {
		o.NAM = &NAM
	}
	
	if PC, ok := StreetaddressMap["PC"].(string); ok {
		o.PC = &PC
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Streetaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
