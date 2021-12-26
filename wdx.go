/**
 * Copyright (c) 2021-present, Rana Jahanzaib
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package wdx

import (
	"fmt"
)

// Log progress / details on the go
func Log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
