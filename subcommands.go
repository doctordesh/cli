package cli

type SubCommands []*SubCommand

func (self *SubCommands) Get(key string) (*SubCommand, bool) {
	var v *SubCommand
	for _, v = range *self {
		if v.Name == key {
			return v, true
		}
	}

	return v, false
}
