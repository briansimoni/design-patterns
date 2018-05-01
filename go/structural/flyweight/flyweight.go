// Flyweight is a pattern which allows sharing the state of a heavy object
// between many instances of some type. Imagine that you have to create
// and store too many objects of some heavy type that are fundamentally
// equal. You'll run out of memory quickly. This problem can be solved with
// the Flyweight pattern, with additional help of the Factory pattern.
// The factory is usually in charge of encapsulating object creation

// Thanks to the Flyweight pattern, we can share all possible states of objects
// in a single common object, and thbus minimize object creation by using pointers
// to already created objects
package flyweight

import "time"

const (
	TEAM_A = iota
	TEAM_B
)

type Team struct {
	ID             uint64
	Name           int
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}

func (t *teamFlyweightFactory) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team

	return t.createdTeams[teamID]
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: TEAM_B,
		}
	default:
		return Team{
			ID:   1,
			Name: TEAM_A,
		}
	}
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
