package goc

import (
	"fmt"

	"github.com/oxtoacart/god"
)

func CallMe() string {
	return fmt.Sprintf("c:%s", god.CallMe())
}
