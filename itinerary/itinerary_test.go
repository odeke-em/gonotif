// Author: Emmanuel Odeke <odeke@ualberta.ca>

package itinerary
import (
    "fmt"
    "testing"
)

func TestInit(t *testing.T) {
    dst := "Cavendish"
    origin := "Staines"
    expiryNS := 10000000

    it := New(100000000, origin, dst, nil)
    if it == nil {
        t.Errorf("Expecting non nil itinerary")
    } else if it.GetExpiry() < 0 {
        t.Errorf("Quick overflow detection")
    }

    retrOrigin := it.GetOrigin()
    if retrOrigin != origin {
        t.Errorf("Expected an origin of %s got instead: %s", origin, retrOrigin)
    }

    retrDst := it.GetDestination()
    if retrDst != dst {
        t.Errorf("Expected a destination of %s got instead: %s", dst, retrDst)
    }

    // Swap of the destination and origin
    it.SetOrigin(dst)

    retrOrigin = it.GetOrigin()
    if retrOrigin != dst {
        t.Errorf("Performed an origin swap and expected %s but instead got: %s",
                                                                        dst, retrOrigin)
    }

    it.SetDestination(origin)
    retrDst = it.GetDestination()
    if retrDst != origin {
        t.Errorf("Performed a dst swap and expected %s but instead got: %s",
                                                                    origin, retrDst)
    }

    if it.IsExpired() == true {
        t.Errorf("Wow, not a naturally expiry of %d ns expected already", expiryNS)
    }

    it.SetTTLNano(0)
    isExpired := it.IsExpired()
    if isExpired != true {
        t.Errorf("Set ttl to 0 but instead got: %d", it.GetExpiry())
    }
}


func TestComparisons(t *testing.T) {
    it1 := New(1, "Fiji", "Atlanta", "Fij-Atl")
    it2 := New(9, "Edmonton", "Southampton", "Edm-Stn")

    if it2.LessByExpiry(it1) != false {
        t.Errorf("Edm-Stn left later than Fij-Atl!")
    } else if it1.LessByExpiry(it2) != true {
        t.Errorf("Fij-Atl left earlier than  Edm-Stn!")
    } else if it1.LessByExpiry(it1) != false {
        t.Errorf("Comparing self should return false")
    }

    it1.SetTTLNano(10000)
    it2.SetTTLNano(1000)

    if it2.LessByExpiry(it1) != true {
        t.Errorf("Switched up the ttls, Edm-Stn should expire earlier than Fij-Atl!")
    } else if it1.LessByExpiry(it2) != false {
        t.Errorf("Switched up the ttls, Edm-Stn should expire after Fij-Atl!")
    } else if it1.LessByExpiry(it1) != false {
        t.Errorf("Comparing self should return false")
    }
}

func TestClusteringByOrigin(t *testing.T) {
    destClustering := ClusterByOrigin(
        New(800, "Edmonton", "Atlanta", "YEG-ATL"),
        New(6880,  "Oslo", "Sacramento", "TRF-SCN"),
        New(6880,  "Edmonton", "Calgary", "YEG-YYZ"),
        New(4590,  "Edmonton", "AbuDhabi", "YYZ-AUH"),
        New(8790,  "Kilimanjaro", "Amsterdam", "JRO-AMS"),
        New(9990,  "Kilimanjaro", "Amsterdam", "JRO-AMS"),
        New(590,  "Kilimanjaro", "Jomo-Kenyatta", "JRO-NBO"),
        New(190,  "Anchorage", "Los Angeles", "ANC-LAX"),
        New(990,  "Luxemborg", "Los Angeles", "LUX-LAX"),
        New(6890,  "Amsterdam", "Ontario", "AMS-ONT"),
        New(7890,  "Amsterdam", "Los Angeles", "AMS-LAX"),
    )

    fromKilimanjaro, ok := destClustering["Kilimanjaro"]
    if ok != true {
        t.Errorf("Kilimanjaro origins are expected")
    } else if len(fromKilimanjaro) < 2 {
        t.Errorf("Kilimanjaro had more than one origin")
    }

    _, illOK := destClustering["illuminati"]
    if illOK != false {
        t.Errorf("No such key was entered previously")
    }

    fmt.Printf("fromKilimanjaro: %v\n", fromKilimanjaro)
}


func TestClusteringByDestination(t *testing.T) {
    destClustering := ClusterByDestination(
        New(80, "Greenland", "Uganda", nil),
        New(10,  "Alaska", "Russia", "AlaRus"),
        New(4,  "Boston", "Seattle", "BosSea"),
        New(1,  "Vancouver", "Seattle", "VanSea"),
    )

    toSeattle, ok := destClustering["Seattle"]
    if ok != true {
        t.Errorf("Seattle final destinations are expected")
    } else if len(toSeattle) < 2 {
        t.Errorf("Seattle as a final destination was present more than once")
    }

    _, illOK := destClustering["NewOrleans"]
    if illOK != false {
        t.Errorf("No such key was entered previously")
    }

    fmt.Printf("toSeattle: %v\n", toSeattle)
}
