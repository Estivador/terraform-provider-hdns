package schema

import "time"

// Record defines the schema of DNS record which may be included
// in responses.
type Record struct {
	Type     string    `json:"type"`
	ID       string    `json:"id"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	ZoneID   string    `json:"zone_id"`
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	TTL      uint64    `json:"ttl"`
}

// RecordGetResponse defines the schema of the response when retrieving
// a single record.
type RecordGetResponse struct {
	Record Record `json:"record"`
}

// RecordListResponse defines the schema of the response when listing
// all DNS records.
type RecordListResponse struct {
	Records []Record `json:"records"`
}

// RecordCreateRequest defines the schema of the request to create a record.
type RecordCreateRequest struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	TTL    uint64 `json:"ttl"`
	Value  string `json:"value"`
	ZoneID string `json:"zone_id"`
}

// RecordCreateResponse defines the schema of the response for record creation.
type RecordCreateResponse struct {
	Record Record `json:"record"`
}

// RecordUpdateRequest defines the schema of the request to update a record.
type RecordUpdateRequest struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	TTL    uint64 `json:"ttl"`
	Value  string `json:"value"`
	ZoneID string `json:"zone_id"`
}

// RecordUpdateResponse defines the schema of the response for zone update.
type RecordUpdateResponse struct {
	Record Record `json:"record"`
}
