// Author: Emmanuel Odeke <odeke@ualberta.ca>

package itinerary

import (
    "time"
)

type Itinerary struct {
    meta interface{}
    expiryNanos uint64
    arrivalNanos uint64
    origin interface{}
    destination interface{}
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
