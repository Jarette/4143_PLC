// Mascot Test File
package mascot_test

// importing test package and the mascot package
import (
	"testing"

	"github.com/4143_PLC/P01/mascot"
)

// TestMascot function tests the mascot package
func TestMascot(t *testing.T) {
	// checks if the mascot package's "BestMascot" function
	// returns the correct output
	if mascot.BestMascot() != "Go Gopher" {
		//display this message if the wrong output is produced
		t.Fatal("Wrong Mascot")
	}
}
