/*

Package templates set type for message structure

*/

package templates

type MessageBox struct {
	MesQuestion bool
	MesWarning  bool
	MesError    bool
	MessageText string
	MessageHref string
}
