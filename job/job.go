// Author: Emmanuel Odeke <odeke@ualberta.ca>

package job

import (
    "github.com/odeke-em/gonotif/itinerary"
)

const (
    DefaultExpiryNS = 1000000
)

type Job struct {
    id int64
    itin *itinerary.Itinerary
}

func New(id int64) *Job {
    return new(Job).Init(id)
}

func (j *Job) Init(id int64, itinArgs ...interface{}) *Job {
    j.id = id
    j.itin = itinerary.New(DefaultExpiryNS, nil, nil, nil)
    return j
}


func (j *Job) SetId(id int64) {
    j.id = id
}


func (j *Job) GetId() int64 {
    return j.id
}
