package utils

import (
	"errors"
	"fmt"
	"time"
)	

func TimeDelta(t2 *time.Time) (float64, error) {

	// Only for my physical machine...
	t1 := (time.Now().UTC()).Add(500 * time.Millisecond)
	// t1 := time.Now().UTC() //
	fmt.Println(t1.Format("15:04:05.000000000") + "    " + fmt.Sprintf((*t2).Format("15:04:05.000000000")))
	if (*t2).IsZero() {
		return 0, errors.New("Error: func TimeDelta(t2 *time.Time).")
	}
    duration := t1.Sub(*t2).Seconds()
    return duration, nil
}