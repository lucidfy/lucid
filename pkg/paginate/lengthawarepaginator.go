package paginate

import (
	"math"
	"net/url"
	"reflect"
	"strconv"

	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/logger"
	"github.com/lucidfy/lucid/pkg/facade/response"
)

const DefaultView = "pkg/pagination/tailwind.go.html"

func Construct(items interface{}, total int, perPage int, currentPage int) *Paginate {
	p := Paginate{
		PerPage:     perPage,
		CurrentPage: currentPage,
	}
	p.Reconstruct(items, total)
	return &p
}

func (p *Paginate) Reconstruct(items interface{}, total int) *Paginate {
	p.Items = items
	p.Total = total
	p.LastPage = int(math.Ceil(float64(total) / float64(p.PerPage)))

	p.OnEachSide = 3
	p.Fragment = nil

	return p
}

func (p *Paginate) Links() string {
	return p.Render(nil)
}

func (p *Paginate) Render(view *string /*, data array*/) string {
	if view == nil {
		dv := DefaultView
		view = &dv
	}

	result, err := response.Render(
		[]string{*view},
		map[string]interface{}{
			"currentPage":     p.CurrentPage,
			"total":           p.Total,
			"firstItem":       p.FirstItem(),
			"hasMorePages":    p.HasMorePages(),
			"hasPages":        p.HasPages(),
			"lastItem":        p.LastItem(),
			"nextPageURL":     p.NextPageURL(),
			"onFirstPage":     p.OnFirstPage(),
			"previousPageURL": p.PreviousPageURL(),

			// here we provide the $elements
			"elements": p.Elements(),
		},
	)
	if err != nil {
		logger.Error("lengthawarepaginator.Render error: ", err)
	}
	return result
}

func (p *Paginate) HasMorePages() bool {
	return p.CurrentPage < p.LastPage
}

func (p *Paginate) NextPageURL() *string {
	if p.HasMorePages() {
		s := p.URL(p.CurrentPage + 1)
		return &s
	}
	return nil
}

func (p *Paginate) ToArray() map[string]interface{} {
	return map[string]interface{}{
		// use for pagination
		"first_item":     p.FirstItem(),
		"has_more_pages": p.HasMorePages(),
		"has_pages":      p.HasPages(),
		"last_item":      p.LastItem(),
		"on_first_page":  p.OnFirstPage(),
		"elements":       p.Elements(),

		// default data
		"current_page":   p.CurrentPage,
		"data":           p.Items,
		"first_page_url": p.URL(1),
		"from":           p.FirstItem(),
		"last_page":      p.LastPage,
		"last_page_url":  p.URL(p.LastPage),
		"next_page_url":  p.NextPageURL(),
		"per_page":       p.PerPage,
		"prev_page_url":  p.PreviousPageURL(),
		"to":             p.LastItem(),
		"total":          p.Total,
	}
}

func (p *Paginate) Elements() map[int]string {
	window := URLWindow(*p).Get()
	elems := window.first
	p.elementsLoop(&elems, window.slider)
	p.elementsLoop(&elems, window.last)
	return elems
}

func (p *Paginate) elementsLoop(elems *map[int]string, m map[int]string) map[int]string {
	if len(m) > 0 {
		(*elems)[len(*elems)+1] = "..."

		for _, value := range m {
			(*elems)[len(*elems)+1] = value
		}
	}
	return *elems
}

func (p Paginate) PreviousPageURL() *string {
	if p.CurrentPage > 1 {
		url := p.URL(p.CurrentPage - 1)
		return &url
	}
	return nil
}

func (p Paginate) GetURLRange(start int, end int) map[int]string {
	r := map[int]string{}
	for i := start; i <= end; i++ {
		r[i] = p.URL(i)
	}
	return r
}

func (p *Paginate) URL(page int) string {
	URL, err := url.Parse(p.BaseURL)
	if errors.Handler("url.Parse error", err) {
		return ""
	}

	q := URL.Query()
	q.Set("page", strconv.Itoa(page))

	URL.RawQuery = q.Encode()
	URL.Fragment = p.buildFragment()
	return URL.String()
}

func (p *Paginate) GetFragment(fragment *string) *Paginate {
	p.Fragment = fragment
	return p
}

func (p Paginate) FirstItem() *int {
	if p.Count() > 0 {
		computed := (p.CurrentPage-1)*p.PerPage + 1
		return &computed
	}
	return nil
}

func (p Paginate) LastItem() *int {
	if p.Count() > 0 {
		firstItem := p.FirstItem()
		computed := *firstItem + p.Count() - 1
		return &computed
	}
	return nil
}

func (p Paginate) HasPages() bool {
	return p.CurrentPage != 1 || p.HasMorePages()
}

func (p Paginate) OnFirstPage() bool {
	return p.CurrentPage <= 1
}

func (p *Paginate) SetOnEachSide(count int) *Paginate { // this is onEachSide() in illuminate
	p.OnEachSide = count
	return p
}

func (p *Paginate) IsEmpty() bool {
	return p.Count() == 0
}

func (p *Paginate) IsNotEmpty() bool {
	return !p.IsEmpty()
}

func (p *Paginate) Count() int {
	// since the p.Items is a pointer
	v := reflect.ValueOf(p.Items)

	// we need to indirect the value, as if it was like *v
	return reflect.Indirect(v).Len()
}

func (p *Paginate) ToHtml() string {
	return p.Render(nil)
}

func (p *Paginate) buildFragment() string {
	fragment := p.Fragment
	if fragment != nil {
		return "#" + *fragment
	}
	return ""
}

func (p *Paginate) GetTotal() int {
	return p.Total
}

func (p *Paginate) GetPerPage() int {
	return p.PerPage
}

func (p *Paginate) GetCurrentPage() int {
	return p.CurrentPage
}

func (p *Paginate) GetLastPage() int {
	return p.LastPage
}

func (p *Paginate) GetBaseURL() string {
	return p.BaseURL
}

func (p *Paginate) GetItems() interface{} {
	return p.Items
}
