package campaign

import "time"

type Campaign struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	GoalAMount       int       `json:"goal_amount"`
	CurrentAmount    int       `json:"current_amount"`
	Perks            string    `json:"perks"`
	BakerCount       int       `json:"baker_count"`
	Slug             string    `json:"slug"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         int       `json:"id"`
	CampaignId int       `json:"campaign_id"`
	Filename   string    `json:"filename"`
	IsPrimary  int       `json:"is_primary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
