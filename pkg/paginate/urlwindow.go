package paginate

type Window struct {
	paginate Paginate
	first    map[int]string
	slider   map[int]string
	last     map[int]string
}

func UrlWindow(p Paginate) *Window {
	return &Window{paginate: p}
}

func (w *Window) Get() *Window {
	onEachSide := w.paginate.OnEachSide
	if w.paginate.LastPage < (onEachSide*2)+8 {
		return w.getSmallSlider()
	}
	return w.getUrlSlider(onEachSide)
}

func (w *Window) GetAdjacentUrlRange(onEachSide int) map[int]string {
	cp := w.paginate.CurrentPage
	return w.paginate.GetUrlRange(
		cp-onEachSide,
		cp+onEachSide,
	)
}

func (w *Window) GetStart() map[int]string {
	return w.paginate.GetUrlRange(1, 2)
}

func (w *Window) GetFinish() map[int]string {
	lp := w.paginate.LastPage
	return w.paginate.GetUrlRange(
		lp-1,
		lp,
	)
}

func (w *Window) HasPages() bool {
	return w.paginate.LastPage > 1
}

func (w *Window) getSmallSlider() *Window {
	w.first = w.paginate.GetUrlRange(1, w.paginate.LastPage)
	w.slider = make(map[int]string)
	w.last = make(map[int]string)
	return w
}

func (w *Window) getUrlSlider(onEachSide int) *Window {
	window := onEachSide + 4

	if !w.HasPages() {
		w.first = make(map[int]string)
		w.slider = make(map[int]string)
		w.last = make(map[int]string)
		return w
	}

	// If the current page is very close to the beginning of the page range, we will
	// just render the beginning of the page range, followed by the last 2 of the
	// links in this list, since we will not have room to create a full slider.
	if w.paginate.CurrentPage <= window {
		return w.getSliderTooCloseToBeginning(window, onEachSide)
	} else if w.paginate.CurrentPage > (w.paginate.LastPage - window) {
		// If the current page is close to the ending of the page range we will just get
		// this first couple pages, followed by a larger window of these ending pages
		// since we're too close to the end of the list to create a full on slider.
		return w.getSliderTooCloseToEnding(window, onEachSide)
	}

	// If we have enough room on both sides of the current page to build a slider we
	// will surround it with both the beginning and ending caps, with this window
	// of pages in the middle providing a Google style sliding paginator setup.
	return w.getFullSlider(onEachSide)
}

func (w *Window) getSliderTooCloseToBeginning(window int, onEachSide int) *Window {
	w.first = w.paginate.GetUrlRange(1, window+onEachSide)
	w.slider = make(map[int]string)
	w.last = w.GetFinish()
	return w
}

func (w *Window) getSliderTooCloseToEnding(window int, onEachSide int) *Window {
	lp := w.paginate.LastPage
	last := w.paginate.GetUrlRange(
		lp-(window+(onEachSide-1)),
		lp,
	)

	w.first = w.GetStart()
	w.slider = make(map[int]string)
	w.last = last
	return w
}

func (w *Window) getFullSlider(onEachSide int) *Window {
	w.first = w.GetStart()
	w.slider = w.GetAdjacentUrlRange(onEachSide)
	w.last = w.GetFinish()
	return w
}
