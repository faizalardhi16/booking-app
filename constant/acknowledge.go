package constant

type Acknowledge struct {
	Success int
	Failed  int
}

var AcknowledgeEnum = Acknowledge{
	Success: 1,
	Failed:  0,
}
