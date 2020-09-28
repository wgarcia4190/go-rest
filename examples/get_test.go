package examples

import (
	"fmt"
	"testing"
)

func TestGetEndpoints(t *testing.T) {
	// Initialization

	//Executions
	endpoints, err := GetEndpoints()

	//Validation
	fmt.Println(err)
	fmt.Println(endpoints)
}
