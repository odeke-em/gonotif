// Author: Emmanuel Odeke <odeke@ualberta.ca>

package itinerary

import (
    "fmt"
    "time"
)

type Itinerary struct {
    meta interface{}
    expiryNanos uint64
    arrivalNanos uint64
    origin interface{}
    destination interface{}
}

type Attribute int

const (
    Meta Attribute = 1
    Origin Attribute = 2
    Destination Attribute = 3
)


func (self Itinerary) String() string {
    return fmt.Sprintf("<%s:%s>::%d",
                        self.origin, self.destination, self.expiryNanos)
}

func (self *Itinerary) Init(ttl uint64, origin, dst, meta interface{}) *Itinerary {
    self.meta = meta
    self.origin = origin
    self.arrivalNanos = uint64(time.Now().UnixNano())
    self.SetTTLNano(ttl)
    self.destination = dst

    return self
}


func (self *Itinerary) SetTTLNano(ttl uint64) {
    self.expiryNanos = self.arrivalNanos + ttl
}


func New(ttl uint64, origin, dst, meta interface{}) *Itinerary {
    return new(Itinerary).Init(ttl, origin, dst, meta)
}


func (self *Itinerary) IsExpired() bool {
    return self.expiryNanos < uint64(time.Now().UnixNano())
}


func (self *Itinerary) SetOrigin(newOrigin interface{}) {
    self.origin = newOrigin
}


func (self *Itinerary) GetOrigin() interface{} {
    return self.origin
}


func (self *Itinerary) SetDestination(newDestination interface{}) {
    self.destination = newDestination
}


func (self *Itinerary) GetDestination() interface{} {
    return self.destination
}


func (self *Itinerary) GetMeta() interface{} {
    return self.meta
}

func (self *Itinerary) GetExpiry() uint64 {
    return self.expiryNanos
}


func (self *Itinerary) GetArrival() uint64 {
    return self.arrivalNanos
}


func (self *Itinerary) SetMeta(newMeta interface{}) {
    self.meta = newMeta
}


func (self *Itinerary) LessByExpiry(other *Itinerary) bool {
    return self.GetExpiry() < other.GetExpiry()
}


func (self *Itinerary) LessByArrival(other *Itinerary) bool {
    return self.GetArrival() < other.GetArrival()
}


func groupBy(attr Attribute, itins []*Itinerary) map[interface{}][]*Itinerary {
    bucketMap := make(map[interface{}][]*Itinerary)
    var retrAttr interface{}
    for _, it := range itins {
        switch attr {
        case Meta:
            retrAttr = it.meta
        case Origin:
            retrAttr = it.origin
        case Destination:
            retrAttr = it.destination
        default:
            retrAttr = nil
        }

        bucket, ok := bucketMap[retrAttr]
        if ok == false { // First time being entered
            bucket = []*Itinerary{it}
        } else {
            bucket = append(bucket, it)
        }

        bucketMap[retrAttr] = bucket
    }

    return bucketMap
}


func ClusterByDestination(itins ...*Itinerary) map[interface{}][]*Itinerary {
    return groupBy(Destination, itins)
}


func ClusterByOrigin(itins ...*Itinerary) map[interface{}][]*Itinerary {
    return groupBy(Origin, itins)
}
