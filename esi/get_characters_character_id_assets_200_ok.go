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
type GetCharactersCharacterIdAssets200Ok struct {

	// is_singleton boolean
	IsSingleton bool `json:"is_singleton,omitempty"`

	// item_id integer
	ItemId int64 `json:"item_id,omitempty"`

	// location_flag string
	LocationFlag string `json:"location_flag,omitempty"`

	// location_id integer
	LocationId int64 `json:"location_id,omitempty"`

	// location_type string
	LocationType string `json:"location_type,omitempty"`

	// quantity integer
	Quantity int32 `json:"quantity,omitempty"`

	// type_id integer
	TypeId int32 `json:"type_id,omitempty"`
}
