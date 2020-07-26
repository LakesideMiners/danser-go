package movers

import (
	"github.com/wieku/danser-go/animation/easing"
	"github.com/wieku/danser-go/beatmap/objects"
	"github.com/wieku/danser-go/bmath"
	"github.com/wieku/danser-go/bmath/curves"
)

type LinearMover struct {
	line               curves.Linear
	beginTime, endTime int64
}

func NewLinearMover() MultiPointMover {
	return &LinearMover{}
}

func (bm *LinearMover) Reset() {

}

func (bm *LinearMover) SetObjects(objs []objects.BaseObject) {
	end, start := objs[0], objs[1]
	endPos := end.GetBasicData().EndPos
	endTime := end.GetBasicData().EndTime
	startPos := start.GetBasicData().StartPos
	startTime := start.GetBasicData().StartTime

	bm.line = curves.NewLinear(endPos, startPos)

	bm.endTime = bmath.MaxI64(endTime, start.GetBasicData().StartTime-380)
	bm.beginTime = startTime
}

func (bm LinearMover) Update(time int64) bmath.Vector2f {
	t := float64(time-bm.endTime) / float64(bm.beginTime-bm.endTime)
	t = bmath.ClampF64(t, 0, 1)
	return bm.line.PointAt(float32(easing.OutQuad(t)))
}

func (bm *LinearMover) GetEndTime() int64 {
	return bm.beginTime
}
