/*******************************************************************************
 * Copyright 2020 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package unmarshaler

// callback provides a way to create limited,
// simple Unmarshalers for use cases like single instance unmarshalers or for testing.
type callback struct {
	callbackFunc CallbackFunc
}

func NewCallback(cbf func(data []byte, v interface{}) error) callback {
	return callback{cbf}
}

// Unmarshal calls the callback function and returns its result
func (cb *callback) Unmarshal(data []byte, v interface{}) error {
	return cb.callbackFunc(data, v)
}

// CallbackFunc is the function that is going to do the unmarshal.
type CallbackFunc func(data []byte, v interface{}) error

// Error provides the required implementation for the Error interface.
func (cbf CallbackFunc) Error() string {
	return "unexpected error trying to unmarshal"
}
