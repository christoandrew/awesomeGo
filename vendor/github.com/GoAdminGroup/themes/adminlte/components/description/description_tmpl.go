package description

var List = map[string]string{
	"description": `{{define "description"}}
<div class="description-block border-{{.Border}}">
    <span class="description-percentage text-{{.Color}}"><i class="fa fa-caret-{{.Arrow}}"></i>{{langHtml .Percent}}%</span>
    <h5 class="description-header">{{langHtml .Number}}</h5>
    <span class="description-text">{{langHtml .Title}}</span>
</div>
{{end}}`,
}