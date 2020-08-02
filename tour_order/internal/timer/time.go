package timer

import "time"


func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)//可以方便支持多种时间类型
	if err != nil {
		return time.Time{}, nil
	}

	return currentTimer.Add(duration), nil
}



