package port

type AdminNotification interface {
	Notify(level string, employeeAbbreviation string, message string) error
}
