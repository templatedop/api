package config

type Options struct {
	FileName  string
	FilePaths []string
}

func DefaultConfigOptions() Options {
	opts := Options{
		FileName: "config",
		FilePaths: []string{
			".",
			"./configs",
		},
	}

	return opts
}

type ConfigOption func(o *Options)

func WithFileName(n string) ConfigOption {
	return func(o *Options) {
		o.FileName = n
	}
}

func WithFilePaths(p ...string) ConfigOption {
	return func(o *Options) {
		o.FilePaths = p
	}
}
