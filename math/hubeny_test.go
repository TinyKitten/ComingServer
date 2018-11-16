package math

import "testing"

func TestHubenyDistance(t *testing.T) {
	tokyo := Coordinate{
		Latitude:  35.65500,
		Longitude: 139.74472,
	}
	tsukuba := Coordinate{
		Latitude:  36.10056,
		Longitude: 140.09111,
	}

	d := HubenyDistance(tokyo, tsukuba)
	if d != 58502.4589312431 {
		t.Fatalf("Tsukuba to Tokyo, not matched: %f", d)
	}
}

func TestHubenyDistance2(t *testing.T) {
	tokyo := Coordinate{
		Latitude:  35.65500,
		Longitude: 139.74472,
	}
	fukuoka := Coordinate{
		Latitude:  33.59532,
		Longitude: 130.36208,
	}

	d := HubenyDistance(tokyo, fukuoka)
	if d != 890233.0643098591 {
		t.Fatalf("Fukuoka to Tokyo, not matched: %f", d)
	}
}
