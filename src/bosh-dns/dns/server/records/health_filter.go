package records

import (
	"bosh-dns/dns/server/criteria"
	"bosh-dns/dns/server/record"
)

type healthFilter struct {
	nextFilter  reducer
	health      chan<- record.Host
	w           healthWatcher
	shouldTrack bool
	domain      string
}

type reducer interface {
	Filter(criteria.Criteria, []record.Record) []record.Record
}

type healthTracker interface {
	MonitorRecordHealth(ip, fqdn string)
}

type healthWatcher interface {
	IsHealthy(ip string) bool
}

func NewHealthFilter(nextFilter reducer, health chan<- record.Host, w healthWatcher, shouldTrack bool) healthFilter {
	return healthFilter{
		nextFilter:  nextFilter,
		health:      health,
		w:           w,
		shouldTrack: shouldTrack,
	}
}

func (q *healthFilter) Filter(crit criteria.Criteria, recs []record.Record) []record.Record {
	records := q.nextFilter.Filter(crit, recs)
	var healthyRecords, unhealthyRecords []record.Record
	for _, r := range records {
		if q.shouldTrack {
			if fqdn, ok := crit["fqdn"]; ok {
				q.health <- record.Host{IP: r.IP, FQDN: fqdn[0]}
			}
		}
		if q.w.IsHealthy(r.IP) {
			healthyRecords = append(healthyRecords, r)
		} else {
			unhealthyRecords = append(unhealthyRecords, r)
		}
	}

	healthStrategy := "0"
	if len(crit["s"]) > 0 {
		healthStrategy = crit["s"][0]
	}

	switch healthStrategy {
	case "1": // unhealthy ones
		return unhealthyRecords
	case "3": // healthy
		return healthyRecords
	case "4": // all
		return append(healthyRecords, unhealthyRecords...)
	default: // smart strategy
		if len(healthyRecords) == 0 {
			return unhealthyRecords
		}

		return healthyRecords
	}
}
