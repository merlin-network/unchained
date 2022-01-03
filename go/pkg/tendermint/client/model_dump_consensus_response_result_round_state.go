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

// DumpConsensusResponseResultRoundState struct for DumpConsensusResponseResultRoundState
type DumpConsensusResponseResultRoundState struct {
	Height string `json:"height"`
	Round int32 `json:"round"`
	Step int32 `json:"step"`
	StartTime string `json:"start_time"`
	CommitTime string `json:"commit_time"`
	Validators DumpConsensusResponseResultRoundStateValidators `json:"validators"`
	LockedRound int32 `json:"locked_round"`
	ValidRound string `json:"valid_round"`
	Votes []DumpConsensusResponseResultRoundStateVotes `json:"votes"`
	CommitRound int32 `json:"commit_round"`
	LastCommit NullableDumpConsensusResponseResultRoundStateLastCommit `json:"last_commit"`
	LastValidators DumpConsensusResponseResultRoundStateValidators `json:"last_validators"`
	TriggeredTimeoutPrecommit bool `json:"triggered_timeout_precommit"`
}

// NewDumpConsensusResponseResultRoundState instantiates a new DumpConsensusResponseResultRoundState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDumpConsensusResponseResultRoundState(height string, round int32, step int32, startTime string, commitTime string, validators DumpConsensusResponseResultRoundStateValidators, lockedRound int32, validRound string, votes []DumpConsensusResponseResultRoundStateVotes, commitRound int32, lastCommit NullableDumpConsensusResponseResultRoundStateLastCommit, lastValidators DumpConsensusResponseResultRoundStateValidators, triggeredTimeoutPrecommit bool) *DumpConsensusResponseResultRoundState {
	this := DumpConsensusResponseResultRoundState{}
	this.Height = height
	this.Round = round
	this.Step = step
	this.StartTime = startTime
	this.CommitTime = commitTime
	this.Validators = validators
	this.LockedRound = lockedRound
	this.ValidRound = validRound
	this.Votes = votes
	this.CommitRound = commitRound
	this.LastCommit = lastCommit
	this.LastValidators = lastValidators
	this.TriggeredTimeoutPrecommit = triggeredTimeoutPrecommit
	return &this
}

// NewDumpConsensusResponseResultRoundStateWithDefaults instantiates a new DumpConsensusResponseResultRoundState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDumpConsensusResponseResultRoundStateWithDefaults() *DumpConsensusResponseResultRoundState {
	this := DumpConsensusResponseResultRoundState{}
	return &this
}

// GetHeight returns the Height field value
func (o *DumpConsensusResponseResultRoundState) GetHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetHeightOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *DumpConsensusResponseResultRoundState) SetHeight(v string) {
	o.Height = v
}

// GetRound returns the Round field value
func (o *DumpConsensusResponseResultRoundState) GetRound() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Round
}

// GetRoundOk returns a tuple with the Round field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetRoundOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Round, true
}

// SetRound sets field value
func (o *DumpConsensusResponseResultRoundState) SetRound(v int32) {
	o.Round = v
}

// GetStep returns the Step field value
func (o *DumpConsensusResponseResultRoundState) GetStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Step
}

// GetStepOk returns a tuple with the Step field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetStepOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Step, true
}

// SetStep sets field value
func (o *DumpConsensusResponseResultRoundState) SetStep(v int32) {
	o.Step = v
}

// GetStartTime returns the StartTime field value
func (o *DumpConsensusResponseResultRoundState) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetStartTimeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *DumpConsensusResponseResultRoundState) SetStartTime(v string) {
	o.StartTime = v
}

// GetCommitTime returns the CommitTime field value
func (o *DumpConsensusResponseResultRoundState) GetCommitTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitTime
}

// GetCommitTimeOk returns a tuple with the CommitTime field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetCommitTimeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommitTime, true
}

// SetCommitTime sets field value
func (o *DumpConsensusResponseResultRoundState) SetCommitTime(v string) {
	o.CommitTime = v
}

// GetValidators returns the Validators field value
func (o *DumpConsensusResponseResultRoundState) GetValidators() DumpConsensusResponseResultRoundStateValidators {
	if o == nil {
		var ret DumpConsensusResponseResultRoundStateValidators
		return ret
	}

	return o.Validators
}

// GetValidatorsOk returns a tuple with the Validators field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetValidatorsOk() (*DumpConsensusResponseResultRoundStateValidators, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Validators, true
}

// SetValidators sets field value
func (o *DumpConsensusResponseResultRoundState) SetValidators(v DumpConsensusResponseResultRoundStateValidators) {
	o.Validators = v
}

// GetLockedRound returns the LockedRound field value
func (o *DumpConsensusResponseResultRoundState) GetLockedRound() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LockedRound
}

// GetLockedRoundOk returns a tuple with the LockedRound field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetLockedRoundOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LockedRound, true
}

// SetLockedRound sets field value
func (o *DumpConsensusResponseResultRoundState) SetLockedRound(v int32) {
	o.LockedRound = v
}

// GetValidRound returns the ValidRound field value
func (o *DumpConsensusResponseResultRoundState) GetValidRound() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidRound
}

// GetValidRoundOk returns a tuple with the ValidRound field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetValidRoundOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValidRound, true
}

// SetValidRound sets field value
func (o *DumpConsensusResponseResultRoundState) SetValidRound(v string) {
	o.ValidRound = v
}

// GetVotes returns the Votes field value
func (o *DumpConsensusResponseResultRoundState) GetVotes() []DumpConsensusResponseResultRoundStateVotes {
	if o == nil {
		var ret []DumpConsensusResponseResultRoundStateVotes
		return ret
	}

	return o.Votes
}

// GetVotesOk returns a tuple with the Votes field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetVotesOk() (*[]DumpConsensusResponseResultRoundStateVotes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Votes, true
}

// SetVotes sets field value
func (o *DumpConsensusResponseResultRoundState) SetVotes(v []DumpConsensusResponseResultRoundStateVotes) {
	o.Votes = v
}

// GetCommitRound returns the CommitRound field value
func (o *DumpConsensusResponseResultRoundState) GetCommitRound() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CommitRound
}

// GetCommitRoundOk returns a tuple with the CommitRound field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetCommitRoundOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommitRound, true
}

// SetCommitRound sets field value
func (o *DumpConsensusResponseResultRoundState) SetCommitRound(v int32) {
	o.CommitRound = v
}

// GetLastCommit returns the LastCommit field value
// If the value is explicit nil, the zero value for DumpConsensusResponseResultRoundStateLastCommit will be returned
func (o *DumpConsensusResponseResultRoundState) GetLastCommit() DumpConsensusResponseResultRoundStateLastCommit {
	if o == nil || o.LastCommit.Get() == nil {
		var ret DumpConsensusResponseResultRoundStateLastCommit
		return ret
	}

	return *o.LastCommit.Get()
}

// GetLastCommitOk returns a tuple with the LastCommit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DumpConsensusResponseResultRoundState) GetLastCommitOk() (*DumpConsensusResponseResultRoundStateLastCommit, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastCommit.Get(), o.LastCommit.IsSet()
}

// SetLastCommit sets field value
func (o *DumpConsensusResponseResultRoundState) SetLastCommit(v DumpConsensusResponseResultRoundStateLastCommit) {
	o.LastCommit.Set(&v)
}

// GetLastValidators returns the LastValidators field value
func (o *DumpConsensusResponseResultRoundState) GetLastValidators() DumpConsensusResponseResultRoundStateValidators {
	if o == nil {
		var ret DumpConsensusResponseResultRoundStateValidators
		return ret
	}

	return o.LastValidators
}

// GetLastValidatorsOk returns a tuple with the LastValidators field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetLastValidatorsOk() (*DumpConsensusResponseResultRoundStateValidators, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastValidators, true
}

// SetLastValidators sets field value
func (o *DumpConsensusResponseResultRoundState) SetLastValidators(v DumpConsensusResponseResultRoundStateValidators) {
	o.LastValidators = v
}

// GetTriggeredTimeoutPrecommit returns the TriggeredTimeoutPrecommit field value
func (o *DumpConsensusResponseResultRoundState) GetTriggeredTimeoutPrecommit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TriggeredTimeoutPrecommit
}

// GetTriggeredTimeoutPrecommitOk returns a tuple with the TriggeredTimeoutPrecommit field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundState) GetTriggeredTimeoutPrecommitOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TriggeredTimeoutPrecommit, true
}

// SetTriggeredTimeoutPrecommit sets field value
func (o *DumpConsensusResponseResultRoundState) SetTriggeredTimeoutPrecommit(v bool) {
	o.TriggeredTimeoutPrecommit = v
}

func (o DumpConsensusResponseResultRoundState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["round"] = o.Round
	}
	if true {
		toSerialize["step"] = o.Step
	}
	if true {
		toSerialize["start_time"] = o.StartTime
	}
	if true {
		toSerialize["commit_time"] = o.CommitTime
	}
	if true {
		toSerialize["validators"] = o.Validators
	}
	if true {
		toSerialize["locked_round"] = o.LockedRound
	}
	if true {
		toSerialize["valid_round"] = o.ValidRound
	}
	if true {
		toSerialize["votes"] = o.Votes
	}
	if true {
		toSerialize["commit_round"] = o.CommitRound
	}
	if true {
		toSerialize["last_commit"] = o.LastCommit.Get()
	}
	if true {
		toSerialize["last_validators"] = o.LastValidators
	}
	if true {
		toSerialize["triggered_timeout_precommit"] = o.TriggeredTimeoutPrecommit
	}
	return json.Marshal(toSerialize)
}

type NullableDumpConsensusResponseResultRoundState struct {
	value *DumpConsensusResponseResultRoundState
	isSet bool
}

func (v NullableDumpConsensusResponseResultRoundState) Get() *DumpConsensusResponseResultRoundState {
	return v.value
}

func (v *NullableDumpConsensusResponseResultRoundState) Set(val *DumpConsensusResponseResultRoundState) {
	v.value = val
	v.isSet = true
}

func (v NullableDumpConsensusResponseResultRoundState) IsSet() bool {
	return v.isSet
}

func (v *NullableDumpConsensusResponseResultRoundState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDumpConsensusResponseResultRoundState(val *DumpConsensusResponseResultRoundState) *NullableDumpConsensusResponseResultRoundState {
	return &NullableDumpConsensusResponseResultRoundState{value: val, isSet: true}
}

func (v NullableDumpConsensusResponseResultRoundState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDumpConsensusResponseResultRoundState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


