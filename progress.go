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

type ProgressOption func(*con, *spinner.Spinner)

func (c *con) StartProgress(label string, opts ...ProgressOption) {
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

	for _, opt := range opts {
		opt(c, sp)
	}

	sp.Start()
	c.progress = sp
}

func (c *con) StopProgress() {
	c.progressLock.Lock()
	defer c.progressLock.Unlock()

	if c.progress == nil {
		return
	}

	if c.progressMin != nil {
		<-c.progressMin
		c.progressMin = nil
	}

	c.progress.Stop()
	c.progress = nil
}

func WithMinimum(d time.Duration) ProgressOption {
	return func(c *con, _ *spinner.Spinner) {
		c.progressMin = time.After(d)
	}
}

func WithProgressStyle(style ProgressStyle) ProgressOption {
	return func(_ *con, sp *spinner.Spinner) {
		cs := spinner.CharSets[int(style)]
		sp.UpdateCharSet(cs)
	}
}
