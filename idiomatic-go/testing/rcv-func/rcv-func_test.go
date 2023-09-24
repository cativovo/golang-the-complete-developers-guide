// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package testing

import "testing"

func TestApplyDamage(t *testing.T) {
	player := New("John doe")
	player.applyDamage(70)

	expectedHealth := Health(30)

	if expectedHealth != player.health {
		t.Errorf("incorrect healt: %v, expected: %v", player.health, expectedHealth)
	}
}

func TestHeal(t *testing.T) {
	player := New("John doe")
	player.applyDamage(80)
	player.heal(30)

	expectedHealth := Health(50)

	if expectedHealth != player.health {
		t.Errorf("incorrect healt: %v, expected: %v", player.health, expectedHealth)
	}
}

func TestConsumeEnery(t *testing.T) {
	player := New("John doe")
	player.consumeEnergy(50)

	expectedEnergy := Energy(50)

	if expectedEnergy != player.energy {
		t.Errorf("incorrect healt: %v, expected: %v", player.energy, expectedEnergy)
	}
}

func TestRestoreEnergy(t *testing.T) {
	player := New("John doe")
	player.consumeEnergy(50)
	player.restoreEnergy(10)

	expectedEnergy := Energy(60)

	if expectedEnergy != player.energy {
		t.Errorf("incorrect healt: %v, expected: %v", player.energy, expectedEnergy)
	}
}
