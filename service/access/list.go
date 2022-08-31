package access

import (
	"fmt"
	"github.com/LittleGriseo/GriseoProxy/common/set"
	"github.com/LittleGriseo/GriseoProxy/config"
)

func GetTargetList(listName string) (*set.StringSet, error) {
	set, ok := config.Lists[listName]
	if ok {
		return set, nil
	}
	return nil, fmt.Errorf("list %q not found", listName)
}
