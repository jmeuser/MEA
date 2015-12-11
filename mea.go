// Package cost calculates the Montly Expense Allocation of a manager (MMEA) or department (DMEA) depending on the structure and positions of employees as described in the basic personal files (which are assumed to be stored in .json format, and to have at least those fields which are used in the Employee type).
package cost

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Employee struct {
	ID, ManagerID int
	Department, Position string
}
