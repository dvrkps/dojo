package medicament

import "time"

type expire struct {
	ExpireDate   time.Time
	DaysToExpire int
}

func newExpire(t time.Time, r refill, ratio float64) expire {
	const hoursInDay = 24

	// normalize today
	t = midnight(t)

	// medicament stock in days
	msid := r.Quantity / ratio

	// days diff between refill date and today
	diff := t.Sub(r.Date).Hours() / hoursInDay

	// days to expire
	dte := int(msid - diff)

	// expire date
	dur := time.Duration(dte*hoursInDay) * time.Hour
	ed := t.Add(dur)
	// expire
	e := expire{
		ExpireDate:   ed,
		DaysToExpire: dte,
	}

	return e
}
