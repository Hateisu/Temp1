package consts

import "encoding/json"

type response struct {
	Statuscode status `json:"state"`
	Body       any    `json:"results"`
}

type status struct {
	STATUS             int    `json:"code"`
	STATUS_DESCRIPTION string `json:"message"`
}

func newstatus(a int, b string) *status {
	return &status{a, b}
}

var (
	sTATUS_OK_200 status = *newstatus(200, "Sucessful")
	sTATUS_OK_201 status = *newstatus(201, "Created")
	sTATUS_OK_202 status = *newstatus(202, "Accepted")
	sTATUS_OK_203 status = *newstatus(203, "Non-Authoritative Information")
	sTATUS_OK_204 status = *newstatus(204, "No Content")
	sTATUS_OK_205 status = *newstatus(205, "Reset Content")
)

func RESPONSE_OK_200(response_body any) *json.RawMessage {
	ret := response{sTATUS_OK_200, response_body}
	data, err := json.Marshal(ret)
	if err != nil {
		//Throw error
	}
	ans := json.RawMessage(data)
	return &ans
}
