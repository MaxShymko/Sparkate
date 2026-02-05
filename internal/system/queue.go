package system

func IsInQueueSmb() (bool, int64) {
	UserMap.mu.Lock()
	defer UserMap.mu.Unlock()

	for _, user := range UserMap.Map {
		if user.State == StateQueue {
			return true, user.ID
		}
	}
	return false, 0
}
