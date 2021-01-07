package procedures

type ProcedureInterface interface {
	drop() error
	dropIfExist() error
	create() error
}

func CreateProcedure(p ProcedureInterface) error {
	return p.create()
}

func DropProcedure(p ProcedureInterface) error {
	return p.drop()
}

func DropProcedureIfExist(p ProcedureInterface) error {
	return p.dropIfExist()
}

func Initialize() {
	_ = DropProcedureIfExist(GetUsersCount{})
	_ = CreateProcedure(GetUsersCount{})
}
