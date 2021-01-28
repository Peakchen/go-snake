package datastatistics

/*
	for Counter
*/

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type AKCounter struct {
	obj    *prometheus.CounterVec
	titles []string
}

func NewAKCounter() *AKCounter {
	return &AKCounter{
		obj:    nil,
		titles: []string{},
	}
}

var (
	_akcounter *AKCounter
)

func GetAKCounter() *AKCounter {
	if _akcounter == nil {
		_akcounter = NewAKCounter()
	}
	return _akcounter
}

func (this *AKCounter) Init(strName, strHelp string, titles []string) {
	this.obj = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: strName,
		Help: strHelp,
	}, titles)
	this.titles = titles
	prometheus.MustRegister(this.obj)
}

func (this *AKCounter) DoAdd(title string, val float64) (err error) {
	err = IsExistStatisticsTitle(this.titles, title)
	if err != nil {
		return
	}
	this.obj.WithLabelValues(title).Add(val)
	return
}

func (this *AKCounter) DoInc(title string) (err error) {
	err = IsExistStatisticsTitle(this.titles, title)
	if err != nil {
		return
	}
	this.obj.WithLabelValues(title).Inc()
}

func RegCounter(model string) {
	http.Handle("/"+model, promhttp.Handler())
}
