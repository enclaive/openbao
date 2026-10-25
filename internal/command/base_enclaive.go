package command

import (
	"github.com/enclaive/openbao/external/vhsm/command"
	"github.com/hashicorp/cli"
	"github.com/openbao/openbao/api/v2"
	"github.com/posener/complete"
)

func (c *BaseCommand) Ui() cli.Ui {
	return c.UI
}

func (c *BaseCommand) Format() string {
	return Format(c.UI)
}

func (c *BaseCommand) FlagSet() command.FlagSets {
	return &EnclaiveFlagSets{c.flagSet(FlagSetHTTP | FlagSetOutputFormat)}
}

func (c *BaseCommand) ExtractListData(secret *api.Secret) ([]any, bool) {
	return extractListData(secret)
}

func (c *BaseCommand) SanitizePath(s string) string {
	return sanitizePath(s)
}

func (c *BaseCommand) OutputData(m map[string]any) {
	OutputData(c.UI, m)
}

func (c *BaseCommand) OutputList(secret *api.Secret) int {
	return OutputList(c.UI, secret)
}

func (c *BaseCommand) OutputSecret(secret *api.Secret) int {
	return OutputSecret(c.UI, secret)
}

func (c *BaseCommand) HandleWriteSecretOutput(s string, secret *api.Secret, err error) int {
	return handleWriteSecretOutput(c, s, secret, err)
}

func (c *BaseCommand) FlagField() string {
	return c.flagField
}

func (c *BaseCommand) PrintRawField(secret *api.Secret, s string) int {
	return PrintRawField(c.UI, secret, s)
}

type EnclaiveFlagSets struct {
	*FlagSets
}

func (f *EnclaiveFlagSets) Parse(args []string) error {
	return f.FlagSets.Parse(args)
}

func (f *EnclaiveFlagSets) NewFlagSet(name string) command.FlagSet {
	return &EnclaiveFlagSet{f.FlagSets.NewFlagSet(name)}
}

type EnclaiveFlagSet struct {
	*FlagSet
}

func (f *EnclaiveFlagSet) BoolVar(name, usage string, fallback bool, target *bool, completion complete.Predictor) {
	f.FlagSet.BoolVar(&BoolVar{
		Name:       name,
		Usage:      usage,
		Default:    fallback,
		Target:     target,
		Completion: completion,
	})
}

func (f *EnclaiveFlagSet) StringVar(name, usage string, fallback string, target *string, completion complete.Predictor) {
	f.FlagSet.StringVar(&StringVar{
		Name:       name,
		Usage:      usage,
		Default:    fallback,
		Target:     target,
		Completion: completion,
	})
}
