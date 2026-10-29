package main

import (
	"github.com/Mikedev115/Aetox/internal/provider"
	"strings"
	"testing"
)

func TestDevinAppearsAsOfficialAccountAgent(t *testing.T) {
	app := &App{}
	for _, method := range app.SignInMethods() {
		if method.Provider == "devin" {
			if method.Kind != "terminal" || !strings.Contains(method.Note, "own agent") || !strings.Contains(method.Note, "quota") {
				t.Fatalf("missing limits: %+v", method)
			}
			spec, ok := provider.Lookup("devin")
			if !ok || spec.AcceptsAPIKey || spec.Runtime != provider.RuntimeExternalCLI {
				t.Fatalf("incorrect runtime: %+v", spec)
			}
			return
		}
	}
	t.Fatal("Devin absent from account onboarding")
}
