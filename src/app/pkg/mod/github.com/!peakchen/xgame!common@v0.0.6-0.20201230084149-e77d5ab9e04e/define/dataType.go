package define

/*
	add by stefan
	date: 20191107 14:36
	purpose: common config data type use.
*/

import (
	"bytes"
	"github.com/Peakchen/xgameCommon/akLog"
	"strconv"
)

// array int with "1,2,3"
type Int32Array []int32

func (this *Int32Array) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 {
		err = nil
		return
	}

	data = data[1 : len(data)-1]
	this = &Int32Array{}
	err = this.UnmarshalJSONEx(data)
	return
}

func (this *Int32Array) UnmarshalJSONEx(data []byte) (err error) {
	data = bytes.TrimSpace(data)
	intarr := bytes.Split(data, []byte(","))
	for _, sval := range intarr {
		sval = bytes.TrimSpace(sval)
		nval, converr := strconv.Atoi(string(sval))
		if converr != nil {
			err = akLog.RetError("IntArry unmarshal json fail, err: %v, val: %v.", converr, sval)
			return
		}
		*this = append(*this, int32(nval))
	}
	akLog.FmtPrintln("Int32Array: ", *this)
	err = nil
	return
}

// array int with "1,2,3;4,5,6"
type Int32Array2D []Int32Array

func (this *Int32Array2D) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 {
		return
	}

	this = &Int32Array2D{}
	data = bytes.TrimSpace(data)
	data = data[1 : len(data)-1]
	splitItems := bytes.Split(data, []byte(";"))
	for _, items := range splitItems {
		childs := &Int32Array{}
		childerr := childs.UnmarshalJSONEx(items)
		if childerr != nil {
			err = akLog.RetError("Int32Array2D Unmashal fail, err: %v.", childerr)
			return
		}

		if len(*childs) > 0 {
			*this = append(*this, *childs)
		}
	}
	akLog.FmtPrintln("Int32Array2D: ", *this)
	err = nil
	return
}

// prop struct
type Property struct {
	PropID  int32
	PropNum int32
}

func (this *Property) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 {
		return
	}

	data = bytes.TrimSpace(data)
	data = data[1 : len(data)-1]
	err = this.UnmarshalJSONEx(data)
	return
}

func (this *Property) UnmarshalJSONEx(data []byte) (err error) {
	if !bytes.Contains(data, []byte(",")) {
		err = akLog.RetError("Invalid data: %v.", data)
		return
	}

	splitItems := bytes.Split(data, []byte(","))
	if len(splitItems) != 2 {
		err = akLog.RetError("Property childitems Invalid, data: %v, len is not equal 2.", string(data))
		return
	}

	this = &Property{}
	byItem1 := bytes.TrimSpace(splitItems[0])
	propid, covererr := strconv.Atoi(string(byItem1))
	if covererr != nil {
		err = akLog.RetError("Property items cover fail, data1: %v, err: %v.", string(data[0]), covererr)
		return
	}

	(*this).PropID = int32(propid)
	byItem2 := bytes.TrimSpace(splitItems[1])
	propnum, covererr := strconv.Atoi(string(byItem2))
	if covererr != nil {
		err = akLog.RetError("Property items cover fail, data2: %v, err: %v.", string(data[1]), covererr)
		return
	}

	(*this).PropNum = int32(propnum)
	akLog.FmtPrintln("Property: ", *this)
	err = nil
	return
}

// propery array
type PropertyArray []*Property

func (this *PropertyArray) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 {
		return
	}

	data = bytes.TrimSpace(data)
	data = data[1 : len(data)-1]
	splitItems := bytes.Split(data, []byte(";"))
	if len(splitItems) == 0 {
		return
	}

	this = &PropertyArray{}
	for _, childitems := range splitItems {
		propitem := &Property{}
		childerr := propitem.UnmarshalJSONEx(childitems)
		if childerr != nil {
			err = akLog.RetError("PropertyArray child Unmarshl fail, er: %v.", childerr)
			return
		}
		*this = append(*this, propitem)
	}
	akLog.FmtPrintln("PropertyArray: ", *this)
	err = nil
	return
}
