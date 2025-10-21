package uv1schedulehttp

type (
	schedule__Reqp struct {
		CinemaId  uint   `json:"cinema_id" form:"cinema_id" xml:"cinema_id" validate:"required"`
		MovieId   uint   `json:"movie_id" form:"movie_id" xml:"movie_id" validate:"required"`
		ShowDate  string `json:"show_date" form:"show_date" xml:"show_date" validate:"required"`
		StartTime string `json:"start_time" form:"start_time" xml:"start_time" validate:"required"`
		EndTime   string `json:"end_time" form:"end_time" xml:"end_time" validate:"required"`
	}

	list__Reqp struct {
		/// The dataTable search request.
		Draw   int    `json:"draw" form:"draw" xml:"draw"`       /// The dataTable draw flag.
		Search string `json:"search" form:"search" xml:"search"` /// The dataTable search request.
		Length int    `json:"length" form:"length" xml:"length"` /// The dataTable offset request.
		Offset int    `json:"offset" form:"offset" xml:"offset"` // The dataTable length or limit request.
	}
)
