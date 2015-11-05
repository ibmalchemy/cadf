// Copyright 2015 IBM Corp.
// 
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package cadf

import (
	"encoding/json"
	"time"
)

// | Name          | Type           | Required | Description     |
// |---------------+----------------+----------+-----------------|
// | id            | string         | Yes      | id              |
// | typeURI       | string         | Yes      | type URI        |
// | eventType     | string         | Yes      | event type      |
// | eventTime     | Timestamp      | Yes      | event time      |
// | action        | string         | Yes      | action          |
// | outcome       | string         | Yes      | outcome         |
// | message       | string         | Yes      | message         |
// | initiator     | Resource       | Yes      | initiator       |
// | target        | Resource       | Yes      | target          |
// | observer      | Resource       | Yes      | observer        |
// | reason        | Reason         | No       | reason          |
// | severity      | string         | No       | event severity  |
// | name          | string         | No       | event name      |
// | tags          | []string       | No       | tags            |
// | attachments   | []Attachment   | No       | attachments     |
// | reporterchain | []Reporterstep | No       | reporterchain   |
// | api           | Api            | No       | api info        |
// | latencies     | Latencies      | No       | latencies       |
type Event struct {
	Id            string         `json:"id,omitempty"`
	TypeURI       string         `json:"typeURI,omitempty"`
	EventType     string         `json:"eventType,omitempty"`
	EventTime     Timestamp      `json:"eventTime,omitempty"`
	Action        string         `json:"action,omitempty"`
	Outcome       string         `json:"outcome,omitempty"`
	Message       string         `json:"message,omitempty"`
	Initiator     Resource       `json:"initiator,omitempty"`
	Target        Resource       `json:"target,omitempty"`
	Observer      Resource       `json:"observer,omitempty"`
	Reason        Reason         `json:"reason,omitempty"`
	Severity      string         `json:"severity,omitempty"`
	Name          string         `json:"name,omitempty"`
	Tags          []string       `json:"tags,omitempty"`
	Attachments   []Attachment   `json:"attachments,omitempty"`
	Reporterchain []Reporterstep `json:"reporterchain,omitempty"`
	Api           Api            `json:"api,omitempty"`
	Latencies     Latencies      `json:"latencies,omitempty"`
}

// | Name       | Type       | Required | Description                        |
// |------------+------------+----------+------------------------------------|
// | id         | string     | Yes      | id of resource                     |
// | typeURI    | string     | Yes      | typeURI of resource                |
// | name       | string     | No       | name of resource                   |
// | host       | Host       | No       | domain to qualify name of resource |
// | credential | Credential | No       | security credential of resource    |
type Resource struct {
	Id         string     `json:"id,omitempty"`
	TypeURI    string     `json:"typeURI,omitempty"`
	Name       string     `json:"name,omitempty"`
	Host       Host       `json:"host,omitempty"`
	Credential Credential `json:"credential,omitempty"`
}

// | Name     | Type   | Required | Description        |
// |----------+--------+----------+--------------------|
// | id       | string | Yes      | id of host         |
// | address  | string | No       | address of host    |
// | agent    | string | No       | agent name of host |
// | platform | string | No       | platform of host   |
type Host struct {
	Id      string `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
	Agent   string `json:"agent,omitempty"`
}

// | Name  | Type   | Required | Description    |
// |-------+--------+----------+----------------|
// | token | string | Yes      | value of token |
// | type  | string | Yes      | type of token  |
type Credential struct {
	Token string `json:"token,omitempty"`
	Type  string `json:"type,omitempty"`
}

type Timestamp struct {
	time.Time
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(t.Format(`"2006-01-02T15:04:05.999Z"`)), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = time.Parse(`"2006-01-02T15:04:05.999Z"`, string(data))
	return
}

// | Name       | Type   | Required | Description |
// |------------+--------+----------+-------------|
// | reasonCode | string | No       | reason code |
// | reasonType | string | No       | reason type |
type Reason struct {
	ReasonCode string `json:"reasonCode,omitempty"`
	ReasonType string `json:"reasonType,omitempty"`
}

// | Name        | Type   | Required | Description   |
// |-------------+--------+----------+---------------|
// | typeURI     | string | Yes      | type URI      |
// | contentType | string | Yes      | content type  |
// | content     | string | Yes      | content value |

type Attachment struct {
	TypeURI     string `json:"typeURI,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Content     string `json:"content,omitempty"`
}

// | Name         | Type      | Required | Description   |
// |--------------+-----------+----------+---------------|
// | role         | string    | Yes      | role          |
// | reporterId   | string    | Yes      | reporter id   |
// | reporterTime | Timestamp | Yes      | reporter time |
type Reporterstep struct {
	Role         string    `json:"role,omitempty"`
	ReporterId   string    `json:"reporterId,omitempty"`
	ReporterTime Timestamp `json:"reportedTime,omitempty"`
}

type Latencies struct {
	Request  time.Duration `json:"request,omitempty"`
	Response time.Duration `json:"response,omitempty"`
	Proxy    time.Duration `json:"proxy,omitempty"`
}

type Api struct {
	TargetUrl string    `json:"targetUrl,omitempty"`
	PublicDns string    `json:"publicDns,omitempty"`
	Name      string    `json:"name,omitempty"`
	Id        string    `json:"id,omitempty"`
	CreatedAt Timestamp `json:"createdAt,omitempty"`
}

func (event *Event) String() string {
	if b, e := json.Marshal(event); e == nil {
		return string(b)
	}
	return ""
}
