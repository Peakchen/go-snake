package datastatistics

/*
	for Histogram
*/

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type AKHistogram struct {
	obj    *prometheus.HistogramVec
	titles []string
}

func NewAKHistogram() *AKHistogram {
	return &AKHistogram{
		obj:    nil,
		titles: []string{},
	}
}

var (
	_akHistogram *AKHistogram
)

func GetAKHistogram() *AKHistogram {
	if _akHistogram == nil {
		_akHistogram = NewAKHistogram()
	}
	return _akHistogram
}

func (this *AKHistogram) Init(strName, strHelp string, titles []string) {
	this.obj = prometheus.NewHistogramVec(prometheus.CounterOpts{
		Name: strName,
		Help: strHelp,
	}, titles)
	this.titles = titles
	prometheus.MustRegister(this.obj)
}

func (this *AKHistogram) DoObserve(title string, val float64) (err error) {
	err = IsExistStatisticsTitle(this.titles, title)
	if err != nil {
		return
	}
	this.obj.WithLabelValues(title).Add(val)
	return
}

func RegHistogram(model string) {
	http.Handle("/"+model, promhttp.Handler())
}
