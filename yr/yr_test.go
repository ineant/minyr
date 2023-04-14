package yr

import (
	"testing"
)

func TestCelsiusToFarhrenheitString(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "6", want: "42.8"},
		{input: "0", want: "32.0"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFarhrenheitString(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperatrmaaling i grader celsius
func TestCelsiusToFahrenheitLine(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;18.03.2022 01:50", want: ""},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
	}
	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitLine(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}

}

func TestAverageCelsius(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "c", want: "8.56"},
	}
	for _, tc := range tests {
		got, _ := AverageCelsius(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}
}

func TestCountLinesInFiles(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "../kjevik-temp-fahr-20220318-20230318.csv", want: "16756"},
	}
	for _, tc := range tests {
		got, _ := CountLinesInFile(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}
}

/*func TestReadLastLine(t *testing.T) {
	input := "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;"
	want := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Ine Antonsen"
	got, err := readLastLine("../Kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Errorf("readLastLine error: %v", err)
	}
	if got != want {
		t.Errorf("readLastLine(%q) = %q, want %q", input, got, want)
	}
}*/
