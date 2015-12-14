// Package cost gives the Monthly Expense Allocation of a Manager (MMEA) or Department (DMEA).
package cost

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Employee has a unique ID, their Manager's ID, and a Pos(ition) in their Dep(artment).
type Employee struct{ ID, MID, Pos, Dep string }

// mea gives an INDIVIDUAL *Employee's MEA.
func (e *Employee) mea() (int, error) {
	switch e.Pos {
	case "Developer":
		return 1000, nil
	case "QA Tester":
		return 500, nil
	case "Manager":
		return 300, nil
	default:
		return -1, fmt.Errorf("The MEA for the position %v is unspecified.", e.Pos)
	}
}

// byID maps ID to *Employee
var byID map[string]*Employee

// byMID (byDep) maps MID (Dep) to IDs.
var byMID, byDep map[string][]string

// LoadBPF unmarshals Basic Personnel Files, from fileName.json, into byID, byMID, and byDep.
func LoadBPF(fileName string) error {
	data, err := ioutil.ReadFile(fileName + ".json")
	if err != nil {
		return err
	}
	var employees []*Employee
	if err := json.Unmarshal(data, &employees); err != nil {
		return fmt.Errorf("Unmarshal fail: %s", err)
	}
	byID = make(map[string]*Employee)
	byMID, byDep = make(map[string][]string), make(map[string][]string)
	for _, e := range employees {
		byID[e.ID] = e
		byMID[e.MID] = append(byMID[e.MID], e.ID)
		byDep[e.Dep] = append(byDep[e.Dep], e.ID)
	}
	return nil
}

func employeeByID(id string) (*Employee, error) {
	if byID == nil {
		return nil, fmt.Errorf("Basic Personal Files have not been loaded: use func LoadBPF")
	}
	e := byID[id]
	if e == nil {
		return nil, fmt.Errorf("There is not an employee with ID %v.", id)
	}
	return e, nil
}

func idsByMID(mid string) ([]string, error) {
	if byMID == nil {
		return nil, fmt.Errorf("Basic Personal Files have not been loaded: use func LoadBPF")
	}
	ids := byMID[mid]
	if ids == nil {
		return nil, fmt.Errorf("There is not an Employee with MID %v.", mid)
	}
	return ids, nil
}

func idsByDep(dep string) ([]string, error) {
	if byDep == nil {
		return nil, fmt.Errorf("Basic Personal Files have not been loaded: use func LoadBPF")
	}
	ids := byDep[dep]
	if ids == nil {
		return nil, fmt.Errorf("There are no employees in the %v department.", dep)
	}
	return ids, nil
}

func MMEA(mid string) (int, error) {
	m, err := employeeByID(mid)
	if err != nil {
		return -1, err
	}
	if m.Pos != "Manager" {
		return -1, fmt.Errorf("The employee with ID %v is not a manager.", mid)
	}
	s, err := m.mea()
	if err != nil {
		return -1, err
	}
	ids, err := idsByMID(mid)
	if err != nil {
		return -1, err
	}
	for _, id := range ids {
		e, err := employeeByID(id)
		if err != nil {
			return -1, err
		}
		if e.Pos == "Manager" {
			x, err := MMEA(id)
			if err != nil {
				return -1, err
			}
			s += x
		} else {
			x, err := e.mea()
			if err != nil {
				return -1, err
			}
			s += x
		}
	}
	return s, nil
}

// DMEA gives a Department's MEA
func DMEA(dep string) (int, error) {
	ids, err := idsByDep(dep)
	if err != nil {
		return -1, err
	}
	s := 0
	for _, id := range ids {
		e, err := employeeByID(id)
		if err != nil {
			return -1, err
		}
		x, err := e.mea()
		if err != nil {
			return -1, err
		}
		s += x
	}
	return s, nil
}
