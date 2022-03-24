// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
// Returns a pointer to a `Resident` struct that holds this information.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

/*
HasRequiredInfo determines if a given resident has all of the required information.

In order to be counted, a resident must provide a non-zero value for their name
and an address that contains a non-zero value for the `street` key.
All other information, such as the resident's age, is optional.
Returns a boolean indicating if the resident has provided the required information. */
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" || r.Address == nil || r.Address["street"] == "" {
		return false
	}
	return true
}

// Delete deletes a resident's information.
// Sets all of the fields the resident to their zero value.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

/*
Count counts all residents that have provided the required information.
Accepts one argument:

- A slice of pointers to `Resident` structs.

The function should return an integer indicating the number of residents that were counted in the census.
A resident can only be counted if they provided the required information to the census worker. */
func Count(residents []*Resident) int {
	countedResidents := 0
	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			countedResidents++
		}
	}
	return countedResidents
}
