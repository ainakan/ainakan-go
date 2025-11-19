#ifndef __AUTHENTICATION_SERVICE_H__
#define __AUTHENTICATION_SERVICE_H__

#include <ainakan-core.h>

#define AINAKAN_TYPE_GO_AUTHENTICATION_SERVICE (ainakan_go_authentication_service_get_type ())
G_DECLARE_FINAL_TYPE (GoAuthenticationService, ainakan_go_authentication_service, AINAKAN, GO_AUTHENTICATION_SERVICE, GObject)

GoAuthenticationService * ainakan_go_authentication_service_new (void * callback);

#endif 