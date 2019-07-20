package display

type Spinner interface {
	Spin() string
}

type spinner struct {
	set []string
	state int
}

func NewDefaultSpinner() Spinner {
	// {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	// {"⬖", "⬘", "⬗", "⬙"}
	// {"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"}
	// {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁", " "}
	// {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	// {"⠋", "⠙", "⠸", "⠴", "⠦", "⠇"}
	// {"◢", "◣", "◤", "◥"}
	// {"◜", "◝", "◞", "◟"}
	// []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}

	var set = []string{"-", "\\", "|", "/"}
	return &spinner{set, 0}
}

func NewCustomSpinner(chars []string) Spinner {
	return &spinner{chars, 0}
}

func (s *spinner) Spin() string {
	var result = s.set[s.state]
	s.state = (s.state+1)%len(s.set)
	return result
}