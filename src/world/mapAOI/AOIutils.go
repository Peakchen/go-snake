package mapAOI


type AOIutils struct {
	NbAOIhorizontal int
	LastAOIid 		int
}

func NewAOIutils()*AOIutils{
	return &AOIutils{

	}
}

func (this *AOIutils) ListAdjacentAOIs(current int)(aois []int){
	
	var isAtTop = (current < this.NbAOIhorizontal)
	var isAtBottom = (current > this.LastAOIid - this.NbAOIhorizontal)
	var isAtLeft = (current % this.NbAOIhorizontal == 0)
	var isAtRight = (current % this.NbAOIhorizontal == this.NbAOIhorizontal - 1)
	
	aois = append(aois, current)

	if !isAtTop {
		aois = append(aois, current - this.NbAOIhorizontal)
	}

	if !isAtBottom {
		aois = append(aois, current + this.NbAOIhorizontal)
	}

	if !isAtLeft {
		aois = append(aois, current - 1)
	}

	if !isAtRight {
		aois = append(aois, current + 1)
	}

	if !isAtTop && !isAtLeft {
		aois = append(aois, current - 1 - this.NbAOIhorizontal)
	}

	if !isAtTop && !isAtRight {
		aois = append(aois, current + 1 - this.NbAOIhorizontal)
	}

	if !isAtBottom && !isAtLeft {
		aois = append(aois, current - 1 + this.NbAOIhorizontal)
	}

	if !isAtBottom && !isAtRight {
		aois = append(aois, current + 1 + this.NbAOIhorizontal)
	}

	return

}