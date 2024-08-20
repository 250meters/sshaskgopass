package sshaskpass

import "testing"

type testStrBool struct {
	name string
	p    Prompt
	want bool
}

type testStrString struct {
	name string
	p    Prompt
	want string
}

func TestPrompt_IsOTP(t *testing.T) {
	tests := []testStrBool{
		{
			name: "empty",
			p:    "",
			want: false,
		},
		{
			name: "is otp",
			p:    "Please provide (user@server.tld) otp code: ",
			want: true,
		},
		{
			name: "not otp",
			p:    "user@server.tld's password: ",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsOTP(); got != tt.want {
				t.Errorf("Prompt.IsOTP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrompt_RequestedOtp(t *testing.T) {
	tests := []testStrString{
		{
			name: "empty",
			p:    "",
			want: "",
		},
		{
			name: "is otp",
			p:    "Please provide (user@server.tld) otp code: ",
			want: "user@server.tld",
		},
		{
			name: "not otp",
			p:    "user@server.tld's password: ",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.RequestedOtp(); got != tt.want {
				t.Errorf("Prompt.RequestedOtp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrompt_IsPassphraseForKey(t *testing.T) {
	tests := []testStrBool{
		{
			name: "empty",
			p:    "",
			want: false,
		},
		{
			name: "is passphrase",
			p:    "Enter passphrase for key '/path/to/key/file': ",
			want: true,
		},
		{
			name: "example dir passphrase",
			p:    "Enter passphrase for key '/home/user/.ssh/id_rsa': ",
			want: true,
		},
		{
			name: "not passphrase",
			p:    "user@server.tld's password: ",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsPassphraseForKey(); got != tt.want {
				t.Errorf("Prompt.IsPassphraseForKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrompt_RequestedPassphraseForKey(t *testing.T) {
	tests := []testStrString{
		{
			name: "empty",
			p:    "",
			want: "",
		},
		{
			name: "is passphrase",
			p:    "Enter passphrase for key '/path/to/key/file': ",
			want: "/path/to/key/file",
		},
		{
			name: "not passphrase",
			p:    "user@server.tld's password: ",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.RequestedPassphraseForKey(); got != tt.want {
				t.Errorf("Prompt.RequestedPassphraseForKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrompt_IsPassword(t *testing.T) {
	tests := []testStrBool{
		{
			name: "empty",
			p:    "",
			want: false,
		},
		{
			name: "is password",
			p:    "user@server.tld's password: ",
			want: true,
		},
		{
			name: "not password",
			p:    "Enter passphrase for key '/path/to/key/file': ",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsPassword(); got != tt.want {
				t.Errorf("Prompt.IsPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrompt_RequestedPassword(t *testing.T) {
	tests := []testStrString{
		{
			name: "empty",
			p:    "",
			want: "",
		},
		{
			name: "is password",
			p:    "user@server.tld's password: ",
			want: "user@server.tld",
		},
		{
			name: "not password",
			p:    "Enter passphrase for key '/path/to/key/file': ",
			want: "Enter passphrase for key ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.RequestedPassword(); got != tt.want {
				t.Errorf("Prompt.RequestedPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
