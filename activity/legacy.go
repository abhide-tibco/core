package activity

import (
	"fmt"

	"github.com/project-flogo/core/support/log"
)

type void struct{}

var (
	hasLegacy     = false
	empty         void
	legacyTracker = make(map[string]void)
)

// DEPRECATED
func HasLegacyActivities() bool {
	return hasLegacy
}

// DEPRECATED
func IsLegacyActivity(ref string) bool {
	_, ok := legacyTracker[ref]
	return ok
}

// DEPRECATED
func LegacyRegister(ref string, act Activity) error {

	if ref == "" {
		return fmt.Errorf("'ref' must be specified when registering")
	}

	if act == nil {
		return fmt.Errorf("cannot register 'nil' activity")
	}
	activities := GetActivities()
	if _, dup := activities[ref]; dup {
		return fmt.Errorf("activity already registered: %s", ref)
	}

	log.RootLogger().Debugf("Registering legacy activity: %s", ref)

	hasLegacy = true
	AddtoActivities(ref, act)
	legacyTracker[ref] = empty
	activityLogger := log.CreateLoggerFromRef(GetRootLogger(), "activity", ref)
	AddtoActivityLoggers(ref, activityLogger)
	return nil
}

type LegacyCtx interface {

	// GetOutput gets the value of the specified output attribute
	GetOutput(name string) interface{}
	GetSetting(name string) (value interface{}, exists bool)
}
