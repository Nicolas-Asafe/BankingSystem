package entities

type Economy struct {
	MinValueForAdd           float64
	MinValueForWithDrawMoney float64
	MaxValueforWithDrawMoney float64
	AuthorId                 string
	Description              *string
	Title                    string
	Value                    float64
}

func (e *Economy) AddValueForValue(Value float64) {
	if Value < e.MinValueForAdd {
		return
	}
	e.Value += Value
}

func (e *Economy) WithdrawMoney(Value float64) {
	if Value > e.MaxValueforWithDrawMoney {
		return
	}
	if Value < e.MinValueForWithDrawMoney {
		return
	}
}
