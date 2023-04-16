package constant

type Condition struct {
	Success string
	Failed  string
}

var ConditionEnum = Condition{
	Success: "Success",
	Failed:  "Failed",
}
