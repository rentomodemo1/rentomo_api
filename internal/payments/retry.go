package payments

func Retry(pay PaymentFunc, notify func(string), booking string) error {
	var err error
	for attempt := 0; attempt < 3; attempt++ {
		err = pay(booking)
		if err == nil {
			return nil
		}
	}
	notify(booking)
	return err
}
