package ainakan

//#include <ainakan-core.h>
import "C"
import "unsafe"

// Spawn represents spawn of the device.
type Spawn struct {
	spawn *C.AinakanSpawn
}

// PID returns process id of the spawn.
func (s *Spawn) PID() int {
	return int(C.ainakan_spawn_get_pid(s.spawn))
}

// Identifier returns identifier of the spawn.
func (s *Spawn) Identifier() string {
	return C.GoString(C.ainakan_spawn_get_identifier(s.spawn))
}

// Clean will clean the resources held by the spawn.
func (s *Spawn) Clean() {
	clean(unsafe.Pointer(s), unrefAinakan)
}
