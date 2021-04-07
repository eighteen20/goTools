package timer

import "time"

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

/*
 时间推算
*/
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	// 从字符串解析出持续时间（duration）
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	// 传入 duration，得到当前Timer时间加上duration后的最终时间
	return currentTimer.Add(duration), nil
}
