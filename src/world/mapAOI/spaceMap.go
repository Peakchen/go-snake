package mapAOI

import (
	"sort"
	"reflect"
)

// A space map is a custom data struture, similar to a sparse 2D array. Entities are stored according to their coordinates;
// that is, two keys are needed to fetch entities, the x position and the y position. This allows fast look-up based on position.


type SpaceMap struct {
	spaces map[interface{}]map[interface{}][]interface{}
}

func NewSpaceMap()*SpaceMap{
	return &SpaceMap{
		spaces: map[interface{}]map[interface{}][]interface{}{},
	}
}

func find(src []interface{}, dst interface{}) int{

	dstT := reflect.TypeOf(dst)

	return sort.Search(len(src), func(i int)bool{
		
		itT := reflect.TypeOf(src[i])
		
		if dstT != itT {
			return false
		}

		srcf := reflect.New(reflect.TypeOf(reflect.ValueOf(src[i]).Elem().Interface())).Interface()
		dstf := reflect.New(reflect.TypeOf(reflect.ValueOf(dst).Elem().Interface())).Interface()

		return reflect.DeepEqual(srcf, dstf)
	})

}

func (this *SpaceMap) Add(x, y, obj interface{}){

	item := this.spaces[x]
	if item == nil {
		item = map[interface{}][]interface{}{}
	}

	if item[y] == nil {
		item[y] = []interface{}{}
	}

	item[y] = append(item[y], obj)

	this.spaces[x] = item

}

func (this *SpaceMap) Delete(x, y interface{})bool{

	item := this.spaces[x]
	if item == nil {
		return false
	}

	if item[y] == nil {
		return false
	}

	delete(item, y)

	if len(item) == 0 {
		delete(this.spaces, x)
	}

	return true

}

func (this *SpaceMap) Move(x1, x2, y1, y2, obj interface{}){
	this.Delete(x1,y1)
	this.Add(x2,y2,obj)
}

func (this *SpaceMap) Get(x, y interface{})[]interface{}{
	
	item := this.spaces[x]
	if item == nil {
		return nil
	}

	if item[y] == nil {
		return nil
	}

	return item[y]

}

func (this *SpaceMap) GetFirst(x, y interface{})interface{}{

	item := this.spaces[x]
	if item == nil {
		return nil
	}

	if item[y] == nil {
		return nil
	}

	return item[y][0]

}

func (this *SpaceMap) GetFirstFiltered(x, y, filters, notFilters interface{})interface{}{

	// filters is an array of property names that need to be true
    // notFilters is an array of property names that need to be false
    // Returns the first entity at the given position, for which the values in filters are true and the values in notFilters are false
    // e.g. return the first item on a given cell that is visible but is not a chest

	if reflect.TypeOf(filters).Kind() != reflect.Slice {
		return nil
	}

	sliceFilters := filters.([]int)

	var slicenotFilters []int
	if notFilters != nil {
		slicenotFilters = notFilters.([]int)
	}

	var objs = this.Get(x,y)
	if objs == nil {
		return nil
	}

	for i := 0; i < len(objs); i++{

		var ok bool = true

		for f := 0; f < len(sliceFilters); f++ {
			if objs[i] == nil {
				ok = false
				break
			}

			if reflect.TypeOf(objs[i]) != reflect.TypeOf([]interface{}{}){
				ok = false
				break
			}
			
			sliceObjs := objs[i].([]interface{})
			if sliceObjs[sliceFilters[f]] == 0 {
				ok = false
				break
			}

		}

		if !ok {
			break
		}

		for f := 0; f < len(slicenotFilters); f++ {
			if objs[i] == nil {
				ok = false
				break
			}

			if reflect.TypeOf(objs[i]) != reflect.TypeOf([]interface{}{}){
				ok = false
				break
			}
			
			sliceObjs := objs[i].([]interface{})
			if sliceObjs[slicenotFilters[f]] == nil {
				ok = false
				break
			}

		}

		if ok {
			return objs[i]
		}

	}

	return nil
}

func (this *SpaceMap) GetAll(fnCall int)(l []interface{}){

	for _, xitems := range this.spaces {

		for _, yitems := range xitems {

			if (fnCall >= 0){

				for _, v := range yitems {
					fn := v.([]interface{})[fnCall]
					l = append(l, reflect.ValueOf(fn).Call([]reflect.Value{}))
				}

			}else{

				for _, v := range yitems {
					l = append(l, v.([]interface{})...)
				}

			}
		}

	}

	return

}