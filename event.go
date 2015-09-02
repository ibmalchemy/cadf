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
	Id            string         `json:"id"`
	TypeURI       string         `json:"typeURI"`
	EventType     string         `json:"eventType"`
	EventTime     Timestamp      `json:"eventTime"`
	Action        string         `json:"action"`
	Outcome       string         `json:"outcome"`
	Message       string         `json:"message"`
	Initiator     Resource       `json:"initiator"`
	Target        Resource       `json:"target"`
	Observer      Resource       `json:"observer"`
	Reason        Reason         `json:"reason"`
	Severity      string         `json:"severity"`
	Name          string         `json:"name"`
	Tags          []string       `json:"tags"`
	Attachments   []Attachment   `json:"attachments"`
	Reporterchain []Reporterstep `json:"reporterchain"`
	Api           Api            `json:"api"`
	Latencies     Latencies      `json:"latencies"`
}

// | Name       | Type       | Required | Description                        |
// |------------+------------+----------+------------------------------------|
// | id         | string     | Yes      | id of resource                     |
// | typeURI    | string     | Yes      | typeURI of resource                |
// | name       | string     | No       | name of resource                   |
// | host       | Host       | No       | domain to qualify name of resource |
// | credential | Credential | No       | security credential of resource    |
type Resource struct {
	Id         string     `json:"id"`
	TypeURI    string     `json:"typeURI"`
	Name       string     `json:"name"`
	Host       Host       `json:"host"`
	Credential Credential `json:"credential"`
}

// | Name     | Type   | Required | Description        |
// |----------+--------+----------+--------------------|
// | id       | string | Yes      | id of host         |
// | address  | string | No       | address of host    |
// | agent    | string | No       | agent name of host |
// | platform | string | No       | platform of host   |
type Host struct {
	Id      string `json:"id"`
	Address string `json:"address"`
	Agent   string `json:"agent"`
}

// | Name  | Type   | Required | Description    |
// |-------+--------+----------+----------------|
// | token | string | Yes      | value of token |
// | type  | string | Yes      | type of token  |
type Credential struct {
	Token string `json:"token"`
	Type  string `json:"type"`
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
	ReasonCode string `json:"reasonCode"`
	ReasonType string `json:"reasonType"`
}

// | Name        | Type   | Required | Description   |
// |-------------+--------+----------+---------------|
// | typeURI     | string | Yes      | type URI      |
// | contentType | string | Yes      | content type  |
// | content     | string | Yes      | content value |

type Attachment struct {
	TypeURI     string `json:"typeURI"`
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

// | Name         | Type      | Required | Description   |
// |--------------+-----------+----------+---------------|
// | role         | string    | Yes      | role          |
// | reporterId   | string    | Yes      | reporter id   |
// | reporterTime | Timestamp | Yes      | reporter time |
type Reporterstep struct {
	Role         string    `json:"role"`
	ReporterId   string    `json:"reporterId"`
	ReporterTime Timestamp `json:"reportedTime"`
}

type Latencies struct {
	Request  time.Duration `json:"request"`
	Response time.Duration `json:"response"`
	Proxy    time.Duration `json:"proxy"`
}

type Api struct {
	TargetUrl string    `json:"targetUrl"`
	PublicDns string    `json:"publicDns"`
	Name      string    `json:"name"`
	Id        string    `json:"id"`
	CreatedAt Timestamp `json:"createdAt"`
}

func (event *Event) String() string {
	if b, e := json.Marshal(event); e == nil {
		return string(b)
	}
	return ""
}
