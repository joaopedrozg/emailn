package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Test Campaign"
	content := "Test Content"
	contacts := []string{"test@example.com", "another@example.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "" {
		t.Errorf("Expected campaign ID to be empty, got %s", campaign.ID)
	} else if campaign.Name != name {
		t.Errorf("Expected campaign Name to be %s, got %s", name, campaign.Name)
	} else if campaign.Content != content {
		t.Errorf("Expected campaign Content to be %s, got %s", content, campaign.Content)
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("Expected campaign Contacts length to be %d, got %d", len(contacts), len(campaign.Contacts))
	}
}
