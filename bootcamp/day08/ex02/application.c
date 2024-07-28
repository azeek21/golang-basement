#include <X11/X.h>
#include <X11/Xlib.h>
#include <err.h>
#include "application.h"

void MakeWindowWithTitle(int x,int y, unsigned int width, unsigned int height,char* title) {
	Window  win;
	Window 	root;
	Display *dsp;
	int 	screen;
        XEvent  ev;
	int running = 1;

	if ((dsp = XOpenDisplay(NULL)) == NULL) {
		err(1, "Can't open display");
	}

	// get default screen and root window
	screen = DefaultScreen(dsp);
	root = RootWindow(dsp, screen);

	// create app window
	win = XCreateSimpleWindow(dsp, root, x, y, width, height, 10, BlackPixel(dsp, screen), WhitePixel(dsp, screen));

	// map app window to xserver
	XMapWindow(dsp, win);
	XStoreName(dsp, win, title);

	// main app loop
	while (running) {
		XNextEvent(dsp, &ev);
		// do smth with event but we don't care
	}

	// exit main loop. Cleanup
	XUnmapWindow(dsp, win);
	XDestroyWindow(dsp, win);
	XCloseDisplay(dsp);
}

