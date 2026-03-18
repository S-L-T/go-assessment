package helper

import "testing"

func TestIsAuthorized(t *testing.T) {
	type args struct {
		t string
		k string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "valid",
			args:    args{
				t: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QifQ.F4Y2QCQyXiLA2UnRGZOme47lK7PbcaK51vQSl9ZeQqE",
				k: "Iamverysecret",
			},
			wantErr: false,
		},
		{
			name:    "empty key",
			args:    args{
				t: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QifQ.F4Y2QCQyXiLA2UnRGZOme47lK7PbcaK51vQSl9ZeQqE",
				k: "",
			},
			wantErr: true,
		},
		{
			name:    "empty token",
			args:    args{
				t: "",
				k: "Iamverysecret",
			},
			wantErr: true,
		},
		{
			name:    "invalid signature",
			args:    args{
				t: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QifQ.yiyw_XrbQjhJw1A-ZYoYJYRaEd89orTgY_Gj1rj_hCI",
				k: "Iamverysecret",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsAuthorized(tt.args.t, tt.args.k); (err != nil) != tt.wantErr {
				t.Errorf("IsAuthorized() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
