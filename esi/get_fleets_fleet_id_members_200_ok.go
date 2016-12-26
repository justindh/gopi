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

import (
	"time"
)

// 200 ok object
type GetFleetsFleetIdMembers200Ok struct {

	// character_id integer
	CharacterId int32 `json:"character_id,omitempty"`

	// join_time string
	JoinTime time.Time `json:"join_time,omitempty"`

	// Member’s role in fleet
	Role string `json:"role,omitempty"`

	// Localized role names
	RoleName string `json:"role_name,omitempty"`

	// ship_type_id integer
	ShipTypeId int32 `json:"ship_type_id,omitempty"`

	// Solar system the member is located in
	SolarSystemId int32 `json:"solar_system_id,omitempty"`

	// ID of the squad the member is in. If not applicable, will be set to -1
	SquadId int64 `json:"squad_id,omitempty"`

	// Station in which the member is docked in, if applicable
	StationId int64 `json:"station_id,omitempty"`

	// Whether the member take fleet warps
	TakesFleetWarp bool `json:"takes_fleet_warp,omitempty"`

	// ID of the wing the member is in. If not applicable, will be set to -1
	WingId int64 `json:"wing_id,omitempty"`
}
