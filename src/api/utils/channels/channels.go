package channels

func Ok(done <-chan bool) bool {

	select {
	case ok := <-done:
		if ok {
			return true
		}
	}
	return false

}
