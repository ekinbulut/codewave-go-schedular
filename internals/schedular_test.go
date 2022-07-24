package schedular

import (
	"testing"
	"time"

	"github.com/ekinbulut/go-schedular/svc/models"
)

func TestSchedular_AddJob(t *testing.T) {
	type fields struct {
		Jobs   []*models.Job
		ticker *time.Ticker
	}
	type args struct {
		job *models.Job
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "add job",
			fields: fields{
				Jobs:   make([]*models.Job, 0),
				ticker: time.NewTicker(time.Duration(5) * time.Second),
			},
			args: args{
				job: NewJob("job1", func(c chan string) {
					c <- "executed"
				}),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schedular{
				Jobs:   tt.fields.Jobs,
				ticker: tt.fields.ticker,
			}
			s.AddJob(tt.args.job)
			if got := s.Size(); got != tt.want {
				t.Errorf("Schedular.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func NewJob(name string, f func(c chan string)) *models.Job {
	m := &models.Job{
		Name: name,
	}
	m.Append(f)
	return m
}

func TestSchedular_Size(t *testing.T) {
	type fields struct {
		Jobs   []*models.Job
		ticker *time.Ticker
	}
	type args struct {
		job *models.Job
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "size",
			fields: fields{
				Jobs:   make([]*models.Job, 0),
				ticker: time.NewTicker(time.Duration(5) * time.Second),
			},
			args: args{
				job: NewJob("job1", func(c chan string) {
					c <- "executed"
				}),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schedular{
				Jobs:   tt.fields.Jobs,
				ticker: tt.fields.ticker,
			}
			s.AddJob(tt.args.job)
			if got := s.Size(); got != tt.want {
				t.Errorf("Schedular.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
