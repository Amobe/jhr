package xls

type AttendanceStatus string

const (
	Normal   AttendanceStatus = "Normal"
	Abnormal AttendanceStatus = "Abnormal"
	TooLow   AttendanceStatus = "TooLow"  // lower than 9 hours
	TooHigh  AttendanceStatus = "TooHigh" // higher than 9.5 hours
)

func GetAttendanceStatus(minutes int) AttendanceStatus {
	switch {
	case minutes < 9*60:
		return TooLow
	case minutes > 9.5*60:
		return TooHigh
	default:
		return Normal
	}
}
