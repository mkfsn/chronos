package jobs

import (
	"testing"
)

func TestUpdateState(t *testing.T) {
	cases := []struct {
		job    *aburiyaJob
		state  map[string]string
		except bool
	}{
		{
			job:    &aburiyaJob{state: make(map[string]string)},
			state:  make(map[string]string),
			except: false,
		},
		{
			job:    &aburiyaJob{state: map[string]string{"x": "y"}},
			state:  map[string]string{},
			except: true,
		},
		{
			job:    &aburiyaJob{state: map[string]string{}},
			state:  map[string]string{"x": "y"},
			except: true,
		},
		{
			job:    &aburiyaJob{state: map[string]string{"a": "b"}},
			state:  map[string]string{"c": "d"},
			except: true,
		},
		{
			job:    &aburiyaJob{state: map[string]string{"a": "b", "c": "d"}},
			state:  map[string]string{"c": "d"},
			except: true,
		},
		{
			job:    &aburiyaJob{state: map[string]string{"a": "b"}},
			state:  map[string]string{"c": "d", "a": "b"},
			except: true,
		},
		{
			job:    &aburiyaJob{state: map[string]string{"a": "b", "c": "d"}},
			state:  map[string]string{"c": "d", "a": "b"},
			except: false,
		},
	}

	for _, testcase := range cases {
		if testcase.except != testcase.job.updateState(testcase.state) {
			if testcase.except {
				t.Error("state should be update")
			} else {
				t.Error("state should not be update")
			}
		}
	}
}
