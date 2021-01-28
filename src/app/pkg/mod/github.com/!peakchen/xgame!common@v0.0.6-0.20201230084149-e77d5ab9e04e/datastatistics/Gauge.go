package datastatistics

/*
	for Gauge
*/

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type AKGauge struct {
	obj    *prometheus.GaugeVec
	titles []string
}

func NewAKGauge() *AKGauge {
	return &AKGauge{
		obj:    nil,
		titles: []string{},
	}
}

var (
	_akgauge *AKGauge
)

func GetAKGauge() *AKGauge {
	if _akgauge == nil {
		_akgauge = NewAKGauge()
	}
	return _akgauge
}

func (this *AKGauge) Init(strName, strHelp string, titles []string) {
	this.obj = prometheus.NewGaugeVec(prometheus.CounterOpts{
		Name: strName,
		Help: strHelp,
	}, titles)
	this.titles = titles
	prometheus.MustRegister(this.obj)
}

func (this *AKGauge) DoSet(title string, val float64) (err error) {
	err = IsExistStatisticsTitle(this.titles, title)
	if err != nil {
		return
	}

	this.obj.WithLabelValues(title).Set(val)
}

func (this *AKGauge) DoInc(title string) (err error) {
	err = IsExistStatisticsTitle(this.titles, title)
	if err != nil {
		return
	}
	this.obj.WithLabelValues(title).Inc()
}

func (this *AKGauge) DoDec(title string) (err error) {
	err = IsExistStatisticsTitle(this.titles, title)
	if err != nil {
		return
	}
	this.obj.WithLabelValues(title).Dec()
}

func (this *AKGauge) DoAdd(title string, val float64) (err error) {
	err = IsExistStatisticsTitle(this.titles, title)
	if err != nil {
		return
	}
	this.obj.WithLabelValues(title).Add(val)
	return
}

func (this *AKGauge) DoSub(title string, val float64) (err error) {
	err = IsExistStatisticsTitle(this.titles, title)
	if err != nil {
		return
	}
	this.obj.WithLabelValues(title).Sub(val)
	return
}

func RegGauge(model string) {
	http.Handle("/"+model, promhttp.Handler())
}
