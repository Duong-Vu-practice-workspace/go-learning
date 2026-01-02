package __enums

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	updateStatusErr := em.recipient.updateStatus(em.status)
	trackErr := a.track(em.status)
	if updateStatusErr != nil {
		return fmt.Errorf("error updating user status: %v", updateStatusErr.Error())
	}
	if trackErr != nil {
		return fmt.Errorf("error tracking user bounce: %v", trackErr.Error())
	}
	return nil
}
