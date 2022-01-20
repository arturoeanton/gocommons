package prolog

import (
	"github.com/mndrix/golog"
	"github.com/mndrix/golog/read"
	"github.com/mndrix/golog/term"
)

type Prolog struct {
	Machine golog.Machine
	Code    string
}

type Register map[string]string
type Vars []string

type Result struct {
	Result bool
	Vars   Vars
	Table  []Register
}

func NewProlog() *Prolog {
	return &Prolog{
		Machine: golog.NewMachine(),
		Code:    ``,
	}
}

func (prolog *Prolog) Load(code string) *Prolog {
	prolog.Code += code + "\n"
	prolog.Machine = golog.NewMachine().Consult(prolog.Code)
	return prolog
}

func (prolog *Prolog) Query(queryParam string) *Result {
	result := Result{}
	query := read.Term_(queryParam)
	vars := term.Variables(query)
	result.Vars = append(result.Vars, vars.Keys()...)
	result.Result = prolog.Machine.CanProve(query)
	solutions := prolog.Machine.ProveAll(query)
	for _, solution := range solutions {
		register := make(Register)
		for _, name := range result.Vars {
			register[name] = solution.ByName_(name).String()
		}
		result.Table = append(result.Table, register)
	}
	return &result
}

func (result *Result) ReverseEach(eachFx func(Register)) *Result {
	for i := len(result.Table) - 1; i >= 0; i-- {
		r := result.Table[i]
		eachFx(r)
	}
	return result
}

func (result *Result) Each(eachFx func(Register)) *Result {
	for _, r := range result.Table {
		eachFx(r)
	}
	return result
}

func (result *Result) Map(mapFx func(Register) Register) *Result {
	for i, r := range result.Table {
		result.Table[i] = mapFx(r)
	}
	return result
}

func (result *Result) Reduce(reduceFx func(Register, Register) Register) *Register {
	reduce := make(Register)
	for _, r := range result.Table {
		reduce = reduceFx(reduce, r)
	}
	return &reduce
}
