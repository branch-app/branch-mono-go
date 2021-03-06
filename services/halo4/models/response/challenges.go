package response

// Challenges defines the structure of the Global Challenges response.
type Challenges struct {
	Response `bson:",inline"`

	// Challenges contains the details of the global halo 4 challenges.
	Challenges []Challenge `json:"challenges" bson:"challenges"`
}

// Challenge defines the structure if a challenge in the global challenges.
type Challenge struct {
	// Name is the name of the challenge.
	Name string `json:"name" bson:"name"`

	// Description is the description of the challenge.
	Description string `json:"description" bson:"description"`

	// CategoryID is the ID of the category this challenge is in.
	CategoryID int `json:"categoryId" bson:"categoryId"`

	// CategoryName is the name of the category this challenge is in.
	CategoryName string `json:"categoryName" bson:"categoryName"`

	// ChallengeIndex is the index of the challenge.
	ChallengeIndex int `json:"challengeIndex" bson:"challengeIndex"`

	// PeriodID is the ID of the period this challenge is valid for.
	PeriodID int `json:"periodId" bson:"periodId"`

	// PeriodNamely is the name of the type of period this challenge is valid for.
	PeriodNamely string `json:"periodNamely" bson:"periodNamely"`

	// BeginDate is the date this challenge started.
	BeginDate string `json:"beginDate" bson:"beginDate"`

	// EndDate is the date this challenge expires.
	EndDate string `json:"endDate" bson:"endDate"`

	// RequiredCount is the tick count reqired to complete the challenge.
	RequiredCount int `json:"requiredCount" bson:"requiredCount"`

	// GameModeName is the name of the game mode the challenge is for.
	GameModeName string `json:"gameModeName" bson:"gameModeName"`

	// XPReward is the XP reward for completing the challenge.
	XPReward int `json:"xpReward" bson:"xpReward"`

	// RequiredSkulls holds the skulls that are required to complete the challenge.
	RequiredSkulls int `json:"requiredSkulls" bson:"requiredSkulls"`

	// Completed shows if the player has completed the challenge (this is unused for
	// global challenges)
	Completed bool `json:"completed" bson:"completed"`

	// Progress shows if the players progression for the challenge (this is unused for
	// global challenges)
	Progress float64 `json:"progress" bson:"progress"`
}
