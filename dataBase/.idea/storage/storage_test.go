package storage
import ("testing"
	"github.com")

func TestNearest(t *testing.T) {
	s := New()
	s.Set(123, &Driver{
		ID: 123,
		LastLocation: Location{
			Lat: 1,
			Lon: 1,
		},
	})
	s.Set(666, &Driver{
		ID: 666,
		LastLocation: Location{
			Lat: 42.875799,
			Lon: 74.588279,
		},
	})
	drivers := s.Nearest(1000, 42.876420, 74.588332)
	assert.Equal(t, len(drivers), 1)
}

func BenchmarkNearest(b *testing.B) {
	s := New()
	for i := 0; i < 100; i++ {
		s.Set(i, &Driver{
			ID: i,
			LastLocation: Location{
				Lat: float64(i),
				Lon: float64(i),
			},
		})
	}
	for i := 0; i < b.N; i++ {
		s.Nearest(1000, 123, 123)
	}
}