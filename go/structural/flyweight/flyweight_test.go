package flyweight

import "testing"

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	// factory := teamFlyweightFactory{}
	factory := NewTeamFactory()

	teamA1 := factory.GetTeam(TEAM_A)
	if teamA1 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	teamA2 := factory.GetTeam(TEAM_A)
	if teamA2 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	if teamA1 != teamA2 {
		t.Error("TEAM_A pointers weren't the same")
	}

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects created was not 1. Got: %d", factory.GetNumberOfObjects())
	}
}
