package pleat

// Hook ...
type Hook interface {
	Levels() []Level
	Fire(*Entry) error
}

// LevelHooks ...
type LevelHooks map[Level][]Hook

// Add ...
func (hooks LevelHooks) Add(hook Hook) {
	panic("LevelHooks.Add not implemented...yet...")
}

// Fire ...
func (hooks LevelHooks) Fire(level Level, entry *Entry) error {
	panic("LevelHooks.Fire not implemented...yet...")

	return nil
}
