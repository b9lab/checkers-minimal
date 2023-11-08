package checkers

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
func (gs *GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	unique := make(map[string]bool)
	for _, indexedStoredGame := range gs.IndexedStoredGameList {
		if length := len([]byte(indexedStoredGame.Index)); MaxIndexLength < length || length < 1 {
			return ErrIndexTooLong
		}
		if _, ok := unique[indexedStoredGame.Index]; ok {
			return ErrDuplicateAddress
		}
		if err := indexedStoredGame.StoredGame.Validate(); err != nil {
			return err
		}
		unique[indexedStoredGame.Index] = true
	}

	return nil
}
