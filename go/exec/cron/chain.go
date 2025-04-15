package cron

type JobWrapper func(Job) Job

type Chain struct {
	wrappers []JobWrapper
}

func NewChain(c ...JobWrapper) Chain {
	return Chain{c}
}

// This:
//
//	NewChain(m1, m2, m3).Then(job)
//
// is equivalent to:
//
//	m1(m2(m3(job)))
func (c Chain) Then(j Job) Job {
	for i := range c.wrappers {
		j = c.wrappers[len(c.wrappers)-i-1](j)
	}
	return j
}

func Recover(logger Logger) JobWrapper {
	return func(j Job) Job {

	}
}
