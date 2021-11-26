package schema

type Clock struct {
	Base        `json:",inline" valid:"required"`
	Type        int    `json:"type" valid:"required"`
	Clock       string `json:"clock" valid:"clock,required"`
	Timezone    string `json:"timezone" valid:"timezone,required"`
	WorkScaleID string `json:"work_scale_id" valid:"uuid"`
}

func NewClock() *Clock {
	return &Clock{}
}
