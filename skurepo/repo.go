package skurepo

//Register -- add a Sku interface object to the Repo
func Register(name string, sku Sku) {
	Repo[name] = sku
}

//GetRegistry -- gets the map of all registered Sku interface objects
func GetRegistry() map[string]Sku {
	return Repo
}
