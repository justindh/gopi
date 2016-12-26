/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.4
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

// 200 ok object
type GetCharactersCharacterIdContacts200Ok struct {

	// contact_id integer
	ContactId int32 `json:"contact_id,omitempty"`

	// contact_type string
	ContactType string `json:"contact_type,omitempty"`

	// Whether this contact is in the blocked list. Note a missing value denotes unknown, not true or false
	IsBlocked bool `json:"is_blocked,omitempty"`

	// Whether this contact is being watched
	IsWatched bool `json:"is_watched,omitempty"`

	// Custom label of the contact
	LabelId int64 `json:"label_id,omitempty"`

	// Standing of the contact
	Standing float32 `json:"standing,omitempty"`
}
