package pezdispenser

//DispenserFinder - interfcae for a object to find Dispensers
type DispenserFinder interface {
	GetAll() []Dispenser
	GetByTypeGUID(string) Dispenser
	GetByItemGUID(string) Dispenser
}

//Dispenser - a interface for leasable inventory
type Dispenser interface {
	GUID() string
	Status() ([]byte, error)
	Lock() ([]byte, error)
	UnLock() ([]byte, error)
	Lease() ([]byte, error)
	Unlease() ([]byte, error)
	Renew() ([]byte, error)
}

//NewFinder - returns a DispenserFinder interface
func NewFinder() (f DispenserFinder) {
	f = new(finder)
	return
}

type finder struct{}

func (s *finder) GetAll() (d []Dispenser) {
	d1 := &fakeDispenser{guid: "abcGUID"}
	d2 := &fakeDispenser{guid: "123GUID"}
	d = []Dispenser{d1, d2}
	return
}

func (s *finder) GetByTypeGUID(guid string) (d Dispenser) {
	d = &fakeDispenser{guid: guid}
	return
}

func (s *finder) GetByItemGUID(guid string) (d Dispenser) {
	d = &fakeDispenser{guid: guid}
	return
}

type fakeDispenser struct {
	guid string
}

func (s *fakeDispenser) GUID() (guid string) {
	guid = s.guid
	return
}

func (s *fakeDispenser) Status() (res []byte, err error) {
	res = []byte(s.guid)
	return
}

func (s *fakeDispenser) Lock() (res []byte, err error) {
	res = []byte(s.guid)
	return
}

func (s *fakeDispenser) UnLock() (res []byte, err error) {
	res = []byte(s.guid)
	return
}

func (s *fakeDispenser) Lease() (res []byte, err error) {
	res = []byte(s.guid)
	return
}

func (s *fakeDispenser) Unlease() (res []byte, err error) {
	res = []byte(s.guid)
	return
}

func (s *fakeDispenser) Renew() (res []byte, err error) {
	res = []byte(s.guid)
	return
}
