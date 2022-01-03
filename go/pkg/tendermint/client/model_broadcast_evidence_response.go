/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BroadcastEvidenceResponse struct for BroadcastEvidenceResponse
type BroadcastEvidenceResponse struct {
	Error *string `json:"error,omitempty"`
	Result *string `json:"result,omitempty"`
	Id int32 `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

// NewBroadcastEvidenceResponse instantiates a new BroadcastEvidenceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastEvidenceResponse(id int32, jsonrpc string) *BroadcastEvidenceResponse {
	this := BroadcastEvidenceResponse{}
	this.Id = id
	this.Jsonrpc = jsonrpc
	return &this
}

// NewBroadcastEvidenceResponseWithDefaults instantiates a new BroadcastEvidenceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastEvidenceResponseWithDefaults() *BroadcastEvidenceResponse {
	this := BroadcastEvidenceResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BroadcastEvidenceResponse) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastEvidenceResponse) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BroadcastEvidenceResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BroadcastEvidenceResponse) SetError(v string) {
	o.Error = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *BroadcastEvidenceResponse) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastEvidenceResponse) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *BroadcastEvidenceResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *BroadcastEvidenceResponse) SetResult(v string) {
	o.Result = &v
}

// GetId returns the Id field value
func (o *BroadcastEvidenceResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BroadcastEvidenceResponse) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BroadcastEvidenceResponse) SetId(v int32) {
	o.Id = v
}

// GetJsonrpc returns the Jsonrpc field value
func (o *BroadcastEvidenceResponse) GetJsonrpc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Jsonrpc
}

// GetJsonrpcOk returns a tuple with the Jsonrpc field value
// and a boolean to check if the value has been set.
func (o *BroadcastEvidenceResponse) GetJsonrpcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Jsonrpc, true
}

// SetJsonrpc sets field value
func (o *BroadcastEvidenceResponse) SetJsonrpc(v string) {
	o.Jsonrpc = v
}

func (o BroadcastEvidenceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["jsonrpc"] = o.Jsonrpc
	}
	return json.Marshal(toSerialize)
}

type NullableBroadcastEvidenceResponse struct {
	value *BroadcastEvidenceResponse
	isSet bool
}

func (v NullableBroadcastEvidenceResponse) Get() *BroadcastEvidenceResponse {
	return v.value
}

func (v *NullableBroadcastEvidenceResponse) Set(val *BroadcastEvidenceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastEvidenceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastEvidenceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastEvidenceResponse(val *BroadcastEvidenceResponse) *NullableBroadcastEvidenceResponse {
	return &NullableBroadcastEvidenceResponse{value: val, isSet: true}
}

func (v NullableBroadcastEvidenceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastEvidenceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


