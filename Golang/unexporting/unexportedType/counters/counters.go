package counters

type alertCounter int

func New(i int) alertCounter {
	return alertCounter(i)
}
