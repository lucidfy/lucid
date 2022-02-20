package paginate

import (
	"net/url"
	"reflect"
	"strconv"

	"github.com/daison12006013/gorvel/pkg/errors"
)

type Paginate struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`

	BaseUrl string      `json:"base_url"`
	Items   interface{} `json:"items"`

	OnEachSide int
	Fragment   *string
}

func (p Paginate) PreviousPageUrl() *string {
	if p.CurrentPage > 1 {
		url := p.Url(p.CurrentPage - 1)
		return &url
	}
	return nil
}

func (p Paginate) GetUrlRange(start int, end int) map[int]string {
	r := map[int]string{}
	for i := start; i <= end; i++ {
		r[i] = p.Url(i)
	}
	return r
}

func (p *Paginate) Url(page int) string {
	URL, err := url.Parse(p.BaseUrl)
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

func (p *Paginate) GetBaseUrl() string {
	return p.BaseUrl
}

func (p *Paginate) GetItems() interface{} {
	return p.Items
}
