package datastatistics

/*
	for Summary
*/

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type AKSummary struct {
	obj    *prometheus.SummaryVec
	titles []string
}

func NewAKSummary() *AKSummary {
	return &AKSummary{
		obj:    nil,
		titles: []string{},
	}
}

var (
	_akSummary *AKSummary
)

func GetAKSummary() *AKSummary {
	if _akSummary == nil {
		_akSummary = NewAKSummary()
	}
	return _akSummary
}

func (this *AKSummary) Init(strName, strHelp string, titles []string) {
	this.obj = prometheus.NewSummaryVec(prometheus.CounterOpts{
		Name: strName,
		Help: strHelp,
	}, titles)
	this.titles = titles
	prometheus.MustRegister(this.obj)
}

func (this *AKSummary) DoObserve(title string, val float64) (err error) {
	err = IsExistStatisticsTitle(this.titles, title)
	if err != nil {
		return
	}
	this.obj.WithLabelValues(title).Add(val)
	return
}

func RegSummary(model string) {
	http.Handle("/"+model, promhttp.Handler())
}
