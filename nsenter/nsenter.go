package nsenter

/*
#define _GNU_SOURCE
#include <unistd.h>
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>

__attribute__((constructor)) void enter_namespace(void) {
	char *parcel_pid;
	parcel_pid = getenv("parcel_pid");
	if (parcel_pid) {
		//fprintf(stdout, "got parcel_pid=%s\n", parcel_pid);
	} else {
		//fprintf(stdout, "missing parcel_pid env skip nsenter");
		return;
	}
	char *parcel_cmd;
	parcel_cmd = getenv("parcel_cmd");
	if (parcel_cmd) {
		//fprintf(stdout, "got parcel_cmd=%s\n", parcel_cmd);
	} else {
		//fprintf(stdout, "missing parcel_cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };

	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", parcel_pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY);

		if (setns(fd, 0) == -1) {
			//fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			//fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}
	int res = system(parcel_cmd);
	exit(0);
	return;
}
*/
import "C"
