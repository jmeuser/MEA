// Package cost gives the Monthly Expense Allocation of a Manager (MMEA) or Department (DMEA).
package cost

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Employee has a unique ID, their Manager's ID, and a Pos(ition) in their Dep(artment).
type Employee struct{ ID, MID, Pos, Dep string }

// byID maps ID to *Employee.
var byID = make(map[string]*Employee)

// byMID (byDep) maps MID (Dep) to IDs.
var byMID = make(map[string][]string)
var byDep = make(map[string][]string)

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
	for _, e := range employees {
		byID[e.ID] = e
		byMID[e.MID] = append(byMID[e.MID], e.ID)
		byDep[e.Dep] = append(byDep[e.Dep], e.ID)
	}
	return nil
}

// MEA calculates an *Employee's Monthly Expense Allocation.
func (e *Employee) MEA() (int, error) {
	switch e.Pos {
	case "Developer":
		return 1000, nil
	case "QA Tester":
		return 500, nil
	case "Manager":
		s := 300 // Manager's MEA
		// Add MEA of Manager's Employees.
		for _, id := range byMID[e.ID] {
			x, err := byID[id].MEA()
			if err != nil {
				return -1, err
			}
			s += x
		}
		return s, nil
	default:
		return -1, fmt.Errorf("The MEA for the position %s has not been specified.", e.Pos)
	}
}

// MMEA gives a Manager's Monthly Expense Allocation.
func MMEA(id string) (int, error) {
	e := byID[id]
	if e == nil {
		return -1, fmt.Errorf("there is not an employee in the Basic Personell Files with ID %d", id)
	}
	if e.Pos != "Manager" {
		return -1, fmt.Errorf("The employee with ID %d is not a manager.")
	}
	return e.MEA()
}

// DMEA gives a Department's Monthly Expense Allocation.
func DMEA(dep string) (int, error) {
	ids := byDep[dep]
	if ids == nil {
		return -1, fmt.Errorf("There are no employees in the %s department.", dep)
	}
	s := 0
	for _, id := range ids {
		x, err := byID[id].MEA()
		if err != nil {
			return -1, err
		}
		s += x
	}
	return s, nil
}
