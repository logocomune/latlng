package latlng

import "testing"

func TestHaversineDistance(t *testing.T) {
	type args struct {
		decPos1 Position
		decPos2 Position
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test #1",
			args: args{
				decPos1: Position{
					Lat: 43.724591,
					Lng: 10.382981,
				},
				decPos2: Position{
					Lat: 41.9027835,
					Lng: 12.4963655,
				},
			},
			want: 266273.0630953111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HaversineDistance(tt.args.decPos1, tt.args.decPos2); got != tt.want {
				t.Errorf("HaversineDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphericalLawDistance(t *testing.T) {
	type args struct {
		decPos1 Position
		decPos2 Position
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test #1",
			args: args{
				decPos1: Position{
					Lat: 43.724591,
					Lng: 10.382981,
				},
				decPos2: Position{
					Lat: 41.9027835,
					Lng: 12.4963655,
				},
			},
			want: 266273.0630953112,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SphericalLawDistance(tt.args.decPos1, tt.args.decPos2); got != tt.want {
				t.Errorf("SphericalLawDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBearing(t *testing.T) {
	type args struct {
		decPos1 Position
		decPos2 Position
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test #1",
			args: args{
				decPos1: Position{
					Lat: 43.724591,
					Lng: 10.382981,
				},
				decPos2: Position{
					Lat: 41.9027835,
					Lng: 12.4963655,
				},
			},
			want: 266273,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bearing(tt.args.decPos1, tt.args.decPos2); got != tt.want {
				t.Errorf("Bearing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBearing1(t *testing.T) {
	type args struct {
		decPos1 Position
		decPos2 Position
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test #1",
			args: args{
				decPos1: Position{
					Lat: 43.724591,
					Lng: 10.382981,
				},
				decPos2: Position{
					Lat: 41.9027835,
					Lng: 12.4963655,
				},
			},
			want: 266273,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bearing(tt.args.decPos1, tt.args.decPos2); got != tt.want {
				t.Errorf("Bearing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decToDMS(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test #1",
			args: args{30.263888889},
			want: "30° 15' 50\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecToDMS(tt.args.f); got != tt.want {
				t.Errorf("DecToDMS() = %v, want %v", got, tt.want)
			}
		})
	}
}
