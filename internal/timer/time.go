package timer

import "time"

/**
 * @Author
 * @Description //TODO
 * @Date 11:52 上午 2020/7/17
 * @Param
 * @return
 **/
func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

/**
 * @Author
 * @Description //TODO
 * @Date 1:10 下午 2020/7/17
 * @Param
 * @return
 **/
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}
