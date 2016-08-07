package gcmjson

/* ----- */

//easyjson:json
type GcmDataSubItem struct {
	SubId    string `json:"sub_id"`
	ChangeTs int64  `json:"ts"`
}

//easyjson:json
type GcmDataSubList struct {
	SubList []GcmDataSubItem `json:"subs"`
}

//easyjson:json
type GcmDataNoSubList struct {
}

//easyjson:json
type GcmPacket struct {
	To                string      `json:"to,omitempty"`
	CollapseKey       string      `json:"collapse_key,omitempty"`
	Priority          string      `json:"priority,omitempty"`
	RestrictedPackage string      `json:"restricted_package_name,omitempty"`
	Data              interface{} `json:"data"`
}

func (p *GcmPacket) SetDataSubList(data *GcmDataSubList) {
	p.Data = data
}

var nothing GcmDataNoSubList = GcmDataNoSubList{}

func (p *GcmPacket) SetDataNothing() {
	p.Data = &nothing
}

/* ----- */

//easyjson:json
type GcmResultItem struct {
	Error          string `json:"error"`
	MessageId      string `json:"message_id"`
	RegistrationId string `json:"registration_id"`
}

//easyjson:json
type GcmResponse struct {
	Success      int             `json:"success"`
	Failure      int             `json:"failure"`
	CanonicalIds int             `json:"canonical_ids"`
	Results      []GcmResultItem `json:"results"`
}
