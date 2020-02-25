
package main

import "github.com/imranh27/snippetbox/pkg/models"

//define a templateData type to act as the holding structure for any holding data to pass to our HTML templates.

type templateData struct {
	Snippet *models.Snippet
}