package client

import (
	"encoding/json"
	"fmt"
)

type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id"`
	RetryAfter      int   `json:"retry_after"`
}

type ErrorPayload struct {
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

type Response struct {
	OK     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
	ErrorPayload
}

func (r Response) Error() string {
	return fmt.Sprintf(`%d | %s`, r.ErrorCode, r.Description)
}