package platformclientv2_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/mypurecloud/platformclientv2"
)

type testConfig struct {
	environment         string
	clientID            string
	clientSecret        string
	debug               bool
	userEmail           string
	usersAPI            *platformclientv2.UsersApi
	userID              string
	userName            string
	userDepartment      string
	userProfileSkill    string
	busyPresenceID      string
	availablePresenceID string
}

type testSerializationStruct struct {
	IntProp          int32     `json:"int,omitempty"`
	IntPropArr       []int32   `json:"intArr,omitempty"`
	IntPropPtr       *int32    `json:"intPtr,omitempty"`
	IntPropArrPtr    *[]int32  `json:"intArrPtr,omitempty"`
	StringProp       string    `json:"string,omitempty"`
	StringPropArr    []string  `json:"stringArr,omitempty"`
	StringPropPtr    *string   `json:"stringPtr,omitempty"`
	StringPropArrPtr *[]string `json:"stringArrPtr,omitempty"`
	BoolProp         bool      `json:"bool,omitempty"`
	BoolPropArr      []bool    `json:"boolArr,omitempty"`
	BoolPropPtr      *bool     `json:"boolPtr,omitempty"`
	BoolPropArrPtr   *[]bool   `json:"boolArrPtr,omitempty"`
}

var config testConfig

func TestEnvVars(t *testing.T) {
	// Get
	config = testConfig{
		environment:         "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT"),
		clientID:            os.Getenv("PURECLOUD_CLIENT_ID"),
		clientSecret:        os.Getenv("PURECLOUD_CLIENT_SECRET"),
		userName:            "GO SDK Tester",
		userDepartment:      "Ministry of Testing",
		userProfileSkill:    "Testmaster",
		busyPresenceID:      "31fe3bac-dea6-44b7-bed7-47f91660a1a0",
		availablePresenceID: "6a3af858-942f-489d-9700-5f9bcdcdae9b",
		debug:               false,
	}
	config.userEmail = fmt.Sprintf("%v@%v", uuid.New().String(), config.environment[12:])

	// Check
	if config.environment == "" {
		t.Error("Not set: PURECLOUD_ENVIRONMENT")
	}
	if config.clientID == "" {
		t.Error("Not set: PURECLOUD_CLIENT_ID")
	}
	if config.clientSecret == "" {
		t.Error("Not set: PURECLOUD_CLIENT_SECRET")
	}
	if config.userEmail == "@"+config.environment[12:] {
		t.Error("Invalid user email")
	}

	// Setup
	platformclientv2.GetDefaultConfiguration().BasePath = config.environment
	platformclientv2.GetDefaultConfiguration().SetDebug(config.debug)
	config.usersAPI = platformclientv2.NewUsersApi()

	// Log
	t.Logf("Enviornment: %v", config.environment)
	t.Logf("clientID: %v", config.clientID)
	t.Logf("userEmail: %v", config.userEmail)
}

func TestDefaultValueSerialization(t *testing.T) {
	expected := `{"intPtr":0,"intArrPtr":[],"stringPtr":"","stringArrPtr":[],"boolPtr":false,"boolArrPtr":[]}`
	intPropArrPtr := make([]int32, 0)
	stringPropArrPtr := make([]string, 0)
	boolPropArrPtr := make([]bool, 0)
	v := testSerializationStruct{
		IntProp:          0,
		IntPropArr:       make([]int32, 0),
		IntPropPtr:       platformclientv2.Int32(0),
		IntPropArrPtr:    &intPropArrPtr,
		StringProp:       "",
		StringPropArr:    make([]string, 0),
		StringPropPtr:    platformclientv2.String(""),
		StringPropArrPtr: &stringPropArrPtr,
		BoolProp:         false,
		BoolPropArr:      make([]bool, 0),
		BoolPropPtr:      platformclientv2.Bool(false),
		BoolPropArrPtr:   &boolPropArrPtr,
	}
	j, _ := json.Marshal(v)
	s := string(j)
	if s != expected {
		t.Log("testSerializationStruct did not serialize correctly")
		t.Logf("Expected: %v", expected)
		t.Logf("Actual:   %v", s)
		t.FailNow()
	}
}

func TestValueSerialization(t *testing.T) {
	expected := `{"int":10,"intArr":[0,0],"intPtr":10,"intArrPtr":[0,0],"string":"asdf","stringArr":["",""],"stringPtr":"asdf","stringArrPtr":["",""],"bool":true,"boolArr":[false,false],"boolPtr":true,"boolArrPtr":[false,false]}`
	intPropArrPtr := make([]int32, 2)
	stringPropArrPtr := make([]string, 2)
	boolPropArrPtr := make([]bool, 2)
	v := testSerializationStruct{
		IntProp:          10,
		IntPropArr:       make([]int32, 2),
		IntPropPtr:       platformclientv2.Int32(10),
		IntPropArrPtr:    &intPropArrPtr,
		StringProp:       "asdf",
		StringPropArr:    make([]string, 2),
		StringPropPtr:    platformclientv2.String("asdf"),
		StringPropArrPtr: &stringPropArrPtr,
		BoolProp:         true,
		BoolPropArr:      make([]bool, 2),
		BoolPropPtr:      platformclientv2.Bool(true),
		BoolPropArrPtr:   &boolPropArrPtr,
	}
	j, _ := json.Marshal(v)
	s := string(j)
	if s != expected {
		t.Log("testSerializationStruct did not serialize correctly")
		t.Logf("Expected: %v", expected)
		t.Logf("Actual:   %v", s)
		t.FailNow()
	}
}

func TestAuthentication(t *testing.T) {
	err := platformclientv2.GetDefaultConfiguration().AuthorizeClientCredentials(config.clientID, config.clientSecret)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateUser(t *testing.T) {
	// Create user
	password := uuid.New().String() + "!@#$1234asdfASDF"
	newUser := platformclientv2.Createuser{Name: &config.userName, Email: &config.userEmail, Password: &password}

	user, response, err := config.usersAPI.PostUsers(newUser)
	if err != nil {
		t.Error(err)
	} else if response != nil && response.Error != nil {
		t.Error(response.Error)
	} else {
		// Validate response
		if *user.Name != config.userName {
			t.Error("Data mismatch: user.Name")
		}
		if *user.Email != config.userEmail {
			t.Error("Data mismatch: user.Email")
		}

		// Store user ID
		config.userID = *user.Id
		t.Logf("New user's ID: %v", *user.Id)
	}
}

func TestUpdateUser(t *testing.T) {
	// Update user
	updateUser := platformclientv2.Updateuser{Department: &config.userDepartment, Version: platformclientv2.Int32(1)}

	user, response, err := config.usersAPI.PatchUser(config.userID, updateUser)
	if err != nil {
		t.Error(err)
	} else if response != nil && response.Error != nil {
		t.Error(response.Error)
	} else {
		// Validate response
		if *user.Name != config.userName {
			t.Error("Data mismatch: user.Name")
		}
		if *user.Email != config.userEmail {
			t.Error("Data mismatch: user.Email")
		}
		if *user.Department != config.userDepartment {
			t.Error("Data mismatch: user.Department")
		}
	}
}

func TestSetProfileSkills(t *testing.T) {
	// Update user
	skills, response, err := config.usersAPI.PutUserProfileskills(config.userID, []string{config.userProfileSkill})
	if err != nil {
		t.Error(err)
	} else if response != nil && response.Error != nil {
		t.Error(response.Error)
	} else {
		// Validate response
		if len(skills) != 1 {
			t.Errorf("Skills array contained the wrong number of elements. Expected 1: %v", skills)
		} else if skills[0] != config.userProfileSkill {
			t.Errorf("Skill did not match. Expected %v, actual: %v", config.userProfileSkill, skills[0])
		}
	}
}

func TestGetUser(t *testing.T) {
	// Get user
	user, response, err := config.usersAPI.GetUser(config.userID, []string{"profileSkills"}, "")
	if err != nil {
		t.Error(err)
	} else if response != nil && response.Error != nil {
		t.Error(response.Error)
	} else {
		// Validate response
		if *user.Name != config.userName {
			t.Error("Data mismatch: user.Name")
		}
		if *user.Email != config.userEmail {
			t.Error("Data mismatch: user.Email")
		}
		if *user.Department != config.userDepartment {
			t.Error("Data mismatch: user.Department")
		}
		if user.ProfileSkills == nil || len(*user.ProfileSkills) != 1 || (*user.ProfileSkills)[0] != config.userProfileSkill {
			t.Error("Data mismatch: user.ProfileSkills")
		}
	}
}

func TestDeleteUser(t *testing.T) {
	// Delete user
	_, response, err := config.usersAPI.DeleteUser(config.userID)
	if err != nil {
		t.Error(err)
	} else if response != nil && response.Error != nil {
		t.Error(response.Error)
	}
}

func Example_authorizeDefaultConfiguration() {
	// Use the default config instance and retrieve settings from env vars
	config := platformclientv2.GetDefaultConfiguration()
	config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com
	clientID := os.Getenv("PURECLOUD_CLIENT_ID")
	clientSecret := os.Getenv("PURECLOUD_CLIENT_SECRET")

	// Authorize the configuration
	err := config.AuthorizeClientCredentials(clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	// Create an API instance using the default config
	usersAPI := platformclientv2.NewUsersApi()
	fmt.Printf("Users API type: %v", reflect.TypeOf(usersAPI).String())
	// Output: Users API type: UsersAPI

	// Make requests using usersAPI
}

func Example_authorizeNewConfiguration() {
	// Create a new config instance and retrieve settings from env vars
	config := platformclientv2.NewConfiguration()
	config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com
	clientID := os.Getenv("PURECLOUD_CLIENT_ID")
	clientSecret := os.Getenv("PURECLOUD_CLIENT_SECRET")

	// Authorize the configuration
	err := config.AuthorizeClientCredentials(clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	// Create an API instance using the config instance
	usersAPI := platformclientv2.NewUsersApiWithConfig(config)
	fmt.Printf("Users API type: %v", reflect.TypeOf(usersAPI).String())
	// Output: Users API type: UsersAPI

	// Make requests using usersAPI
}
