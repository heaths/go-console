package console

import (
	"time"

	"github.com/briandowns/spinner"
)

type ProgressStyle int

const (
	// https://github.com/briandowns/spinner#available-character-sets
	ProgressStyleBars ProgressStyle = 9
	ProgressStyleDots ProgressStyle = 11
)

type ProgressOption func(*spinner.Spinner)

func (c *Console) StartProgress(label string, opts ...ProgressOption) {
	if !c.progressEnabled || !c.IsStderrTTY() {
		return
	}

	c.progressLock.Lock()
	defer c.progressLock.Unlock()

	cs := spinner.CharSets[int(ProgressStyleDots)]
	sp := spinner.New(
		cs,
		120*time.Millisecond,
		spinner.WithWriter(c.stderr),

		// TODO: Allow specifying another color.
		spinner.WithColor("fgCyan"),
	)

	if label != "" {
		sp.Suffix = " " + label
	}

	sp.Start()
	c.progress = sp
}

func (c *Console) StopProgress() {
	c.progressLock.Lock()
	defer c.progressLock.Unlock()

	if c.progress == nil {
		return
	}

	c.progress.Stop()
	c.progress = nil
}

func WithProgressStyle(style ProgressStyle) ProgressOption {
	return func(sp *spinner.Spinner) {
		cs := spinner.CharSets[int(style)]
		sp.UpdateCharSet(cs)
	}
}
