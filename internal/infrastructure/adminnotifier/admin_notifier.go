package adminnotifier

type AdminNotifier struct{}

func (n *AdminNotifier) Notify(severity string, affectedEmployee string, message string) error {
	return nil
}
