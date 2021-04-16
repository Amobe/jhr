package dp

type AttendanceStatus string

const (
	Normal   AttendanceStatus = "Normal"
	Abnormal AttendanceStatus = "Abnormal"
	TooLow   AttendanceStatus = "TooLow"  // lower than 9 hours
	TooHigh  AttendanceStatus = "TooHigh" // higher than 9.5 hours
)
