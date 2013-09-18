package main

import (
	"html/template"
	"io"
	"os"
)

type GoogleAnalytics struct {
	Code string
}

// RenderOn writes the result of executing its template with the receiver as its data
// implements renderbee.Renderable so it can be decorated
//
func (g GoogleAnalytics) RenderOn(w io.Writer) {
	GoogleAnalytics_Template.Execute(w, g)
}

var GoogleAnalytics_Template = template.Must(template.New("GoogleAnalytics").Parse(`
<script>
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
ga('create', '{{.Code}}');
ga('send', 'pageview');
</script>
`))

func main() {
	GoogleAnalytics{"UA999"}.RenderOn(os.Stdout)
}