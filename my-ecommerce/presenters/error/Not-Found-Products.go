package error

import (
	"errors"
	"fmt"
	"strings"
)

func NotFoundProducts(ids []int) error {
	return errors.New(fmt.Sprintf("Not found ids: %s", strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")))
}
