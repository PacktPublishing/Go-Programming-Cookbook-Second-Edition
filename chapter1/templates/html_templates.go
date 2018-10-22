package templates

import (
	"fmt"
	"html/template"
	"os"
)

// HTMLDifferences highlights some of the differences
// between html/template and text/template
func HTMLDifferences() error {
	t := template.New("html")
	t, err := t.Parse("<h1>Hello! {{.Name}}</h1>\n")
	if err != nil {
		return err
	}

	// html/template auto-escapes unsafe operations like javascript injection
	// this is contextually aware and will behave differently
	// depending on where a variable is rendered
	err = t.Execute(os.Stdout, map[string]string{"Name": "<script>alert('Can you see me?')</script>"})
	if err != nil {
		return err
	}

	// you can also manually call the escapers
	fmt.Println(template.JSEscaper(`example <example@example.com>`))
	fmt.Println(template.HTMLEscaper(`example <example@example.com>`))
	fmt.Println(template.URLQueryEscaper(`example <example@example.com>`))

	return nil
}
