package ainakan

//#include <ainakan-core.h>
import "C"
import "unsafe"

// PortalMembership type is used to join portal with session.
type PortalMembership struct {
	mem *C.AinakanPortalMembership
}

// Terminate terminates the session membership
func (p *PortalMembership) Terminate() error {
	var err *C.GError
	C.ainakan_portal_membership_terminate_sync(p.mem, nil, &err)
	return handleGError(err)
}

// Clean will clean the resources held by the portal membership.
func (p *PortalMembership) Clean() {
	clean(unsafe.Pointer(p.mem), unrefAinakan)
}
