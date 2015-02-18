package semantics3

import (
	"github.com/mrjones/oauth"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var mockExpected = make([]interface{}, 2)

//==================================================
// Base Mock OAuthConsumer
//==================================================
type MockOAuthConsumer struct{}

func (m MockOAuthConsumer) Get(string, map[string]string,
	*oauth.AccessToken) (*http.Response, error) {

	return nil, nil
}

//==================================================
// Tests
//==================================================
func Test_AddParams(t *testing.T) {
	mockClient := &Client{}
	params := map[string]interface{}{"key": "value"}
	expected := `{"key":"value"}`
	mockClient.AddParams(params)
	assert.Equal(t, mockClient.dataquery, expected)
}

func Test_NewClient(t *testing.T) {
	ret := NewClient("1", "2", "bags")
	assert.Equal(t, ret.endpoint, "bags")
}

type MockOAuthConsumer_2 struct{ MockOAuthConsumer }

func (m MockOAuthConsumer_2) Get(baseURL string, params map[string]string,
	token *oauth.AccessToken) (*http.Response, error) {

	mockExpected[0] = baseURL
	mockExpected[1] = params
	return nil, nil
}

func Test_Get(t *testing.T) {
	// should return error if oauth is not set
	mockClient := &Client{}
	ret, err := mockClient.Get()
	assert.Nil(t, ret)
	assert.Error(t, err)

	// should return response
	mockOAuth := &MockOAuthConsumer_2{}
	mockClient = &Client{endpoint: "bags", dataquery: "query", oauth: mockOAuth}
	ret, err = mockClient.Get()
	assert.Nil(t, ret)
	assert.Nil(t, err)
	assert.Equal(t, mockExpected[0].(string), "https://api.semantics3.com/v1/bags")
	assert.Equal(t, mockExpected[1].(map[string]string), map[string]string{"q": "query"})
}
