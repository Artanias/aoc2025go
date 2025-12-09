package tools

type Range struct {
	Start int64
	End   int64
}

func (r Range) Len() int64 {
	return r.End - r.Start + 1
}
