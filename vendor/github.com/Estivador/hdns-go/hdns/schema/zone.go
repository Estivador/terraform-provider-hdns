package schema

import "time"

// Zone defines the schema of DNS zone information which may be included
// in responses.
type Zone struct {
	ID              string          `json:"id"`
	Created         time.Time       `json:"created"`
	Modified        time.Time       `json:"modified"`
	LegacyDNSHost   string          `json:"legacy_dns_host"`
	LegacyNs        []string        `json:"legacy_ns"`
	Name            string          `json:"name"`
	Ns              []string        `json:"ns"`
	Owner           string          `json:"owner"`
	Paused          bool            `json:"paused"`
	Permission      string          `json:"permission"`
	Project         string          `json:"project"`
	Registrar       string          `json:"registrar"`
	Status          string          `json:"status"`
	TTL             uint64          `json:"ttl"`
	Verified        time.Time       `json:"verified"`
	RecordsCount    uint64          `json:"records_count"`
	IsSecondaryDNS  bool            `json:"is_sedondary_dns"`
	TxtVerification TxtVerification `json:"txt_verification"`
}

// ZoneGetResponse defines the schema of the response when retrieving
// a single zone.
type ZoneGetResponse struct {
	Zone Zone `json:"zone"`
}

// ZoneListResponse defines the schema of the response when listing
// zones.
type ZoneListResponse struct {
	Zones []Zone `json:"zones"`
}

// ZoneCreateRequest defines the schema of the request to create a zone.
type ZoneCreateRequest struct {
	Name string `json:"name"`
	TTL  uint64 `json:"ttl"`
}

// ZoneCreateResponse defines the schema of the response for zone creation.
type ZoneCreateResponse struct {
	Zone Zone `json:"zone"`
}

// ZoneUpdateRequest defines the schema of the request to update a zone.
type ZoneUpdateRequest struct {
	Name string `json:"name"`
	TTL  uint64 `json:"ttl"`
}

// ZoneUpdateResponse defines the schema of the response for zone update.
type ZoneUpdateResponse struct {
	Zone Zone `json:"zone"`
}
