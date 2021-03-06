package srv

import (
	"calendar/model"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	u := model.User{Name: "Andy"}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Andy", args: struct{ user *model.User }{user: &u},

			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnb2xlbmRhciIsInN1YiI6IkFuZHkifQ.TcqvtFjryWm2O9SoHhXjftgSW_K5WAdk9Fcxx6EEgVw",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := GenToken(tt.args.user); got != tt.want {
				t.Errorf("GenToken() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestReadConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    Config
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	type args struct {
		u   []model.User
		idx int
	}
	tests := []struct {
		name string
		args args
		want []model.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.args.u, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStart(t *testing.T) {
	type args struct {
		port int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Start(tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_auth(t *testing.T) {
	type args struct {
		h http.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := auth(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("auth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dd_Run(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name    string
		i       dd
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.Run(tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_do(t *testing.T) {
	type args struct {
		w myIntreface
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_eventsHandler(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_factory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want DataWriter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEventHandler(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_loginHandler(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_logoutHandler(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_postEventHandler(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_putEventHandler(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		request *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_registerHandlers(t *testing.T) {
	tests := []struct {
		name string
		want *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := registerHandlers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("registerHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registrationHandlers(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_userHandler(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_worker_Error(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &worker{
				name: tt.fields.name,
			}
			if got := s.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worker_Run(t *testing.T) {
	type fields struct {
		name string
	}
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &worker{
				name: tt.fields.name,
			}
			if err := s.Run(tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_worker_String(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &worker{
				name: tt.fields.name,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worker_myfucn(t *testing.T) {
	type fields struct {
		name string
	}
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &worker{
				name: tt.fields.name,
			}
			if err := s.myfucn(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("myfucn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenToken1(t *testing.T) {
	type args struct {
		user *model.User
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenToken(tt.args.user); got != tt.want {
				t.Errorf("GenToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadConfig1(t *testing.T) {
	tests := []struct {
		name    string
		want    Config
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveElement1(t *testing.T) {
	type args struct {
		u   []model.User
		idx int
	}
	tests := []struct {
		name string
		args args
		want []model.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.args.u, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStart1(t *testing.T) {
	type args struct {
		port int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Start(tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_auth1(t *testing.T) {
	type args struct {
		h http.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := auth(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("auth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dd_Run1(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name    string
		i       dd
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.Run(tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_do1(t *testing.T) {
	type args struct {
		w myIntreface
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_eventsHandler1(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_factory1(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want DataWriter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEventHandler1(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_loginHandler1(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_logoutHandler1(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_postEventHandler1(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_putEventHandler1(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		request *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_registerHandlers1(t *testing.T) {
	tests := []struct {
		name string
		want *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := registerHandlers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("registerHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registrationHandlers1(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_userHandler1(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_worker_Error1(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &worker{
				name: tt.fields.name,
			}
			if got := s.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worker_Run1(t *testing.T) {
	type fields struct {
		name string
	}
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &worker{
				name: tt.fields.name,
			}
			if err := s.Run(tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_worker_String1(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &worker{
				name: tt.fields.name,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worker_myfucn1(t *testing.T) {
	type fields struct {
		name string
	}
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &worker{
				name: tt.fields.name,
			}
			if err := s.myfucn(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("myfucn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
