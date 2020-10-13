package types

type ExecutionArgs struct {
	Platform string
	APIURL   string
	APIToken string
	Repo     string
	GitUser  string
	GitEmail string
}

type Executor interface {
	Execute(args ExecutionArgs) error
}
