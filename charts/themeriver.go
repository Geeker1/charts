package charts

import (
	"github.com/Geeker1/charts/v2/opts"
	"github.com/Geeker1/charts/v2/render"
	"github.com/Geeker1/charts/v2/types"
)

// ThemeRiver represents a theme river chart.
type ThemeRiver struct {
	BaseConfiguration
	BaseActions
}

// Type returns the chart type.
func (ThemeRiver) Type() string { return types.ChartThemeRiver }

// NewThemeRiver creates a new theme river chart.
func NewThemeRiver() *ThemeRiver {
	c := &ThemeRiver{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasSingleAxis = true
	return c
}

// AddSeries adds new data sets.
func (c *ThemeRiver) AddSeries(name string, data []opts.ThemeRiverData, options ...SeriesOpts) *ThemeRiver {
	cd := make([][3]interface{}, len(data))
	for i := 0; i < len(data); i++ {
		cd[i] = data[i].ToList()
	}
	series := SingleSeries{Name: name, Type: types.ChartThemeRiver, Data: cd}
	series.configureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the ThemeRiver instance.
func (c *ThemeRiver) SetGlobalOptions(options ...GlobalOpts) *ThemeRiver {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// SetDispatchActions sets actions for the ThemeRiver instance.
func (c *ThemeRiver) SetDispatchActions(actions ...GlobalActions) *ThemeRiver {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

// Validate validates the given configuration.
func (c *ThemeRiver) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
