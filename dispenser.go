package pezdispenser

type DispenserFinder interface {
	GetAll() []Dispenser
	GetByTypeGUID(string) Dispenser
	GetByItemGUID(string) Dispenser
}

type Dispenser interface {
	GUID() string
	Status() ([]byte, error)
	Lock() ([]byte, error)
	UnLock() ([]byte, error)
	Lease() ([]byte, error)
	Unlease() ([]byte, error)
	Renew() ([]byte, error)
}

func NewFinder() (f DispenserFinder) {
	f = new(finder)
	return
}

type finder struct{}

func (s *finder) GetAll() (d []Dispenser) {
	d1 := &FakeDispenser{guid: "abcGUID"}
	d2 := &FakeDispenser{guid: "123GUID"}
	d = []Dispenser{d1, d2}
	return
}

func (s *finder) GetByTypeGUID(guid string) (d Dispenser) {
	d = &FakeDispenser{guid: guid}
	return
}

func (s *finder) GetByItemGUID(guid string) (d Dispenser) {
	d = &FakeDispenser{guid: guid}
	return
}

type FakeDispenser struct {
	guid string
}

func (s *FakeDispenser) GUID() (guid string) {
	guid = s.guid
	return
}

func (s *FakeDispenser) Status() (res []byte, err error) {
	res = []byte(s.guid)
	return
}

func (s *FakeDispenser) Lock() (res []byte, err error) {
	res = []byte(s.guid)
	return
}

func (s *FakeDispenser) UnLock() (res []byte, err error) {
	res = []byte(s.guid)
	return
}

func (s *FakeDispenser) Lease() (res []byte, err error) {
	res = []byte(s.guid)
	return
}

func (s *FakeDispenser) Unlease() (res []byte, err error) {
	res = []byte(s.guid)
	return
}

func (s *FakeDispenser) Renew() (res []byte, err error) {
	res = []byte(s.guid)
	return
}
