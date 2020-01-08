package flyweight

type Team struct {
	Name string
}

type teamFlyWeigthFactory struct {
	//本质上共享了元数据，要是没有则添加更新。注意：元数据的数量必须要有限。内存消耗很大
	createdTeams map[string]*Team
}

func (tf *teamFlyWeigthFactory) GetTeamByName(name string) *Team {
	if tf.createdTeams[name] != nil {
		return tf.createdTeams[name]
	}
	team := getTeamFactory(TeamA)
	tf.createdTeams["Team"] = &team
	return &team
}

func (tf *teamFlyWeigthFactory) GetNumOfObjects() int {
	return len(tf.createdTeams)
}

const (
	TeamA = iota
	TeamB
)

func getTeamFactory(team int) Team {
	switch team {
	case TeamA:
		return Team{
			Name: "TeamA",
		}
	case TeamB:
		return Team{
			Name: "TeamB",
		}
	default:
		return Team{
			Name: "TeamB",
		}
	}
}

func Newfactory() teamFlyWeigthFactory {
	return teamFlyWeigthFactory{
		createdTeams: make(map[string]*Team),
	}
}
