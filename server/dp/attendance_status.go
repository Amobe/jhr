package dp

type AttendanceStatus string

const (
	Normal   AttendanceStatus = "Normal"
	Abnormal AttendanceStatus = "Abnormal"
	TooLow   AttendanceStatus = "TooLow"  // lower than 9 hours
	TooHigh  AttendanceStatus = "TooHigh" // higher than 9.5 hours
)

func GetAttendanceStatus(start, end TimeUnit) (status AttendanceStatus, diff TimeUnit) {
	if !start.IsValid() || !end.IsValid() {
		return Abnormal, NewInvalidTimeUnit()
	}
	diff = diffTimeUnit(start, end)
	return getAttendanceStatus(diff.ToMinutes()), diff
}

func getAttendanceStatus(minutes int) AttendanceStatus {
	switch {
	case minutes < 9*60:
		return TooLow
	case minutes > 9.5*60:
		return TooHigh
	default:
		return Normal
	}
}

func diffTimeUnit(start, end TimeUnit) TimeUnit {
	diffMinutes := end.ToMinutes() - start.ToMinutes()
	return NewTimeUnit(diffMinutes)
}
