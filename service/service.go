package service

type Unprepared struct {
	Data []byte
	Id   uint32
}

var (
	CRUDS map[string]func(unprepared Unprepared) (data []byte)
)

func init() {
	CRUDS = make(map[string]func(unprepared Unprepared) (data []byte))
}
