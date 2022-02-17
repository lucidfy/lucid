package paginate

import (
	"reflect"
	"strconv"
	"strings"
)

type Paginate struct {
	Total       int
	PerPage     int
	CurrentPage int
	LastPage    int

	BaseUrl    string
	Items      interface{}
	OnEachSide int

	fragment *string
	// query    map[string]interface{} // * Still unused!
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
	c := "?"
	if strings.Contains(p.BaseUrl, "?") {
		c = "&"
	}
	return p.BaseUrl + c + "page=" + strconv.Itoa(page) + p.buildFragment()
}

func (p *Paginate) Fragment(fragment *string) *Paginate {
	p.fragment = fragment
	return p
}

// ! NOT APPLICABLE: func (p *Paginate) Appends(key string, $value = null) {}
// ! NOT APPLICABLE: func (p *Paginate) WithQueryString() {}
// ! NOT APPLICABLE: func (p *Paginate) LoadMorph($relation, $relations) {}
// ! NOT APPLICABLE: func (p *Paginate) LoadMorphCount($relation, $relations) {}

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

// ! NOT APPLICABLE: func (p *Paginate) Through(callable $callback) {}
// ! NOT APPLICABLE: func (p *Paginate) PerPage() {}

func (p Paginate) HasPages() bool {
	return p.CurrentPage != 1 || p.HasMorePages()
}

func (p Paginate) OnFirstPage() bool {
	return p.CurrentPage <= 1
}

// ! NOT APPLICABLE: func (p *Paginate) CurrentPage() {}
// ! NOT APPLICABLE: func (p *Paginate) GetPageName() {}
// ! NOT APPLICABLE: func (p *Paginate) SetPageName($name) {}
// ! NOT APPLICABLE: func (p *Paginate) WithPath($path) {}
// ! NOT APPLICABLE: func (p *Paginate) SetPath($path) {}

func (p *Paginate) SetOnEachSide(count int) *Paginate { // this is onEachSide() in illuminate
	p.OnEachSide = count
	return p
}

// ! NOT APPLICABLE: func (p *Paginate) Path() {}
// ! NOT APPLICABLE: func (p *Paginate) GetIterator() {}

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

// ! NOT APPLICABLE: func (p *Paginate) GetCollection() {}
// ! NOT APPLICABLE: func (p *Paginate) SetCollection(Collection $collection) {}
// ! NOT APPLICABLE: func (p *Paginate) GetOptions() {}
// ! NOT APPLICABLE: func (p *Paginate) OffsetExists(key string) {}
// ! NOT APPLICABLE: func (p *Paginate) OffsetGet(key string) {}
// ! NOT APPLICABLE: func (p *Paginate) OffsetSet(key string, $value) {}
// ! NOT APPLICABLE: func (p *Paginate) OffsetUnset(key string) {}

func (p *Paginate) ToHtml() string {
	return p.Render(nil)
}

func (p *Paginate) buildFragment() string {
	fragment := p.fragment
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

/*
* Still unused!
func (p *Paginate) isValidPageNumber(page int) bool {
	return page >= 1
}

func (p *Paginate) appendArray(values map[string]interface{}) *Paginate {
	for key, value := range values {
		p.addQuery(key, value)
	}
	return p
}

func (p *Paginate) addQuery(key string, value interface{}) *Paginate {
	p.query[key] = value
	return p
}
*/
