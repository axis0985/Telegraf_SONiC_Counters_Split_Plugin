package split

// split.go

import (
    "strings"

    metric "github.com/influxdata/telegraf/metric"
    "github.com/influxdata/telegraf"
    "github.com/influxdata/telegraf/plugins/processors"
)

type Splitter struct {
    Log telegraf.Logger `toml:"-"`
}

var sampleConfig = `
`

func (p *Splitter) SampleConfig() string {
    return sampleConfig
}

func (p *Splitter) Description() string {
    return "Split metrics with certain pattern."
}

// Init is for setup, and validating config.
func (p *Splitter) Init() error {
    return nil
}

func (p *Splitter) Apply(in ...telegraf.Metric) []telegraf.Metric {
    var splitted []telegraf.Metric
    for _, oldMetric := range in {
        fields := oldMetric.Fields()
        splitMap := make(map[string]([]*telegraf.Field))
        for key, val := range fields {
            splittedKey := strings.SplitN(key, "_", 2)
            // if _, ok := splitMap[splittedKey[0]]; !ok {
            //     splitMap[splittedKey[0]] = make([]*telegraf.Field)
            // }
            splitMap[splittedKey[0]] = append(splitMap[splittedKey[0]],
                    &telegraf.Field{
                        Key: splittedKey[1],
                        Value: val,
                    })
        }
        for key, fields := range splitMap {
            newMetric := metric.New(
                oldMetric.Name(),
                oldMetric.Tags(),
                nil,
                oldMetric.Time(),
                oldMetric.Type(),
            )
            newMetric.AddTag("port", key)
            for _, field := range fields {
                newMetric.AddField(
                    field.Key,
                    field.Value,
                )
            }

            splitted = append(splitted, newMetric)
        }
    }
    return splitted
}

func init() {
    processors.Add("split", func() telegraf.Processor {
        return &Splitter{}
    })
}