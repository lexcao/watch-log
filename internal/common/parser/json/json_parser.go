package json

import (
	"encoding/json"

	"github.com/lexcao/watch-log/pkg/model"
)

type Parser struct {
}

func (p Parser) Parse(entry *model.Entry) {
	entry.Err = json.Unmarshal([]byte(entry.Origin), &entry.ParsedObject)
}
